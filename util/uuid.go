package util

import "github.com/twinj/uuid"

func GetUuid() string {
	return uuid.NewV4().String()
}
