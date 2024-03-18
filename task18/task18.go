package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

type counter struct {
	// мьютекс для работы нашего счетчика в конкурентной среде
	sync.Mutex
	// наш счетчик
	count int
}

func main() {
	c := new(counter)
	// добавляем вейтгруппу для демонстрации работы структуры
	wg:= new(sync.WaitGroup)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			c.Lock()
			c.count++
			c.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}