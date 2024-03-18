package main

import (
	"fmt"
	"math/rand"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func quicksort(step []int) []int {
	if len(step) < 1 {
		return step
	}

	pos := rand.Intn(len(step))
	pivot := step[pos]

	var less, great []int
	for i, j := range step {
		if i == pos {
			continue
		}
		if j < pivot {
			less = append(less, j)
		} else {
			great = append(great, j)
		}
	}

	return append(append(quicksort(less), pivot), quicksort(great)...) 
}


func initData(arr []int) {
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
}

func main() {
	data := make([]int, 20)
	initData(data)
	fmt.Println(data)
	fmt.Println(quicksort(data))
}