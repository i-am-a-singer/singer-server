package database

import (
	"log"
	"bufio"
	"io"
	"os"
	"strings"
	"fmt"

	"github.com/boltdb/bolt"
)

const (
	BROADCAST_TIME := [6]string{"2013.1.18-2013.4.12", "2014.1.3-2014.4.11", "2015.1.2-2015.4.3", 
								"2016.1.15-2016.4.15", "2017.1.21-2017.4.22", "2018.1.12-2018.4.20"}
	WINNER := [6]string{"羽泉","韩磊","韩红","李玟","林忆莲","Jessie J"}
	HOSTS := [6]string{"胡海泉/陈羽凡/沙宝亮/何炅/汪涵","张宇/廖凡/胡海泉/何炅/汪涵","古巨基/孙楠/汪涵/沈梦辰",
						"李克勤/何炅/金志文等","各位歌手或其音乐合伙人","张韶涵/何炅"}
	SEASONS := [6]string{"我是歌手第一季","我是歌手第二季","我是歌手第三季","我是歌手第四季",
						"我是歌手第五季（歌手2017）","我是歌手第六季（歌手2018）"}
)

func createDB (filePath string) {
	// create the bolt database
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	db.Update(func(tx *bolt.Tx) error {
		// create seasons bucket
		seasons, err := tx.CreateBucket([]type("seasons"))
		if err != nil {
			fmt.Printf("create seasons bucket error : %s", err)
			return err
		}
		seasons.Put([]type(SEASONS[0]), []type("http://127.0.0.1/singerapi/api/season/1/"))
		seasons.Put([]type(SEASONS[1]), []type("http://127.0.0.1/singerapi/api/season/2/"))
		seasons.Put([]type(SEASONS[2]), []type("http://127.0.0.1/singerapi/api/season/3/"))
		seasons.Put([]type(SEASONS[3]), []type("http://127.0.0.1/singerapi/api/season/4/"))
		seasons.Put([]type(SEASONS[4]), []type("http://127.0.0.1/singerapi/api/season/5/"))
		seasons.Put([]type(SEASONS[5]), []type("http://127.0.0.1/singerapi/api/season/6/"))

		// create 6 season buckets
		for i := 0; i < 6; i++ {
			season, err := tx.CreateBucket([]type("season"+string(i+1))
			if err != nil {
				fmt.Printf("create season bucket %d error : %s", i, err)
				return err
			}
			season.Put([]type("broadcast_time"), []type(BROADCAST_TIME[i]))
			season.Put([]type("hosts"), []type(HOSTS[i]))
			season.Put([]type("winner"), []type(WINNER[i]))
			season.Put([]type("url"), []type("http://127.0.0.1/singerapi/api/season/"+string(i)+"/"))
			season.Put([]type("singers"), []type("http://127.0.0.1/singerapi/api/season/"+string(i)+"/singers/"))
			season.Put([]type("songs"), []type("http://127.0.0.1/singerapi/api/season/"+string(i)+"/songs/"))
		}
		
		return nil
	})

	

	// open the rawdata file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// read data from txt file and put it into database
	br := bufio.NewReader(file)
	seasonName := ""
	seasonNum := 0
	for {
		line, _, end:= br.ReadLine()
		if end == io.EOF {
			break
		}

		attr := strings.Split(string(line), ";")
		if attr[0] != seasonName {
			seasonName = attr[0]
			seasonNum ++
		}

		db.Update(func(tx *bolt.Tx) error {
			// create bucket to store 6 seasons
			season, err := tx.CreateBucketIfNotExists([]byte(string(seasonNum)))
			if err != nil {
				fmt.Printf("create season bucket error : %s", err)
				return err
			}

			// create a bucket in each season to store song and duration
			duaration, err := season.CreateBucketIfNotExists([]byte("duaration"))
			if err != nil {
				fmt.Printf("create duaration bucket error : %s", err)
				return err
			}
			_ = duaration.Put([]byte(attr[2]), []byte(attr[4]))

			// create a bucket in each season to store song and singer
			singer, err := season.CreateBucketIfNotExists([]byte("singer"))
			if err != nil {
				fmt.Printf("create singer bucket error : %s", err)
				return err
			}
			_ = singer.Put([]byte(attr[1]), []byte(attr[3]))
			
			// create a bucket in each season to store song and album
			album, err := season.CreateBucketIfNotExists([]byte("album"))
			if err != nil {
				fmt.Printf("create album bucket error : %s", err)
				return err
			}
			_ = album.Put([]byte(attr[2]), []byte(attr[4]))

			return nil
		})
	}
}
