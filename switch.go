package main

import "fmt"

func main() {
	nilai := 80

	switch {
	case nilai >= 80:
		fmt.Println("A")
	case nilai >= 70:
		fmt.Println("B")
	case nilai >= 60:
		fmt.Println("C")
	case nilai >= 50:
		fmt.Println("D")
	default:
		fmt.Println("E")

	}
}
