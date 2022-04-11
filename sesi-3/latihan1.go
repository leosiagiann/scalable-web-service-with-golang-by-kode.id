package main

import "fmt"

func main() {
	// call the function
	if check("Leonardo Siagian", "Laki laki", 21) {
		fmt.Println("User boleh bekerja")
	} else {
		fmt.Println("User tidak boleh bekerja")
	}
}

// Function to check if user is eligible to work
func check(nama, gender string, umur int) bool {
	// check gender is boy and age more than 17
	if gender == "Laki laki" && umur > 17 {
		return true
	}
	// check gender is gitl and age more than 17
	if gender == "Perempuan" && umur > 20 {
		return true
	}
	return false
}
