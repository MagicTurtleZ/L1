package main

import (
	"fmt"
	"sort"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	union := make(map[int][]float32)
	seqFluc := [8]float32{-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}
	for _, j := range seqFluc {
		// опредлеям в какую группу поместить текущее колебание
		group := int(j)
		group /= 10
		group %= 10
		group *= 10
		// добавляем текущее колебание в объединение
		union[group] = append(union[group], j)
	}
	// сортируем ключи
	keys := make([]int, 0, len(union))
	for k := range union {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// выводим группы в порядке возрастания
	for _, k := range keys {
		fmt.Println(k, ":", union[k])
	}
}
