package rawdata

import (
	"testing"

	"github.com/boltdb/bolt"
)

func TestLoadRawData(t *testing.T) {
	type args struct {
		tx *bolt.Tx
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
			if err := LoadRawData(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("LoadRawData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
