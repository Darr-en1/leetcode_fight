package leet

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "两数相加",
			args: args{
				l1: &ListNode{ // 3->6->5
					Val: 3,
					Next: &ListNode{
						Val:  6,
						Next: &ListNode{Val: 5},
					}},
				l2: &ListNode{ // 7->3->4
					Val: 7,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4},
					}},
			},
			want: &ListNode{ // 0->0->0->1
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
