package turno

var (
	QuertyInsertTurno        = `INSERT INTO turno (fecha_hora, descripcion, odontologo_id, paciente_id) VALUES (?,?,?,?)`
	QueryGetTurnoById        = `SELECT id, fecha_hora, descripcion, odontologo_id, paciente_id FROM turno WHERE id = ?`
	QueryUpdateTurno         = `UPDATE turno SET fecha_hora = ?, descripcion = ?, odontologo_id = ?, paciente_id = ? WHERE id = ?`
	QueryDeleteTurno = `DELETE FROM clinica_dental.turno WHERE id = ?;`
)
