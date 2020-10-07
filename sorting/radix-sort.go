package sorting

// RadixSort 基数排序
// @param array:待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#基数排序
func RadixSort(array []int) []int {
	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	radix := make([][]int, 10)
	i := 1
	for max/i != 0 {
		for _, num := range array {
			radix[num/i%10] = append(radix[num/i%10], num)
		}
		array = make([]int, 0)
		for i := 0; i < 10; i++ {
			array = append(array, radix[i]...)
			radix[i] = make([]int, 0)
		}
		i *= 10
	}
	return array
}
