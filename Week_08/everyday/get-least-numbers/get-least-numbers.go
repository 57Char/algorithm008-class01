package get_least_numbers

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

// TODO: 优化
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	heap := []int{}
	for i := 0; i < len(arr); i++ {
		if len(heap) < k {
			heap = buildHeapify(heap, arr[i])
		} else if arr[i] < heap[0] {
			heap = heapifyUp(heap, arr[i])
		}
	}
	return heap
}

func buildHeapify(heap []int, val int) []int {
	heap = append(heap, val)
	maxPos := len(heap) - 1
	parent := (maxPos - 1) / 2
	for parent >= 0 && heap[parent] < heap[maxPos] {
		heap[maxPos], heap[parent] = heap[parent], heap[maxPos]
		maxPos = parent
		parent = (parent - 1) / 2
	}
	return heap
}

func heapifyUp(heap []int, val int) []int {
	heap[0] = val
	pos := 0
	for {
		tmp := pos
		child1 := 2*pos + 1
		child2 := child1 + 1
		if child1 < len(heap) && heap[child1] > heap[tmp] {
			tmp = child1
		}
		if child2 < len(heap) && heap[child2] > heap[tmp] {
			tmp = child2
		}
		if tmp == pos {
			break
		}
		heap[tmp], heap[pos] = heap[pos], heap[tmp]
		pos = tmp
	}
	return heap
}
