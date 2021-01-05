package main

import "fmt"

var size = 0
var queue = new(Node)

func main() {
	Push(queue, 1)
	Push(queue, 2)
	Push(queue, 3)
	Push(queue, 4)
	Push(queue, 5)
	Push(queue, 6)
	Push(queue, 7)
	Print(queue)
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	Print(queue)
	Push(queue, 1)
	Push(queue, 2)
	Push(queue, 3)
	Push(queue, 4)
	Push(queue, 5)
	Push(queue, 6)
	Push(queue, 7)
	Print(queue)
}

type Node struct {
	value int
	next *Node
	prev *Node
}

func Print(queue *Node) {
	q := queue

	for q.next != nil {
		fmt.Println("value: ", q.next.value)
		q = q.next
	}
}

func Push(queue *Node, value int) bool {
	n := &Node{
		value: value,
		next:  nil,
		prev: nil,
	}

	if size == 0 {
		queue.next = n
		size++
	} else {

		for queue.next != nil {
			queue = queue.next
		}
		queue.next = n
		n.prev = queue
		size += 1
	}
	return true
}

func Pop(queue *Node) (int, bool) {
 	result := 0 
 	if size == 0{
  		return 0, false
	}
	if size == 1 {
		result = queue.next.value
		queue.prev = nil
		queue.next = nil
		return result, true
	}

	for queue.next != nil {
		queue = queue.next
		if queue.next.next == nil {
			result = queue.next.value
			queue.next = nil
			break
		}
	}
	queue.prev = nil
	t := queue
	queue = t.next
	size--
	return result, true
}