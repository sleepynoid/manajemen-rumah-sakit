package susterviews

import (
	"fmt"
	"main/entities"
)

type SusterView struct{}

func (view *SusterView) TampilkanSuster(suster *entities.Suster) {
	if suster == nil {
		fmt.Println("Suster tidak ditemukan")
	} else {
		fmt.Printf("ID: %d\n", suster.ID)
		fmt.Printf("Nama: %s\n", suster.Nama)
		fmt.Printf("Telepon: %s\n", suster.Tlp)
		fmt.Printf("Shift: %s\n", suster.Shift)
		fmt.Printf("Tugas: %s\n", suster.Tugas)
	}
}

func (v *SusterView) TampilUpdateSuster(status bool) {
	if status {
		fmt.Println("Data suster berhasil diupdate")
	} else {
		fmt.Println("Data suster gagal diupdate")
	}
}
func (v *SusterView) TampilDeleteSuster(status bool) {
	if status {
		fmt.Println("Data suster berhasil Delete")
	} else {
		fmt.Println("Data suster gagal Delete")
	}
}
