package main

import (
	// "fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// структура для конкурентной записи в хеш-таблицу
type mapWithMutex struct {
	sync.RWMutex
	mp	map[int]int
}

func newMWM() *mapWithMutex {
	return &mapWithMutex{
		sync.RWMutex{},
		// для примера сразу выделяем в нашей мапе место под 5 ключей
		make(map[int]int, 5),
	}
}

func main() {
	// в пакете sync имееется тип map поддерживающий конкурентый доступ к данным
	sm := new(sync.Map)
	// создаем экземпляр нашей структуры MapWithMutex
	mm := newMWM()

	wg := new(sync.WaitGroup)
	// добавим в каждую мапу по 5 ключей
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			// добавляем ключ i записывая случайный инт
			sm.Store(i, rand.Int())
			// лочим мутекс для записи в обычную мапу
			mm.Lock()
			// анлочим по завершению записи
			defer mm.Unlock()
			// записываем в обычную мапу по ключу i случайный инт
			mm.mp[i] = rand.Int()
		}(i)
	}

	wg.Wait()
	/*
	итерация по синкмапе для проверки записи
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
	итерация по обычной мапе для провреки записи 
	for i, j := range mm.mp {
		fmt.Println(i, j)
	}
	*/
}