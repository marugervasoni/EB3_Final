package paciente

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jum8/EBE3_Final.git/internal/domain"
	"strings"
	"log"
	"fmt"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)

type repository struct {
	db *sql.DB
}

func NewRepositoryPaciente(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) (*[]domain.Paciente, error) {
	rows, err := r.db.QueryContext(ctx, QueryGetAllPacientes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []domain.Paciente
	for rows.Next() {
		var paciente domain.Paciente
		if err := rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta); err != nil {
			return nil, err
		}
		pacientes = append(pacientes, paciente)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &pacientes, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*domain.Paciente, error) {
	row := r.db.QueryRowContext(ctx, QueryGetPacienteById, id)

	var paciente domain.Paciente
	if err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &paciente, nil
}

func (r *repository) Create(ctx context.Context, paciente domain.Paciente) (*domain.Paciente, error) {
	statement, err := r.db.PrepareContext(ctx, QueryInsertPaciente)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.ExecContext(ctx, paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaDeAlta)
	if err != nil {
		log.Printf("SQL error: %v", err)
		return nil, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, ErrLastInsertedId
	}

	paciente.Id = int(lastId)
	return &paciente, nil
}

func (r *repository) Update(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
    statement, err := r.db.PrepareContext(ctx, QueryUpdatePaciente)
    if err != nil {
        return nil, fmt.Errorf("error al preparar la declaración SQL: %w", err)
    }
    defer statement.Close()

    result, err := statement.ExecContext(ctx, paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaDeAlta, id)
    if err != nil {
        return nil, fmt.Errorf("error al ejecutar la declaración SQL: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return nil, fmt.Errorf("error al verificar las filas afectadas: %w", err)
    }
    if rowsAffected == 0 {
        return nil, fmt.Errorf("ninguna fila afectada, es posible que el ID no exista")
    }

    return &paciente, nil
}


func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, QueryDeletePaciente, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Patch(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	query, args := buildPatchQuery(id, paciente)
	if query == "" {
		return nil, errors.New("no fields to update")
	}

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, id)
}


func buildPatchQuery(id int, paciente domain.Paciente) (string, []interface{}) {
	var parts []string
	var args []interface{}

	if paciente.Nombre != "" {
		parts = append(parts, "nombre = ?")
		args = append(args, paciente.Nombre)
	}
	if paciente.Apellido != "" {
		parts = append(parts, "apellido = ?")
		args = append(args, paciente.Apellido)
	}
	if paciente.Domicilio != "" {
		parts = append(parts, "domicilio = ?")
		args = append(args, paciente.Domicilio)
	}

	if paciente.DNI != 0 {
		parts = append(parts, "dni = ?")
		args = append(args, paciente.DNI)
	}

	if len(parts) == 0 {
		return "", nil 
	}

	query := "UPDATE paciente SET " + strings.Join(parts, ", ") + " WHERE id = ?"
	args = append(args, id)

	return query, args
}




