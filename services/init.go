package services

import (
	"fmt"

	"github.com/huynhtrongtien/dove/pkg/jwt"
	"github.com/spf13/viper"
)

var (
	JWTMaker *jwt.JWTMaker
)

func InitServices() error {
	JWTMaker = &jwt.JWTMaker{
		SecretKey: viper.GetString("jwt.key"),
		Lifetime:  4320,
		Issuer:    "",
	}

	if JWTMaker.SecretKey == "" {
		return fmt.Errorf("jwt key is empty")
	}

	return nil
}
