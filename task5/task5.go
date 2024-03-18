package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func sender(secCount int, ch chan int) {
	// создаем таймер для отсчета когда программа должна завершиться 
	timer := time.NewTimer(time.Duration(secCount * int(time.Second)))
	for {
		// каждую секунду отправляем случайное число от 0 до 1000 в канал пока таймер не завершится
		select {
		case <-timer.C:
			fmt.Println("Timer expired!")
			// по истечению времени закрываем канал
			close(ch)
			return
		case <-time.After(time.Second): 
			ch <- rand.Intn(1000)
		}
	}
}

func main() {
	var secCount int
	fmt.Println("Enter the program lifetime in seconds:")
	fmt.Scan(&secCount) 
	
	ch := make(chan int)
	// функция для последовательной отправки в канал
	go sender(secCount, ch)

	// чтение канала пока тот не закроется 
	for i := range ch {
		fmt.Println(i)
	}
}