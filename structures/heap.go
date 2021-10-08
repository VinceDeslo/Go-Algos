package main

import "fmt"

type Heap struct {
	values []int
}

// Constructor for heap structure
func createHeap(vals []int) *Heap {
	return &Heap{
		values: vals,
	}
}

// Build phase method
func (h *Heap) build() {
	size := len(h.values)
	for idx := (size / 2) - 1; idx >= 0; idx-- {
		h.heapify(size, idx)
	}
}

// Sort phase method
func (h *Heap) sort() {
	for idx := len(h.values) - 1; idx > 0; idx-- {
		h.swap(idx, 0)
		h.heapify(idx, 0)
	}
}

// Method to rearrange the heap (downwards)
func (h *Heap) heapify(size, root int) {
	max := root
	left := h.getLeftChild(root)
	right := h.getRightChild(root)

	if left < size && h.values[left] > h.values[max] {
		max = left
	}
	if right < size && h.values[right] > h.values[max] {
		max = right
	}
	if max != root {
		h.swap(root, max)
		h.heapify(size, max)
	}
}

// Method to fetch left child index
func (h *Heap) getLeftChild(idx int) int {
	return 2*idx + 1
}

// Method to fetch right child index
func (h *Heap) getRightChild(idx int) int {
	return 2*idx + 2
}

// Method to swap values in the heap
func (h *Heap) swap(first, second int) {
	h.values[first], h.values[second] = h.values[second], h.values[first]
}

// Main entry point
func main() {
	vals := []int{8, 16, 3, 7, 5, 1, 12, 11, 4, 13}
	heap := createHeap(vals)
	fmt.Println("Init", heap.values)
	heap.build()
	fmt.Println("Build", heap.values)
	heap.sort()
	fmt.Println("Sort", heap.values)
}
