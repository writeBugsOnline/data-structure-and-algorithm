package main

import (
	"strings"
)


type Queue struct {
	front int
	rare int
	que []*Node
}

func NewQueue() *Queue {
	q :=Queue{
		front: 0,
		rare: 0,
		que: nil,
	}
	return &q
}

func (q *Queue)InQue(n *Node){
	q.rare++
	q.que=append(q.que,n)
}

func (q *Queue)OutQue() *Node {
	if len(q.que)<=0 {
		return nil
	}
	n := q.que[q.front]
	q.que=q.que[1:]
	return n
}





func Test1(){
	pre :=strings.Split("A B D H E I C F J K G"," ")
	mid :=strings.Split("D H B E I A J F K C G"," ")

	tree :=buildTree(pre,mid)

	PreTravel(tree)
	println()
	MidTravel(tree)
}



func main(){
	Test1()
}

