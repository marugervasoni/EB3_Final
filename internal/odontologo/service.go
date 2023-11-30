package odontologo

import (
	"context"
	"errors"

	"github.com/jum8/EBE3_Final.git/internal/domain"
)

var (
	ErrInvalidAttributes  = errors.New("invalid attributes")
	ErrDuplicateMatricula = errors.New("this matricula already exists")
)

type service struct {
	repository Repository
}

func NewServiceOdontologo(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Odontologo, error) {
	listOdontologos, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return listOdontologos, nil
}

func (s *service) GetById(ctx context.Context, id int) (*domain.Odontologo, error) {
	odontologo, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return odontologo, nil
}

func (s *service) GetByMatricula(ctx context.Context, matricula string) (*domain.Odontologo, error) {
	odontologo, err := s.repository.GetByMatricula(ctx, matricula)
	if err != nil {
		return nil, err
	}

	return odontologo, nil
}

func (s *service) Create(ctx context.Context, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	err := s.validateOdontologoAttributes(odontologo)
	if err != nil {
		return nil, err
	}

	err = s.validateMatricula(ctx, odontologo)
	if err != nil {
		return nil, err
	}

	odontologoCreated, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		return nil, err
	}

	return odontologoCreated, nil
}

func (s *service) Update(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	_, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	odontologo.Id = id

	err = s.validateOdontologoAttributes(odontologo)
	if err != nil {
		return nil, err
	}

	err = s.validateMatricula(ctx, odontologo)
	if err != nil {
		return nil, err
	}

	odontologoUpdated, err := s.repository.Update(ctx, id, odontologo)
	if err != nil {
		return nil, err
	}

	return odontologoUpdated, nil
}

func (s *service) Patch(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	odontologoSaved, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	odontologo.Id = id

	odontologoPatch, err := s.validatePatch(ctx, *odontologoSaved, odontologo)
	if err != nil {
		return nil, err
	}

	odotologoPatched, err := s.repository.Patch(ctx, id, *odontologoPatch)
	if err != nil {
		return nil, err
	}

	return odotologoPatched, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) validateOdontologoAttributes(odontologoUpdate domain.Odontologo) error {
	if odontologoUpdate.Apellido == "" || odontologoUpdate.Nombre == "" || odontologoUpdate.Matricula == "" {
		return ErrInvalidAttributes
	}

	return nil
}

func (s *service) validateMatricula(ctx context.Context, odontologo domain.Odontologo) error {
	odontologoSaved, err := s.GetByMatricula(ctx, odontologo.Matricula)

	if err == nil && odontologoSaved.Id != odontologo.Id {
		return ErrDuplicateMatricula
	}

	return nil
}

func (s *service) validatePatch(ctx context.Context, odontologoSaved, odontologoPatch domain.Odontologo) (*domain.Odontologo, error) {
	if odontologoPatch.Apellido != "" {
		odontologoSaved.Apellido = odontologoPatch.Apellido
	}
	if odontologoPatch.Nombre != "" {
		odontologoSaved.Nombre = odontologoPatch.Nombre
	}
	if odontologoPatch.Matricula != "" {
		err := s.validateMatricula(ctx, odontologoPatch)
		if err != nil {
			return nil, err
		}
		odontologoSaved.Matricula = odontologoPatch.Matricula
	}

	return &odontologoSaved, nil
}
