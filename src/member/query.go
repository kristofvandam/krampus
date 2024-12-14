package member

import (
	"context"

	"github.com/google/uuid"
)

func (ctrl *Controller) GetMember(uuid uuid.UUID) (*Member, error) {
	member := Member{}
	err := ctrl.DB.NewSelect().
		Model(&member).
		Where("uuid = ?", uuid).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return &member, nil
}
