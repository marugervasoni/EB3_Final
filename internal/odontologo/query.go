package odontologo

var (
	QueryGetAllOdontologos        = `SELECT id, apellido, nombre, matricula FROM odontologo`
	QueryGetOdontologoById        = `SELECT id, apellido, nombre, matricula FROM odontologo WHERE id = ?`
	QueryGetOdontologoByMatricula = `SELECT id, apellido, nombre, matricula FROM odontologo WHERE matricula = ?`
	QuertyInsertOdontologo        = `INSERT INTO odontologo (apellido, nombre, matricula) VALUES (?,?,?)`
	QueryUpdateOdontologo         = `UPDATE odontologo SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?`
	QueryDeleteOdontologo         = `DELETE FROM odontologo WHERE id = ?`
)
