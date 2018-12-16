package entity

import "testing"

func Test_hasID(t *testing.T) {
	type args struct {
		idList []int
		id     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasID(tt.args.idList, tt.args.id); got != tt.want {
				t.Errorf("hasID() = %v, want %v", got, tt.want)
			}
		})
	}
}
