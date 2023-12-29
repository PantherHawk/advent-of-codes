package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FormatQuery(t *testing.T) {
	tests := map[string]struct {
		day      string
		year     string
		filepath string
		base     string
		want     string
	}{
		"happy path": {
			day:      "2",
			year:     "2023",
			filepath: "../",
			base:     "https://adventofcode.com",
			want:     "https://adventofcode.com/2023/day/2/input",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := formatQuery(test.day)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_FetchInput(t *testing.T) {
	tests := map[string]struct {
		day     string
		want    interface{}
		wantErr error
	}{
		"200 status": {
			day:     "2",
			want:    nil,
			wantErr: nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := fetchInput(test.day)
			assert.Equal(t, test.want, got)
			assert.ErrorIs(t, test.wantErr, err)

		})
	}
}
