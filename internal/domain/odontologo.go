package domain

type Odontologo struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Matricula string `json:"matricula"`
}