package service

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/singerapi/singer-server/entity"
)

func Test_parseRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantID  int
		wantOpt PageOptions
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, gotOpt := parseRequest(tt.args.r)
			if gotID != tt.wantID {
				t.Errorf("parseRequest() gotID = %v, want %v", gotID, tt.wantID)
			}
			if !reflect.DeepEqual(gotOpt, tt.wantOpt) {
				t.Errorf("parseRequest() gotOpt = %v, want %v", gotOpt, tt.wantOpt)
			}
		})
	}
}

func Test_handleSeasons(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		filter func(s *entity.Season) bool
		opt    *PageOptions
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleSeasons(tt.args.w, tt.args.r, tt.args.filter, tt.args.opt)
		})
	}
}

func Test_handleSongs(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		filter func(s *entity.Song) bool
		opt    *PageOptions
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleSongs(tt.args.w, tt.args.r, tt.args.filter, tt.args.opt)
		})
	}
}

func Test_handleSingers(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		filter func(s *entity.Singer) bool
		opt    *PageOptions
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleSingers(tt.args.w, tt.args.r, tt.args.filter, tt.args.opt)
		})
	}
}

func Test_handleAlbums(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		filter func(s *entity.Album) bool
		opt    *PageOptions
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleAlbums(tt.args.w, tt.args.r, tt.args.filter, tt.args.opt)
		})
	}
}

func Test_notFindErrorHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notFindErrorHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_internalErrorHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internalErrorHandler(tt.args.w, tt.args.r)
		})
	}
}
