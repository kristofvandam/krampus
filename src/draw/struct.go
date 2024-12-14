package draw

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"go.hagfi.sh/krampus/database"
	"go.hagfi.sh/krampus/member"
)

func Migrate(db *bun.DB) error {
	if _, err := db.NewCreateTable().Model((*Drawing)(nil)).Exec(context.Background()); err != nil {
		return err
	}

	if _, err := db.NewCreateTable().Model((*DrawingGroup)(nil)).Exec(context.Background()); err != nil {
		return err
	}

	return nil
}

type Drawing struct {
	bun.BaseModel `bun:"table:drawings"`
	database.Model
	Name      string                  `json:"name"`
	Slug      string                  `json:"slug"`
	Members   []member.MemberRedacted `json:"members" bun:",rel:has-many,join:uuid=drawing_uuid"`
	Groups    []DrawingGroup          `json:"groups" bun:",rel:has-many,join:uuid=drawing_uuid"`
	Config    DrawingConfig           `json:"config" bun:",type:jsonb"`
}

type DrawingConfig struct {
	Chained             bool `json:"chained"`
	PreventGroupMembers bool `json:"prevent_group_members"`
	Price               int  `json:"price"`
	RenewLink           bool `json:"renew_link"`
}

type DrawingGroup struct {
	bun.BaseModel `bun:"table:drawing_groups"`
	database.Model
	DrawingUUID uuid.UUID       `json:"drawing_uuid" bun:",type:uuid,notnull"`
	Name        string          `json:"name"`
	Members     []member.Member `json:"members" bun:",rel:has-many,join:uuid=uuid"`
}

