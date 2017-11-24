package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_CombineBraces(t *testing.T) {
	type args struct {
		s      string
		opened int
		closed int
		n      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_CombineBraces_1",
			args: args{s: "", opened: 0, closed: 0, n: 1},
			want: "()\n",
		},
		{
			name: "Test_CombineBraces_2",
			args: args{s: "", opened: 0, closed: 0, n: 2},
			want: "(())\n()()\n",
		},
		{
			name: "Test_CombineBraces_3",
			args: args{s: "", opened: 0, closed: 0, n: 3},
			want: "((()))\n(()())\n(())()\n()(())\n()()()\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CombineBraces(tt.args.s, tt.args.opened, tt.args.closed, tt.args.n)
			res, _ := ioutil.ReadAll(os.Stdout)
			if string(res) != tt.want {
				t.Errorf("got = %v | want %v", res, tt.want)
			}
		})
	}
}

func Test_ParseBraces(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_ParseBraces",
			args: args{s: "()"},
			want: "() => true",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: "(())"},
			want: "(()) => true",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: "()()"},
			want: "()() => true",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: "())"},
			want: "()) => false",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: "(()"},
			want: "(() => false",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: "("},
			want: "( => false",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: ")"},
			want: ") => false",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: ""},
			want: "=> false",
		},
		{
			name: "Test_ParseBraces",
			args: args{s: "fail"},
			want: "unsupported char: f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ParseBraces(tt.args.s)
			if res != tt.want {
				t.Errorf("got = %v | want %v", res, tt.want)
			}
		})
	}
}
