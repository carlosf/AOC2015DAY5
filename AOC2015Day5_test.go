package main

import "testing"

func Test_isLineNice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"vowel1", args{s: "aei"}, true},
		{"noVowels", args{s: "bcdfg"}, false},
		{"repeatedChars", args{s: "aabbcc"}, true},
		{"lessThanThreeVowels", args{s: "ae"}, false},
		{"bothCriteria", args{s: "baae"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLineNice(tt.args.s); got != tt.want {
				t.Errorf("isLineNice() = %v, want %v", got, tt.want)
			}
		})
	}
}
