package config

type APISignature struct {
	Skip            bool   `json:"skip"`
	DashboardSecret string `json:"dashboardSecret" validate:"required"`
	ConsumerSecret  string `json:"consumerSecret" validate:"required"`
}

type API struct {
	Port      int          `json:"port" validate:"required,number,gte=0,lte=65535"`
	Signature APISignature `json:"signature" validate:"required"`
}

type Mongo struct {
	Host     string `json:"host" validate:"required"`
	Port     uint16 `json:"port" validate:"required,number,gte=0,lte=65535"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Database string `json:"database" validate:"required"`
}

type Redis struct {
	Host     string `json:"host" validate:"required,hostname"`
	Port     uint16 `json:"port" validate:"required,number,gte=0,lte=65535"`
	Password string `json:"password"`
	DB       int    `json:"db" validate:"number,gte=0"`
}

type JWT struct {
	Secret string `json:"secret" validate:"required"`
	Issuer string `json:"issuer" validate:"required"`
}

type Dashboard struct {
	BaseURL string `json:"baseURL" validate:"required,url"`
}
