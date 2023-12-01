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