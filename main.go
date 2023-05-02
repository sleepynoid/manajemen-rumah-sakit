package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"main/controllers/doktercontrollers"
	"main/models/doktermodels"
	"main/views/dokterviews"

	"main/controllers/pasiencontrollers"
	"main/models/pasienmodels"
	"main/views"
	"main/views/pasienviews"

	"main/controllers/sustercontrollers"
	"main/models/sustermodels"
	"main/views/susterviews"
)

func clear() {
	// hapus isi layar konsol
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	dokterModel := &doktermodels.DokterModel{}
	dokterView := &dokterviews.DokterView{}
	dokterController := &doktercontrollers.DokterController{
		Model: dokterModel,
		View:  dokterView,
	}

	susterModel := &sustermodels.SusterModel{}
	susterView := &susterviews.SusterView{}
	susterController := &sustercontrollers.SusterController{
		Model: susterModel,
		View:  susterView,
	}

	pasienModel := &pasienmodels.PasienModel{}
	pasienView := &pasienviews.PasienView{}
	pasienController := &pasiencontrollers.PasienController{
		Model: pasienModel,
		View:  pasienView,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		clear()
		fmt.Println("Selamat datang di program Manajemen Rumah Sakit")
		fmt.Println("-----------------------------------------------")
		views.MenuUtama()
		menu, _ := reader.ReadString('\n')
		menu = strings.TrimSpace(menu)

		switch menu {
		case "1":
			clear()
			for {
				views.SubMenu()

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "1" {
					fmt.Println("Masukkan informasi dokter:")
					fmt.Print("ID: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					fmt.Print("Nama: ")
					nama, _ := reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Telepon: ")
					tlp, _ := reader.ReadString('\n')
					tlp = strings.TrimSpace(tlp)

					fmt.Print("Jam Kerja: ")
					jamKerja, _ := reader.ReadString('\n')
					jamKerja = strings.TrimSpace(jamKerja)

					fmt.Print("Spesialis: ")
					spesialis, _ := reader.ReadString('\n')
					spesialis = strings.TrimSpace(spesialis)

					dokterController.TambahDokter(id, nama, tlp, jamKerja, spesialis)
					fmt.Println("Dokter berhasil ditambahkan.")
					time.Sleep(2 * time.Second)

				} else if input == "2" {
					fmt.Println("UPDATE")
					fmt.Print("ID: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					fmt.Print("Nama: ")
					nama, _ := reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Telepon: ")
					tlp, _ := reader.ReadString('\n')
					tlp = strings.TrimSpace(tlp)

					fmt.Print("Jam Kerja: ")
					jamKerja, _ := reader.ReadString('\n')
					jamKerja = strings.TrimSpace(jamKerja)

					fmt.Print("Spesialis: ")
					spesialis, _ := reader.ReadString('\n')
					spesialis = strings.TrimSpace(spesialis)

					dokterController.UpdateDokter(id, nama, tlp, jamKerja, spesialis)

				} else if input == "3" {
					fmt.Println("DELETE")
					fmt.Print("Masukkan ID dokter yang ini di delete: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					dokterController.HapusDokter(id)

				} else if input == "4" {
					fmt.Println("VIEW ALL")
					dokterController.TampilSemuaDokter()

				} else if input == "5" {
					fmt.Print("Masukkan ID dokter: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					dokterController.AmbilDokter(id)

				} else if input == "6" {
					break

				} else {
					fmt.Println("Pilihan tidak valis")
				}
				fmt.Println("-----------------------------------------")
			}

		case "2":
			clear()
			for {
				views.SubMenu()

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "1" {
					fmt.Println("Masukkan informasi suster:")
					fmt.Print("ID: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					fmt.Print("Nama: ")
					nama, _ := reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Telepon: ")
					tlp, _ := reader.ReadString('\n')
					tlp = strings.TrimSpace(tlp)

					fmt.Print("Shift: ")
					shift, _ := reader.ReadString('\n')
					shift = strings.TrimSpace(shift)

					fmt.Print("Tugas: ")
					tugas, _ := reader.ReadString('\n')
					tugas = strings.TrimSpace(tugas)

					susterController.TambahSuster(id, nama, tlp, shift, tugas)
					fmt.Println("Suster berhasil ditambahkan.")

				} else if input == "2" {
					fmt.Print("ID: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					fmt.Print("Nama: ")
					nama, _ := reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Telepon: ")
					tlp, _ := reader.ReadString('\n')
					tlp = strings.TrimSpace(tlp)

					fmt.Print("Shift: ")
					shift, _ := reader.ReadString('\n')
					shift = strings.TrimSpace(shift)

					fmt.Print("Tugas: ")
					tugas, _ := reader.ReadString('\n')
					tugas = strings.TrimSpace(tugas)

					susterController.UpdateSuster(id, nama, tlp, shift, tugas)

				} else if input == "3" {
					fmt.Println("DELETE")
					fmt.Print("Masukan ID suster yang ingin di delete: ")
					var id int
					fmt.Scan(&id)
					susterController.HapusSuster(id)

				} else if input == "4" {
					fmt.Println("VIEW ALL")
					susterController.TampilSemuaSuster()

				} else if input == "5" {
					fmt.Print("Masukkan ID suster: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					susterController.AmbilSuster(id)

				} else if input == "6" {
					break
				} else {
					fmt.Println("Pilihan tidak valid")
				}
				fmt.Println("-----------------------------------------")
			}

		case "3":
			clear()
			views.SubMenu()

			for {
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "1" {
					fmt.Print("ID: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					fmt.Print("Nama: ")
					nama, _ := reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Telepon: ")
					tlp, _ := reader.ReadString('\n')
					tlp = strings.TrimSpace(tlp)

					fmt.Print("Kondisi Penyakit: ")
					kondisiPenyakit, _ := reader.ReadString('\n')
					kondisiPenyakit = strings.TrimSpace(kondisiPenyakit)

					fmt.Print("Tanggal: ")
					tglKondisi, _ := reader.ReadString('\n')
					tglKondisi = strings.TrimSpace(tglKondisi)

					fmt.Print("Riwayat Penyakit: ")
					riwayatPenyakit, _ := reader.ReadString('\n')
					riwayatPenyakit = strings.TrimSpace(riwayatPenyakit)

					fmt.Print("Tanggal: ")
					tglRiwayat, _ := reader.ReadString('\n')
					tglRiwayat = strings.TrimSpace(tglRiwayat)

					pasienController.TambahPasien(id, nama, tlp, tglKondisi, kondisiPenyakit, tglRiwayat, riwayatPenyakit)

				} else if input == "2" {
					fmt.Print("ID: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					fmt.Print("Nama: ")
					nama, _ := reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Telepon: ")
					tlp, _ := reader.ReadString('\n')
					tlp = strings.TrimSpace(tlp)

					fmt.Print("Kondisi Penyakit: ")
					kondisiPenyakit, _ := reader.ReadString('\n')
					kondisiPenyakit = strings.TrimSpace(kondisiPenyakit)

					fmt.Print("Tanggal: ")
					tglKondisi, _ := reader.ReadString('\n')
					tglKondisi = strings.TrimSpace(tglKondisi)

					fmt.Print("Riwayat Penyakit: ")
					riwayatPenyakit, _ := reader.ReadString('\n')
					riwayatPenyakit = strings.TrimSpace(riwayatPenyakit)

					fmt.Print("Tanggal: ")
					tglRiwayat, _ := reader.ReadString('\n')
					tglRiwayat = strings.TrimSpace(tglRiwayat)

					pasienController.UpdatePasien(id, nama, tlp, tglKondisi, kondisiPenyakit, tglRiwayat, riwayatPenyakit)

				} else if input == "3" {
					fmt.Println("DELETE")
					fmt.Print("Masukan ID pasien yang ingin di delete: ")
					var id int
					fmt.Scan(&id)

					pasienController.HapusPasien(id)

				} else if input == "4" {
					fmt.Println("VIEW ALL")
					susterController.TampilSemuaSuster()

				} else if input == "5" {
					fmt.Print("Masukkan ID pasien: ")
					idInput, _ := reader.ReadString('\n')
					idInput = strings.TrimSpace(idInput)
					id, _ := strconv.Atoi(idInput)

					pasienController.AmbilPasien(id)

				} else if input == "6" {
					break
				} else {
					fmt.Println("Pilihan tidak valid")
				}
				fmt.Println("-----------------------------------------")

			}

		case "4":
			return
		}
	}
}
