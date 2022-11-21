package main

import (
	"log"
)

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func main() {
	Tree := &Node{Key: 100}
	Tree.Insert(223)
	Tree.Insert(13)
	Tree.Insert(543)
	Tree.Insert(281)
	Tree.Insert(938)
	log.Println(Tree.Search(938))
	log.Println(count)
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}
