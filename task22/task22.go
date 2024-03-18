package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

func calc(a, b float64, oper string) *big.Float {
	// конвертируем числа в big.Float
	bigA := big.NewFloat(a)
	bigB := big.NewFloat(b)
	// для каждой арифметической операции создаем анонимную функцию, 
	// которая проводит соответствующую операцию и устанавливает точность результата
	m := map[string]func(a, b *big.Float) *big.Float {
		"+": func(a, b *big.Float) *big.Float {return new(big.Float).Add(a, b).SetPrec(20)},
		"-": func(a, b *big.Float) *big.Float {return new(big.Float).Sub(a, b).SetPrec(20)},
		"*": func(a, b *big.Float) *big.Float {return new(big.Float).Mul(a, b).SetPrec(20)},
		"/": func(a, b *big.Float) *big.Float {return new(big.Float).Quo(a, b).SetPrec(20)},
	}

	return m[oper](bigA, bigB)
}

func main() {
	var (
		a, b	float64
		oper	string
	)
	fmt.Scan(&a, &oper, &b)
	fmt.Println(calc(a, b, oper))
}