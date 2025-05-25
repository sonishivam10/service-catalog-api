package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonishivam10/service-catalog/internal/model"
)

type FilterParams struct {
	Search    string
	SortBy    string
	SortOrder string
	Page      int
	PageSize  int
}

type ServiceRepository interface {
	GetServiceByID(ctx context.Context, id uuid.UUID) (model.Service, error)
	ListServices(ctx context.Context, params FilterParams) ([]model.Service, error)
	GetVersionsByServiceID(ctx context.Context, serviceID uuid.UUID) ([]model.Version, error)
}
