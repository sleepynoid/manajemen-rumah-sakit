package controllers

import (
	"regexp"
	"strconv"
	"strings"
	"tugas-uts/entities"
	"tugas-uts/model/doktermodel"
)

// fungsi insert data dokter
func InsertDataDokter(g *entities.Dokter, id int, nama, Tlp, jamKerja, spesialis string) {
	// mengubah inputan spesialis dana nama dari user menjadi huruf kapital semua
	spesialis = strings.ToUpper(spesialis)
	nama = strings.ToUpper(nama)
	// memanggil fungsi InsertDataDokter di model
	doktermodel.InsertDataDokter(g, id, nama, Tlp, jamKerja, spesialis)
}

// Fungsi menampilkan data
func GetListDataDokter(g *entities.Dokter, header []string) string {
	table := doktermodel.GetListDataDokter(g, header)
	if table == " " {
		return "\nDATA KOSONG\n"
	}
	return table
}

// Fungsi search data by Id
func Search(g *entities.Dokter, ID int) *entities.Dokter {
	return doktermodel.Search(g, ID)
}

// Fungsi Update Data
func UpdateDataDokter(g *entities.Dokter, nama, Tlp, jamKerja, spesialis string) {
	spesialis = strings.ToUpper(spesialis)
	nama = strings.ToUpper(nama)
	doktermodel.Update(g, nama, Tlp, jamKerja, spesialis)
}

func IsName(input string) bool {
	// Nama harus terdiri dari satu atau lebih kata
	if len(strings.Fields(input)) < 1 {
		return false
	}
	// Nama tidak boleh mengandung angka atau karakter khusus
	regex := `^[a-zA-Z\s]+$`
	match, _ := regexp.MatchString(regex, input)
	return match
}

// Fungsi untuk validasi ID
func IsValidID(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
}

// Fungsi untuk validasi input No. Telpon
func IsValidTlp(tlp string) bool {
	regex := `^\d{10,}$`
	match, _ := regexp.MatchString(regex, tlp)
	return match
}

// Fungsi untuk validasi input Jam Kerja
func IsValidJamKerja(jamKerja string) bool {
	regex := `^\d{2}:\d{2}-\d{2}:\d{2}$`
	match, _ := regexp.MatchString(regex, jamKerja)
	return match
}

/*

func DeleteDataDokter(g *entities.Dokter {
	doktermodel.Delete(g,
}

func GetlistDataDokterByNip(g *entities.Dokter {
	doktermodel.GetListByNip(g,
}

// fungsi dokter end

// fungsi mahasiswa
func InsertDataMahasiswa(g *entities.GerbongMhs, data entities.Mahasiswa) {
	mahasiswa.InsertDataMahasiswa(g, data)
}

func UpdateDataMahasiswa(g *entities.GerbongMhs, npm string) {
	mahasiswa.Update(g, npm)
}

func DeleteDataMahasiswa(g *entities.GerbongMhs, npm string) {
	mahasiswa.Delete(g, npm)
}

func GetListDataMahasiswa(g *entities.GerbongMhs) {
	mahasiswa.GetListMahasiswa(g)
}

func GetlistDatMahasiswaByNpm(g *entities.GerbongMhs, npm string) {
	mahasiswa.GetlistMahasiswaByNpm(g, npm)
}

//fungsi mahasiswa end
*/
