package dokterviews

import (
	"fmt"
	"main/entities"
)

type DokterView struct{}

func (view *DokterView) TampilkanDokter(dokter *entities.Dokter) {
	if dokter == nil {
		fmt.Println("Dokter tidak ditemukan")
	} else {
		fmt.Printf("ID: %d\n", dokter.ID)
		fmt.Printf("Nama: %s\n", dokter.Nama)
		fmt.Printf("Telepon: %s\n", dokter.Tlp)
		fmt.Printf("Jam Kerja: %s\n", dokter.JamKerja)
		fmt.Printf("Spesialis: %s\n", dokter.Spesialis)
	}
}

func (v *DokterView) TampilUpdateDokter(status bool) {
	if status {
		fmt.Println("Data dokter berhasil diupdate")
	} else {
		fmt.Println("Data dokter gagal diupdate")
	}
}

func (v *DokterView) TampilDeleteDokter(status bool) {
	if status {
		fmt.Println("Data Dokter berhasil Delete")
	} else {
		fmt.Println("Data Dokter gagal Delete")
	}
}
