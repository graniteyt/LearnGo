package main

import "fmt"

func main() {
	num := 3     // переменная num
	if num > 0 { // if - если
		fmt.Println("Number is greater than zero")
	} else if num < 0 { // else if - иначе если
		fmt.Println("Number is lower than zero")
	} else if num == 3 { // иначе (не проверяется)
		fmt.Println("Number equals three")
	}
}
