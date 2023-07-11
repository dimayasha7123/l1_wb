package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
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
			name: "а",
			s:    "а",
			want: "а",
		},
		{
			name: "аа",
			s:    "аа",
			want: "аа",
		},
		{
			name: "аб",
			s:    "аб",
			want: "ба",
		},
		{
			name: "главрыба",
			s:    "главрыба",
			want: "абырвалг",
		},
		{
			name: "unicode trash",
			s:    "😼AÅÄÖ\u20A9",
			want: "₩ÖÄÅA😼",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.s)

			assert.Equalf(t, tt.want, got, "got %q, want %q for %q", got, tt.want, tt.s)
		})
	}
}
