package main

import "fmt"

func main() {

	var person map[string]string = map[string]string{
		"nama":   "Budi",
		"alamat": "Indonesia",
	}

	person["pekerjaan"] = "Karyawan Swasta"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person["pekerjaan"])

	book := make(map[string]string)
	book["judul"] = "Belajar Golang"
	book["pengarang"] = "Budi Pambudi"
	book["penerbit"] = "Gramed"

	fmt.Println(book)

	delete(book, "penerbit")

	fmt.Println(book)
	fmt.Println(len(book))

}
