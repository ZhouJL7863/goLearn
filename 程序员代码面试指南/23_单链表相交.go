package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLoopNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	var n1 *ListNode = head.Next      // slow node
	var n2 *ListNode = head.Next.Next //fast node

	for n1 != n2 {
		if n2.Next == nil || n2.Next.Next == nil {
			return nil
		}
		n2 = n2.Next.Next
		n1 = n1.Next
	}
	n2 = head
	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}

func noLoop(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	var cur1 *ListNode = head1
	var cur2 *ListNode = head2

	var n int = 0
	for cur1 != nil {
		n++
		cur1 = cur1.Next
	}
	for cur2 != nil {
		n--
		cur2 = cur2.Next
	}
	if cur1 != cur2 {
		return nil
	}

	if n > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}
	n = math.Abs(n)
	for n != 0 {
		n--
		cur1 = cur1.Next
	}
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next

	}
	return cur1
}

func bothLoop(head1, head2, loop1, loop2 *ListNode) *ListNode {
	var cur1 *ListNode
	var cur2 *ListNode
	if loop1 == loop2 {
		cur1 = head1
		cur2 = head2
		var n int = 0
		for cur1 != loop1 {
			n++
			cur1 = cur1.Next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.Next
		}
		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head2
		}
		n = math.Abs(n)
		for n != 0 {
			n--
			cur1 = cur1.Next
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur1 = loop1.Next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.Next
		}
		return nil
	}
}

func getIntersectNode(head1, head2 *ListNode) bool {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(loop1, loop2)
	}
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, head2, loop1, loop2)
	}
	return nil
}
