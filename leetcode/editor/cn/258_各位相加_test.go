package leet

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "各位相加", args: args{num: 38}, want: 2},
		{name: "各位相加", args: args{num: 465}, want: 6},
		{name: "各位相加", args: args{num: 234}, want: 9},
		{name: "各位相加", args: args{num: 235}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
