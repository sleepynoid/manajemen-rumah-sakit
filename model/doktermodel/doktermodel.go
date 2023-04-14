package doktermodel

import (
	"bytes"
	"fmt"
	"os"
	"tugas-uts/entities"

	"github.com/olekukonko/tablewriter"
)

func InsertDataDokter(g *entities.Dokter, id int, nama string, Tlp string, jamKerja string, spesialis string) {

	newDokter := entities.Dokter{
		ID:        id,
		Nama:      nama,
		Tlp:       Tlp,
		JamKerja:  jamKerja,
		Spesialis: spesialis,
	}

	temp := g
	if temp.Next == nil {
		temp.Next = &newDokter
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newDokter
	}
}

func GetListDataDokter(dataDokter *entities.Dokter, header []string) string {
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

func Search(g *entities.Dokter, ID int) *entities.Dokter {
	for g != nil {
		if g.ID == ID {
			return g
		}
		g = g.Next
	}
	return nil
}

func Update(g *entities.Dokter, nama, Tlp, jamKerja, spesialis string) {
	temp := g
	temp.Nama = nama
	temp.Tlp = Tlp
	temp.JamKerja = jamKerja
	temp.Spesialis = spesialis
}

/*

func Delete(g *entities.Dokter, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	}
	// mengecek apakah data yang di hapus data pertama
	if temp.Dok.Nip == nip {
		*temp = *temp.Next
		return
	}

	for temp.Next != nil {
		if temp.Next.Dok.Nip == nip {
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}

	fmt.Println("Delete gagal, ID tidak tersedia")
}

// Fungsi untuk mendapatkan daftar NIP dan Nama dok pada suatu gerbong
func GetList(g *entities.Dokter) {
	// Inisialisasi tabel dan header kolom
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NIP", "Nama"})

	// Iterasi melalui setiap elemen GerbongMhs dan menambahkan data ke dalam tabel
	current := g
	for current.Next != nil {
		dsn := current.Next.Dok
		row := []string{dsn.Nip, dsn.Nama}
		table.Append(row)
		current = current.Next
	}

	// Render tabel ke terminal
	table.Render()
}

func GetListByNip(g *entities.Dokter, nip string) {
	current := g

	for current.Next != nil {
		if current.Next.Dok.Nip == nip {
			dsn := current.Next.Dok
			// Inisialisasi tabel dan header kolom
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"NIP", "Nama"})
			row := []string{dsn.Nip, dsn.Nama}
			table.Append(row)
			// Render tabel ke terminal
			table.Render()
			return
		}
		current = current.Next
	}
	fmt.Printf("Data dengan NIP %s tidak tersedia", nip)
}
*/
