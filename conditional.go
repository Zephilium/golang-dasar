package main

import "fmt"

func main() {
	nilai := 80

	if nilai >= 80 {
		fmt.Println("A")
	} else if nilai >= 70 {
		fmt.Println("B")
	} else if nilai >= 60 {
		fmt.Println("C")
	} else if nilai >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

}
