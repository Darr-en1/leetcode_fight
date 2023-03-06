package sort

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, left int, right int)  {
	if left < right {
		partition_index := partition(arr, left, right)
		quickSort(arr, left, partition_index-1)
		quickSort(arr, partition_index+1, right)

	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[left]
	j := left
	for i := left + 1; i <= right; i++ {
		if arr[i] < pivot {
			j += 1
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[left], arr[j] = arr[j], arr[left]
	return j
}
