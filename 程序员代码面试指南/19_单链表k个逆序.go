package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKNode2(head *ListNode,k int){
	if k <2 {
		return head 
	}
	var cur *ListNode = head
	var start *ListNode 
	var next *ListNode 
	var pre *ListNode
	var count int =1
	for cur !=nil{
		next = cur.Next 
		if count == k {
			if pre == nil {
				start = head 
			}else {
				start = head.Next
			}
			resign2(pre,start,cur,next)
			pre = start
			count = 0
		}
		count++
		cur = next
	}
	return head
}

func resign2(left,start,end,right *ListNode){
	var pre *ListNode = start
	cur := start.Next
	var next *ListNode
	for cur !
}