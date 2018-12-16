package service

import (
	"net/http"
	"testing"
)

func Test_handleAlbumsID(t *testing.T) {
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
			handleAlbumsID(tt.args.w, tt.args.r)
		})
	}
}

func Test_handleAlbumsAll(t *testing.T) {
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
			handleAlbumsAll(tt.args.w, tt.args.r)
		})
	}
}

func Test_handleAlbumsSeasons(t *testing.T) {
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
			handleAlbumsSeasons(tt.args.w, tt.args.r)
		})
	}
}

func Test_handleAlbumsSongs(t *testing.T) {
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
			handleAlbumsSongs(tt.args.w, tt.args.r)
		})
	}
}

func Test_handleAlbumsSingers(t *testing.T) {
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
			handleAlbumsSingers(tt.args.w, tt.args.r)
		})
	}
}
