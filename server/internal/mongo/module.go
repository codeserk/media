package mongo

import (
	"media/internal/config"
	"net/url"
	"strconv"
)

func New(conf *config.Mongo) (*Connection, error) {
	connection := Connection{conf: conf}
	err := connection.Connect("")
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func FromConnectionString(connectionURL string) (*Connection, error) {
	parsedURL, err := url.Parse(connectionURL)
	if err != nil {
		return nil, err
	}

	host := parsedURL.Hostname()
	portStr := parsedURL.Port()
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, err
	}

	user := parsedURL.User.Username()
	password, _ := parsedURL.User.Password()
	database := parsedURL.Path[1:]

	conf := &config.Mongo{
		Host:     host,
		Port:     uint16(port),
		User:     user,
		Password: password,
		Database: database,
	}

	connection := Connection{conf: conf}
	err = connection.Connect(connectionURL)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}
