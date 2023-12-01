package turno

import (
	"context"
	"github.com/jum8/EBE3_Final.git/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, turno domain.Turno) (*domain.Turno, error)
	GetById(ctx context.Context, id int) (*domain.Turno, error)
	Update(ctx context.Context, id int, turno domain.Turno) (*domain.Turno, error)
}

type Service interface {
	Create(ctx context.Context, turno domain.Turno) (*domain.Turno, error)
	GetById(ctx context.Context, id int) (*domain.Turno, error)
	Update(ctx context.Context, id int, turno domain.Turno) (*domain.Turno, error)
}