package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binary(mas []int, target int) int {
	left := 0
	right := len(mas) - 1

	for left <= right {
		mid := left + (right - left)/2

		if mas[mid] == target {
			return mid
		}
		if mas[mid] < target {
			left = mid + 1
		}
		if mas[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func initData(arr []int) {
	for i := range arr {
		arr[i] = i + 1
	}
}

func main() {
	data := make([]int, 100)
	initData(data)
	sort.Ints(data)
	fmt.Println(binary(data, 42))
}