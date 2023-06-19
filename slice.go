package main

import "fmt"

func main() {
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
	fmt.Println("Slice hari kerja :", sliceHariKerja)

	sliceHariKerja[0] = "Seninn" //Ketika mengubah slice, maka array juga ikut berubah
	fmt.Println("Array hari :", hari)

	fmt.Println("Slice hari kerja :", sliceHariKerja)
	fmt.Println("Panjang slice :", len(sliceHariKerja))
	fmt.Println("Kapasitas slice :", cap(sliceHariKerja))

	fmt.Println("")

	var sliceTiga = hari[2:5]
	fmt.Println("Slice tiga :", sliceTiga)
	fmt.Println("Panjang slice tiga :", len(sliceTiga))
	fmt.Println("Kapasitas slice tiga :", cap(sliceTiga))

	fmt.Println("")
	sliceBaru := make([]string, 2, 4)
	sliceBaru[0] = "Budi"
	sliceBaru[1] = "Pambudi"

	fmt.Println(sliceBaru)
	fmt.Println(len(sliceBaru))
	fmt.Println(cap(sliceBaru))

	fmt.Println("")
	newSlice := make([]string, len(sliceBaru), cap(sliceBaru)) //Jika menggunakan fungsi make maka slice akan menjadi slice baru

	copy(newSlice, sliceBaru) //Copy slice lama ke slice baru sehingga slice baru tidak mengubah array lama

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "Maulana")
	fmt.Println(newSlice)

	fmt.Println("Slice baru :", sliceBaru)
}
