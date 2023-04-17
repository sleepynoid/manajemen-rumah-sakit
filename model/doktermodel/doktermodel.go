package doktermodel

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"tugas-uts/entities"
)

type NodeDokter struct {
	head *entities.Dokter
	tail *entities.Dokter
}

func (dk *NodeDokter) InsertDokter(dokter *entities.Dokter) {
	if dk.head == nil {
		fmt.Println("first")
		dk.head = dokter
		dk.tail = dokter
	} else {
		fmt.Println("NOTfirst")
		dk.tail.Next = dokter
		dk.tail = dokter
		fmt.Println(dk.head.ID, dk.head.Next.ID)
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

func (dl *NodeDokter) Update(update *entities.Dokter, target int) bool {
	current := dl.head
	for current != nil {
		if current.ID == target {
			current.Nama = update.Nama
			current.Tlp = update.Tlp
			current.JamKerja = update.JamKerja
			current.Spesialis = update.Spesialis
			return true
		}
	}
	return false
}

func (dl *NodeDokter) Delete(dataDokter *entities.Dokter, target int) bool {
	if dl.head == nil {
		fmt.Println("1")
		return false
	}
	if dl.head.ID == target {
		fmt.Println("2")
		fmt.Println(dl.head.ID)
		dl.head = dl.head.Next
		fmt.Println(dl.head.ID)
		return true
	}
	fmt.Println("3")
	temp := dl.head
	for temp.Next != nil {
		fmt.Println("loop")
		if temp.Next.ID == target {
			fmt.Println("found")
			temp.Next = temp.Next.Next
			return true
		}
		temp = temp.Next
	}
	return false
}
