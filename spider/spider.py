#!/usr/bin/env python
# -*- coding:utf-8 -*-

import json
import re
from multiprocessing.pool import Pool
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from bs4 import BeautifulSoup

import requests
from requests.exceptions import RequestException
num = 1

browser = webdriver.Chrome()
wait = WebDriverWait(browser, 6)


def get_page(url):
    global num
    try:
        browser.get(url)
        browser.switch_to.frame('g_iframe')
        response = browser.page_source
        print("第" + str(num) + "歌单")
        num += 1
        return response
    except RequestException:
        return None


def parse_page(html_cont):
    pattern = re.compile(r'.*?<span class="u-dur ">(.*?)</span>.*?'
                         + 'data-res-name="(.*?)" data-res-author="(.*?)" data-res-pic.*?'
                         + '<a href="/album\?id=.*?" title="(.*?)">', re.S)
    items = re.findall(pattern, html_cont)
    soup = BeautifulSoup(browser.page_source, 'lxml')
    info = soup.select('.f-ff2.f-brk')
    info = info[0].get_text()
    for item in items:
        yield info + ";" + item[1] + ";" + item[0] + ";" + item[2] + ";" + item[3]


def write2file(content):
    with open('result.txt', 'a', encoding='utf-8') as f:
        # f.write(json.dumps(content, ensure_ascii=False) + '\n')
        f.write(content + '\n')
        f.close()


def main(offset):
    url = 'https://music.163.com/playlist?id=' + str(offset)
    html_cont = get_page(url)
    for item in parse_page(html_cont):
        write2file(item)


if __name__ == '__main__':
    for i in [4981163, 158743851, 42180665, 154852370, 706104823, 2055513280]:
        main(i)
    # 使用进程池，可能存在抓取后信息写入丢失情况,需加锁
    # pool = Pool()
    # pool.map(main, [i for i in range(2110, 2120)])
    # pool.close()
    # pool.join()
    # r = requests.get('https://music.163.com/artist?id=2110')
    # print(r.text)


