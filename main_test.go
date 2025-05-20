package main

import (
	"testing"
)

/*
3.141 59265
35897 93238
46264 33832
79502 88419
71693 99375
10582 09749
*/

func Test_search(t *testing.T) {
	type args struct {
		digits string
		cache  int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			args: args{
				digits: "6535",
				cache:  100,
			},
			want:    8,
			wantErr: false,
		},
		{
			args: args{
				digits: "6535",
				cache:  10,
			},
			want:    8,
			wantErr: false,
		},
		{
			args: args{
				digits: "8979",
				cache:  10,
			},
			want:    12,
			wantErr: false,
		},
		{
			args: args{
				digits: "3383",
				cache:  10,
			},
			want:    25,
			wantErr: false,
		},
		{
			args: args{
				digits: "71693",
				cache:  10,
			},
			want:    40,
			wantErr: false,
		},
		{
			args: args{
				digits: "7169399375",
				cache:  10,
			},
			want:    40,
			wantErr: false,
		},
		{
			args: args{
				digits: "7169399375105",
				cache:  10,
			},
			want:    40,
			wantErr: false,
		},
		{
			args: args{
				digits: "12345",
				cache:  10,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := search("pi_1k.txt", tt.args.digits, tt.args.cache)
			if (err != nil) != tt.wantErr {
				t.Errorf("search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
