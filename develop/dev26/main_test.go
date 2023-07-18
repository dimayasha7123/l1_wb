package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniqSymbolsCaseInsensitive(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty",
			s:    "",
			want: true,
		},
		{
			name: "abcd",
			s:    "abcd",
			want: true,
		},
		{
			name: "abCdefAaf",
			s:    "abCdefAaf",
			want: false,
		},
		{
			name: "aabcd",
			s:    "aabcd",
			want: false,
		},
		{
			name: "aA",
			s:    "aA",
			want: false,
		},
		{
			name: "aa",
			s:    "aa",
			want: false,
		},
		{
			name: "a",
			s:    "a",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqSymbolsCaseInsensitive(tt.s)

			assert.Equalf(t, tt.want, got, "got %v, want %v for %q", got, tt.want, tt.s)
		})
	}
}
