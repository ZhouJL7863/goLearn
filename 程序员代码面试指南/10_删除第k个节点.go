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
	res := &ListNode{Val: arr[0]}
	temp := res
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return res
}

func removeLastKthNode(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return head
	}
	var cur *ListNode
	cur = head
	for cur != nil {
		k--
		cur = cur.Next
	}
	fmt.Println(k)
	if k == 0 {
		head = head.Next
	}
	if k < 0 {
		cur = head
		k = k + 1 //这一步需要想清楚，当你复制的时候，已经默认了为第一个节点，因此需要加1,还是java中++i好用啊
		for ; k != 0; k++ {
			cur = cur.Next

		}
		cur.Next = cur.Next.Next

	}
	return head
}

func main() {
	arr := []int{1, 2, 3, 4}
	head := makeListNode(arr)
	res := removeLastKthNode(head, 1)
	//fmt.Println(res)
	for res != nil {
		fmt.Println(res)
		res = res.Next

	}
}
