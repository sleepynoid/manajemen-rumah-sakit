package dokterview

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tugas-uts/controllers"
	"tugas-uts/entities"
)

func MenuInsertDataDokter(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 1: INSERT")
	scanner := bufio.NewScanner(os.Stdin)
	var id int
	var nama string
	var Tlp string
	var jamKerja string
	var spesialis string

	// Validasi untuk input ID
	for {
		fmt.Print("Masukkan ID: ")
		scanner.Scan()
		idInput := scanner.Text()
		if !controllers.IsValidID(idInput) {
			fmt.Println("ID harus berupa angka positif.")
			continue
		}
		id, _ = strconv.Atoi(idInput)
		break
	}
	// Validasi untuk input Nama
	for {
		fmt.Print("Masukkan Nama: ")
		scanner.Scan()
		nama = scanner.Text()
		if !controllers.IsName(nama) {
			fmt.Println("Nama tidak valid.")
			continue
		}
		break
	}
	// Validasi unutuk input spesialis
	for {
		fmt.Print("Masukkan Spesialis: ")
		scanner.Scan()
		spesialis = scanner.Text()
		if !controllers.IsName(spesialis) {
			fmt.Println("Inputan tidak valid.")
			continue
		}
		break
	}
	// Validasi untuk input Telepon
	for {
		fmt.Print("Masukkan No. Telpon: ")
		scanner.Scan()
		Tlp = scanner.Text()
		if !controllers.IsValidTlp(Tlp) {
			fmt.Println("No. Telpon tidak valid.")
			continue
		}
		break
	}
	// Validasi untuk input Jam Kerja
	for {
		fmt.Print("Masukkan Jam Kerja (contoh: 07:00-16:00): ")
		scanner.Scan()
		jamKerja = scanner.Text()
		if !controllers.IsValidJamKerja(jamKerja) {
			fmt.Println("Jam Kerja tidak valid.")
			continue
		}
		break
	}

	controllers.InsertDokter(dataDokter, id, nama, Tlp, jamKerja, spesialis)
	fmt.Println("Data berhasil di tambahkan")
}

func MenuUpdateDokter(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 2: UPDATE")
	var ID int
	fmt.Println("Masukan ID yang ingin di Update : ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&ID)
	scanner.Scan()
	if found := controllers.Search(dataDokter, ID); found != nil {
		var nama string
		var Tlp string
		var jamKerja string
		var spesialis string
		// Validasi untuk input Nama
		for {
			fmt.Print("Masukkan Nama: ")
			scanner.Scan()
			nama = scanner.Text()
			if !controllers.IsName(nama) {
				fmt.Println("Nama tidak valid.")
				continue
			}
			break
		}
		// Validasi unutuk input spesialis
		for {
			fmt.Print("Masukkan Spesialis: ")
			scanner.Scan()
			spesialis = scanner.Text()
			if !controllers.IsName(spesialis) {
				fmt.Println("Inputan tidak valid.")
				continue
			}
			break
		}
		// Validasi untuk input Telepon
		for {
			fmt.Print("Masukkan No. Telpon: ")
			scanner.Scan()
			Tlp = scanner.Text()
			if !controllers.IsValidTlp(Tlp) {
				fmt.Println("No. Telpon tidak valid.")
				continue
			}
			break
		}
		// Validasi untuk input Jam Kerja
		for {
			fmt.Print("Masukkan Jam Kerja (contoh: 07:00-16:00): ")
			scanner.Scan()
			jamKerja = scanner.Text()
			if !controllers.IsValidJamKerja(jamKerja) {
				fmt.Println("Jam Kerja tidak valid.")
				continue
			}
			break
		}
		controllers.UpdateDataDokter(found, nama, Tlp, jamKerja, spesialis)
		fmt.Printf("Data berhasil di update\n\n")
	} else {
		fmt.Printf("data dengan ID %d  tidak ada\n\n", ID)
	}
}

func MenuDeleteDokter(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 2: DELETE")
	fmt.Println("Masukan ID yang ingin di Delete : ")
	var ID int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&ID)
	scanner.Scan()
	if found := controllers.Search(dataDokter, ID); found != nil {
		var confirmation string
		header := []string{"ID", "Nama", "Telepon", "Jam Kerja", "Sepesialis"}
		controllers.GetListDokterById(found, header)
		fmt.Printf("Apakah anda yakin ingin menghapus data ini? (y/t): ")
		fmt.Scan(&confirmation)
		if confirmation == "y" {
			controllers.DeleteDataDokter(dataDokter, ID)
			fmt.Printf("Data berhasil di Delete\n\n")
		} else {
			fmt.Printf("Data gagal di Delete\n\n")
		}
	} else {
		fmt.Printf("Data dengan ID %d tidak tersedia\n\n", ID)
	}
}

func MenuViewAll(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 4: VIEW ALL")
	header := []string{"ID", "Nama", "Telepon", "Jam Kerja", "Sepesialis"}
	dataTable := controllers.GetListDokter(dataDokter, header)
	fmt.Println(dataTable)
}

func MenuViewById(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 5: VIEW BY ID")
	header := []string{"ID", "Nama", "Telepon", "Jam Kerja", "Sepesialis"}
	var ID int
	fmt.Print("masukkan NIP : ")
	fmt.Scan(&ID)
	if found := controllers.Search(dataDokter, ID); found != nil {
		controllers.GetListDokterById(found, header)
	} else {
		fmt.Printf("Data dengan ID %d tidak tersedia\n\n", ID)
	}
}
