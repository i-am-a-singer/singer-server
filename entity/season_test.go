package entity

import (
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewSeason(t *testing.T) {
	type args struct {
		id        int
		ti        string
		singersID []int
		songsID   []int
		albumsID  []int
	}
	tests := []struct {
		name string
		args args
		want Season
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeason(tt.args.id, tt.args.ti, tt.args.singersID, tt.args.songsID, tt.args.albumsID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeason_Type(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Season.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeason_Encode(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			got, err := s.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Season.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Season.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeason_Decode(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		b []byte
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
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if err := s.Decode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Season.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSeason_HasSinger(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasSinger(tt.args.id); got != tt.want {
				t.Errorf("Season.HasSinger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeason_HasSong(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasSong(tt.args.id); got != tt.want {
				t.Errorf("Season.HasSong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeason_HasAlbum(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasAlbum(tt.args.id); got != tt.want {
				t.Errorf("Season.HasAlbum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeason_AddSingerID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddSingerID(tt.args.id)
		})
	}
}

func TestSeason_AddSongID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddSongID(tt.args.id)
		})
	}
}

func TestSeason_AddAlbumID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddAlbumID(tt.args.id)
		})
	}
}

func TestSeason_APIFormatConstruct(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SingersID []int
		SongsID   []int
		AlbumsID  []int
	}
	type args struct {
		HostPrefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Season{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SingersID: tt.fields.SingersID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.APIFormatConstruct(tt.args.HostPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Season.APIFormatConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteSeasonToBucket(t *testing.T) {
	type args struct {
		entityList []Season
		bucket     *bolt.Bucket
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteSeasonToBucket(tt.args.entityList, tt.args.bucket)
		})
	}
}
