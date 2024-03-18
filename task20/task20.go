package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseWords1(o string) string {
	// разбиваем строку o на отдельные слова, используя пробел как разделитель
	s := strings.Split(o, " ")
	var res string
	// собираем строку в обратном порядке
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] + " "
	}
	// возвращаем строку res, но перед этим удаляем пробел справа от строки с помощью функции TrimRight из пакета strings.
	return strings.TrimRight(res, " ")
}

func reverseWords2(o string) string {
	// разбиваем входную строку на слова
	s := strings.Fields(o)
	revO := make([]string, 0, len(o))
	// записываем слова в обратно порядке
	for i := len(s) - 1; i >= 0; i-- {
		revO = append(revO, s[i])
	}
	// объединяем revO в одну строку с пробелами и возвращаем результат
	return strings.Join(revO, " ")
}

func main() {
	original := "snow dog sun"
	fmt.Println(reverseWords2(original))
}