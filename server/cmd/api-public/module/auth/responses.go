package auth

import "media/cmd/api-public/module/user"

// type UserResponse struct {
// 	Id    string `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`

// 	CreatedAt time.Time `json:"createdAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// }

// func ToUserResponse(u *user.Entity) Response {
// 	return Response{
// 		Id:        u.Id,
// 		Name:      u.Name,
// 		Email:     u.Email,
// 		CreatedAt: u.CreatedAt,
// 		UpdatedAt: u.UpdatedAt,
// 	}
// }

type loginResponse struct {
	User  user.Response `json:"user" swagger:"required"`
	Token string        `json:"token" swagger:"required"`
}
