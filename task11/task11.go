package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
    set1 := map[int]struct{}{
		2: {},
        1: {},
        3: {},
    }

    set2 := map[int]struct{}{
        2: {},
		4: {},
        3: {},
    }

    intersection := make(map[int]struct{})

    // находим пересечение множеств заполняя множество пересечений 
    for key := range set1 {
        if _, ok := set2[key]; ok {
            intersection[key] = struct{}{}
        }
    }

    // выводим результат
    for key := range intersection {
        fmt.Println(key)
    }
}