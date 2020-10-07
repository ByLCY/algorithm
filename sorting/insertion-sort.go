package sorting

// InsertionSort 插入排序
// @param array是待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#插入排序
func InsertionSort(array []float32) []float32 {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for ; j >= 0 && key < array[j]; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = key
	}
	return array
}
