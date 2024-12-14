package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Model struct {
	UUID      uuid.UUID `json:"uuid" bun:",pk,type:uuid"`
	CreatedAt time.Time `json:"created_at" bun:",default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:",default:current_timestamp"`
	VisitedAt time.Time `json:"visited_at" bun:",default:current_timestamp"`
	DeletedAt time.Time `json:"deleted_at" bun:",soft_delete"`
}

func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.SelectQuery:
		m.VisitedAt = time.Now()
	case *bun.InsertQuery:
		m.UUID = uuid.New()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
