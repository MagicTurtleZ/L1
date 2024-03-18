package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func method1(a, b int) {
	const op = "method1"
	fmt.Println(op, "before:", a, b)
	a, b = b, a
	fmt.Println(op, "after: ", a, b)
}

func method2(a, b int) {
	const op = "method2"
	fmt.Println(op, "before:", a, b)
	a = a + b // 10 + 20 = 30
    b = a - b // 30 - 20 = 10
    a = a - b // 30 - 19 = 20
	fmt.Println(op, "after: ", a, b)
}

func method3(a, b int) {
	const op = "method3"
	fmt.Println(op, "before:", a, b)
	a = a ^ b // 00011110
	b = a ^ b // 00001010 (a ^ b ^ b = a)
	a = a ^ b // 00010100 (a ^ b ^ (a ^ b ^ b) = b)
	fmt.Println(op, "after: ", a, b)
}

func main() {
	a := 10
    b := 20

    method1(a, b)
	method2(a, b)
	method3(a, b)
}
