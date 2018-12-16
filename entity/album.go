package entity

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

// AlbumType for bolt dababase key refers to bucket
var AlbumType = "album"

// Album implements interface of Entity
type Album struct {
	ID        int
	Title     string
	SeasonsID []int
	SongsID   []int
	SingersID []int
}

// NewAlbum create new Season entity
func NewAlbum(id int, ti string, seasonsID []int, songsID []int, singersID []int) Album {
	return Album{
		id,
		ti,
		seasonsID,
		songsID,
		singersID,
	}
}

// Type return AlbumType
func (s *Album) Type() string {
	return AlbumType
}

// Encode encode
func (s *Album) Encode() ([]byte, error) {
	return json.Marshal(s)
}

// Decode decode
func (s *Album) Decode(b []byte) error {
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	return nil
}

// HasSinger check singer id list contain id
func (s *Album) HasSinger(id int) bool {
	return hasID(s.SingersID, id)
}

// HasSong check song id list contain id
func (s *Album) HasSong(id int) bool {
	return hasID(s.SongsID, id)
}

// HasSeason check season id list contain id
func (s *Album) HasSeason(id int) bool {
	return hasID(s.SeasonsID, id)
}

// AddSongID check duplicate singer id
func (s *Album) AddSongID(id int) {
	if !s.HasSong(id) {
		s.SongsID = append(s.SongsID, id)
	}
}

// AddSingerID check duplicate singer id
func (s *Album) AddSingerID(id int) {
	if !s.HasSinger(id) {
		s.SingersID = append(s.SingersID, id)
	}
}

// AddSeasonID check duplicate singer id
func (s *Album) AddSeasonID(id int) {
	if !s.HasSeason(id) {
		s.SeasonsID = append(s.SeasonsID, id)
	}
}

// APIFormatConstruct construct struct with api prefix
func (s *Album) APIFormatConstruct(HostPrefix string) interface{} {
	var apiItem = struct {
		ID        int      `json:"id"`
		Title     string   `json:"title"`
		SeasonsID []string `json:"seasons"`
		SongsID   []string `json:"songs"`
		SingersID []string `json:"singers"`
		URL       string   `json:"url"`
	}{
		s.ID,
		s.Title,
		[]string{},
		[]string{},
		[]string{},
		HostPrefix + "/" + s.Type() + "s/" + strconv.Itoa(s.ID) + "/",
	}
	for _, ID := range s.SeasonsID {
		apiItem.SeasonsID = append(apiItem.SeasonsID, HostPrefix+"/"+SeasonType+"s/"+strconv.Itoa(ID)+"/")
	}
	for _, ID := range s.SongsID {
		apiItem.SongsID = append(apiItem.SongsID, HostPrefix+"/"+SongType+"s/"+strconv.Itoa(ID)+"/")
	}
	for _, ID := range s.SingersID {
		apiItem.SingersID = append(apiItem.SingersID, HostPrefix+"/"+SingerType+"s/"+strconv.Itoa(ID)+"/")
	}
	return apiItem
}

// WriteAlbumToBucket write entity to bucket with key of entity ID
func WriteAlbumToBucket(entityList []Album, bucket *bolt.Bucket) {
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
