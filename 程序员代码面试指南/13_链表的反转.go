package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	var pre *ListNode
// 	var next *ListNode
// 	for head != nil {
// 		next = head.Next
// 		head.Next = pre
// 		pre = head
// 		head = next
// 	}
// 	return pre
// }
func makeListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	ret := &ListNode{Val: arr[0]}
	temp := ret
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return ret
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := makeListNode(arr)
	res := reverseList(head)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

type DoubleNode struct {
	Val  int
	Last *DoubleNode
	Next *DoubleNode
}

func DoubleNode(head *DoubleNode) *DoubleNode {
	var pre *DoubleNode
	var next *DoubleNode
	for head != nil {
		next = head.Next
		head.Next = pre
		head.Last = next
		head = next
		pre = head
	}
	return pre
}
