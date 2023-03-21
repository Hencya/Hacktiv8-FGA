package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Implement interface {
	CheckAbsent() bool
	PrintBio() string
	GenerateBio() []Domain
}

type Domain struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type List struct {
	Nama   []string
	Domain []Domain
	Absent int
}

func (ls List) PrintBio() string {
	d := ls.Domain
	absent := ls.Absent - 1
	return fmt.Sprintf("Nama \t\t\t\t: %v\nAlamat \t\t\t\t: %v\nPekerjaan \t\t\t: %v\nAlasan memilih kelas Golang \t: %v", d[absent].Nama, d[absent].Alamat, d[absent].Pekerjaan, d[absent].Alasan)
}

func (ls List) CheckAbsent() bool {
	return ls.Absent <= len(ls.Domain) && ls.Absent > 0
}

func (ls List) GenerateBio() []Domain {
	listDomain := ls.Nama
	alamat := []string{"Palembang", "Surakarta", "Jakarta", "Bandung"}
	pekerjaan := []string{"Mahasiswa", "Backend Engineer", "Fullstack Developer"}
	alasan := []string{"Menambah skill", "Mengisi Waktu Luang", "Menambah Portofolio"}

	generate := make([]Domain, 0)

	var data Domain
	for _, value := range listDomain {
		data.Nama = value
		data.Alamat = alamat[rand.Intn(len(alamat))]
		data.Pekerjaan = pekerjaan[rand.Intn(len(pekerjaan))]
		data.Alasan = alasan[rand.Intn(len(alasan))]
		generate = append(generate, data)
	}
	return generate
}

func main() {
	fmt.Printf("Hello World\n")

	args, _ := strconv.Atoi(os.Args[1])
	listNama := []string{"Faiz", "Rofi", "Hencya"}
	var generateBiodata Implement = List{Nama: listNama}
	var dataKelas Implement = List{Nama: listNama, Domain: generateBiodata.GenerateBio(), Absent: args}

	if dataKelas.CheckAbsent() {
		fmt.Println(dataKelas.PrintBio())
		os.Exit(0)
	} else {
		err := fmt.Errorf("siswa dengan absen berikut %d tidak ditemukan", args)
		fmt.Println(err)
	}

}
