package model

import "time"

type Config struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
	CacheConfig    CacheConfig    `mapstructure:"cache"`
	LoggerConfig   LoggerConfig   `mapstructure:"logger"`
	ZipkinConfig   ZipkinConfig   `mapstructure:"zipkin"`
	AuthConfig     AuthConfig     `mapstructure:"auth"`
}

// ServerConfig has only server specific configuration
type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	CloseTimeout time.Duration `mapstructure:"closeTimeout"`
}

// DatabaseConfig has database related configuration.
type DatabaseConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	DbName   string `mapstructure:"dbName"`
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

// CacheConfig has cache related configuration.
type CacheConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// ZipkinConfig has zipkin related configuration.
type ZipkinConfig struct {
	IsEnable bool   `mapstructure:"isEnable"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

// LoggerConfig has logger related configuration.
type LoggerConfig struct {
	LogFilePath string `mapstructure:"file"`
}

// AuthConfig has logger related configuration.
type AuthConfig struct {
	HmacSecret string `mapstructure:"hmacSecret"`
}
