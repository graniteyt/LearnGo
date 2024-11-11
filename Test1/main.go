package main // пакет в котором находится программа

import "fmt" // импортировать fmt

func main() { // точка входа  в программу
	// fmt - встроенная функция в Go
	fmt.Println("Hello World!") // Println === print / console.log
	fmt.Println("Привет от Golang!")

	var age int8 = 16 //int - integer(тип переменной), var - переменная
	fmt.Println(age)

	var number float64 = 275.656 //float - число с плавающей точкой
	fmt.Println(number)

	age1 := 15.5 // краткий формат создания переменных
	fmt.Println(age1)

	var name string = "Ivan"
	fmt.Println(name)
	fmt.Println(len(name)) // len - длина строки

	// name := = "Ivan"
	// fmt.Println(name)

	var nm string
	fmt.Println("What is your name?")
	fmt.Scan(&nm) // ввод пользователя
	fmt.Println("Hello " + nm + "!")
	fmt.Println("How old might you be?")
	var age3 int8
	fmt.Scan(&age3)
	fmt.Println("You're " + fmt.Sprint(age3) + " years old!") // Sprint - преобразование числа в переменную
}
