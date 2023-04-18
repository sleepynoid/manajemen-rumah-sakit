package sustercontrollers

import (
	"main/entities"
	"main/models/sustermodels"
	"main/views/susterviews"
)

type SusterController struct {
	Model *sustermodels.SusterModel
	View  *susterviews.SusterView
}

func (controller *SusterController) TambahSuster(id int, nama string, tlp string, shift string, tugas string) {
	suster := &entities.Suster{
		ID:    id,
		Nama:  nama,
		Tlp:   tlp,
		Shift: shift,
		Tugas: tugas,
	}
	controller.Model.TambahSuster(suster)
}

func (c *SusterController) UpdateSuster(id int, nama string, tlp string, shift string, tugas string) {
	status := c.Model.UpdateSuster(id, nama, tlp, shift, tugas)
	c.View.TampilUpdateSuster(status)
}

func (c *SusterController) HapusSuster(id int) {
	status := c.Model.HapusSuster(id)
	c.View.TampilDeleteSuster(status)
}

func (c *SusterController) TampilSemuaSuster() {
	c.Model.TampilSemuaSuster()
}

func (controller *SusterController) AmbilSuster(id int) {
	suster := controller.Model.AmbilSuster(id)
	controller.View.TampilkanSuster(suster)
}
