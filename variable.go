package main

import "fmt"

func main() {
	var nama string

	var (
		umur   = 20
		tinggi = 170
	)

	nama = "Budi"

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(tinggi)

	var (
		uts = 80
		uas = 50
	)

	var lulusUTS bool = uts >= 70
	var lulusUAS bool = uas >= 70

	var lulus = lulusUTS && lulusUAS

	fmt.Println("Lulus :", lulus)

}
