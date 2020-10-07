package sorting

// Heapsort 堆排序
// @param array:待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#堆排序
func Heapsort(array []int) []int {
	n := len(array)
	for i := n / 2; i >= 0; i-- {
		MaxHeap(array, i, n)
	}
	for i := n - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		n--
		MaxHeap(array, 0, n)
	}
	return array
}

// MaxHeap 构建最大堆
// @param array:待构成的数组
// @param i:第i个子树
// @param n:数组长度
func MaxHeap(array []int, i, n int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i // 最大值的下标

	if left < n && array[left] > array[max] {
		max = left
	}

	if right < n && array[right] > array[max] {
		max = right
	}

	if max != i {
		array[i], array[max] = array[max], array[i]
		MaxHeap(array, max, n)
	}
}
