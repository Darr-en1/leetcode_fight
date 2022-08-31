package leet

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	ringLinkedList := ToLinkedList([]int{1, 2, 3})
	ringLinkedList.Next.Next = ringLinkedList

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"环形链表(不存在)",
			args{head: ToLinkedList([]int{1, 2, 3, 4, 5})},
			false,
		},
		{
			"环形链表(存在)",
			args{head: ringLinkedList},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
