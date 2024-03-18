package main

import (
	"fmt"
	"time"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/


// интерфейс классов адаптера
type Smart interface {
	Watch() time.Time
}

// адаптируемый класс
type Rolex struct {
}

// Метод адаптируемого класса, который нужно вызвать где-то
func (r *Rolex) ViewTime() string {
	return time.Now().Format(time.TimeOnly)
}

// класс конкретного адаптера
type MiBand struct {
	*Rolex
}

// реализация метода интерфейса, реализующего вызов адаптируемого класса
func (mi *MiBand) Watch() time.Time {
	mi.ViewTime()
	return time.Now()
}

// конструктор адаптера
func NewBand(r *Rolex) Smart {
	return &MiBand{Rolex: r}
}

// основной метод для демонстрации
func main() {
	b := NewBand(&Rolex{})
	fmt.Println(b.Watch().Format(time.DateTime))
}
