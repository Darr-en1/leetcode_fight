package leet

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "链表的中间结点(基数)",
			args: args{head: ToLinkedList([]int{1, 2, 3, 4, 5})},
			want: ToLinkedList([]int{3, 4, 5}),
		},
		{
			name: "链表的中间结点(偶数)",
			args: args{head: ToLinkedList([]int{1, 2, 3, 4, 5, 6})},
			want: ToLinkedList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
