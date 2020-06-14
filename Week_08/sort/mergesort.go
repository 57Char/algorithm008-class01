package sort

func mergeSort(array []int, left, right int) {
	if right <= left {
		return
	}
	mid := left + (right-left)>>1 // (left + right) / 2
	mergeSort(array, left, mid)
	mergeSort(array, mid+1, right)
	merge(array, left, mid, right)
}

func merge(array []int, left, mid, right int) {
	// temp := make([]int, 0, right-left+1) // 中间数组
	var temp []int
	i, j := left, mid+1

	for i <= mid && j <= right {
		if array[i] <= array[j] {
			temp = append(temp, array[i])
			i++
		} else {
			temp = append(temp, array[j])
			j++
		}
	}

	for i <= mid {
		temp = append(temp, array[i])
		i++
	}
	for j <= right {
		temp = append(temp, array[j])
		j++
	}

	for i, v := range temp {
		array[left+i] = v
	}
}
