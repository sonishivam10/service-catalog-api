package model

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Version struct {
	ID        uuid.UUID `db:"id" json:"id"`
	ServiceID uuid.UUID `db:"service_id" json:"service_id"`
	Version   string    `db:"version" json:"version"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
