package entity

import (
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewAlbum(t *testing.T) {
	type args struct {
		id        int
		ti        string
		seasonsID []int
		songsID   []int
		singersID []int
	}
	tests := []struct {
		name string
		args args
		want Album
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlbum(tt.args.id, tt.args.ti, tt.args.seasonsID, tt.args.songsID, tt.args.singersID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlbum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_Type(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Album.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_Encode(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			got, err := s.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Album.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Album.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_Decode(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			if err := s.Decode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Album.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlbum_HasSinger(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			if got := s.HasSinger(tt.args.id); got != tt.want {
				t.Errorf("Album.HasSinger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_HasSong(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			if got := s.HasSong(tt.args.id); got != tt.want {
				t.Errorf("Album.HasSong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_HasSeason(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			if got := s.HasSeason(tt.args.id); got != tt.want {
				t.Errorf("Album.HasSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_AddSongID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			s.AddSongID(tt.args.id)
		})
	}
}

func TestAlbum_AddSingerID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			s.AddSingerID(tt.args.id)
		})
	}
}

func TestAlbum_AddSeasonID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			s.AddSeasonID(tt.args.id)
		})
	}
}

func TestAlbum_APIFormatConstruct(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		SeasonsID []int
		SongsID   []int
		SingersID []int
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
			s := &Album{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				SingersID: tt.fields.SingersID,
			}
			if got := s.APIFormatConstruct(tt.args.HostPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Album.APIFormatConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteAlbumToBucket(t *testing.T) {
	type args struct {
		entityList []Album
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
			WriteAlbumToBucket(tt.args.entityList, tt.args.bucket)
		})
	}
}
