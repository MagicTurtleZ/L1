package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func uniq(s string) bool {
	// создаем мапу для првоерки уникальности символа
	m := make(map[rune]struct{}, len(s))
	for _, j := range s {
		// если символ встречается более 1 раза возвращаем false
		if _, ok := m[j]; ok {
			return false
		} else {
			m[j] = struct{}{}
		}
	}

	// иначе возвращаем true
	return true
}

func main() {
	test := [...]string{"abcd", "abCdefAaf", "aabcd"}

	for _, j := range test {
		fmt.Println(j, "-", uniq(j))
	}
}