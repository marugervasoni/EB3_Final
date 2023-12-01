package domain

import "time"

type Turno struct {
	Id           int       `json:"id"`
	FechaHora    time.Time `json:"fecha_hora"`
	Descripcion  string    `json:"descripcion"`
	OdontologoId int       `json:"odontologo_id"`
	PacienteId   int       `json:"paciente_id"`
}
