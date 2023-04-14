package entities

type Dokter struct {
	ID        int
	Nama      string
	Tlp       string
	JamKerja  string
	Spesialis string
	Next      *Dokter
}
type Suster struct {
	ID    int
	Nama  string
	Tlp   string
	Shift string
	Tugas string
	Next  *Suster
}
type Medis struct {
	Tanggal string
	Peyakit string
}
type Pasien struct {
	ID      int
	Nama    string
	Tlp     string
	Kondisi Medis
	Riwayat Medis
	Next    *Pasien
}
