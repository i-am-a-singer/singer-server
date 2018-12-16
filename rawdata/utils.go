package rawdata

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/singerapi/singer-server/entity"
)

// RawDataPath is path to rawdata
const RawDataPath = "./rawdata/rawdata.txt"

// LoadRawData preprocessing and load rawdata from rawdata.txt
// And it will be refered in bolt database Update method
// Help convert all seasons, singers, songs and albums rawdata to bolt database bucket.
func LoadRawData(tx *bolt.Tx) error {
	// creating seasons, singers, songs, albums bucket
	seasonBk, _ := tx.CreateBucket([]byte(entity.SeasonType))
	songBk, _ := tx.CreateBucket([]byte(entity.SongType))
	singerBk, _ := tx.CreateBucket([]byte(entity.SingerType))
	albumBk, _ := tx.CreateBucket([]byte(entity.AlbumType))

	seasons := []entity.Season{}
	songs := []entity.Song{}
	singers := []entity.Singer{}
	albums := []entity.Album{}

	seasonsTitleMap := map[string]int{}
	songsTitleMap := map[string]int{}
	singersNameMap := map[string]int{}
	albumsTitleMap := map[string]int{}

	// open the rawdata file
	file, err := os.Open(RawDataPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	// read data from txt file and put it into database
	br := bufio.NewReader(file)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}

		// the data format : season;song;duration;singer;album
		attr := strings.Split(string(line), ";")
		nameList := strings.Split(attr[3], "/")

		var season *entity.Season
		var song *entity.Song
		var singerList = make([]*entity.Singer, len(nameList))
		var album *entity.Album

		// Season instance
		if ind, ok := seasonsTitleMap[attr[0]]; ok {
			season = &seasons[ind]
		} else {
			seasons = append(seasons, entity.NewSeason(len(seasons), attr[0], []int{}, []int{}, []int{}))
			season = &seasons[len(seasons)-1]
			seasonsTitleMap[attr[0]] = len(seasons) - 1
		}
		// Song instance
		if ind, ok := songsTitleMap[attr[1]]; ok {
			song = &songs[ind]
		} else {
			songs = append(songs, entity.NewSong(len(songs), attr[1], attr[2], []int{}, []int{}, []int{}))
			song = &songs[len(songs)-1]
			songsTitleMap[attr[1]] = len(songs) - 1
		}
		// Singer instances
		for i, name := range nameList {
			if ind, ok := singersNameMap[name]; ok {
				singerList[i] = &singers[ind]
			} else {
				singers = append(singers, entity.NewSinger(len(singers), name, []int{}, []int{}, []int{}))
				singerList[i] = &singers[len(singers)-1]
				singersNameMap[name] = len(singers) - 1
			}
		}
		// Album instance
		if ind, ok := albumsTitleMap[attr[4]]; ok {
			album = &albums[ind]
		} else {
			albums = append(albums, entity.NewAlbum(len(albums), attr[4], []int{}, []int{}, []int{}))
			album = &albums[len(albums)-1]
			albumsTitleMap[attr[4]] = len(albums) - 1
		}

		// season related entity id
		season.AddSongID(song.ID)
		season.AddAlbumID(album.ID)

		// song related entity id
		song.AddSeasonID(season.ID)
		song.AddAlbumID(album.ID)

		// album related entity id
		album.AddSongID(song.ID)
		album.AddSeasonID(season.ID)

		// singer related entity id
		for i := range nameList {
			singerList[i].AddSongID(song.ID)
			singerList[i].AddSeasonID(season.ID)
			singerList[i].AddAlbumID(album.ID)
			season.AddSingerID(singerList[i].ID)
			song.AddSingerID(singerList[i].ID)
			album.AddSingerID(singerList[i].ID)
		}
	}

	// Save all entity to Bucket
	entity.WriteSeasonToBucket(seasons, seasonBk)
	entity.WriteSongToBucket(songs, songBk)
	entity.WriteSingerToBucket(singers, singerBk)
	entity.WriteAlbumToBucket(albums, albumBk)

	return nil
}
