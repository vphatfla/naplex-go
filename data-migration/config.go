package main

import (
	"fmt"
	"os"
)

type Config struct {
	SrcDBConfig *DBConfig
	DstDBConfig *DBConfig
}

type DBConfig struct {
	Username   string
	Password   string
	DBName     string
	Host       string
	Port       string
	ConnString string
}

func (c *DBConfig) ToURLString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func LoadConfig() *Config {
	SrcDBConfig := &DBConfig{
		Username: os.Getenv("SRC_POSTGRES_USER"),
		Password: os.Getenv("SRC_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("SRC_POSTGRES_DB"),
		Host:     os.Getenv("SRC_POSTGRES_HOST"),
		Port:     os.Getenv("SRC_POSTGRES_PORT"),
	}
	DstDBConfig := &DBConfig{
		Username: os.Getenv("DST_POSTGRES_USER"),
		Password: os.Getenv("DST_POSTGRES_PASSWORD"),
		DBName:   os.Getenv("DST_POSTGRES_DB"),
		Host:     os.Getenv("DST_POSTGRES_HOST"),
		Port:     os.Getenv("DST_POSTGRES_PORT"),
	}

	SrcDBConfig.ConnString = SrcDBConfig.ToURLString()
	DstDBConfig.ConnString = DstDBConfig.ToURLString()

	return &Config{
		SrcDBConfig: SrcDBConfig,
		DstDBConfig: DstDBConfig,
	}
}
