package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/


// устанавливаем указанный бит в 0
func setBit0(num int64, pos uint8) int64 {
    var set int64
    set = 1 << pos
    num &= ^set
    return num
}

// устанавливаем указанный бит в 1
func setBit1(num int64, pos uint8) int64 {
	var set int64
    set = 1 << pos
    num |= set
    return num
}

func main() {
    var (
        num int64
        pos uint8
        set int8
    )
    fmt.Println("Введите число: ")
    fmt.Scan(&num)
    fmt.Println("Выберите бит для утсановки: ")
    fmt.Scan(&pos)
    fmt.Println("Заменить на? (1 или 0)")
    fmt.Scan(&set)

    if pos > 63 {
        fmt.Printf("Невозможно изменить %d бит: ", pos)
        return
    }

    switch set {
    case 0:
        num = setBit0(num, pos)
        fmt.Println(num)
    case 1:
        num = setBit1(num, pos)
        fmt.Println(num)
    default: 
        fmt.Println("Бит может быть только 1 или 0")
    }

}