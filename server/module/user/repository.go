package user

type FilterParams struct {
	Name      string
	NameExact string

	Email      string
	EmailExact string
}

type CreateParams struct {
	Name        string
	Email       string
	EmailVerify EmailVerify
	Password    string

	Role Role
}

type UpdateParams struct {
	Name string
}

type Repository interface {
	GetById(id string) (*Entity, error)
	GetByEmail(email string) (*Entity, error)

	Create(params CreateParams) (*Entity, error)

	Update(id string, params UpdateParams) (*Entity, error)
	UpdateEmailVerify(u *Entity) (*Entity, error)
	VerifyEmail(id string) (*Entity, error)

	Archive(id string) (*Entity, error)
	Delete(id string) error
}
