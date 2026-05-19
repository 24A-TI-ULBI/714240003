package model

type Pendaftaran struct {

	NamaMahasiswa string `json:"nama"`
	NPM           string `json:"npm"`
	Email         string `json:"email"`
	Semester      string `json:"semester"`
	Prodi         string `json:"prodi"`
	IPK           string `json:"ipk"`

	Beasiswa      string `json:"beasiswa"`
	Status        string `json:"status"`
}