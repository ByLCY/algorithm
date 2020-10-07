package sorting

// BubbleSort 插入排序
// @param array是待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#冒泡排序
func BubbleSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
