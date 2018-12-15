package database

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

// constant
var (
	BroadcastTime = [6]string{"2013.1.18-2013.4.12", "2014.1.3-2014.4.11", "2015.1.2-2015.4.3",
		"2016.1.15-2016.4.15", "2017.1.21-2017.4.22", "2018.1.12-2018.4.20"}
	WINNER = [6]string{"羽泉", "韩磊", "韩红", "李玟", "林忆莲", "Jessie J"}
	HOSTS  = [6]string{"胡海泉/陈羽凡/沙宝亮/何炅/汪涵", "张宇/廖凡/胡海泉/何炅/汪涵", "古巨基/孙楠/汪涵/沈梦辰",
		"李克勤/何炅/金志文等", "各位歌手或其音乐合伙人", "张韶涵/何炅"}
	TITLES = [6]string{"我是歌手第一季", "我是歌手第二季", "我是歌手第三季", "我是歌手第四季",
		"我是歌手第五季（歌手2017）", "我是歌手第六季（歌手2018）"}
	ROOT = "http://127.0.0.1/singerapi/api/"
)

// CreateDB create DB from raw data.
func CreateDB(filePath string) {
	// create the bolt database
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		// create seasons bucket
		seasons, err := tx.CreateBucket([]byte("seasons"))
		if err != nil {
			log.Printf("create seasons bucket error : %s", err)
			return err
		}
		seasons.Put([]byte(TITLES[0]), []byte(ROOT+"season/1/"))
		seasons.Put([]byte(TITLES[1]), []byte(ROOT+"season/2/"))
		seasons.Put([]byte(TITLES[2]), []byte(ROOT+"season/3/"))
		seasons.Put([]byte(TITLES[3]), []byte(ROOT+"season/4/"))
		seasons.Put([]byte(TITLES[4]), []byte(ROOT+"season/5/"))
		seasons.Put([]byte(TITLES[5]), []byte(ROOT+"season/6/"))

		// create 6 season buckets
		for i := 0; i < 6; i++ {
			season, err := tx.CreateBucket([]byte("season" + string(i+1)))
			if err != nil {
				log.Printf("create season bucket %d error : %s", i, err)
				return err
			}
			season.Put([]byte("title"), []byte(TITLES[i]))
			season.Put([]byte("BroadcastTime"), []byte(BroadcastTime[i]))
			season.Put([]byte("hosts"), []byte(HOSTS[i]))
			season.Put([]byte("winner"), []byte(WINNER[i]))
			season.Put([]byte("url"), []byte(ROOT+"season/"+string(i)+"/"))
			season.Put([]byte("singers"), []byte(ROOT+"season/"+string(i)+"/singers/"))
			season.Put([]byte("songs"), []byte(ROOT+"season/"+string(i)+"/songs/"))
		}

		// create buckets to store songs info and singers
		_, err = tx.CreateBucket([]byte("songAndAlbum"))
		_, err = tx.CreateBucket([]byte("songAndSinger"))
		_, err = tx.CreateBucket([]byte("songAndDuration"))
		_, err = tx.CreateBucket([]byte("songAndSeason"))
		_, err = tx.CreateBucket([]byte("singerAndSong"))
		if err != nil {
			fmt.Printf("create bucket error : %s", err)
			return err
		}

		// open the rawdata file
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer file.Close()

		// read data from txt file and put it into database
		br := bufio.NewReader(file)
		seasonName := ""
		seasonCnt := 0

		for {
			line, _, end := br.ReadLine()
			if end == io.EOF {
				break
			}

			// the data format : season;song;duration;singer;album
			attr := strings.Split(string(line), ";")
			if attr[0] != seasonName {
				seasonName = TITLES[seasonCnt]
				seasonCnt++
			}

			season := tx.Bucket([]byte("season" + string(seasonCnt)))

			// songs bucket in each season
			songs, err := season.CreateBucketIfNotExists([]byte("songs"))
			if err != nil {
				fmt.Printf("create songs bucket in season %d error : %s", seasonCnt, err)
				return err
			}
			songs.Put([]byte(attr[1]), []byte(ROOT+"songs/?name="+attr[1]))

			// singer bucket in each season
			singer, err := season.CreateBucketIfNotExists([]byte("singer"))
			if err != nil {
				fmt.Printf("create singer bucket in season %d error : %s", seasonCnt, err)
				return err
			}
			// split singers
			singers := make([]string, 0, 10)
			flag := strings.Contains(attr[3], "/")
			if flag == true {
				singers = strings.Split(string(attr[3]), "/")
			} else {
				singers = append(singers, attr[3])
			}

			for i := 0; i < len(singers); i++ {
				if singers[i] == "" {
					break
				}
				// put singer info into season bucket
				singer.Put([]byte(singers[i]), []byte(ROOT+"singers/?name="+singers[i]))
				// put the singer and song into singerAndSong bucket in root
				singerAndSong := tx.Bucket([]byte("singerAndSong"))
				singerName, err := singerAndSong.CreateBucketIfNotExists([]byte(singers[i]))
				if err != nil {
					fmt.Printf("create name bucket in singerAndSong error : %s", err)
					return err
				}
				singerName.Put([]byte(attr[1]), []byte(ROOT+"songs/?name="+attr[1]))
			}

			// handle the song info and put it into corrsponding bucket
			songAndSeason := tx.Bucket([]byte("songAndSeason"))
			songAndSeason.Put([]byte(attr[1]), []byte(string(seasonCnt)))

			songAndDuration := tx.Bucket([]byte("songAndDuration"))
			songAndDuration.Put([]byte(attr[1]), []byte(attr[2]))

			songAndSinger := tx.Bucket([]byte("songAndSinger"))
			songAndSinger.Put([]byte(attr[1]), []byte(attr[3]))

			songAndAlbum := tx.Bucket([]byte("songAndAlbum"))
			songAndAlbum.Put([]byte(attr[1]), []byte(attr[4]))

		}

		return nil
	})

}

// Query query resource from db
func Query(api int, param string) string {
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var jsonStr string
	switch api {
	case 1:
		//api : /singerapi/api/seasons/{id}
		db.View(func(tx *bolt.Tx) error {
			season := tx.Bucket([]byte("season" + param))

			// title := season.Get([]byte("title"))
			// BroadcastTime := season.Get([]byte("BroadcastTime"))
			// winner := season.Get([]byte("winner"))
			// hosts  := season.Get([]byte("hosts"))
			// url  := season.Get([]byte("url"))
			// singers  := season.Get([]byte("singers"))
			// songs  := season.Get([]byte("songs"))

			cur := season.Cursor()

			jsonStr = "{"
			for k, v := cur.First(); k != nil; k, v = cur.Next() {
				jsonStr += "\"" + string(k) + "\" : \"" + string(v) + "\","
			}
			jsonStr += "}"

			return nil
		})
	case 2:
		//api : /singerapi/api/seasons/{id}/singers/
		db.View(func(tx *bolt.Tx) error {
			season := tx.Bucket([]byte("season" + param))
			singers := season.Bucket([]byte("singers"))

			cur := singers.Cursor()

			jsonStr = "{"
			for k, v := cur.First(); k != nil; k, v = cur.Next() {
				jsonStr += "\"" + string(k) + "\" : \"" + string(v) + "\","
			}
			jsonStr += "}"
			return nil
		})
	case 3:
		//api : /singerapi/api/seasons/{id}/songs
		db.View(func(tx *bolt.Tx) error {
			season := tx.Bucket([]byte("season" + param))
			songs := season.Bucket([]byte("songs"))

			cur := songs.Cursor()

			jsonStr = "{"
			for k, v := cur.First(); k != nil; k, v = cur.Next() {
				jsonStr += "\"" + string(k) + "\" : \"" + string(v) + "\","
			}
			jsonStr += "}"
			return nil
		})
	case 4:
		//api : /singerapi/api/singers/?singer={singer}
		db.View(func(tx *bolt.Tx) error {
			singers := tx.Bucket([]byte("singers"))
			cur := singers.Cursor()

			jsonStr = "{"
			for k, v := cur.Seek([]byte(param)); k != nil && string(k) == param; k, v = cur.Next() {
				jsonStr += "\"" + string(k) + "\" : \"" + string(v) + "\","
			}
			jsonStr += "}"

			return nil
		})
	case 5:
		//api : /singerapi/api/songs/?song={song}
		db.View(func(tx *bolt.Tx) error {
			songs := tx.Bucket([]byte("songs"))
			cur := songs.Cursor()

			jsonStr = "{"
			for k, v := cur.Seek([]byte(param)); k != nil && string(k) == param; k, v = cur.Next() {
				jsonStr += "\"" + string(k) + "\" : \"" + string(v) + "\","
			}
			jsonStr += "}"

			return nil
		})
	default:
		return ""

	}
	return ""
}
