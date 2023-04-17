package doktermodels

import (
	"fmt"
	"main/entities"
)

type DokterModel struct {
	DaftarDokter *entities.Dokter
}

func (model *DokterModel) TambahDokter(dokter *entities.Dokter) {
	if model.DaftarDokter == nil {
		model.DaftarDokter = dokter
	} else {
		last := model.DaftarDokter
		for last.Next != nil {
			last = last.Next
		}
		last.Next = dokter
	}
}

func (m *DokterModel) UpdateDokter(id int, nama string, tlp string, jamKerja string, spesialis string) bool {
	current := m.DaftarDokter
	for current != nil {
		if current.ID == id {
			current.Nama = nama
			current.Tlp = tlp
			current.JamKerja = jamKerja
			current.Spesialis = spesialis
			return true
		}
		current = current.Next
	}
	return false
}

func (m *DokterModel) HapusDokter(id int) bool {
	if m.DaftarDokter == nil {
		return false
	}

	if m.DaftarDokter.ID == id {
		m.DaftarDokter = m.DaftarDokter.Next
		return true
	}

	prev := m.DaftarDokter
	for prev.Next != nil {
		if prev.Next.ID == id {
			prev.Next = prev.Next.Next
			return true
		}
		prev = prev.Next
	}
	return false
}

func (m *DokterModel) TampilSemuaDokter() {
	if m.DaftarDokter == nil {
		fmt.Println("Tidak ada data dokter")
		return
	}

	current := m.DaftarDokter
	fmt.Println("ID\tNama\t\tTlp\t\tJamKerja\tSpesialis")
	for current != nil {
		fmt.Printf("%d\t%s\t\t%s\t%s\t%s\n", current.ID, current.Nama, current.Tlp, current.JamKerja, current.Spesialis)
		current = current.Next
	}
}

func (model *DokterModel) AmbilDokter(id int) *entities.Dokter {
	current := model.DaftarDokter
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}
