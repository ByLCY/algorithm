package sorting

// SelectionSort 选择排序
// @param array:待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#选择排序
func SelectionSort(array []int) []int {
	n := len(array)
	index := 0
	for i := 0; i < n; i++ {
		index = i
		for j := i + 1; j < n; j++ {
			if array[index] > array[j] {
				index = j
			}
		}
		array[i], array[index] = array[index], array[i]
	}
	return array
}
