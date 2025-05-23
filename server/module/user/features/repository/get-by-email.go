package repository

import (
	"media/module/user"
)

func (r *repository) GetByEmail(email string) (*entity, error) {
	ctx, cancel := r.db.Context()
	defer cancel()

	return r.base.GetOne(ctx, user.FilterParams{EmailExact: email})
}
