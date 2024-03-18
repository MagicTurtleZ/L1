package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2 + 3^2 + 4^2…) с использованием конкурентных вычислений.
*/

// Task3_1 - синзронизация с помощью вейтгруппы и атомиков.
func Task3_1() {
	mas := [...]int{2, 4, 6, 8, 10}
	res := new(int32)
	// объявляем вейтгруппу
	wg := sync.WaitGroup{}

	for _, j := range mas {
		// увеличиваем счетчик для ожидания завершения каждой горутины
		wg.Add(1)
		// создаем горутинку с вычеслением квадрата числа,
		// передаем число как копию в анонимную функцию, чтобы избежать состояния гонки
		go func(j int) {
			// уменьшаем счетчик WaitGroup при завершении горутины
			defer wg.Done()
			tmp := j * j
			// атомарно добавляем квадрат текущего числа к сумме
			atomic.AddInt32(res, int32(tmp))
		}(j)
	}
	// ждем завершения всех горутин, чтобы преждевременно не выйти из функции
	wg.Wait()

	// вывод суммы квадратов
	fmt.Println(*res)
}

// Task3_2 - синзронизация с помощью канала.
func Task3_2() {
	mas := []int{2, 4, 6, 8, 10}
	res := 0
	ch := make(chan int)

	for _, j := range mas {
		// создаем горутинку с вычеслением квадрата числа,
		// передаем число как копию в анонимную функцию, чтобы избежать состояния гонки
		go func(j int) {
			// записываем в канал квадрат числа
			ch <- j * j
		}(j)
	}

	// читаем из канала столько раз сколько у нас есть чисел в массиве
	for range mas {
		res += <-ch
	}

	// вывод суммы квадратов
	fmt.Println(res)

	// закрываем канал
	close(ch)
}

// Task3_3 - конвейер для вычислений суммы квадратов чисел
func Task3_3() {
	upCh := make(chan int)
	dwCh := make(chan int)
	res := 0

	// первая горутина отправляет числа которые нужно возвести в квадрат в канал upCh
	go func() {
		mas := [...]int{2, 4, 6, 8, 10}

		for _, j := range mas {
			upCh <- j
		}
		// когда все числа отправлены закрываем канал
		close(upCh)
	}()

	// вторая горутина получает числа из канала upCh, возведит их в квадрат
	// и отправляет результаты в канал dwCh
	go func() {
		for i := range upCh {
			dwCh <- i * i
		}
		// когда квадраты всех чисел посчитаны закрываем канал
		close(dwCh)
	}()

	// читаем канал dwCh пока в не закроется и добаавляем полученные числа к результату
	for i := range dwCh {
		res += i
	}

	// вывод суммы квадратов
	fmt.Println(res)
}

// Task3_4 - синзронизация с помощью вейтгруппы и мьютексов.
func Task3_4() {
	// объявление структуры sqrt, содержащей вложенный Mutex и переменную res для хранения результата
	var sqrt struct {
		sync.Mutex
		res int
	}
	// объявляем вейтгруппу
	wg := sync.WaitGroup{}

	for _, j := range []int{2, 4, 6, 8, 10} {
		// увеличиваем счетчик для ожидания завершения каждой горутины
		wg.Add(1)
		// создаем горутинку с вычеслением квадрата числа,
		// передаем число как копию в анонимную функцию, чтобы избежать состояния гонки
		go func(j int) {
			// уменьшаем счетчик WaitGroup при завершении горутины
			defer wg.Done()
			// лочим мьютекс для безопасного доступа к переменной res
			sqrt.Lock()
			sqrt.res += j * j
			// разблокировка мьютекса 
			sqrt.Unlock()
		}(j)
	}

	// ждем завершения всех горутин, чтобы преждевременно не выйти из функции
	wg.Wait()

	// вывод суммы квадратов
	fmt.Println(sqrt.res)
}

func main() {
	Task3_1()
	fmt.Println("=========================================")
	time.Sleep(time.Second)
	Task3_2()
	fmt.Println("=========================================")
	time.Sleep(time.Second)
	Task3_3()
	fmt.Println("=========================================")
	time.Sleep(time.Second)
	Task3_4()
}
