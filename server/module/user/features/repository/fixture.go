package repository

import "media/module/user"

type Fixture struct {
	repository *repository
}

func NewFixture(repository *repository) *Fixture {
	return &Fixture{repository}
}

func (s *Fixture) Create(params *createParams) (*entity, error) {
	return s.repository.Create(TestCreateParams(params))
}

func TestCreateParams(params *createParams) createParams {
	result := createParams{
		Name:     "user-1",
		Email:    "user-1@mail.com",
		Password: "password",
		Role:     user.RoleUser,
	}
	if params == nil {
		return result
	}

	if params.Name != "" {
		result.Name = params.Name
	}

	return result
}
