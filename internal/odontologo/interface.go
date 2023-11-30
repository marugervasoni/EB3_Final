package odontologo

import (
	"context"

	"github.com/jum8/EBE3_Final.git/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) (*[]domain.Odontologo, error)
	GetById(ctx context.Context, id int) (*domain.Odontologo, error)
	GetByMatricula(ctx context.Context, matricula string) (*domain.Odontologo, error)
	Create(ctx context.Context, odontologo domain.Odontologo) (*domain.Odontologo, error)
	Update(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error)
}

type Service interface {
	GetAll(ctx context.Context) (*[]domain.Odontologo, error)
	GetById(ctx context.Context, id int) (*domain.Odontologo, error)
	GetByMatricula(ctx context.Context, matricula string) (*domain.Odontologo, error)
	Create(ctx context.Context, odontologo domain.Odontologo) (*domain.Odontologo, error)
	Update(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error)
}