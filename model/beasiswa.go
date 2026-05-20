package model

type Beasiswa struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Syarat    string `json:"syarat"`
	Deadline  string `json:"deadline"`
}