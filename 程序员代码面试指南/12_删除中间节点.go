package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeByRatio(head *ListNode, a, b int) *ListNode {
	if a < 1 || a > b {
		return head
	}
	var n int = 0
	var cur *ListNode
	for cur != nil {
		n++
		cur = cur.Next
	}

	var c int
	c = a * n / b
	if c == 1 {
		head = head.Next
	}
	if n > 1 {
		cur = head
		n--
		for n != 1 {
			n--
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}

	return head
}

//关于删除链表节点的问题都是找到链表的前一个节点
