package member

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"go.hagfi.sh/krampus/database"
)

func Migrate(db *bun.DB) error {
	if _, err := db.NewCreateTable().Model((*Member)(nil)).Exec(context.Background()); err != nil {
		return err
	}

	if _, err := db.NewCreateTable().Model((*MemberWish)(nil)).Exec(context.Background()); err != nil {
		return err
	}

	return nil
}

// This struct is meant to be use ONLY behind specific routes where
// we know the user knows the UUID of the Member
// @Summary Get a member by UUID
type Member struct {
	bun.BaseModel `bun:"table:drawing_members"`
	database.Model
	DrawingUUID         uuid.UUID    `json:"drawing_uuid" bun:",type:uuid,notnull"`
	DrawingGroupUUID    uuid.UUID    `json:"drawing_group_uuid" bun:",type:uuid,notnull"`
	DrawnMember         *Member      `json:"drawn_member" bun:",rel:has-one,join:uuid=uuid"`
	DrawnMemberRevealed bool         `json:"drawn_member_reveiled" bun:",notnull,default:false"`
	Name                string       `json:"name"`
	Wishes              []MemberWish `json:"wishes" bun:",rel:has-many,join:uuid=drawing_member_uuid"`
}

// MemberRedacted is a redacted version of the Member struct
// It is used in the Drawing struct so that we can give a list of members without their wishes
// and the member that they have drawn
type MemberRedacted struct {
	bun.BaseModel `bun:"table:drawing_members"`
	database.Model
	DrawingUUID      uuid.UUID `json:"-"`
	DrawingGroupUUID uuid.UUID `json:"-"`
	Name             string    `json:"name"`
}

type MemberWish struct {
	bun.BaseModel `bun:"table:drawing_member_wishes"`
	database.Model
	DrawingMemberUUID uuid.UUID `json:"drawing_member_uuid" bun:",type:uuid,notnull"`
	Title             string    `json:"title"`
}
