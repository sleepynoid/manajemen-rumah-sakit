package pasien

import (
	"fmt"
	"tugas-uts/entities"
	// ui "github.com/gizak/termui/v3"
	// "github.com/gizak/termui/v3/widgets"
)

type NodePasien struct {
	head *entities.Pasien
	tail *entities.Pasien
}

func (pl *NodePasien) SearchPasien(target int) *entities.Pasien {
	fmt.Println("search")
	temp := pl.head
	for temp != nil {
		fmt.Println("loop")
		if temp.ID == target {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func (pl *NodePasien) InsertPasien(pasien *entities.Pasien) {
	fmt.Println("InsertPasien start!!!")
	if pl.head == nil {
		fmt.Println("first")
		pl.head = pasien
		pl.tail = pasien
	} else {
		fmt.Println("NOTfirst")
		pl.tail.Next = pasien
		pl.tail = pasien
		fmt.Println(pl.head.ID, pl.head.Next.ID)
	}
}

func (pl *NodePasien) DeletePasien(target int) bool {
	fmt.Println("DeletePasien")
	if pl.head == nil {
		fmt.Println("1")
		return false
	}
	if pl.head.ID == target {
		fmt.Println("2")
		fmt.Println(pl.head.ID)
		pl.head = pl.head.Next
		fmt.Println(pl.head.ID)
		return true
	}
	fmt.Println("3")
	temp := pl.head
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

func (pl *NodePasien) UpdatePasien(target int, update *entities.Pasien) bool {
	current := pl.head
	for current != nil {
		if current.ID == target {
			current.Nama = update.Nama
			current.Tlp = update.Tlp
			current.Kondisi = update.Kondisi
			current.Riwayat = update.Riwayat
			return true
		}
	}
	fmt.Println("Not Found")
	return false
}

func (pl *NodePasien) ListPasien() {
	current := pl.head
	// for temp != nil {
	// 	print(temp.ID, "->")
	// 	temp = temp.Next
	// }
	for current != nil {
		fmt.Println("ID Pasien:", current.ID)
		fmt.Println("Nama Pasien:", current.Nama)
		fmt.Println("Nomor Telepon Pasien:", current.Tlp)
		fmt.Println("Kondisi Pasien:", current.Kondisi)
		fmt.Println("Riwayat Pasien:", current.Riwayat)
		fmt.Println()
		current = current.Next
	}
}
