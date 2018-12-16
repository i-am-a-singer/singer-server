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

// OpenDB is interface from DB
func OpenDB(loader func(*bolt.Tx) error) error {
	return singerDB.OpenDB(loader)
}

// Close release all bolt.db resource
func Close() error {
	return singerDB.Close()
}

// CURD
// For our service require we only implement Request

// Query Entity
func (sdb *SingerDB) Query(id int, e Entity) error {
	return sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(e.Type()))
		data := bk.Get([]byte(strconv.Itoa(id)))
		if err := e.Decode(data); err != nil {
			return err
		}
		return nil
	})
}

// Query Entity
func Query(id int, e Entity) error {
	return singerDB.Query(id, e)
}

// QuerySeasonList query all entity of season
func (sdb *SingerDB) QuerySeasonList(e *[]entity.Season, filter func(s *entity.Season) bool) error {
	return sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.SeasonType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Season{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				*e = append(*e, s)
			}
			return nil
		}); err != nil {
			return err
		}

		return nil
	})
}

// QuerySongList query all entity of Song
func (sdb *SingerDB) QuerySongList(e *[]entity.Song, filter func(s *entity.Song) bool) error {
	return sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.SongType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Song{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				*e = append(*e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QuerySingerList query all entity of Singer
func (sdb *SingerDB) QuerySingerList(e *[]entity.Singer, filter func(s *entity.Singer) bool) error {
	return sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.SingerType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Singer{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				*e = append(*e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QueryAlbumList query all entity of Album
func (sdb *SingerDB) QueryAlbumList(e *[]entity.Album, filter func(s *entity.Album) bool) error {
	return sdb.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(entity.AlbumType))
		if err := bk.ForEach(func(k, v []byte) error {
			s := entity.Album{}
			if err := s.Decode(v); err != nil {
				return err
			}
			if filter(&s) {
				*e = append(*e, s)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}

// QuerySeason query season by id
func (sdb *SingerDB) QuerySeason(id int, season *entity.Season) error {
	return sdb.Query(id, season)
}

// QuerySong query song by id
func (sdb *SingerDB) QuerySong(id int, song *entity.Song) error {
	return sdb.Query(id, song)
}

// QuerySinger query season by id
func (sdb *SingerDB) QuerySinger(id int, singer *entity.Singer) error {
	return sdb.Query(id, singer)
}

// QueryAlbum query season by id
func (sdb *SingerDB) QueryAlbum(id int, album *entity.Album) error {
	return sdb.Query(id, album)
}

// QuerySeasonList query all entity of Season
func QuerySeasonList(e *[]entity.Season, filter func(s *entity.Season) bool) error {
	return singerDB.QuerySeasonList(e, filter)
}

// QuerySongList query all entity of Song
func QuerySongList(e *[]entity.Song, filter func(s *entity.Song) bool) error {
	return singerDB.QuerySongList(e, filter)
}

// QuerySingerList query all entity of Singer
func QuerySingerList(e *[]entity.Singer, filter func(s *entity.Singer) bool) error {
	return singerDB.QuerySingerList(e, filter)
}

// QueryAlbumList query all entity of Album
func QueryAlbumList(e *[]entity.Album, filter func(s *entity.Album) bool) error {
	return singerDB.QueryAlbumList(e, filter)
}

// QuerySeason query season by id
func QuerySeason(id int, season *entity.Season) error {
	return singerDB.QuerySeason(id, season)
}

// QuerySong query song by id
func QuerySong(id int, song *entity.Song) error {
	return singerDB.QuerySong(id, song)
}

// QuerySinger query season by id
func QuerySinger(id int, singer *entity.Singer) error {
	return singerDB.QuerySinger(id, singer)
}

// QueryAlbum query season by id
func QueryAlbum(id int, album *entity.Album) error {
	return singerDB.QueryAlbum(id, album)
}
