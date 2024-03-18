package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func switchCase(x interface{}) {
	const op = "switchCase"
	switch v := x.(type) {
	case int:
		fmt.Println(op, "int:", v)
	case string:
		fmt.Println(op, "string:", v)
	case bool:
		fmt.Println(op, "bool:", v)
	//case chan struct{}: при таком подходе необходимо для каждого типа данных канала делать свой кейс
	default:
		fmt.Println(op, "undefined type:", v)
	}
}

func ref(x interface{}) {
	const op = "ref: "
	fmt.Println(op, reflect.TypeOf(x), reflect.ValueOf(x))
}

func main() {
	var (
		a int
		b string
		c bool
		d chan struct{}
	)

	types := []interface{}{a, b, c, d}

	for _, j := range types {
		switchCase(j)
		ref(j)
	}
}