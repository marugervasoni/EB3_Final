package paciente

import (
	"context"
	"errors"
	"github.com/jum8/EBE3_Final.git/internal/domain"
	"database/sql"
	"time"
)

var (
	ErrInvalidAttributes  = errors.New("invalid attributes")
	ErrDuplicateDNI       = errors.New("this DNI already exists")
)

type service struct {
	repository Repository
}

func NewServicePaciente(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Paciente, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetById(ctx context.Context, id int) (*domain.Paciente, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) Create(ctx context.Context, paciente domain.Paciente) (*domain.Paciente, error) {
	if err := s.validatePacienteAttributes(paciente); err != nil {
		return nil, err
	}

	if err := s.checkDuplicateDNI(ctx, paciente.DNI); err != nil {
		return nil, err
	}

	if paciente.FechaDeAlta.IsZero() {
        paciente.FechaDeAlta = time.Now()
    }

	return s.repository.Create(ctx, paciente)
}

func (s *service) Update(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	if err := s.validatePacienteAttributes(paciente); err != nil {
		return nil, err
	}

	if err := s.checkDuplicateDNIOnUpdate(ctx, id, paciente.DNI); err != nil {
		return nil, err
	}

	if paciente.FechaDeAlta.IsZero() {
        paciente.FechaDeAlta = time.Now()
    }

	return s.repository.Update(ctx, id, paciente)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) Patch(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	return s.repository.Patch(ctx, id, paciente)
}

func (s *service) validatePacienteAttributes(paciente domain.Paciente) error {
	if paciente.Nombre == "" || paciente.Apellido == "" || paciente.DNI == "" {
		return ErrInvalidAttributes
	}
	return nil
}

func (s *service) checkDuplicateDNI(ctx context.Context, dni string) error {
	paciente, err := s.repository.GetByDNI(ctx, dni)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return err
	}
	if paciente != nil {
		return ErrDuplicateDNI
	}
	return nil
}

func (s *service) checkDuplicateDNIOnUpdate(ctx context.Context, id int, dni string) error {
	paciente, err := s.repository.GetByDNI(ctx, dni)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return err
	}
	if paciente != nil && paciente.Id != id {
		return ErrDuplicateDNI
	}
	return nil
}

func (r *repository) GetByDNI(ctx context.Context, dni string) (*domain.Paciente, error) {
    const query = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta FROM paciente WHERE dni = ?`
    row := r.db.QueryRowContext(ctx, query, dni)
    var paciente domain.Paciente
    err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaDeAlta)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotFound
        }
        return nil, err
    }
    return &paciente, nil
}


