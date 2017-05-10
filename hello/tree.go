package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

//1, 1, 2, 3, 5, 8, 13.
/*
    3
   /  \
  1    8
 / \  / \
1   2 5 13
*/
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(root *Tree, ch chan int) {
	//stack := make([]*Tree, 1)
	stack := list.New()
	current := root
	for current != nil {
		stack.PushBack(current)
		current = current.Left
	}
	for stack.Len() > 0 || current != nil {
		current = stack.Back().Value.(*Tree)
		stack.Remove(stack.Back())
		ch <- current.Value
		//fmt.Println("DEBUG: ", current.Value)
		current = current.Right
		for current != nil {
			stack.PushBack(current)
			current = current.Left
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go Walk(t1, chan1)
	go Walk(t2, chan2)
	s1 := ToSlice(chan1)
	s2 := ToSlice(chan2)
	//fmt.Println(s1)
	//fmt.Println(s2)
	return reflect.DeepEqual(s1, s2)
}

func ToSlice(c chan int) []int {
	s := make([]int, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}

func main() {
	tree1 := &Tree{3, nil, nil}
	tree1.Left = &Tree{1, nil, nil}
	tree1.Right = &Tree{8, nil, nil}
	tree1.Left.Left = &Tree{1, nil, nil}
	tree1.Left.Right = &Tree{2, nil, nil}
	tree1.Right.Left = &Tree{5, nil, nil}
	tree1.Right.Right = &Tree{13, nil, nil}

	tree2 := &Tree{8, nil, nil}
	tree2.Left = &Tree{3, nil, nil}
	tree2.Right = &Tree{13, nil, nil}
	tree2.Left.Left = &Tree{1, nil, nil}
	tree2.Left.Right = &Tree{5, nil, nil}
	tree2.Left.Left.Left = &Tree{1, nil, nil}
	tree2.Left.Left.Right = &Tree{2, nil, nil}

	if Same(tree1, tree2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
