package generator

import "testing"

func TestToCamel(t *testing.T) {
	tests := []struct {
		orig string
		want string
	}{
		{
			orig: "small_char",
			want: "SmallChar",
		},
		{
			orig: "ALLBIG",
			want: "Allbig",
		},
		{
			orig: "ALL_BIG_WITH_UNDERbar",
			want: "AllBigWithUnderbar",
		},
	}
	for _, tt := range tests {
		got := toCamel(tt.orig)
		if got != tt.want {
			t.Errorf("exected %s but got %s", tt.want, got)
		}
	}
}
