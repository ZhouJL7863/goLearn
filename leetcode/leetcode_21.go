package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	arr := []int{}
	for l1 != nil {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		arr = append(arr, l2.Val)
		l2 = l2.Next
	}

	sort.Ints(arr)

	return makeListNode(arr)
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
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := makeListNode(arr1)
	l2 := makeListNode(arr2)
	ret := mergeTwoLists(l1, l2)
	for ret != nil {
		fmt.Println(ret)
		ret = ret.Next
	}
}
