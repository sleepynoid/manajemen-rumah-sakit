package sustermodels

import (
	"fmt"
	"main/entities"
)

type SusterModel struct {
	DaftarSuster *entities.Suster
}

func (model *SusterModel) TambahSuster(suster *entities.Suster) {
	if model.DaftarSuster == nil {
		model.DaftarSuster = suster
	} else {
		last := model.DaftarSuster
		for last.Next != nil {
			last = last.Next
		}
		last.Next = suster
	}
}

func (m *SusterModel) UpdateSuster(id int, nama string, tlp string, shift string, tugas string) bool {
	current := m.DaftarSuster
	for current != nil {
		if current.ID == id {
			current.Nama = nama
			current.Tlp = tlp
			current.Shift = shift
			current.Tugas = tugas
			return true
		}
		current = current.Next
	}
	return false
}

func (m *SusterModel) HapusSuster(id int) bool {
	if m.DaftarSuster == nil {
		return false
	}

	if m.DaftarSuster.ID == id {
		m.DaftarSuster = m.DaftarSuster.Next
		return true
	}

	prev := m.DaftarSuster
	for prev.Next != nil {
		if prev.Next.ID == id {
			prev.Next = prev.Next.Next
			return true
		}
		prev = prev.Next
	}

	return false
}

func (m *SusterModel) TampilSemuaSuster() {
	if m.DaftarSuster == nil {
		fmt.Println("Tidak ada data suster")
		return
	}

	current := m.DaftarSuster
	fmt.Println("ID\tNama\tTlp\t\tShift\tTugas")
	for current != nil {
		fmt.Printf("%d\t%s\t%s\t%s\t%s\n", current.ID, current.Nama, current.Tlp, current.Shift, current.Tugas)
		current = current.Next
	}
}

func (model *SusterModel) AmbilSuster(id int) *entities.Suster {
	current := model.DaftarSuster
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}
