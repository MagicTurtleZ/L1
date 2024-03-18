package main

import "fmt"

/*
Удалить i-ый элемент из слайса
*/

//O(1) без сохранения порядка
func fast(s []int, i int) []int {
	// копируем последний элемент в индекс i.
	s[i] = s[len(s) - 1]
	// усекаем срез
	s = s[:len(s) - 1]
	return s
}

// O(n) с сохранением порядка
func long(s []int, i int) []int {
	// выполняем сдвиг влево на один индекс.
	copy(s[i:], s[i+1:])
	// усекаем срез
	s = s[:len(s) - 1]
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = fast(s, 2)
	fmt.Println(s)
	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = long(s, 2)
	fmt.Println(s)
}