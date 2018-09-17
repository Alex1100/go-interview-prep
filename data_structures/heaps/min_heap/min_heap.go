package min_heap

import "math"

type MinHeap struct {
	heap_array []int
	heap_nodes map[int]bool
}

func InitMinHeap() *MinHeap {
	return &MinHeap{
		heap_array: make([]int, 0),
		heap_nodes: make(map[int]bool, 0),
	}
}

func (mh *MinHeap) GetParentIndex(child_index int) int {
	converted_index := float64(child_index)
	return int(math.Floor(converted_index-1.0) / 2.0)
}

func (mh *MinHeap) GetLeftChild(parent_index int) int {
	return ((parent_index * 2) + 1)
}

func (mh *MinHeap) GetRightChild(parent_index int) int {
	return ((parent_index * 2) + 2)
}

func (mh *MinHeap) Add(data int) {
	mh.heap_array = append(mh.heap_array, data)
	mh.BubbleUp(len(mh.heap_array)-1, data)
	mh.heap_nodes[data] = true
}

func (mh *MinHeap) RemoveHead() (int, error) {
	head_node := mh.heap_array[0]
	tail_node := mh.heap_array[len(mh.heap_array)-1]

	mh.heap_array = mh.heap_array[0 : len(mh.heap_array)-1]

	if len(mh.heap_array) > 0 {
		mh.heap_array[0] = tail_node
		mh.BubbleDown(0, tail_node)
	}

	delete(mh.heap_nodes, head_node)
	return head_node, nil
}

func (mh *MinHeap) BubbleDown(parent_index int, parent_data int) {
	if parent_index < len(mh.heap_array) {
		target_index := parent_index
		target_data := parent_data
		left_child_index := mh.GetLeftChild(parent_index)
		right_child_index := mh.GetRightChild(parent_index)

		if left_child_index < len(mh.heap_array) {
			mh.Swap(left_child_index, target_data, left_child_index, mh.heap_array)
		}

		if right_child_index < len(mh.heap_array) {
			mh.Swap(right_child_index, target_data, right_child_index, mh.heap_array)
		}

		if target_index != parent_index {
			mh.heap_array[parent_index] = target_data
			mh.heap_array[target_index] = parent_data
			mh.BubbleDown(target_index, parent_data)
		}

	}
}

func (mh *MinHeap) BubbleUp(child_index int, child_data int) {
	if child_index > 0 {
		parent_index := mh.GetParentIndex(child_index)
		parent_data := mh.heap_array[parent_index]

		if mh.ShouldSwap(child_data, parent_data) {
			mh.heap_array[parent_index] = child_data
			mh.heap_array[child_index] = parent_data
			mh.BubbleUp(parent_index, child_data)
		}
	}
}

func (mh *MinHeap) HasNode(data int) bool {
	return mh.heap_nodes[data] == true
}

func (mh *MinHeap) Contains(data int) bool {
	return mh.heap_nodes[data] == true
}

func (mh *MinHeap) ShouldSwap(left int, right int) bool {
	return left < right
}

func (mh *MinHeap) Swap(target_index int, target_data int, index int, arr []int) {
	data := arr[index]
	if mh.ShouldSwap(data, target_data) {
		target_index = index
		target_data = index
	}
}

func (mh *MinHeap) HeapArray() []int {
	return mh.heap_array
}

func (mh *MinHeap) Size() int {
	return len(mh.heap_array)
}
