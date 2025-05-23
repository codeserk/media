package repository

import (
	"media/module/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) VerifyEmail(id string) (*user.Entity, error) {
	ctx, cancel := r.db.Context()
	defer cancel()

	return r.base.UpdateByIdUsingQuery(ctx, id, bson.M{
		"$set": bson.M{"emailVerify.isVerified": true},
	})
}
