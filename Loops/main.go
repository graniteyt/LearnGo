package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { // цикл for
		fmt.Printf("Hello  %d\n", 5)
	}

	for i := 0; i <= 100; i++ { // четные числа от 0 до 100
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 100; i++ { // четные числа от 0 до 100
		if i%2 != 0 {
			continue // пропуск итерации
		}
		fmt.Println(i)
	}

	for i := 0; i <= 100; i++ { // четные числа от 0 до 100
		if i == 50 {
			break // остановка цикла
		}
		fmt.Println(i)
	}

	nums := []int{1, 2, 3, 4}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	for _, element := range nums {
		fmt.Printf("Element %d\n", element)
	}
}
