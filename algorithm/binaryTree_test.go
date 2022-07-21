package algorithm

import (
	"reflect"
	"testing"
)

/*
          1
	     / \
	    3   4
	   / \
      5   6
	 / \
	7   8
         \
         10
*/
var root = TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:  8,
				Left: nil,
				Right: &TreeNode{
					Val:   10,
					Left:  nil,
					Right: nil,
				},
			},
		},
	},
	Right: &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	},
}

func Test_preOrderTraverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "前序遍历", args: args{root: &root}, want: []int{1, 3, 5, 7, 6, 8, 10, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inOrderTraverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "中序遍历", args: args{root: &root}, want: []int{7, 5, 3, 6, 8, 10, 1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postOrderTraverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "后序遍历", args: args{root: &root}, want: []int{7, 5, 10, 8, 6, 3, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
