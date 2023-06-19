package main

import "fmt"

func main() {
	for angka := 1; angka <= 10; angka++ {
		fmt.Println(angka)
	}

	nama := [...]string{
		"Budi",
		"Pambudi",
	}

	for index, nama := range nama {
		fmt.Println("index", index, "=", nama)
	}

	for index, _ := range nama {
		fmt.Println("index", index)
	}

	for i := 0; i < 100; i++ {
		if i == 10 {
			break
		}

		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("angka ganjil :", i)
	}
}
