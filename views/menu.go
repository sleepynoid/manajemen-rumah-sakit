package views

import "fmt"

func MenuUtama() {
	fmt.Println("1. Dokter")
	fmt.Println("2. Suster")
	fmt.Println("3. Pasien")
	fmt.Println("4. Exit")
	fmt.Print("masukan pilihan menu : ")
}

func SubMenu() {
	fmt.Println("menu pilihan manipulasi data")
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View All")
	fmt.Println("5. View By Id")
	fmt.Println("6. BACK <-")
	fmt.Print("masukkan pilihan menu :")
}
