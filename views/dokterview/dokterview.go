package dokterview

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tugas-uts/controllers"
	"tugas-uts/entities"
)

func MenuInsertDokter(dataDokter *entities.Dokter) {
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

	controllers.InsertDataDokter(dataDokter, id, nama, Tlp, jamKerja, spesialis)
}

func MenuViewAll(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 4: VIEW ALL")
	header := []string{"ID", "Nama", "Telepon", "Jam Kerja", "Sepesialis"}
	dataTable := controllers.GetListDataDokter(dataDokter, header)
	fmt.Println(dataTable)
}

func MenuUpdateDokter(dataDokter *entities.Dokter) {
	fmt.Println("anda masuk pilihan 2: UPDATE")
	var ID int
	fmt.Println("Masukan ID yang ingin di Update : ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&ID)
	scanner.Scan()
	if isEmpty := controllers.Search(dataDokter, ID); isEmpty != nil {
		var nama string
		var Tlp string
		var jamKerja string
		var spesialis string
		fmt.Println("masukan nama baru : ")
		scanner.Scan()
		nama = scanner.Text()
		fmt.Println("masukan no telepon baru : ")
		fmt.Scan(&Tlp)
		scanner.Scan()
		fmt.Println("masukan jam kerja baru : ")
		fmt.Scan(&jamKerja)
		scanner.Scan()
		fmt.Println("masukan spesialis baru : ")
		fmt.Scan(&spesialis)
		scanner.Scan()
		controllers.UpdateDataDokter(isEmpty, nama, Tlp, jamKerja, spesialis)
	} else {
		fmt.Printf("data dengan ID %d  tidak ada\n", ID)
	}
}

/*

func MenuDeleteDosen(dataDsn *entities.Dokter) {
	fmt.Println("anda masuk pilihan 3: DELETE")
	var nip string
	fmt.Print("Masukan ID yang ingin delete : ")
	fmt.Scan(&nip)
	controllers.DeleteDataDosen(dataDsn, nip)
}



func MenuViewByNip(dataDsn *entities.Dokter) {
	fmt.Println("anda masuk pilihan 5: VIEW BY NIP")
	var nip string
	fmt.Print("masukkan NIP : ")
	fmt.Scan(&nip)
	controllers.GetlistDataDosenByNip(dataDsn, nip)
}
*/
