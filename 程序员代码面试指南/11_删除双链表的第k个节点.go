//双向链表也叫双链表，它的每个数据节点都有两个指针，分别指向直接后继和直接前驱
package main

type DoubleNode struct {
	Val  int
	Last *DoubleNode
	Next *DoubleNode
}

func removeLastKthNode(head DoubleNode, n int) {
	if head == nil || n < 1 {
		return head
	}

	cur := head
	for cur != nil {
		n--
		cur = cur.Next
	}

	if n == 0 {
		head = head.Next
		head.Last = nil
	}
	if n < 0 {
		cur = head
		n++
		for n != 0 {
			n++
			cur = cur.Next
		}
		var newNext DoubleNode
		newNext = cur.Next.Next
		if newNext != nil {
			newNext.Last = cur
		}
	}
	return head
}
