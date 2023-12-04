package domain

import (
    "time"
)

type Paciente struct {
    Id           int       `json:"id"`
    Nombre       string    `json:"nombre"`
    Apellido     string    `json:"apellido"`
    Domicilio    string    `json:"domicilio"`
    DNI          int    `json:"dni"`
    FechaDeAlta  time.Time `json:"fecha_de_alta"`
}
