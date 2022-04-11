package main

import "fmt"

func main() {
	if check("Leonardo Siagian", "Laki laki", 21) {
		fmt.Println("User boleh bekerja")
	} else {
		fmt.Println("User tidak boleh bekerja")
	}
}

func check(nama, gender string, umur int) bool {
	if gender == "Laki laki" && umur > 17 {
		return true
	}
	if gender == "Perempuan" && umur > 20 {
		return true
	}
	return false
}
