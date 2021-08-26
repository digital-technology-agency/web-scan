package services

import (
	"testing"
)

func TestGen(t *testing.T) {
	type args struct {
		alphabet string
		len      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Name generator",
			args: args{
				alphabet: "abc",
				len:      1,
			},
			want: 3,
		},
		{
			name: "Name generator",
			args: args{
				alphabet: "abc",
				len:      2,
			},
			want: 12,
		},
		{
			name: "Name generator",
			args: args{
				alphabet: "abc",
				len:      0,
			},
			want: 0,
		},
		{
			name: "Name generator",
			args: args{
				alphabet: "",
				len:      1,
			},
			want: 0,
		},
		{
			name: "Name generator",
			args: args{
				alphabet: "",
				len:      100,
			},
			want: 0,
		},
		{
			name: "Name generator",
			args: args{
				alphabet: "abcde",
				len:      2,
			},
			want: 30,
		},
		{
			name: "Name generator",
			args: args{
				alphabet: "abcdef",
				len:      2,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			incremet := 0
			for range Gen(tt.args.alphabet, tt.args.len) {
				incremet += 1
			}
			if tt.want != incremet {
				t.Errorf("Gen() = %v, want %v", incremet, tt.want)
			}
		})
	}
}
