package database

import (
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/singerapi/singer-server/entity"
)

// SingerDB implements DB interface
type SingerDB struct {
	db *bolt.DB
}

var (
	distPath = "data.db"
	singerDB = SingerDB{nil}
)

// OpenDB is interface from DB
func (sdb *SingerDB) OpenDB(loader func(*bolt.Tx) error) error {
	// create the bolt database
	os.Remove(distPath)
	dbptr, err := bolt.Open(distPath, 0600, nil)
	if err != nil {
		return err
	}
	sdb.db = dbptr
	sdb.db.Update(loader)
	return nil
}

// Close release all bolt.db resource
func (sdb *SingerDB) Close() error {
	return sdb.db.Close()
}

// CURD
// For our service require we only implement Request

// Query Entity
func (sdb *SingerDB) Query(id int, e Entity) {
	sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(e.Type()))
		data := bk.Get([]byte(strconv.Itoa(id)))
		if err := e.Decode(data); err != nil {
			return err
		}
		return nil
	})
}

// QuerySeasonList query all entity of season
func (sdb *SingerDB) QuerySeasonList(e []entity.Season, filter func(s *entity.Season) bool) {
	sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.SeasonType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Season{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				e = append(e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QuerySongList query all entity of Song
func (sdb *SingerDB) QuerySongList(e []entity.Song, filter func(s *entity.Song) bool) {
	sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.SongType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Song{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				e = append(e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QuerySingerList query all entity of Singer
func (sdb *SingerDB) QuerySingerList(e []entity.Singer, filter func(s *entity.Singer) bool) {
	sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.SingerType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Singer{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				e = append(e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QueryAlbumList query all entity of Album
func (sdb *SingerDB) QueryAlbumList(e []entity.Album, filter func(s *entity.Album) bool) {
	sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.AlbumType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Album{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				e = append(e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QuerySeason query season by id
func (sdb *SingerDB) QuerySeason(id int, season *entity.Season) {
	sdb.Query(id, season)
}

// QuerySong query song by id
func (sdb *SingerDB) QuerySong(id int, song *entity.Song) {
	sdb.Query(id, song)
}

// QuerySinger query season by id
func (sdb *SingerDB) QuerySinger(id int, singer *entity.Singer) {
	sdb.Query(id, singer)
}

// QueryAlbum query season by id
func (sdb *SingerDB) QueryAlbum(id int, album *entity.Album) {
	sdb.Query(id, album)
}

// QuerySeasonList query all entity of Season
func QuerySeasonList(e []entity.Season, filter func(s *entity.Season) bool) {
	singerDB.QuerySeasonList(e, filter)
}

// QuerySongList query all entity of Song
func QuerySongList(e []entity.Song, filter func(s *entity.Song) bool) {
	singerDB.QuerySongList(e, filter)
}

// QuerySingerList query all entity of Singer
func QuerySingerList(e []entity.Singer, filter func(s *entity.Singer) bool) {
	singerDB.QuerySingerList(e, filter)
}

// QueryAlbumList query all entity of Album
func QueryAlbumList(e []entity.Album, filter func(s *entity.Album) bool) {
	singerDB.QueryAlbumList(e, filter)
}

// QuerySeason query season by id
func QuerySeason(id int, season *entity.Season) {
	singerDB.QuerySeason(id, season)
}

// QuerySong query song by id
func QuerySong(id int, song *entity.Song) {
	singerDB.QuerySong(id, song)
}

// QuerySinger query season by id
func QuerySinger(id int, singer *entity.Singer) {
	singerDB.QuerySinger(id, singer)
}

// QueryAlbum query season by id
func QueryAlbum(id int, album *entity.Album) {
	singerDB.QueryAlbum(id, album)
}
