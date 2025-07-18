package hashmap

import "testing"

func TestCanConstruct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
		{
			ransomNote: "aa",
			magazine:   "aa",
			want:       true,
		},
		{
			ransomNote: "alabama",
			magazine:   "balabol",
			want:       false,
		},
		{
			ransomNote: "letter",
			magazine:   "rettel",
			want:       true,
		},
	}

	for _, test := range tests {
		if got := canConstruct(test.ransomNote, test.magazine); got != test.want {
			t.Errorf("canConstruct(%v, %v) = %v, want %v", test.ransomNote, test.magazine, got, test.want)
		}
	}
}
