package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Println("Реши квадратное уравнение")
	fmt.Println("Введи a:")
	fmt.Scan(&a)
	fmt.Println("Введи b:")
	fmt.Scan(&b)
	fmt.Println("Введи c:")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет два корня \nD = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2" + fmt.Sprint(x2))
	} else if D == 0 {
		var x1 float64

		x1 = -b / a * 2
		fmt.Println("Ваше уравнение имеет один корень")
		fmt.Println("X: " + fmt.Sprint(x1))
	} else if D < 0 {
		fmt.Println("Ваше уравнение не имеет корней \nD < 0\nD = " + fmt.Sprint(D))
	}

	var quit string

	fmt.Println("Закрыть программу?")
	fmt.Scan(&quit)

	if quit == "нет" {
		fmt.Scan(&a)
		fmt.Println("Alr")
	} else {
		os.Exit(0)
	}
}
