package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

// реализация через таймер
func customSleep(ns time.Duration) {
	// проверяем, что ns больше нуля. Если меньше, то просто возвращаемся
	if ns <= 0 {
		return
	}

	// создаем таймер на указаное кол-во времени
	t := time.NewTimer(ns)
	// просто ждем конца таймера
	<-t.C
}

func main() {
	fmt.Println("Hello, World!")
	customSleep(time.Second * 5)
	fmt.Println("Goodby, World!")
}