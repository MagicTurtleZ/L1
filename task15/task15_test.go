package main

import (
	"math/rand"
	"testing"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func main() {
	someFunc()
}
*/

var justString string
const asciiChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

// для примера создаем строку только из аски
func createHugeString(size int) string {
	result := make([]byte, size)
    for i := range result {
        result[i] = asciiChars[rand.Intn(len(asciiChars))]
    }
	
	return string(result)
}

func WithoutBuf() {
	v := createHugeString(1 << 10)
	/* т.к. строка в го это слайс с ограничением на чтение, 
	тогда после выхода из функции появляется вероятность утечки, 
	поскольку она будет ссылаться на 100 элементов от нашей оригинальной строки и тогда
	строка из 1024 элементов будет хранится в памяти.*/
	justString = v[:100]
}

func WithBuf() {
	v := createHugeString(1 << 10)
	/* чтобы garbage collector мог добраться до оригинальной строки из 1024 символов нам необходимо, 
	чтобы по выходу из функции на нее не осталось ссылающихся элементов. 
	Этого можно добиться созданием нового среза на 100 байт со значениями 100 байт из оригинальной строки, 
	для этого используется функция copy что не создает зависимости между buf и v, а просто копирует её значения.
	Теперь по выходу из функции justString будет ссылаться на массив из 100 элементов, а оригинал из 1024 удалится из памяти*/
	buf := make([]rune, 100)
	copy(buf, []rune(v))
	justString = string(buf)
}

// на данных тестах можно пронаблюдать, что для избежания утечки мы алоцируем в 2 раза больше памяти.
func BenchmarkWithoutBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithoutBuf() // 514kB
	}
}

func BenchmarkWithBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithBuf() // 1028.38kB
	}
}