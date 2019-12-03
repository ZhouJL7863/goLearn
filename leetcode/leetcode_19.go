package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//我自己想的是最常规的方法，因此还是采用别人的双指针的方法。
	// var length int
	// length = 0
	// temp := head
	// for temp != nil {
	// 	temp = temp.Next
	// 	length++
	// }

	// fmt.Println(length)
	// if n == length {
	// 	return head.Next
	// }

	// temp1 := head
	// for i := 0; i < length-n-1; i++ {
	// 	head = head.Next
	// }
	// head.Next = (head.Next).Next
	// return temp1
	//这个就是双指针的方法，太秒了，其中dummy是我们采用的哑节点，但是在leetcode上提交的结果来看，
	//并没有什么性能上的提升，smile death。但这个空间复杂度还是比上面这种常规的方法减少了O(L)次数。
	dummy := &ListNode{}
	dummy.Next = head
	first := dummy
	second := dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
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
func main() {
	arr := []int{1, 2, 3, 5}
	ret := makeListNode(arr)

	res := removeNthFromEnd(ret, 3)
	//fmt.Println(res)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
