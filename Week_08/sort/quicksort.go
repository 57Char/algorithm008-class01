package sort

func quickSort(array []int, left, right int) {
	if right <= left {
		return
	}
	pivot := partition(array, left, right)
	quickSort(array, left, pivot-1)
	quickSort(array, pivot+1, right)
}

func partition(array []int, left, right int) int {
	// pivot: 标杆位置，counter: 小于pivot的元素的个数
	pivot, counter := right, left
	for i := left; i < right; i++ {
		if array[i] < array[pivot] {
			array[i], array[counter] = array[counter], array[i]
			counter++
		}
	}
	array[pivot], array[counter] = array[counter], array[pivot]
	return counter
}
