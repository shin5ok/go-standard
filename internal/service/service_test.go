package service

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want params
	}{
		{
			name: "たこ",
			args: args{name: "tako"},
			want: params{
				Name: "tako",
			},
		},
		{
			name: "いか",
			args: args{name: "ika"},
			want: params{
				Name: "ika",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
