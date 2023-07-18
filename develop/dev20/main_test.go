package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWordsInString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty",
			s:    "",
			want: "",
		},
		{
			name: "a",
			s:    "a",
			want: "a",
		},
		{
			name: "a a",
			s:    "a a",
			want: "a a",
		},
		{
			name: "a b",
			s:    "a b",
			want: "b a",
		},
		{
			name: "abcd efgh",
			s:    "abcd efgh",
			want: "efgh abcd",
		},
		{
			name: "with unicode",
			s:    "₩ÖÄÅA😼 abcd",
			want: "abcd ₩ÖÄÅA😼",
		},
		{
			name: "snow dog sun",
			s:    "snow dog sun",
			want: "sun dog snow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWordsInString(tt.s)

			assert.Equalf(t, tt.want, got, "got %q, want %q for %q", got, tt.want, tt.s)
		})
	}
}
