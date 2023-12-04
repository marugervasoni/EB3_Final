package paciente

import (
	"context"
	"github.com/jum8/EBE3_Final.git/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) (*[]domain.Paciente, error)
	GetById(ctx context.Context, id int) (*domain.Paciente, error)
	Create(ctx context.Context, paciente domain.Paciente) (*domain.Paciente, error)
	Update(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error)
    GetByDNI(ctx context.Context, dni int) (*domain.Paciente, error)
}


type Service interface {
	GetAll(ctx context.Context) (*[]domain.Paciente, error)
	GetById(ctx context.Context, id int) (*domain.Paciente, error)
	Create(ctx context.Context, paciente domain.Paciente) (*domain.Paciente, error)
	Update(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error)
}
