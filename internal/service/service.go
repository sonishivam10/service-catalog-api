package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonishivam10/service-catalog/internal/model"
	"github.com/sonishivam10/service-catalog/internal/repository"
)

type ServiceUsecase interface {
	GetService(ctx context.Context, id uuid.UUID) (model.Service, error)
	ListServices(ctx context.Context, params repository.FilterParams) ([]model.Service, error)
	GetVersionsByServiceID(ctx context.Context, id uuid.UUID) ([]model.Version, error)
}

type serviceUsecase struct {
	repo repository.ServiceRepository
}

func NewServiceUsecase(repo repository.ServiceRepository) ServiceUsecase {
	return &serviceUsecase{repo: repo}
}

func (s *serviceUsecase) GetService(ctx context.Context, id uuid.UUID) (model.Service, error) {
	return s.repo.GetServiceByID(ctx, id)
}

func (s *serviceUsecase) ListServices(ctx context.Context, params repository.FilterParams) ([]model.Service, error) {
	return s.repo.ListServices(ctx, params)
}

func (s *serviceUsecase) GetVersionsByServiceID(ctx context.Context, id uuid.UUID) ([]model.Version, error) {
	return s.repo.GetVersionsByServiceID(ctx, id)
}
