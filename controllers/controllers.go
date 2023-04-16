package controllers

import (
	"regexp"
	"strconv"
	"strings"
	"tugas-uts/entities"
	"tugas-uts/model/doktermodel"
	pasien "tugas-uts/model/pasienmodel"
)

// fungsi insert data dokter
func InsertDokter(dataDokter *entities.Dokter, id int, nama, Tlp, jamKerja, spesialis string) {
	// Mengubah spesialis dan nama menjadi huruf besar
	spesialis = strings.ToUpper(spesialis)
	nama = strings.ToUpper(nama)
	// memanggil fungsi InsertDataDokter di model
	doktermodel.InsertDokter(dataDokter, id, nama, Tlp, jamKerja, spesialis)
}

// Fungsi search data by Id
func Search(dataDokter *entities.Dokter, ID int) *entities.Dokter {
	return doktermodel.Search(dataDokter, ID)
}

// Fungsi Update Data
func UpdateDataDokter(dataDokter *entities.Dokter, nama, Tlp, jamKerja, spesialis string) {
	// Mengubah spesialis dan nama menjadi huruf besar
	spesialis = strings.ToUpper(spesialis)
	nama = strings.ToUpper(nama)
	// Memanggil fungsi Update di doktermodel dengan parameter yang diberikan
	doktermodel.Update(dataDokter, nama, Tlp, jamKerja, spesialis)
}

func DeleteDataDokter(dataDokter *entities.Dokter, ID int) {
	doktermodel.Delete(dataDokter, ID)
}

// Fungsi GetListDokter mengambil daftar dokter dari database dan mengembalikan string berupa tabel
func GetListDokter(dataDokter *entities.Dokter, header []string) string {
	table := doktermodel.GetListDokter(dataDokter, header)
	if table == " " {
		return "\nDATA KOSONG\n"
	}
	return table
}

func GetListDokterById(dataDokter *entities.Dokter, header []string) string {
	return doktermodel.GetListDokterById(dataDokter, header)
}

func IsName(input string) bool {
	// Nama harus terdiri dari satu atau lebih kata
	if len(strings.Fields(input)) < 1 {
		return false
	}
	// Nama tidak boleh mengandung angka atau karakter khusus
	regex := `^[a-zA-Z\s]+$` //`^` merepresentasikan awal dari string. `$` merepresentasikan akhir dari string.
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

type PasienController struct {
	PasienList *pasien.NodePasien
}

func (pc *PasienController) SearchPasien(ID int) *entities.Pasien {
	return pc.PasienList.SearchPasien(ID)
}

func (pc *PasienController) TambahPasien(input entities.Pasien) {
	temp := &entities.Pasien{
		ID:      input.ID,
		Nama:    input.Nama,
		Tlp:     input.Tlp,
		Kondisi: input.Kondisi,
		Riwayat: input.Riwayat,
		Next:    nil,
	}
	pc.PasienList.InsertPasien(temp)
}

func (pc *PasienController) DeletePasien(ID int) bool {
	return pc.PasienList.DeletePasien(ID)
}

func (pc *PasienController) UpdatePasien(ID int, input entities.Pasien) bool {
	update := &entities.Pasien{
		ID:      input.ID,
		Nama:    input.Nama,
		Tlp:     input.Tlp,
		Kondisi: input.Kondisi,
		Riwayat: input.Riwayat,
		Next:    nil,
	}
	return pc.PasienList.UpdatePasien(ID, update)
}

func (pc *PasienController) ListPasien() {
	pc.PasienList.ListPasien()
}
