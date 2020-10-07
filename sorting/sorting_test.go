package sorting

import (
	"reflect"
	"testing"
)

// TestInsertionSort 测试插入排序结果
func TestInsertionSort(t *testing.T) {
	array := []float32{1.1, 1.21, 1, 10, 8}
	check := []float32{1, 1.1, 1.21, 8, 10}
	array = InsertionSort(array)
	if !reflect.DeepEqual(array, check) {
		t.Error("插入排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试希尔排序
func TestShellSort(t *testing.T) {
	array := []int{2, 54, 6, 5, 45, 15, 5, 5, 6, 4, 787, 5, 641, 2, 54, 7, 9, 867, 8}
	check := []int{2, 2, 4, 5, 5, 5, 5, 6, 6, 7, 8, 9, 15, 45, 54, 54, 641, 787, 867}
	array = ShellSort(array)
	if !intEq(array, check) {
		t.Error("希尔排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试选择排序
func TestSelectionSort(t *testing.T) {
	array := []int{4536, 4, 6, 9, 5, 5, 46, 4, 6, 545, 4, 54, 64, 5467, 84, 5}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = SelectionSort(array)
	if !intEq(array, check) {
		t.Error("选择排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试堆排序
func TestHeapsort(t *testing.T) {
	array := []int{454, 4, 5, 4, 545, 6, 5, 345, 2, 44, 54, 5, 544, 336, 998, 83, 7, 34}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = Heapsort(array)
	if !intEq(array, check) {
		t.Error("堆排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	array := []int{45, 3, 6, 46, 5, 4, 8, 745, 6, 4389, 78, 9, 12, 3, 5, 775}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = BubbleSort(array)
	if !intEq(array, check) {
		t.Error("冒泡排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试快速排序
func TestQuickSort(t *testing.T) {
	array := []int{546, 54, 65, 43, 45, 43, 41, 25, 7, 623, 41445, 78, 6, 2}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = QuickSort(array, 0, len(array)-1)
	if !intEq(array, check) {
		t.Error("快速排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试归并排序
func TestMergeSort(t *testing.T) {
	array := []int{456, 4434, 4, 56, 4, 45, 5, 6356, 7, 8, 1, 94, 89, 6215, 4, 0, 49, 4546, 6, 54}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = MergeSort(array)
	if !intEq(array, check) {
		t.Error("归并排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试计数排序
func TestCountingSort(t *testing.T) {
	array := []int{456, 4434, 4, 56, 4, 45, 5, 6356, 7, 8, 1, 94, 89, 6215, 4, 0, 49, 4546, 6, 54}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = CountingSort(array)
	if !intEq(array, check) {
		t.Error("计数排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试桶排序
func TestBucketSort(t *testing.T) {
	array := []int{456, 4, 1, 568, 74, 57, 8, 5, 413, 894, 6, 52, 8645, 1, 3, 98, 65, 22}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = BubbleSort(array)
	if !intEq(array, check) {
		t.Error("桶排序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

// 测试基数排序
func TestRadixSort(t *testing.T) {
	array := []int{487, 96, 538, 51, 235, 67, 8, 01, 6, 50, 484, 648, 798, 486}
	check := make([]int, 0)
	check = append(check, array...)
	check = ShellSort(check)
	array = RadixSort(array)
	if !intEq(array, check) {
		t.Error("基数序错误")
		t.Error("预期结果：", check)
		t.Error("实际结果：", array)
	}
}

func intEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
