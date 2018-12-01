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


func createDB (filePath string) {
	// open the rawdata file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// create the bolt database
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// read data from txt file and put it into database
	br := bufio.NewReader(file)
	seasonName := ""
	seasonCnt := 0
	for {
		line, _, end:= br.ReadLine()
		if end == io.EOF {
			break
		}

		attr := strings.Split(string(line), ";")
		if attr[0] != seasonName {
			seasonName = attr[0]
			seasonCnt ++
		}

		db.Update(func(tx *bolt.Tx) error {
			// create bucket to store 6 seasons
			season, err := tx.CreateBucketIfNotExists([]byte(string(seasonCnt)))
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
