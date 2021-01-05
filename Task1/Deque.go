package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
	Prev *Node
}

type Queue struct {
	Size int
	Head *Node
	Tail *Node
}

func (q *Queue) Push_back(value int) bool {
	n := &Node {
		Value: value,
		Prev: nil,
	}
	if q.Size == 0 {
		n.Next = nil
		q.Head = n
		q.Tail = n
	} else {
		n.Prev = q.Tail
		q.Tail.Next = n
		q.Tail = n
	}

	q.Size++

	return true
}

func (q *Queue) Push_front(value int) bool {
	n := &Node {
		Value: value,
		Prev: nil,
	}
	if q.Size == 0 {
		n.Next = nil
		q.Head = n
		q.Tail = n
	} else {
		n.Next = q.Head
		q.Head.Prev = n
		q.Head = n
	}

	q.Size++

	return true
}

func (q *Queue) Pop_back() (int, bool) {
	if q.Size == 0 {
		return 0, false
	} else 
	if q.Size == 1 {
		result := q.Tail.Value
		q.Tail = nil
		q.Head = nil
		q.Size--
		return result, true
	} else {
		res := q.Tail.Value

		q.Tail.Prev.Next = nil 
		q.Tail = q.Tail.Prev
		q.Size--

		return res, true
	}
}

func (q *Queue) Pop_front() (int, bool) {
	if q.Size == 0 {
		return 0, false
	} else 
	if q.Size == 1 {
		result := q.Head.Value
		q.Head = nil
		q.Tail = nil
		q.Size--
		return result, true
	} else {
		res := q.Head.Value

		q.Head.Next.Prev = nil 
		q.Head = q.Head.Next
		q.Size--

		return res, true
	}
}

func (q *Queue) Peek(way string) (int, bool) {
	if q.Size == 0 {
		return 0, false
	}

	switch way {
		case "front":
			return q.Head.Value,true
		case "back":
			return q.Tail.Value,true
		default:
			return 0,false
	}

}

func main() {
	q := &Queue {}
	q.Push_front(3)
	q.Push_front(2)
	q.Push_front(1)
	q.Push_back(-1)
	q.Push_back(-2)
	q.Push_back(-3)


	fmt.Println(q.Head.Value) 			
	fmt.Println(q.Head.Next.Value)		
	fmt.Println(q.Head.Next.Next.Value) 
	fmt.Println(q.Head.Next.Next.Next.Value) 
	fmt.Println(q.Head.Next.Next.Next.Next.Value) 
	fmt.Println(q.Head.Next.Next.Next.Next.Next.Value) 
	fmt.Println("Peek")
	fmt.Println(q.Peek("back"))
	fmt.Println(q.Peek("front"))
	fmt.Println(q.Pop_front())
	fmt.Println(q.Pop_back())
	fmt.Println(q.Pop_back())
	fmt.Println(q.Pop_back())
	fmt.Println(q.Pop_back())
	fmt.Println(q.Pop_back())
	fmt.Println(q.Pop_back())
	fmt.Println(q.Size)
	fmt.Println(q.Head) 			
	fmt.Println(q.Tail) 			

}



