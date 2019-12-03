package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePart(head *ListNode, from, to int) *ListNode {
	n := 0
	node1 := head
	var fpre *ListNode
	var tpos *ListNode
	var next *ListNode
	for node1 != nil {
		n++
		if n == from-1 {
			fpre = node1
		}
		if n == to+1 {
			tpos = node1
		}
		node1 = node1.Next
	}
	if from > to || from < 1 || to > n {
		return head
	}
	if fpre == nil {
		node1 = head
	} else {
		node1 = fpre.Next
	}
	node2 := node1.Next
	node1.Next = tpos
	for node2 != tpos {
		next = node2.Next
		node2.Next = node1
		node1 = node2
		node2 = next
	}
	if fpre != nil {
		fpre.Next = node1
		return head
	}
	return node1

}

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
	res := reversePart(head, 2, 4)
	for res != nil{
		fmt.Println(res)
		res = res.Next
	}
}
