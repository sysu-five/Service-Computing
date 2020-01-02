package main

import "fmt"

// 分组
func partition(array []int, left, right int) int {
	// 选取最左边的数为基数
	base_number := array[left]
	// 先从右往左找，将第一个小于基数的数放到左边，再从左往右找，将第一个大于基数的数放到右边
	// 不断循环直到左边的数都小于基数，右边的数都大于基数，此时将基数放置在中间并返回其下标
	for left < right {
		for (array[right] >= base_number && right > left) {
			right --
		}
		array[left] = array[right]
		for (array[left] <= base_number && left < right) {
			left ++
		}
		array[right] = array[left]
	}
	array[right] = base_number
	return right
}

// 快速排序
func quickSort(array []int, left, right int) {
	// 数组为空时直接返回
	if array == nil {
		return
	}
	// 当分区只有一个数时该分区排序完成
	if left >= right {
		return
	}
	// 分区
	index := partition(array,left,right)
	// 对左分区进一步进行快排
	quickSort(array, left, index-1)
	// 对右分区进一步进行快排
	quickSort(array, index+1, right)
}

// main函数用于测试，排序结果应从小到大
func main() {
	array1 := []int {}
	array2 := []int {17,33,67,5,12,0,9,111}
	fmt.Println(array1)
	quickSort(array1, 0, len(array1)-1)
	fmt.Println(array1)
	fmt.Println(array2)
	quickSort(array2, 0, len(array2)-1)
	fmt.Println(array2)
}