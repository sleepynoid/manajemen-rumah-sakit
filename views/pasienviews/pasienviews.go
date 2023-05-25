package pasienviews

import (
	"fmt"
	"main/entities"
)

type PasienView struct{}

func (view *PasienView) TampilkanPasien(pasien *entities.Pasien) {
	if pasien == nil {
		fmt.Println("Pasien tidak ditemukan")
	} else {
		fmt.Printf("ID: %d\n", pasien.ID)
		fmt.Printf("Nama: %s\n", pasien.Nama)
		fmt.Printf("Telepon: %s\n", pasien.Tlp)
		fmt.Printf("Kondisi: %s\n", pasien.Kondisi.Peyakit)
		fmt.Printf("Tanggal: %s\n", pasien.Kondisi.Tanggal)
		fmt.Printf("Riwayat: %s\n", pasien.Riwayat.Peyakit)
		fmt.Printf("Tanggal: %s\n", pasien.Riwayat.Tanggal)
	}
}

func (v *PasienView) TampilUpdatePasien(status bool) {
	if status {
		fmt.Println("Data pasien berhasil diupdate")
	} else {
		fmt.Println("Data pasien gagal diupdate")
	}
}

func (v *PasienView) TampilDeletePasien(status bool) {
	if status {
		fmt.Println("Data pasien berhasil Delete")
	} else {
		fmt.Println("Data pasien gagal Delete")
	}
}
