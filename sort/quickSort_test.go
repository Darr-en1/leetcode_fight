package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "快排",
			args: args{arr: []int{7, 2, 4, 3, 8, 9, 5}},
			want: []int{2,3,4,5,7,8,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
