package controllers

import (
	"strings"
	"tugas-uts/entities"
	"tugas-uts/model/doktermodel"
)

// fungsi doktermodel
func InsertDataDokter(g *entities.Dokter, id int, nama, Tlp, jamKerja, spesialis string) {
	spesialis = strings.ToUpper(spesialis)
	nama = strings.ToUpper(nama)
	doktermodel.InsertDataDokter(g, id, nama, Tlp, jamKerja, spesialis)
}

func GetListDataDokter(g *entities.Dokter, header []string) string {
	table := doktermodel.GetListDataDokter(g, header)
	if table == " " {
		return "\nDATA KOSONG\n"
	}
	return table
}

func Search(g *entities.Dokter, ID int) *entities.Dokter {
	return doktermodel.Search(g, ID)
}

func UpdateDataDokter(g *entities.Dokter, nama, Tlp, jamKerja, spesialis string) {
	spesialis = strings.ToUpper(spesialis)
	nama = strings.ToUpper(nama)
	doktermodel.Update(g, nama, Tlp, jamKerja, spesialis)
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
