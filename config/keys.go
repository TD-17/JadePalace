package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func GetKey(keyType string) []byte {
	var key string
	switch keyType {
	case "public_key":
		key = viper.GetString("public_key")
	case "private_key":
		key = viper.GetString("private_key")
	default:
		log.Panic("key not found in config")
	}
	key = strings.ReplaceAll(key, `\n`, "\n")
	return []byte(key)
}
