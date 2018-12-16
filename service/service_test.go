package service

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/urfave/negroni"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want *negroni.Negroni
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleAPI(t *testing.T) {
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
			handleAPI(tt.args.w, tt.args.r)
		})
	}
}
