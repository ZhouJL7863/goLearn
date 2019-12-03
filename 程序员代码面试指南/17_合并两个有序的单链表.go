package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		if head1 == nil {
			return head2
		}
		return head1
	}
	var head *ListNode
	if head1.Val < head2.Val {
		head = head1
	} else {
		head = head2
	}

	var cur1 *ListNode
	var cur2 *ListNode
	if head == head1 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}
	var pre *ListNode = head
	var next *ListNode
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			pre = cur1
			cur1 = cur1.Next
		} else {
			next = cur2.Next
			pre.Next = cur2
			cur2.Next = cur1
			pre = cur2
			cur2 = next
		}

	}
	if cur1 == nil {
		pre.Next = cur2
	} else {
		pre.Next = cur1
	}
	return head
}

func makeListNode(arr []int) *ListNode {
	var head *ListNode = &ListNode{Val: arr[0]}
	temp := head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return head
}

func main() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	head1 := makeListNode(arr1)
	head2 := makeListNode(arr2)

	res := merge(head1, head2)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
