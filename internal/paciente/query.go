package paciente

const (
	QueryGetAllPacientes        = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta FROM paciente`
	QueryGetPacienteById        = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta FROM paciente WHERE id = ?`
	QueryInsertPaciente         = `INSERT INTO paciente (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?)`
	QueryUpdatePaciente         = `UPDATE paciente SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?`
	QueryDeletePaciente         = `DELETE FROM paciente WHERE id = ?`
	QueryPatchPaciente          = `UPDATE paciente SET {column} = ? WHERE id = ?` 
)
