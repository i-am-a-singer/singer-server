package entity

import (
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewSong(t *testing.T) {
	type args struct {
		id        int
		ti        string
		du        string
		seasonsID []int
		singersID []int
		albumsID  []int
	}
	tests := []struct {
		name string
		args args
		want Song
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSong(tt.args.id, tt.args.ti, tt.args.du, tt.args.seasonsID, tt.args.singersID, tt.args.albumsID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_Type(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Song.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_Encode(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			got, err := s.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Song.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_Decode(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if err := s.Decode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Song.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSong_HasSinger(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasSinger(tt.args.id); got != tt.want {
				t.Errorf("Song.HasSinger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_HasAlbum(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasAlbum(tt.args.id); got != tt.want {
				t.Errorf("Song.HasAlbum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_HasSeason(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasSeason(tt.args.id); got != tt.want {
				t.Errorf("Song.HasSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_AddSeasonID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddSeasonID(tt.args.id)
		})
	}
}

func TestSong_AddAlbumID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddAlbumID(tt.args.id)
		})
	}
}

func TestSong_AddSingerID(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddSingerID(tt.args.id)
		})
	}
}

func TestSong_APIFormatConstruct(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Duration  string
		SeasonsID []int
		SingersID []int
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
			s := &Song{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Duration:  tt.fields.Duration,
				SeasonsID: tt.fields.SeasonsID,
				SingersID: tt.fields.SingersID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.APIFormatConstruct(tt.args.HostPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.APIFormatConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteSongToBucket(t *testing.T) {
	type args struct {
		entityList []Song
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
			WriteSongToBucket(tt.args.entityList, tt.args.bucket)
		})
	}
}
