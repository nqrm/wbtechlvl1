package main

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}
*/

/*
Переменная justString будет ссылаться на весь участок памяти,
который выделяется в результате создания большой строки в функции createHugeString,
Это приводит к тому, что GC не может его освободить до тех пор пока переменная justString использует его
*/

var justString string

func createHugeString(length int) string {
	s := "a"
	for i := 0; i < length; i++ {
		s += string(i)
	}
	return s
}

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = strings.Clone(v[:100]) 	// Вариант 1
	justString = string([]byte(v[:100])) //	Вариант 2
}
func main() {
	someFunc()
}
