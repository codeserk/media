package email

type SMTP struct {
	Host     string `json:"host" validate:"required"`
	Port     int    `json:"port" validate:"required,min=1,max=65535"`
	User     string `json:"user" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Product struct {
	Name string `json:"name" validate:"required"`
	Link string `json:"link" validate:"url"`
	Logo string `json:"logo" validate:"url"`
}

type Config struct {
	SMTP      `json:"smtp" validate:"required"`
	Product   `json:"product" validate:"required"`
	FromEmail string `json:"fromEmail" validate:"required,email"`
	FromName  string `json:"fromName" validate:"required"`
}
