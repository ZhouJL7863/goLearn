package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	left  *Node
	right *Node
}

type retrunType struct {
	start *Node
	end   *Node
}

func convert2(head *Node) *Node {
	if head == nil {
		return nil
	}
	return process(head).start
}

func process(head *Node) *retrunType {
	if head == nil {
		return retrunType{nil, nil}
	}
	leftList := process(head.left)  //左节点
	rightList := proces(head.right) //右节点
	if leftList.end != nil {
		leftList.end.right = head
	}
	head.left = leftList.end
	head.right = rightList.start
	if rightList.start != nil {
		rightList.start.left = head
	}
	if leftList.start != nil {
		l
	}
}
