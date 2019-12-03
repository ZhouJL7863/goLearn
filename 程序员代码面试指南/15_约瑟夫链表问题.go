package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func josephusKill(head *ListNode, m int) *ListNode {
	if head == nil || head.Next == head || m < 1 {
		return head
	}
	var last *ListNode = head
	for last.Next != head {
		last = last.Next
	}

	count := 0
	for head != last {
		count++
		if count == m {
			last.Next = head.Next
			count = 0
		} else {
			last = last.Next
		}
		head = last.Next
	}
	return head

}

func makeListNode(arr []int) *ListNode {
	var head *ListNode
	var temp *ListNode
	head = &ListNode{Val: arr[0]}
	temp = head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	temp.Next = head
	return head
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := makeListNode(arr)
	res := josephusKill(head, 2)
	fmt.Println(res)
}

//只有看着别人的代码才能写出代码来，自己就是写不出代码

func josephusKill2(head *ListNode, m int) *ListNode {
	if head == nil || head.Next == head || m < 1 {
		return head
	}
	var cur *ListNode = head.Next
	var tmp int = 1
	for cur != head {
		tmp++
		cur = cur.Next
	}
	tmp = getLive(tmp, m)
	tmp--
	for tmp != 0 {
		tmp--
		head = head.Next
	}
	head.Next = head
	return head
}

func getLive(i, m int) int {
	if i == 1 {
		return 1
	}
	return (getLive(i-1, m)+m-1)%i + 1
}
