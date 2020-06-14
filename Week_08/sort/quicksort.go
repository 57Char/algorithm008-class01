package sort

func quickSort(array []int, begin, end int) {
	if end <= begin {
		return
	}
	pivot := partition(array, begin, end)
	quickSort(array, begin, pivot-1)
	quickSort(array, pivot+1, end)
}

func partition(array []int, begin, end int) int {
	// pivot: 标杆位置，counter: 小于pivot的元素的个数
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if array[i] < array[pivot] {
			array[i], array[counter] = array[counter], array[i]
			counter++
		}
	}
	array[pivot], array[counter] = array[counter], array[pivot]
	return counter
}
