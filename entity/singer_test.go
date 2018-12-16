package entity

import (
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewSinger(t *testing.T) {
	type args struct {
		id        int
		na        string
		seasonsID []int
		songsID   []int
		albumsID  []int
	}
	tests := []struct {
		name string
		args args
		want Singer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSinger(tt.args.id, tt.args.na, tt.args.seasonsID, tt.args.songsID, tt.args.albumsID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSinger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinger_Type(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Singer.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinger_Encode(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			got, err := s.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Singer.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Singer.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinger_Decode(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if err := s.Decode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Singer.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSinger_HasSong(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasSong(tt.args.id); got != tt.want {
				t.Errorf("Singer.HasSong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinger_HasAlbum(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasAlbum(tt.args.id); got != tt.want {
				t.Errorf("Singer.HasAlbum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinger_HasSeason(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.HasSeason(tt.args.id); got != tt.want {
				t.Errorf("Singer.HasSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinger_AddSongID(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddSongID(tt.args.id)
		})
	}
}

func TestSinger_AddAlbumID(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddAlbumID(tt.args.id)
		})
	}
}

func TestSinger_AddSeasonID(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			s.AddSeasonID(tt.args.id)
		})
	}
}

func TestSinger_APIFormatConstruct(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		SeasonsID []int
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
			s := &Singer{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				SeasonsID: tt.fields.SeasonsID,
				SongsID:   tt.fields.SongsID,
				AlbumsID:  tt.fields.AlbumsID,
			}
			if got := s.APIFormatConstruct(tt.args.HostPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Singer.APIFormatConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteSingerToBucket(t *testing.T) {
	type args struct {
		entityList []Singer
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
			WriteSingerToBucket(tt.args.entityList, tt.args.bucket)
		})
	}
}
