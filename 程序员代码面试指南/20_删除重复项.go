package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeRep1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var temp *ListNode = head
	var n int
	for temp != nil {
		n++
		temp = temp.Next
	}
	mapSet := make(map[int]int, n)
	mapSet[head.Val] = 0
	cur := head.Next
	pre := head
	for cur != nil {
		_, ok := mapSet[cur.Val]
		if ok {
			pre.Next = cur.Next
		} else {
			mapSet[cur.Val] = 0
			pre = cur
		}
		cur = cur.Next
	}
	return head

}

func main() {
	arr := []int{1, 2, 3, 3, 4, 4, 2, 1, 1}
	head := makeListNode(arr)
	res := removeRep1(head)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
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
