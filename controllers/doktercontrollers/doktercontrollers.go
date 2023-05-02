package doktercontrollers

import (
	"main/entities"
	"main/models/doktermodels"
	"main/views/dokterviews"
)

type DokterController struct {
	Model *doktermodels.DokterModel
	View  *dokterviews.DokterView
}

func (controller *DokterController) TambahDokter(id int, nama string, tlp string, jamKerja string, spesialis string) {
	dokter := &entities.Dokter{
		ID:        id,
		Nama:      nama,
		Tlp:       tlp,
		JamKerja:  jamKerja,
		Spesialis: spesialis,
	}
	controller.Model.TambahDokter(dokter)
}

func (c *DokterController) UpdateDokter(id int, nama string, tlp string, jamKerja string, spesialis string) {
	status := c.Model.UpdateDokter(id, nama, tlp, jamKerja, spesialis)
	c.View.TampilUpdateDokter(status)
}

func (c *DokterController) HapusDokter(id int) {
	status := c.Model.HapusDokter(id)
	c.View.TampilDeleteDokter(status)
}

func (c *DokterController) TampilSemuaDokter() {
	c.Model.TampilSemuaDokter()
}

func (controller *DokterController) AmbilDokter(id int) {
	dokter := controller.Model.AmbilDokter(id)
	controller.View.TampilkanDokter(dokter)
}
