package dokterview

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Print("masukkan nip: ")
	fmt.Scan(&id)
	scanner.Scan()
	fmt.Print("masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("masukan No. Telpon: ")
	fmt.Scan(&Tlp)
	fmt.Print("masukan Jam Kerja (contoh: 07:00-16:00): ")
	fmt.Scan(&jamKerja)
	scanner.Scan()
	fmt.Print("masukan Spesialis: ")
	fmt.Scan(&spesialis)
	scanner.Scan()

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
		fmt.Printf("data dengan nip %d  tidak ada\n", ID)
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
