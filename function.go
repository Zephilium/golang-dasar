package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World")
}

func sum(a1 int, a2 int) {
	fmt.Println(a1 + a2)
}

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	helloWorld()
	sum(5, 10)
	fmt.Println(sayHello("Pambudi"))
}
