package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/singerapi/singer-server/rawdata"
)

// constant
const (
	PORT string = "8080"
	PATH string = "rawdata/rawdata.txt"
)

func main() {
	// create the bolt database
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(rawdata.PreprocessRawData)

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("season"))
		cur := bucket.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			fmt.Printf("%s %s\n", k, v)
		}
		return nil
	})
}
