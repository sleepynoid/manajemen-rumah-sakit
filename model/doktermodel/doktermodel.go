package doktermodel

import (
	"bytes"
	"fmt"
	"os"
	"tugas-uts/entities"

	"github.com/olekukonko/tablewriter"
)

func InsertDokter(dataDokter *entities.Dokter, id int, nama string, Tlp string, jamKerja string, spesialis string) {

	newDokter := entities.Dokter{
		ID:        id,
		Nama:      nama,
		Tlp:       Tlp,
		JamKerja:  jamKerja,
		Spesialis: spesialis,
	}

	temp := dataDokter
	if temp.Next == nil {
		temp.Next = &newDokter
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newDokter
	}
}

func GetListDokter(dataDokter *entities.Dokter, header []string) string {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	if dataDokter.Next == nil {
		return " "
	}
	for dataDokter.Next != nil {
		dktr := dataDokter.Next
		row := []string{fmt.Sprint(dktr.ID), dktr.Nama, dktr.Tlp, dktr.JamKerja, dktr.Spesialis}
		table.Append(row)
		dataDokter = dataDokter.Next
	}
	// menyimpan tabel ke dalam buffer
	buf := new(bytes.Buffer)
	table.Render()

	// mengembalikan hasil tabel dalam bentuk string
	return buf.String()
}

func GetListDokterById(dataDokter *entities.Dokter, header []string) string {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	dktr := dataDokter
	// menyimpan semua field di row slice
	row := []string{fmt.Sprint(dktr.ID), dktr.Nama, dktr.Tlp, dktr.JamKerja, dktr.Spesialis}
	table.Append(row)

	// menyimpan tabel ke dalam buffer
	buf := new(bytes.Buffer)
	table.Render()

	// mengembalikan hasil tabel dalam bentuk string
	return buf.String()
}

func Search(dataDokter *entities.Dokter, ID int) *entities.Dokter {
	temp := dataDokter
	var found *entities.Dokter = nil
	for temp != nil {
		if ID == temp.ID {
			found = temp
		}
		temp = temp.Next
	}
	return found
}

func Update(dataDokter *entities.Dokter, nama, Tlp, jamKerja, spesialis string) {
	temp := dataDokter
	temp.Nama = nama
	temp.Tlp = Tlp
	temp.JamKerja = jamKerja
	temp.Spesialis = spesialis
}

func Delete(dataDokter *entities.Dokter, ID int) {
	temp := dataDokter

	// mengecek apakah data yang di hapus data pertama
	if temp.ID == ID {
		temp = temp.Next
		return
	}

	for temp.Next != nil {
		if temp.Next.ID == ID {
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}
