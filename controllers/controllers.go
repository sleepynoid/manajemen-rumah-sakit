package controllers

import (
	"regexp"
	"strconv"
	"strings"
	"tugas-uts/entities"
	dokter "tugas-uts/model/doktermodel"
	pasien "tugas-uts/model/pasienmodel"
)

type DokterController struct {
	DokterList *dokter.NodeDokter
}

// fungsi insert data dokter
func (dl *DokterController) InsertDokter(input *entities.Dokter) {
	temp := &entities.Dokter{
		ID:        input.ID,
		Nama:      input.Nama,
		Tlp:       input.Tlp,
		JamKerja:  input.JamKerja,
		Spesialis: input.Spesialis,
		Next:      nil,
	}
	dl.DokterList.InsertDokter(temp)
}

// Fungsi search data by Id
func (dl *DokterController) Search(dataDokter *entities.Dokter, ID int) *entities.Dokter {
	return dokter.Search(dataDokter, ID)
}

// Fungsi Update Data
func (dl *DokterController) UpdateDataDokter(dataDokter *entities.Dokter, ID int, nama, Tlp, jamKerja, spesialis string) {
	temp := entities.Dokter{
		Nama:      nama,
		Tlp:       Tlp,
		JamKerja:  jamKerja,
		Spesialis: spesialis,
	}
	dl.DokterList.Update(&temp, ID)
}

func (dl *DokterController) DeleteDataDokter(dataDokter *entities.Dokter, ID int) {
	dl.DokterList.Delete(dataDokter, ID)
}

// Fungsi GetListDokter mengambil daftar dokter dari database dan mengembalikan string berupa tabel
func GetListDokter(dataDokter *entities.Dokter, header []string) string {
	table := dokter.GetListDokter(dataDokter, header)
	if table == " " {
		return "\nDATA KOSONG\n"
	}
	return table
}

func GetListDokterById(dataDokter *entities.Dokter, header []string) string {
	return dokter.GetListDokterById(dataDokter, header)
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
	PasienList pasien.NodePasien
}

func (pc *PasienController) SearchPasien(ID int) *entities.Pasien {
	return pc.PasienList.SearchPasien(ID)
}

func (pc *PasienController) TambahPasien(ID int, Nama string, Tlp string, Tgl string, pnykt string) {
	riwayat := entities.Medis{
		Tanggal: Tgl,
		Peyakit: pnykt,
	}
	temp := &entities.Pasien{
		ID:      ID,
		Nama:    Nama,
		Tlp:     Tlp,
		Riwayat: riwayat,
	}
	pc.PasienList.InsertPasien(temp)
}

func (pc *PasienController) DeletePasien(ID int) bool {
	return pc.PasienList.DeletePasien(ID)
}

func (pc *PasienController) UpdatePasien(ID int, Nama string, Tlp string, Tgl string, pnykt string) bool {
  riwayat := entities.Medis{
		Tanggal: Tgl,
		Peyakit: pnykt,
	}
	update := &entities.Pasien{
		ID:      ID,
		Nama:    Nama,
		Tlp:     Tlp,
		Riwayat: riwayat,
		Next:    nil,
	}
	return pc.PasienList.UpdatePasien(ID, update)
}

func (pc *PasienController) ListPasien() {
	pc.PasienList.ListPasien()
}
