package main

import "fmt"

/*
Slice implementation
Enqueue: O(n)
Dequeue: O(1)
*/

type Queue struct {
	values []int
}

// Constructor of the queue
func createQueue() *Queue {
	return &Queue{
		values: make([]int, 0),
	}
}

// Add to the queue
func (q *Queue) enqueue(value int) {
	q.values = append(q.values, value)
}

// Remove from the queue
func (q *Queue) dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue empty")
		return -1
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value
}

// Check if queue is empty
func (q *Queue) isEmpty() bool {
	if len(q.values) == 0 {
		return true
	}
	return false
}

// Utility to print a queue
func printQueue(q *Queue) {
	fmt.Println(q.values)
}

// Main entry point
func main() {
	queue := createQueue()
	queue.enqueue(2)
	queue.enqueue(7)
	queue.enqueue(9)
	printQueue(queue)

	fmt.Println("Is empty:", queue.isEmpty())

	fmt.Println("Removing:", queue.dequeue())
	fmt.Println("Removing:", queue.dequeue())
	fmt.Println("Removing:", queue.dequeue())

	fmt.Println("Is empty:", queue.isEmpty())
	printQueue(queue)
}
