package susterview

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tugas-uts/controllers"
	"tugas-uts/entities"
	"tugas-uts/model/mahasiswa"
)

func MenuInsertMhs(dataMhs *entities.GerbongMhs) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("anda masuk pilihan 1: INSERT")
	var jurusan string
	var npm string
	var nama string
	var ipk float64

	fmt.Print("masukkan nama jurusan: ")
	scanner.Scan()
	jurusan = scanner.Text()

	fmt.Print("masukkan NPM: ")
	fmt.Scan(&npm)
	scanner.Scan()

	fmt.Print("masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Print("masukkan IPK: ")
	scanner.Scan()
	temp := scanner.Text()
	ipk, _ = strconv.ParseFloat(temp, 64)

	data := entities.Mahasiswa{
		Jurusan: jurusan,
		Npm:     npm,
		Nama:    nama,
		Ipk:     ipk,
	}
	mahasiswa.InsertDataMahasiswa(dataMhs, data)
}

func MenuUpdatetMhs(dataMhs *entities.GerbongMhs) {
	fmt.Println("anda masuk pilihan 2: UPDATE")
	var npm string
	fmt.Println("Masukan NPM : ")
	fmt.Scan(&npm)
	controllers.UpdateDataMahasiswa(dataMhs, npm)
}
func MenuDeletetMhs(dataMhs *entities.GerbongMhs) {
	fmt.Println("anda masuk pilihan 3: DELETE")
	var npm string
	fmt.Scan(&npm)
	controllers.DeleteDataMahasiswa(dataMhs, npm)
}
func MenuViewAllMhs(dataMhs *entities.GerbongMhs) {
	fmt.Println("anda masuk pilihan 4. VIEW ALL")
	controllers.GetListDataMahasiswa(dataMhs)
}
func MenuViewByNpmMhs(dataMhs *entities.GerbongMhs) {
	fmt.Println("anda masuk pilihan 5")
	var npm string
	fmt.Print("masukkan npm : ")
	fmt.Scan(&npm)
	controllers.GetlistDatMahasiswaByNpm(dataMhs, npm)
}
