package sorting

// BucketSort 桶排序
// @param array:待排序数组
// 对应笔记 https://0w0.pub/sorting-algorithm/#桶排序
func BucketSort(array []int) []int {
	buckets := make(map[int][]int)
	for _, v := range array {
		if bucket, ok := buckets[v/10]; ok {
			bucket = append(bucket, v)
			continue
		}
		buckets[v/10] = []int{v}
	}
	record := make([]int, 0) // 记录桶详情
	for key, bucket := range buckets {
		record = append(record, key)
		bucket = Heapsort(bucket) // 使用堆排序对桶内元素排序
	}
	record = Heapsort(record) // map是不保证遍历顺序的
	array = make([]int, 0)
	for _, v := range record {
		array = append(array, buckets[v]...)
	}
	return array
}
