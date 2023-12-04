package turno

var (
	QuertyInsertTurno           = `INSERT INTO turno (fecha_hora, descripcion, odontologo_id, paciente_id) VALUES (?,?,?,?)`
	QueryGetTurnoById           = `SELECT id, fecha_hora, descripcion, odontologo_id, paciente_id FROM turno WHERE id = ?`
	QueryUpdateTurno            = `UPDATE turno SET fecha_hora = ?, descripcion = ?, odontologo_id = ?, paciente_id = ? WHERE id = ?`
	QueryDeleteTurno            = `DELETE FROM turno WHERE id = ?;`
	QueryGetTurnosByPacienteDNI = `SELECT turno.id, turno.fecha_hora, turno.descripcion,
	odontologo.*, paciente.*
	FROM turno
    JOIN paciente ON turno.paciente_id = paciente.id
    JOIN odontologo ON turno.odontologo_id = odontologo.id
    WHERE paciente.dni = ?;`
)
