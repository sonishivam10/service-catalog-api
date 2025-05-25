package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/sonishivam10/service-catalog/internal/model"
	"github.com/sonishivam10/service-catalog/internal/repository"
	"github.com/sonishivam10/service-catalog/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetServiceByID(ctx context.Context, id uuid.UUID) (model.Service, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.Service), args.Error(1)
}

func (m *mockRepo) ListServices(ctx context.Context, params repository.FilterParams) ([]model.Service, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]model.Service), args.Error(1)
}

func (m *mockRepo) GetVersionsByServiceID(ctx context.Context, serviceID uuid.UUID) ([]model.Version, error) {
	args := m.Called(ctx, serviceID)
	return args.Get(0).([]model.Version), args.Error(1)
}

func TestListServices_Success(t *testing.T) {
	mockRepo := new(mockRepo)
	usecase := service.NewServiceUsecase(mockRepo)

	expected := []model.Service{
		{ID: uuid.New(), Name: "Test Service", Description: "A test service"},
	}

	params := repository.FilterParams{
		Page:     1,
		PageSize: 10,
	}

	mockRepo.On("ListServices", mock.Anything, params).Return(expected, nil)

	result, err := usecase.ListServices(context.Background(), params)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetService_NotFound(t *testing.T) {
	mockRepo := new(mockRepo)
	usecase := service.NewServiceUsecase(mockRepo)

	id := uuid.New()
	mockRepo.On("GetServiceByID", mock.Anything, id).Return(model.Service{}, errors.New("not found"))

	result, err := usecase.GetService(context.Background(), id)

	assert.Error(t, err)
	assert.Equal(t, model.Service{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetVersions_Success(t *testing.T) {
	mockRepo := new(mockRepo)
	usecase := service.NewServiceUsecase(mockRepo)

	id := uuid.New()
	versions := []model.Version{
		{ID: uuid.New(), ServiceID: id, Version: "v1.0.0"},
		{ID: uuid.New(), ServiceID: id, Version: "v1.1.0"},
	}

	mockRepo.On("GetVersionsByServiceID", mock.Anything, id).Return(versions, nil)

	result, err := usecase.GetVersionsByServiceID(context.Background(), id)

	assert.NoError(t, err)
	assert.Equal(t, versions, result)
	mockRepo.AssertExpectations(t)
}
