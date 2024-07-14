package app

import (
	"github.com/core-go/log"
)

type Config struct {
	Credentials string     `mapstructure:"credentials"`
	Log         log.Config `mapstructure:"log"`
}
