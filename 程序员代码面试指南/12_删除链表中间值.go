package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeMidNode(head *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	//当链表长度为2时，直接删除第一个
	if head.Next.Next == nil {
		return head.Next
	}
	var pre *ListNode = head
	var cur *ListNode = head.Next.Next

	for cur.Next != nil && cur.Next.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
	}
	pre.Next = pre.Next.Next
	return head
}
