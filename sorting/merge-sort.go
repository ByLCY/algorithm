package sorting

// MergeSort 归并排序
// @param array是待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#归并排序
func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	mid := n / 2
	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	array := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] > right[j] {
			array[index] = right[j]
			j++
			index++
			if j == len(right) {
				copy(array[index:], left[i:])
				break
			}
		} else {
			array[index] = left[i]
			i++
			index++
			if i == len(left) {
				copy(array[index:], right[j:])
				break
			}
		}
	}
	return array
}
