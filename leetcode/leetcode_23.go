package main

import (
	"fmt"
)

//合并K个链表,这里采用的分治法

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return startMergeLists(lists, 0, len(lists)-1)
}

func mergeNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		if head1 == nil {
			return head2
		} else {
			return head1
		}
	}
	var head *ListNode
	var cur1 *ListNode
	var cur2 *ListNode

	if head1.Val < head2.Val {
		head = head1
	} else {
		head = head2
	}
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
	if len(arr) == 0 {
		return nil
	}
	var head *ListNode = &ListNode{Val: arr[0]}
	var temp *ListNode = head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return head
}

func main() {
	arr1 := []int{}
	arr2 := []int{1, 3, 4}
	head1 := makeListNode(arr1)
	head2 := makeListNode(arr2)
	var lists []*ListNode
	lists = append(lists, head1)
	lists = append(lists, head2)

	res := mergeNode(head1, head2)
	for res != nil {
		fmt.Println(res)

		res = res.Next
	}
}

func startMergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	var mid int
	mid = (left + right) / 2
	l1 := startMergeLists(lists, left, mid)
	l2 := startMergeLists(lists, mid+1, right)
	return mergeNode(l1, l2)
}
