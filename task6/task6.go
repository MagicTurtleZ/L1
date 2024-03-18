package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// отмена с помощью канала
func withChanCancel(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	const op = "withChanCancel"
	for {
		select {
		case <-ch:
			return
		default:
			fmt.Printf("%s: imitation of work\n", op)
			time.Sleep(time.Second * 5)
		}
	}
}

// отмена с помощью контекста
func withContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	const op = "withContext"
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("%s: imitation of work\n", op)
			time.Sleep(time.Second * 5)
		}
	}
}

// отмена с помощью runtime.Goexit
func withGoexit() {
	const op = "withGoexit"
	for {
		fmt.Printf("%s: imitation of work\n", op)
		time.Sleep(time.Second * 5)
		runtime.Goexit()
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})
	// создаем по экземпляру каждой функции
	go withChanCancel(ch, &wg)
	go withContext(ctx, &wg)
	go withGoexit()

	// ждем сообщения от каждой горутины
	time.Sleep(time.Second * 5)
	// посылаем в канал отмены
	ch <- struct{}{}
	// вызываем отмену контекста
	cancel()
	// 3 горутина сама завершится, когда код дойдет до функции runtime.Goexit, дожидаемся этого
	runtime.Gosched()
	// ждём завершения всех горутин 
	wg.Wait()
}