package util

import "github.com/twinj/uuid"

/*
Generates a UUID used for Eureka client id.
*/
func GetUuid() string {
	return uuid.NewV4().String()
}
