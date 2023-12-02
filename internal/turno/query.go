package turno

var (
	QuertyInsertTurno        = `INSERT INTO turno (fecha_hora, descripcion, odontologo_id, paciente_id) VALUES (?,?,?,?)`
	QueryGetTurnoById        = `SELECT id, fecha_hora, descripcion, odontologo_id, paciente_id FROM turno WHERE id = ?`
	QueryUpdateTurno         = `UPDATE turno SET fecha_hora = ?, descripcion = ?, odontologo_id = ?, paciente_id = ? WHERE id = ?`
	QueryDeleteTurno = `DELETE FROM clinica_dental.turno WHERE id = ?;`
	QueryGetTurnosByPacienteDNI = `SELECT turno.*
		FROM turno
    JOIN paciente ON turno.paciente_id = paciente.id
    JOIN odontologo ON turno.odontologo_id = odontologo.id
    WHERE paciente.dni = ?;`
)
