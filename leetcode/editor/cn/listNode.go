package leet

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return head
}
