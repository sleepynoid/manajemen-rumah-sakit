package pasienmodels

import (
	"fmt"
	"main/entities"
)

type PasienModel struct {
	DaftarPasien *entities.Pasien
}

func (model *PasienModel) TambahPasien(pasien *entities.Pasien) {
	if model.DaftarPasien == nil {
		model.DaftarPasien = pasien
	} else {
		last := model.DaftarPasien
		for last.Next != nil {
			last = last.Next
		}
		last.Next = pasien
	}
}

func (m *PasienModel) UpdatePasien(id int, nama, tlp, tglKondisi, kondisiPenyakit, tglRiwayat, riwayatPenyakit string) bool {
	current := m.DaftarPasien
	for current != nil {
		if current.ID == id {
			current.Nama = nama
			current.Tlp = tlp
			current.Kondisi.Tanggal = tglKondisi
			current.Kondisi.Peyakit = kondisiPenyakit
			current.Riwayat.Tanggal = tglRiwayat
			current.Riwayat.Peyakit = riwayatPenyakit
			return true
		}
		current = current.Next
	}
	return false
}

func (m *PasienModel) HapusPasien(id int) bool {
	if m.DaftarPasien == nil {
		return false
	}

	if m.DaftarPasien.ID == id {
		m.DaftarPasien = m.DaftarPasien.Next
		return true
	}

	prev := m.DaftarPasien
	for prev.Next != nil {
		if prev.Next.ID == id {
			prev.Next = prev.Next.Next
			return true
		}
		prev = prev.Next
	}

	return false
}

func (m *PasienModel) TampilSemuaPasien() {
	if m.DaftarPasien == nil {
		fmt.Println("Tidak ada data pasien")
		return
	}

	current := m.DaftarPasien
	fmt.Println("ID\tNama\t\tTlp\t\tShift\tTugas")
	for current != nil {
		fmt.Printf("%d\t%s\t\t%s\n", current.ID, current.Nama, current.Tlp)
		current = current.Next
	}
}

func (model *PasienModel) AmbilPasien(id int) *entities.Pasien {
	current := model.DaftarPasien
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}
