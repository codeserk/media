package user

import (
	"media/module/user"
	"time"
)

// structs

type PermissionsResponse struct {
	Read   bool `json:"read" validate:"required"`
	Update bool `json:"update" validate:"required"`
	Delete bool `json:"delete" validate:"required"`
}

type OrganizationResponse struct {
	Id          string              `json:"id" validate:"required"`
	Permissions PermissionsResponse `json:"permissions" validate:"required"`
}

type ProjectResponse struct {
	Id          string              `json:"id" validate:"required"`
	Permissions PermissionsResponse `json:"permissions" validate:"required"`
}

type Response struct {
	Id              string `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	IsEmailVerified bool   `json:"isEmailVerified" validate:"required"`

	Role          user.Role              `json:"role" validate:"required"`
	Organizations []OrganizationResponse `json:"organizations" validate:"required"`
	Projects      []ProjectResponse      `json:"projects" validate:"required"`

	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// transforms

func ToResponse(u *user.Entity) Response {
	return Response{
		Id:              u.Id,
		Name:            u.Name,
		Email:           u.Email,
		IsEmailVerified: u.EmailVerify.IsVerified,

		Role: u.Role,

		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
