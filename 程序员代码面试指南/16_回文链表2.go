package main

import (
	"fmt"
	"stack"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(arr []int) *ListNode {
	var temp *ListNode
	var head *ListNode

	head = &ListNode{Val: arr[0]}
	temp = head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return head
}
func main() {
	arr := []int{1, 2, 3, 3, 2, 2}
	head := makeListNode(arr)
	res := isPalindrome2(head)
	fmt.Println(res)
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var right *ListNode = head.Next
	cur := head
	for cur.Next != nil && cur.Next.Next != nil {
		right = right.Next
		cur = cur.Next.Next
	}
	var newStack stack.Stack
	for right != nil {
		newStack.Push(right.Val)
		right = right.Next
	}

	for newStack.Len() != 0 {
		if head.Val != newStack.Pop() {
			return false
		}
		head = head.Next
	}
	return true
}
