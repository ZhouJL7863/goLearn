package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
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
	temp.Next = head
	return head

}

func insertNode(head *ListNode, num int) *ListNode {
	var node *ListNode = &ListNode{Val: num}
	if head == nil {
		node.Next = node
		return node
	}

	var pre *ListNode = head
	var cur *ListNode = head.Next
	for cur != head {
		if pre.Val <= num && cur.Val >= num {
			break
		}
		pre = cur
		cur = cur.Next
	}
	pre.Next = node
	node.Next = cur
	if head.Val < num {
		return head
	} else {
		return node
	}
}

func main() {
	arr := []int{1, 3, 4}
	head := makeListNode(arr)
	res := insertNode(head, 2)
	for i := 0; i < 4; i++ {
		fmt.Println(res)
		res = res.Next
	}
}
