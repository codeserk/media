package repository

import (
	"media/internal/util"
	"media/module/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) UpdateEmailVerify(u *user.Entity) (*user.Entity, error) {
	ctx, cancel := r.db.Context()
	defer cancel()

	return r.base.UpdateByIdUsingQuery(ctx, u.Id, bson.M{
		"$set": bson.M{
			"emailVerify": emailVerifyDocument{
				IsVerified: false,
				Token:      util.UniqueRandomString(),
			},
		},
	})
}
