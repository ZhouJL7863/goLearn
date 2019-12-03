package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listPartition1(head *ListNode, pivot int) *ListNode {
	if head == nil {
		return head
	}
	var cur *ListNode = head
	var n int = 0
	for cur != nil {
		n++
		cur = cur.Next
	}
	arr := make([]int, n)
	cur = head
	for i := 0; i < n; i++ {
		arr[i] = cur.Val
		cur = cur.Next
	}
	arrPartition(arr, pivot)
	res := makeListNode(arr)
	return res

}

func arrPartition(arr []int, pivot int) {
	var index int
	n := len(arr)
	for index != n {
		if arr[index] <= pivot {
			index++
		} else {
			n--
			swap(arr, n, index)
		}
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
func makeListNode(arr []int) *ListNode {
	var res *ListNode = &ListNode{Val: arr[0]}
	temp := res
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Val: arr[i]}
		temp = temp.Next
	}
	return res
}

func main() {
	arr := []int{9, 0, 4, 5, 1}
	head := makeListNode(arr)
	res := listPartition2(head, 3)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

func listPartition2(head *ListNode, pivot int) *ListNode {
	var sH *ListNode   //小的头
	var sT *ListNode   //小的尾
	var eH *ListNode   //相等的头
	var eT *ListNode   //相等的尾
	var bH *ListNode   //大的头
	var bT *ListNode   //大的尾
	var next *ListNode //保存下一个节点

	for head != nil {
		next = head.Next
		head.Next = nil
		if head.Val < pivot {
			if sH == nil {
				sH = head
				sT = head
			} else {
				sT.Next = head
				sT = head
			}
		} else if head.Val == pivot {
			if eH == nil {
				eH = head
				eT = head
			} else {
				eT.Next = head
				eT = head
			}
		} else {
			if bH == nil {
				bH = head
				bT = head
			} else {
				bT.Next = head
				bT = head
			}
		}
		head = next
	}

	if sT != nil {
		sT.Next = eH
		if eT == nil {
			eT = sT
		} else {
			eT = eT
		}
	}

	if eT != nil {
		eT.Next = bH
	}
	if sH != nil {
		return sH
	} else {
		if eH != nil {
			return eH
		}
		return bH
	}
}
