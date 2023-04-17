package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"tugas-uts/entities"
	"tugas-uts/views"
	"tugas-uts/views/dokterview"
)

func clear() {
	// hapus isi layar konsol
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	dataDokter := entities.Dokter{}
	// dataSuster := entities.Suster{}
	var menuUtama int = 0

	for menuUtama != 4 {
		clear()
		views.MenuUtama()
		fmt.Scan(&menuUtama)
		scanner.Scan()

		switch menuUtama {
		case 1:
			var input int = 0
			clear()
			for input != 6 {
				views.SubMenu()
				fmt.Scan(&input)
				scanner.Scan()

				switch input {
				case 1:
					dokterview.MenuInsertDataDokter(&dataDokter)
				case 2:
					dokterview.MenuUpdateDokter(&dataDokter)
				case 3:
					dokterview.MenuDeleteDokter(&dataDokter)
				case 4:
					dokterview.MenuViewAll(&dataDokter)
				case 5:
					dokterview.MenuViewById(&dataDokter)
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
				}
			}
		case 2:
			var input int
			for input != 6 {
				views.SubMenu()
				fmt.Scan(&input)
				scanner.Scan()

				switch input {
				case 1:
					// mhsview.MenuInsertSstr(&dataSuster)
				case 2:
					// mhsview.MenuUpdatetSstr(&dataSuster)
				case 3:
					// mhsview.MenuDeletetSstr(&dataSuster)
				case 4:
					// mhsview.MenuViewAllSstr(&dataSuster)
				case 5:
					// mhsview.MenuViewByNpmSstr(&dataSuster)
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
				}
			}
		case 3:
			var input int
			for input != 6 {
				views.SubMenu()
				fmt.Scan(&input)
				scanner.Scan()

				switch input {
				case 1:
					// mhsview.MenuInsertMhs(&dataSuster)
				case 2:
					// mhsview.MenuUpdatetMhs(&dataSuster)
				case 3:
					// mhsview.MenuDeletetMhs(&dataSuster)
				case 4:
					// mhsview.MenuViewAllMhs(&dataSuster)
				case 5:
					// mhsview.MenuViewByNpmMhs(&dataSuster)
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
				}
			}
		default:
			if menuUtama == 3 {
				break
			}
			fmt.Println("pilihan tidak valid")
		}
	}

}
