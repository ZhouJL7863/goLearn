package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addList1(head1, head2 *ListNode) *ListNode {
	head1 = reverseList(head1)
	head2 = reverseList(head2)
	var ca int = 0
	var n1 int = 0
	var n2 int = 0
	var c1 *ListNode = head1
	var c2 *ListNode = head2
	var node *ListNode
	var pre *ListNode
	for c1 != nil || c2 != nil {
		if c1 == nil {
			n1 = 0
		} else {
			n1 = c1.Val
		}

		if c2 == nil {
			n2 = 0

		} else {
			n2 = c2.Val
		}
		n := n1 + n2 + ca
		pre = node
		node = &ListNode{Val: n % 10}
		node.Next = pre
		ca = n / 10

		if c1 == nil {
			c1 = nil
		} else {
			c1 = c1.Next
		}

		if c2 == nil {
			c2 = nil
		} else {
			c2 = c2.Next
		}
	}
	if ca == 1 {
		pre = node
		node = &ListNode{Val: 1}
		node.Next = pre
	}
	return node
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode = head
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func makeListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head *ListNode = &ListNode{Val: arr[0]}
	temp := head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return head

}

func main() {
	arr1 := []int{9, 3, 7}
	arr2 := []int{6, 3}
	head1 := makeListNode(arr1)
	head2 := makeListNode(arr2)
	res := addList1(head1, head2)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
