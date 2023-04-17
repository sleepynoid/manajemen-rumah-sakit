package pasiencontrollers

import (
	"main/entities"
	"main/models/pasienmodels"
	"main/views/pasienviews"
)

type PasienController struct {
	Model *pasienmodels.PasienModel
	View  *pasienviews.PasienView
}

func (controller *PasienController) TambahPasien(id int, nama, tlp, tglKondisi, kondisiPenyakit, tglRiwayat, riwayatPenyakit string) {
	pasien := &entities.Pasien{
		ID:   id,
		Nama: nama,
		Tlp:  tlp,
		Kondisi: entities.Medis{
			Tanggal: tglKondisi,
			Peyakit: kondisiPenyakit,
		},
		Riwayat: entities.Medis{
			Tanggal: tglRiwayat,
			Peyakit: riwayatPenyakit,
		},
	}
	controller.Model.TambahPasien(pasien)
}

func (c *PasienController) UpdatePasien(id int, nama, tlp, tglKondisi, kondisiPenyakit, tglRiwayat, riwayatPenyakit string) {
	status := c.Model.UpdatePasien(id, nama, tlp, tglKondisi, kondisiPenyakit, tglRiwayat, riwayatPenyakit)
	c.View.TampilUpdatePasien(status)
}

func (c *PasienController) HapusPasien(id int) {
	status := c.Model.HapusPasien(id)
	c.View.TampilDeletePasien(status)
}

func (c *PasienController) TampilSemuaPasien() {
	c.Model.TampilSemuaPasien()
}

func (controller *PasienController) AmbilPasien(id int) {
	pasien := controller.Model.AmbilPasien(id)
	controller.View.TampilkanPasien(pasien)
}
