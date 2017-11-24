package main

import (
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_IsPalindrome",
			args: args{s: "Test"},
			want: "Test - is not palindrome.",
		},
		{
			name: "Test_IsPalindrome",
			args: args{s: "Never odd or even"},
			want: "Never odd or even - is palindrome.",
		},
		{
			name: "Test_IsPalindrome",
			args: args{s: "Инна"},
			want: "Инна - is not palindrome.",
		},
		{
			name: "Test_IsPalindrome",
			args: args{s: "А роза упала на лапу Азора"},
			want: "А роза упала на лапу Азора - is palindrome.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsPalindrome(tt.args.s)
			if res != tt.want {
				t.Errorf("got = %v | want %v", res, tt.want)
			}
		})
	}
}
