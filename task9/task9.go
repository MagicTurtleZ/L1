package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

// воркер отправляющий числа из массива
func send(valueList []int, downstream chan int) {
	for _, x := range valueList {
		downstream <- x
	}
	// закрываем канал для записи
	close(downstream)
}

// воркер умножает числа на 2 и отправляет результат во второй канал
func mult(upstream, downstream chan int) {
	for x := range upstream {
		downstream <- x * 2
	}
	// закрываем канал, когда значения для умножения закончатся
	close(downstream)
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	// var valueList []int
	// fmt.Scan(&valueList)
	go send([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, ch1)
	go mult(ch1, ch2)

	// читаем и выводим в stdout результат умножения
	for res := range ch2 {
		fmt.Println(res)
	}
}
