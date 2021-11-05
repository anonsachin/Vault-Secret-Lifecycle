package config

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
)

func GetConfig() *vault.Config{
	fmt.Println("Gettig config")
	return vault.DefaultConfig()
}