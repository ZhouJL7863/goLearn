package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(arr []int) *ListNode {
	var temp *ListNode
	var head *ListNode

	head = &ListNode{Val: arr[0]}
	temp = head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return head
}
func main() {
	arr := []int{1, 2, 3, 3, 2, 2}
	head := makeListNode(arr)
	res := isPalindrome3(head)
	fmt.Println(res)
}

func isPalindrome3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var n1 *ListNode = head
	var n2 *ListNode = head
	//查找中间节点
	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	n2 = n1.Next //右部分第一个节点
	n1.Next = nil
	var n3 *ListNode
	for n2 != nil {
		n3 = n2.Next
		n2.Next = n1
		n1 = n2
		n2 = n3
	}
	n3 = n1
	n2 = head
	var res bool = true
	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			res = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	n1 = n3.Next
	n3.Next = nil
	for n1 != nil {
		n2 = n1.Next
		n1.Next = n3
		n3 = n1
		n1 = n2
	}
	return res

}
