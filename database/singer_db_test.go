package database

import (
	"testing"

	"github.com/boltdb/bolt"
	"github.com/singerapi/singer-server/entity"
)

func TestSingerDB_OpenDB(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		loader func(*bolt.Tx) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.OpenDB(tt.args.loader); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.OpenDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_Close(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.Close(); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOpenDB(t *testing.T) {
	type args struct {
		loader func(*bolt.Tx) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OpenDB(tt.args.loader); (err != nil) != tt.wantErr {
				t.Errorf("OpenDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClose(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_Query(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		id int
		e  Entity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.Query(tt.args.id, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.Query() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuery(t *testing.T) {
	type args struct {
		id int
		e  Entity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Query(tt.args.id, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QuerySeasonList(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		e      *[]entity.Season
		filter func(s *entity.Season) bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QuerySeasonList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QuerySeasonList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QuerySongList(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		e      *[]entity.Song
		filter func(s *entity.Song) bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QuerySongList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QuerySongList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QuerySingerList(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		e      *[]entity.Singer
		filter func(s *entity.Singer) bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QuerySingerList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QuerySingerList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QueryAlbumList(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		e      *[]entity.Album
		filter func(s *entity.Album) bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QueryAlbumList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QueryAlbumList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QuerySeason(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		id     int
		season *entity.Season
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QuerySeason(tt.args.id, tt.args.season); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QuerySeason() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QuerySong(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		id   int
		song *entity.Song
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QuerySong(tt.args.id, tt.args.song); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QuerySong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QuerySinger(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		id     int
		singer *entity.Singer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QuerySinger(tt.args.id, tt.args.singer); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QuerySinger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSingerDB_QueryAlbum(t *testing.T) {
	type fields struct {
		db *bolt.DB
	}
	type args struct {
		id    int
		album *entity.Album
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &SingerDB{
				db: tt.fields.db,
			}
			if err := sdb.QueryAlbum(tt.args.id, tt.args.album); (err != nil) != tt.wantErr {
				t.Errorf("SingerDB.QueryAlbum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuerySeasonList(t *testing.T) {
	type args struct {
		e      *[]entity.Season
		filter func(s *entity.Season) bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuerySeasonList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("QuerySeasonList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuerySongList(t *testing.T) {
	type args struct {
		e      *[]entity.Song
		filter func(s *entity.Song) bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuerySongList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("QuerySongList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuerySingerList(t *testing.T) {
	type args struct {
		e      *[]entity.Singer
		filter func(s *entity.Singer) bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuerySingerList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("QuerySingerList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryAlbumList(t *testing.T) {
	type args struct {
		e      *[]entity.Album
		filter func(s *entity.Album) bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QueryAlbumList(tt.args.e, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("QueryAlbumList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuerySeason(t *testing.T) {
	type args struct {
		id     int
		season *entity.Season
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuerySeason(tt.args.id, tt.args.season); (err != nil) != tt.wantErr {
				t.Errorf("QuerySeason() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuerySong(t *testing.T) {
	type args struct {
		id   int
		song *entity.Song
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuerySong(tt.args.id, tt.args.song); (err != nil) != tt.wantErr {
				t.Errorf("QuerySong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuerySinger(t *testing.T) {
	type args struct {
		id     int
		singer *entity.Singer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuerySinger(tt.args.id, tt.args.singer); (err != nil) != tt.wantErr {
				t.Errorf("QuerySinger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryAlbum(t *testing.T) {
	type args struct {
		id    int
		album *entity.Album
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QueryAlbum(tt.args.id, tt.args.album); (err != nil) != tt.wantErr {
				t.Errorf("QueryAlbum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}