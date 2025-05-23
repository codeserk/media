package util

import (
	"github.com/google/uuid"
)

func UniqueRandomString() string {
	return uuid.New().String()
}

func LargeUniqueRandomString() string {
	return uuid.New().String() + uuid.New().String()
}
