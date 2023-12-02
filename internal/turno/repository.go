package turno

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"github.com/jum8/EBE3_Final.git/internal/domain"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrPrepareStatement = errors.New("error prepare statement turno")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)

type repository struct {
	db *sql.DB
}

func NewRepositoryTurno(db *sql.DB) Repository {
	return &repository{db: db}
}


//Post
func (r *repository) Create(ctx context.Context, turno domain.Turno) (*domain.Turno, error) {
	statement, err := r.db.Prepare(QuertyInsertTurno)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(turno.FechaHora, turno.Descripcion, turno.OdontologoId, turno.PacienteId)
	if err != nil {
		log.Println("Error executing statement:", err)
		return nil, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last inserted ID:", err)
		return nil, ErrLastInsertedId
	}

	turno.Id = int(lastId)

	return &turno, nil
}


//GetById
func (r *repository) GetById(ctx context.Context, id int) (*domain.Turno, error) {
	row := r.db.QueryRow(QueryGetTurnoById, id)

	var turno domain.Turno
	err := row.Scan(
		&turno.Id,
		&turno.FechaHora,
		&turno.Descripcion,
		&turno.OdontologoId,
		&turno.PacienteId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &turno, nil
}


//Put
func (r *repository) Update(ctx context.Context, id int, turno domain.Turno) (*domain.Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(turno.FechaHora, turno.Descripcion, turno.OdontologoId, turno.PacienteId, id)
	if err != nil {
		log.Println("Error executing statement:", err)
		return nil, ErrExecStatement
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &turno, nil
}

//Patch

func (r *repository) Patch(ctx context.Context,	turno domain.Turno,	id int) (*domain.Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		return nil, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(turno.FechaHora, turno.Descripcion, turno.OdontologoId, turno.PacienteId, id)

	if err != nil {
		return nil, ErrExecStatement
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &turno, nil
}

//Delete
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteTurno, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) GetByDNI(ctx context.Context, dni int) ([]domain.TurnoFull, error) {
	rows, err := r.db.Query(QueryGetTurnosByPacienteDNI, dni)
	if err != nil {
		return []domain.TurnoFull{}, err
	}

	defer rows.Close()


	var turnos []domain.TurnoFull

	for rows.Next() {
		var turno domain.TurnoFull
		var odontologo domain.Odontologo
		var paciente domain.Paciente

		err := rows.Scan(
			&turno.Id,
			&turno.FechaHora,
			&turno.Descripcion,
			&odontologo.Id,
			&odontologo.Nombre,
			&odontologo.Apellido,
			&odontologo.Matricula,
			&paciente.Id,
			&paciente.Nombre,
			&paciente.Apellido,
			&paciente.Domicilio,
			&paciente.DNI,
			&paciente.FechaDeAlta,
		)
		if err != nil {
			return []domain.TurnoFull{}, err
		}

		turno.Odontologo = odontologo
		turno.Paciente = paciente

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []domain.TurnoFull{}, err
	}

	return turnos, nil

}