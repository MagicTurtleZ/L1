package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// при получении данных сообщаем какой воркер и какие данные обработал
	for data := range ch {
		fmt.Println("Worker №", id, " received: ", data)
	}
	// сообщаем, что воркер закончил свою работу
	fmt.Println("Worker №", id, "has finished")
}

func writeToChanel(ch chan int) {
	defer close(ch)
	// канал для перехвата сискола о завершении работы
	gs := make(chan os.Signal)
	signal.Notify(gs, os.Interrupt)

	// бесконечный цикл для постоянной записи данных в канал
	for {
		select {
			// каждую секунду отправляем в канал случайное число в диапазоне от 0 до 1000
		case <- time.After(time.Second):
			ch <- rand.Intn(1000)
			// при получении сискола закрываем канал и выходим из функции (предварительно закрыв канал для данных)
		case <- gs:
			close(gs)
			return
		}
	}

}

func main() {
	var maxWorkers int 
	// получение кол-ва воркеров 
	fmt.Scan(&maxWorkers)
	// канал для данных
	ch := make(chan int)
	// веййтгруппа для синхронизации завершения воркеров
	wg := new(sync.WaitGroup)
	wg.Add(maxWorkers)

	// запуск набора воркеров 
	for i := 0; i < maxWorkers; i++ {
		go worker(i, ch, wg)
	}

	// запуск потоковой записи
	go writeToChanel(ch)

	// при получении сигнала завершения ждем окончания работы всех воркеров
	wg.Wait()
	// сообщаем, что все воркеры закочили своё выполнение
	fmt.Println("All workers have finished")
}