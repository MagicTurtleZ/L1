package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

// в данном случае, мы просто записываем символы в обратном порядке
func reverse1(o string) string {
	// используем руны, т.к. есть условие: "Символы могут быть unicode."
	res := make([]rune, len(o))
	for i, j := range []rune(o) {
		res[len(res) - 1 - i] = j
	}
	return string(res)
}

// в данном случае, мы идем от двух концов к середине и меняем символы местами
func reverse2(o string) string {
	// используем руны, т.к. есть условие: "Символы могут быть unicode."
	res := []rune(o)
	i, j := 0, len(res) - 1
	for i <= j{
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}

func main() {
	var original string
	buf := bufio.NewReader(os.Stdin)
	original, err := buf.ReadString('\n')
	if err != nil {
		return
	}
	fmt.Println(reverse2(original))
}