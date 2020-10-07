package sorting

// ShellSort 希尔排序
// @param array:待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#希尔排序
// 本例选择的步长为 N/2^k,目前的最佳步长为2^p3^q https://oeis.org/A003586
func ShellSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
			}
		}
		key = key / 2
	}
	return array
}
