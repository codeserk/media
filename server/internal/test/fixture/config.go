package fixture

import "media/internal/config"

func JWTConfig() config.JWT {
	return config.JWT{
		Secret: "secret",
		Issuer: "media-test",
	}
}

func DashboardConfig() config.Dashboard {
	return config.Dashboard{
		BaseURL: "https://localhost:1234",
	}
}
