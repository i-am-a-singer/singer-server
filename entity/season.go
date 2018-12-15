package entity

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

// SeasonType for bolt dababase key refers to bucket
var SeasonType = "season"

// Season implements interface of Entity
type Season struct {
	ID        int
	Title     string
	SingersID []int
	SongsID   []int
	AlbumsID  []int
}

// NewSeason create new Season entity
func NewSeason(id int, ti string, singersID []int, songsID []int, albumsID []int) Season {
	return Season{
		id,
		ti,
		singersID,
		songsID,
		albumsID,
	}
}

// Type return SeasonType
func (s *Season) Type() string {
	return SeasonType
}

// Encode encode
func (s *Season) Encode() ([]byte, error) {
	return json.Marshal(*s)
}

// Decode decode
func (s *Season) Decode(b []byte) error {
	if err := json.Unmarshal(b, *s); err != nil {
		return err
	}
	return nil
}

// HasSinger check singer id list contain id
func (s *Season) HasSinger(id int) bool {
	return hasID(s.SingersID, id)
}

// HasSong check song id list contain id
func (s *Season) HasSong(id int) bool {
	return hasID(s.SongsID, id)
}

// HasAlbum check album id list contain id
func (s *Season) HasAlbum(id int) bool {
	return hasID(s.AlbumsID, id)
}

// AddSingerID check duplicate singer id
func (s *Season) AddSingerID(id int) {
	if !s.HasSinger(id) {
		s.SingersID = append(s.SingersID, id)
	}
}

// AddSongID check duplicate singer id
func (s *Season) AddSongID(id int) {
	if !s.HasSong(id) {
		s.SongsID = append(s.SongsID, id)
	}
}

// AddAlbumID check duplicate singer id
func (s *Season) AddAlbumID(id int) {
	if !s.HasAlbum(id) {
		s.AlbumsID = append(s.AlbumsID, id)
	}
}

// WriteSeasonToBucket write entity to bucket with key of entity ID
func WriteSeasonToBucket(entityList []Season, bucket *bolt.Bucket) {
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
