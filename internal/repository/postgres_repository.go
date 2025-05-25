package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sonishivam10/service-catalog/internal/model"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetServiceByID(ctx context.Context, id uuid.UUID) (model.Service, error) {
	var service model.Service
	err := r.db.GetContext(ctx, &service, `
        SELECT id, name, description, created_at
        FROM services
        WHERE id = $1
    `, id)
	return service, err
}

func (r *PostgresRepository) ListServices(ctx context.Context, p FilterParams) ([]model.Service, error) {
	var services []model.Service

	baseQuery := `SELECT id, name, description, created_at FROM services`
	args := []interface{}{}
	where := ""
	i := 1 // SQL placeholder index

	if p.Search != "" {
		where = fmt.Sprintf("WHERE name ILIKE $%d", i)
		args = append(args, "%"+p.Search+"%")
		i++
	}

	sortBy := "created_at"
	if p.SortBy != "" {
		switch p.SortBy {
		case "name", "created_at":
			sortBy = p.SortBy
		}
	}

	order := "DESC"
	if p.SortOrder == "asc" {
		order = "ASC"
	}

	limit := 10
	offset := 0
	if p.PageSize > 0 {
		limit = p.PageSize
	}
	if p.Page > 1 {
		offset = (p.Page - 1) * limit
	}

	query := fmt.Sprintf("%s %s ORDER BY %s %s LIMIT %d OFFSET %d",
		baseQuery, where, sortBy, order, limit, offset)

	err := r.db.SelectContext(ctx, &services, query, args...)
	return services, err
}

func (r *PostgresRepository) GetVersionsByServiceID(ctx context.Context, serviceID uuid.UUID) ([]model.Version, error) {
	var versions []model.Version
	err := r.db.SelectContext(ctx, &versions, `
        SELECT id, service_id, version, created_at
        FROM versions
        WHERE service_id = $1
        ORDER BY created_at DESC
    `, serviceID)
	return versions, err
}
