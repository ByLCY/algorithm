package sorting

// CountingSort 计数排序
// @param array:待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#计数排序
func CountingSort(array []int) []int {
	min := array[0]
	max := array[0]
	m := make(map[int]int)
	for _, v := range array {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
		if _, ok := m[v]; !ok {
			m[v] = 1
			continue
		}
		m[v]++
	}
	array = make([]int, 0)
	for ; min <= max; min++ {
		if num, ok := m[min]; ok {
			for ; num > 0; num-- {
				array = append(array, min)
			}
		}
	}
	return array
}
