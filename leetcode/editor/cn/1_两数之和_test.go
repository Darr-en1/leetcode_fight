package leet

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "两数之和",
			args: args{
				nums:   []int{2, 7, 11, 13},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "两数之和",
			args: args{
				nums:   []int{2, 7, 11, 13},
				target: 20,
			},
			want: []int{1, 3},
		},
		{
			name: "两数之和",
			args: args{
				nums:   []int{2, 7, 11, 13},
				target: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
