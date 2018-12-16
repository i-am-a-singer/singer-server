package render

import (
	"bytes"
	"testing"
)

func TestResourceItemJSONRender(t *testing.T) {
	type args struct {
		status     int
		r          Resource
		HostPrefix string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := ResourceItemJSONRender(w, tt.args.status, tt.args.r, tt.args.HostPrefix); (err != nil) != tt.wantErr {
				t.Errorf("ResourceItemJSONRender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ResourceItemJSONRender() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestResourceListJSONRender(t *testing.T) {
	type args struct {
		status       int
		resourceList []Resource
		HostPrefix   string
		APIURL       string
		page         int
		limit        int
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := ResourceListJSONRender(w, tt.args.status, tt.args.resourceList, tt.args.HostPrefix, tt.args.APIURL, tt.args.page, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("ResourceListJSONRender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ResourceListJSONRender() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestMapJSONRender(t *testing.T) {
	type args struct {
		status   int
		respBody map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := MapJSONRender(w, tt.args.status, tt.args.respBody); (err != nil) != tt.wantErr {
				t.Errorf("MapJSONRender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("MapJSONRender() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
