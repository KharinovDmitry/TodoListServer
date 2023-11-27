package config

import "time"

type Config struct {
	ConnStr  string        `yaml:"conn_str"`
	Port     int           `yaml:"port"`
	TokenTTL time.Duration `yaml:"tokenTTL"`
}
