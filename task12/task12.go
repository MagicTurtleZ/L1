package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	seqStr := [...]string{"cat", "cat", "dog", "cat", "tree"}
	uniq := make(map[string]struct{}, len(seqStr))
	set := make([]string, 0, len(seqStr)) 
	for _, j := range seqStr {
		if _, ok := uniq[j]; !ok {
			set = append(set, j)
			uniq[j] = struct{}{}
		}
	}

	fmt.Println(set)
}