package entity

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

// SongType for bolt dababase key refers to bucket
var SongType = "song"

// Song implements interface of Entity
type Song struct {
	ID        int
	Title     string
	Duration  string
	SeasonsID []int
	SingersID []int
	AlbumsID  []int
}

// NewSong create new Season entity
func NewSong(id int, ti string, du string, seasonsID []int, singersID []int, albumsID []int) Song {
	return Song{
		id,
		ti,
		du,
		seasonsID,
		singersID,
		albumsID,
	}
}

// Type return SongType
func (s *Song) Type() string {
	return SongType
}

// Encode encode
func (s *Song) Encode() ([]byte, error) {
	return json.Marshal(*s)
}

// Decode decode
func (s *Song) Decode(b []byte) error {
	if err := json.Unmarshal(b, *s); err != nil {
		return err
	}
	return nil
}

// HasSinger check singer id list contain id
func (s *Song) HasSinger(id int) bool {
	return hasID(s.SingersID, id)
}

// HasAlbum check album id list contain id
func (s *Song) HasAlbum(id int) bool {
	return hasID(s.AlbumsID, id)
}

// HasSeason check season id list contain id
func (s *Song) HasSeason(id int) bool {
	return hasID(s.SeasonsID, id)
}

// AddSeasonID check duplicate singer id
func (s *Song) AddSeasonID(id int) {
	if !s.HasSeason(id) {
		s.SeasonsID = append(s.SeasonsID, id)
	}
}

// AddAlbumID check duplicate singer id
func (s *Song) AddAlbumID(id int) {
	if !s.HasAlbum(id) {
		s.AlbumsID = append(s.AlbumsID, id)
	}
}

// AddSingerID check duplicate singer id
func (s *Song) AddSingerID(id int) {
	if !s.HasSinger(id) {
		s.SingersID = append(s.SingersID, id)
	}
}

// WriteSongToBucket write entity to bucket with key of entity ID
func WriteSongToBucket(entityList []Song, bucket *bolt.Bucket) {
	for i := range entityList {
		data, err := entityList[i].Encode()
		if err != nil {
			log.Fatal(err)
		}
		if err := bucket.Put([]byte(strconv.Itoa(entityList[i].ID)), data); err != nil {
			log.Fatal(err)
		}
	}
}
