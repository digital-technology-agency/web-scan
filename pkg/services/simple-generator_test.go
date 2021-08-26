package services

import (
	"testing"
)

func TestGen(t *testing.T) {
	type fields struct {
		Alphabet string
		Len      int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "abc",
				Len:      1,
			},
			want: 3,
		},
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "abc",
				Len:      2,
			},
			want: 12,
		},
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "abc",
				Len:      0,
			},
			want: 0,
		},
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "",
				Len:      1,
			},
			want: 0,
		},
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "",
				Len:      100,
			},
			want: 0,
		},
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "abcde",
				Len:      2,
			},
			want: 30,
		},
		{
			name: "Name generator",
			fields: fields{
				Alphabet: "abcdef",
				Len:      2,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &SimpleGenerator{
				Alphabet: tt.fields.Alphabet,
				Len:      tt.fields.Len,
			}
			incremet := 0
			for range gen.Gen() {
				incremet += 1
			}
			if tt.want != incremet {
				t.Errorf("Gen() = %v, want %v", incremet, tt.want)
			}
		})
	}
}
