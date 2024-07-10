package http_parser

import (
	"errors"
	"strings"
)

func ParseTokenFromBearerToken(bearerToken string) (token string, err error) {
	splitToken := strings.Split(bearerToken, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1]), nil
	}
	return "", errors.New("bearer token wrong format")
}
