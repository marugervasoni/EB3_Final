package turno

import (
	"context"
	"errors"
	"time"
	"github.com/jum8/EBE3_Final.git/internal/domain"
	Odontologo "github.com/jum8/EBE3_Final.git/internal/odontologo"
	Paciente "github.com/jum8/EBE3_Final.git/internal/paciente"
)

var (
	ErrInvalidAttributes  = errors.New("invalid attributes")
	ErrOdontologoNotFound = errors.New("the specified Odontologo does not exist")
	ErrPacienteNotFound   = errors.New("the specified Paciente does not exist")
)

type OdontologoRepository = Odontologo.Repository
type PacienteRepository = Paciente.Repository

type service struct {
	repository     Repository
	odontologoRepo OdontologoRepository
	pacienteRepo   PacienteRepository
}

func NewServiceTurno(repository Repository, odontologoRepo OdontologoRepository, pacienteRepo PacienteRepository) Service {
	return &service{
		repository:     repository,
		odontologoRepo: odontologoRepo,
		pacienteRepo:   pacienteRepo,
	}
}


// Post
func (s *service) Create(ctx context.Context, turno domain.Turno) (*domain.Turno, error) {
	err := s.validateTurnoAttributes(turno)
	if err != nil {
		return nil, err
	}

	if turno.FechaHora.IsZero() {
        turno.FechaHora = time.Now()
    }

	err = s.validateOdontologo(ctx, turno.OdontologoId)
	if err != nil {
		return nil, err
	}

	err = s.validatePaciente(ctx, turno.PacienteId)
	if err != nil {
		return nil, err
	}

	turnoCreated, err := s.repository.Create(ctx, turno)
	if err != nil {
		return nil, err
	}

	return turnoCreated, nil
}

// GetById
func (s *service) GetById(ctx context.Context, id int) (*domain.Turno, error) {
	turno, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return turno, nil
}

// Put
func (s *service) Update(ctx context.Context, id int, turno domain.Turno) (*domain.Turno, error) {
	_, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	turno.Id = id

	err = s.validateTurnoAttributes(turno)
	if err != nil {
		return nil, err
	}

	if turno.FechaHora.IsZero() {
        turno.FechaHora = time.Now()
    }

	err = s.validateOdontologo(ctx, turno.OdontologoId)
	if err != nil {
		return nil, err
	}

	err = s.validatePaciente(ctx, turno.PacienteId)
	if err != nil {
		return nil, err
	}

	turnoUpdated, err := s.repository.Update(ctx, id, turno)
	if err != nil {
		return nil, err
	}

	return turnoUpdated, nil
}


// validations:
func (s *service) validateTurnoAttributes(turnoUpdate domain.Turno) error {
	if  turnoUpdate.Descripcion == "" || turnoUpdate.OdontologoId == 0 || turnoUpdate.PacienteId == 0 {
		return ErrInvalidAttributes
	}

	return nil
}

func (s *service) validateOdontologo(ctx context.Context, odontologoId int) error {
	odontologoSaved, err := s.odontologoRepo.GetById(ctx, odontologoId)
	if err == nil && odontologoSaved.Id != odontologoId {
		return ErrOdontologoNotFound
	}

	return nil
}

func (s *service) validatePaciente(ctx context.Context, pacienteId int) error {
	pacienteSaved, err := s.pacienteRepo.GetById(ctx, pacienteId)
	if err == nil && pacienteSaved.Id != pacienteId{
		return ErrPacienteNotFound
	}

	return nil
}