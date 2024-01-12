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
		{"3 vowels double letter", args{s: "ugknbfddgicrmopn"}, true},
		{"3 vowels and double letter", args{s: "aaa"}, true},
		{"no double letter", args{s: "jchzalrnumimnmhp"}, false},
		{"string xy", args{s: "haegwjzuvuyypxyu"}, false},
		{"only one vowel", args{s: "dvszwmarrgswjxmb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLineNice(tt.args.s); got != tt.want {
				t.Errorf("isLineNice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLineNice2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"qjappearwqj", args{s: "qjhvhtzxzqqjkmpb"}, true},
		{"xxyxx", args{s: "xxyxx"}, true},
		{"uurcxstgmygtbstg", args{s: "uurcxstgmygtbstg"}, false},
		{"ieodomkazucvgmuy", args{s: "ieodomkazucvgmuy"}, false},
		{"abaaaa", args{s: "abaaaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLineNice2(tt.args.s); got != tt.want {
				t.Errorf("isLineNice2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strRemoveAt(t *testing.T) {
	type args struct {
		s      string
		index  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"remove 2 chars", args{s: "abcdefg", index: 2, length: 2}, "abefg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strRemoveAt(tt.args.s, tt.args.index, tt.args.length); got != tt.want {
				t.Errorf("strRemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
