package main

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

type Human struct {
	field1 string
	field2 string
}

func (h Human) method1(t string) {
	return
}

func (h Human) method2(t string) {
	return
}

type Action struct {
	// анонимное встраивание
	Human // таким образом мы можем использовать данные и методы структуры `Human` так,
	// как если бы они были объявлены для структуры `Action`.

	// именнованное встраивание
	hum Human // в данном случае необходимо явно указывать имя присвоенное структуре `Human` для доступа к ее данным и методам.
}

func main() {
	a := Action{
		Human: Human{
			field1: "aboba",
			field2: "amogus",
		},
		hum: Human{
			field1: "test1",
			field2: "test2",
		},
	}
	// пример использования полей и методов структуры `Human` при анонимном встраивании
	anon := a.field1
	a.method1(anon)

	// пример использования полей и методов структуры `Human` как переменной hum при именованном встраивании
	neanon := a.hum.field1
	a.hum.method2(neanon)
}
