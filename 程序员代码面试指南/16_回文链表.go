package main

//左神利用栈的思想，利用一个栈先来保存这个链表的数组，之后重新浏览一次链表，将栈的弹出结构与链表进行比较，如果不相等的话
//就证明了不是回文链表。
import (
	"fmt"
	"stack"
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
	arr := []int{1, 2, 3, 3, 2, 1}
	head := makeListNode(arr)
	res := isPalindrome1(head)
	fmt.Println(res)
}

func isPalindrome1(head *ListNode) bool {
	var newStack stack.Stack
	var temp *ListNode = head
	var value int
	for temp != nil {
		newStack.Push(temp.Val)
		temp = temp.Next
	}

	temp = head
	for temp != nil {
		value = temp.Val
		if value != newStack.Pop() {
			return false
		}
		temp = temp.Next
	}
	return true
}
