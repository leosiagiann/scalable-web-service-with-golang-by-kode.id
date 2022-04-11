package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	data := []Data{
		{},
		{"Leonardo Siagian", "Bandung", "Back End Engineer", "Tertarik dengan golang"},
		{"Dhandy Muham", "Jakarta", "UI/UX Engineer", "Mudah"},
		{"Vicktor Lambok Desrony", "Medan", "Front End Engineer", "Simple"},
		{"Rikki Naldo Napitupulu", "Tangerang", "Back End Engineer", "Materi enak"},
		{"Muhammad Rusydy M", "Solo", "Front End Engineer", "Mengikuti rules & prosedur perusahaan"},
		{"M. Taufiq Hidayah", "Bekasi", "Back End Engineer", "seru"},
		{"Temmy Alex", "Sitoluama", "UI/UX End Engineer", "gokil"},
		{"Muhamad Ilfani Miftakhudin", "Balige", "Back End Engineer", "Beda dari yang lain"},
		{"Alexander Jordan Himawan", "Tarutung", "Front End Engineer", "Penasaran"},
		{"Clement Prolifel Priyatama", "Sibolga", "Back End Engineer", "Coba coba aja"},
		{"Johannes Pagar M. Manurung", "Siantar", "UI/UX End Engineer", "Mau tau doang"},
		{"Sahir Syatha", "NTT", "Back End Engineer", "Di suruh dosen"},
		{"Andika Satrio Pangestu", "NTB", "Front End Engineer", "Karna ada tugas golang"},
		{"Banta Solagratia", "Sorong", "UI/UX End Engineer", "Mengikuti trend"},
		{"Hosea Felix Hutauruk", "Manokwari", "Front End Engineer", "Powerfull"},
	}
	flag.Parse()
	s := flag.Arg(0)
	index, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Please INput Number")
		os.Exit(0)
	}
	result, success := getData(&data, index)
	if success {
		fmt.Print(result)
	} else {
		fmt.Println("Data Not Found")
	}
}

func getData(data *[]Data, index int) (string, bool) {
	if index > 0 {
		for i, v := range *data {
			if i == index {
				return fmt.Sprintf("Nama\t\t: %s\nAlamat\t\t: %s\nPekerjaan\t: %s\nAlamat\t\t: %s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan), true
			}
		}
	}
	return "", false
}
