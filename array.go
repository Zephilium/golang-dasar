package main

import "fmt"

func main() {
	var nama [2]string

	nama[0] = "Budi"
	nama[1] = "Pambudi"

	fmt.Println(nama)
	fmt.Println(nama[0])
	fmt.Println(nama[1])

	var nilai = [3]int{
		80,
		90,
		70,
	}

	fmt.Println(nilai)

	nilai[2] = 85
	fmt.Println(nilai)

	var arrayKosong [5]int

	fmt.Println(len(arrayKosong))
	fmt.Println(arrayKosong[4])

	fmt.Println("\n")
	var hari = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	fmt.Println(hari)

	var sliceHariKerja = hari[0:5]
	fmt.Println(sliceHariKerja)

}
