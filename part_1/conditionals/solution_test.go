package conditionals

import (
	"reflect"
	"testing"
)

func Test_getMovies(t *testing.T) {
	type args struct {
		age int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "17+",
			args: args{
				age: 21,
			},
			want: movies17up,
		},
		{
			name: "13-17",
			args: args{
				age: 16,
			},
			want: movies13to17,
		},
		{
			name: "10-13",
			args: args{
				age: 12,
			},
			want: movies10to13,
		},
		{
			name: "5-10",
			args: args{
				age: 9,
			},
			want: movies5to10,
		},
		{
			name: "0-5",
			args: args{
				age: 2,
			},
			want: movies0to5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMovies(tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMovies() = %v, want %v", got, tt.want)
			}
		})
	}
}
