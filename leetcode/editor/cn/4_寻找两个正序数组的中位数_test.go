package leet

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "寻找两个正序数组的中位数",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{name: "寻找两个正序数组的中位数",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0.0,
		},
		{name: "寻找两个正序数组的中位数",
			args: args{
				nums1: []int{2},
				nums2: []int{0},
			},
			want: 1.0,
		},
		{name: "寻找两个正序数组的中位数",
			args: args{
				nums1: []int{2},
				nums2: []int{},
			},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
