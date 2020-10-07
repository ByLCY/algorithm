package sorting

// QuickSort 快速排序
// @param array是待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#快速排序
func QuickSort(array []int, left, right int) []int {
	if left >= right {
		return array
	}
	l, r := left, right
	p := array[l]
	for l < r {
		for l < r && array[r] >= p {
			r--
		}
		if l < r {
			array[l] = array[r]
			l++
		}
		for l < r && array[l] < p {
			l++
		}
		if l < r {
			array[r] = array[l]
			r--
		}
	}
	array[l] = p
	QuickSort(array, left+1, right)
	QuickSort(array, left, right-1)
	return array
}
