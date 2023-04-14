package sustermodel

import (
	"fmt"
	"os"
	"tugas-uts/entities"

	"github.com/olekukonko/tablewriter"
)

func InsertDataMahasiswa(gerbong *entities.GerbongMhs, data entities.Mahasiswa) {
	newGerbong := &entities.GerbongMhs{}
	newGerbong.Mahasiswa = data
	if gerbong.Next == nil {
		gerbong.Next = newGerbong
	} else {
		for gerbong.Next != nil {
			gerbong = gerbong.Next
		}
		gerbong.Next = newGerbong
	}
}

func Update(g *entities.GerbongMhs, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	} else {
		for temp.Next != nil {
			if temp.Next.Mahasiswa.Npm == nip {
				fmt.Print("Masukan nama baru : ")
				fmt.Scan(&temp.Next.Mahasiswa.Nama)
				return
			}
			temp = temp.Next
		}
	}

	fmt.Printf("data dengan nip %s  tidak ada\n", nip)
}

func Delete(g *entities.GerbongMhs, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	}

	if temp.Next.Mahasiswa.Npm == nip {
		fmt.Println("data pertama di delete")
		*temp = *temp.Next
		return
	}

	for temp.Next != nil {
		if temp.Next.Mahasiswa.Npm == nip {
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}

	fmt.Printf("Delete gagal, id %s tidak ada", nip)
}

func GetListMahasiswa(g *entities.GerbongMhs) {
	// Inisialisasi tabel dan header kolom
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Jurusan", "NPM", "Nama", "IPK"})

	// Iterasi melalui setiap elemen GerbongMhs dan menambahkan data ke dalam tabel
	current := g
	for current.Next != nil {
		mhs := current.Next.Mahasiswa
		row := []string{mhs.Jurusan, mhs.Npm, mhs.Nama, fmt.Sprintf("%.2f", mhs.Ipk)}
		table.Append(row)
		current = current.Next
	}

	// Render tabel ke terminal
	table.Render()
}

func GetlistMahasiswaByNpm(g *entities.GerbongMhs, npm string) {
	current := g

	for current.Next != nil {
		if current.Next.Mahasiswa.Npm == npm {
			mhs := current.Next.Mahasiswa
			// Inisialisasi tabel dan header kolom
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Jurusan", "NPM", "Nama", "IPK"})
			row := []string{mhs.Jurusan, mhs.Npm, mhs.Nama, fmt.Sprintf("%.2f", mhs.Ipk)}
			table.Append(row)
			// Render tabel ke terminal
			table.Render()
			return
		}
		current = current.Next
	}
	fmt.Printf("data dengan NPM %s tidak ada\n", npm)
}
