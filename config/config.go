// Package config ...
package config

import (
	"attendance/constant"
	"log"

	fileConfig "github.com/lukmanlukmin/go-lib/file"
	kafka "github.com/lukmanlukmin/go-lib/kafka"
)

// Config ...
type Config struct {
	Server           ServerConfig   `yaml:"Server"`
	Security         Security       `yaml:"Security"`
	PostgreSQLConfig PostgresConfig `yaml:"PostgresDBConfig"`
	KafkaConfig      kafka.Config   `yaml:"KafkaConfig"`
	Application      Application    `yaml:"Application"`
}

type (
	// ServerConfig Configuration
	ServerConfig struct {
		HTTPPort string `yaml:"HttpPort"`
	}

	// Security Configuration
	Security struct {
		JWTSecret   string `yaml:"JWTSecret"`
		JWTDuration string `yaml:"JWTDuration"`
	}

	// PostgresConfig Configuration
	PostgresConfig struct {
		DSN           string `yaml:"DSN"`
		RetryInterval int    `yaml:"RetryInterval"`
		MaxIdleConn   int    `yaml:"MaxIdleConn"`
		MaxConn       int    `yaml:"MaxConn"`
	}

	// Application Configuration
	Application struct {
		AllowOverLapPeriod   bool    `yaml:"AllowOverLapPeriod"`
		StartWOrkingHour     int     `yaml:"StartWOrkingHour"`
		EndWorkingHour       int     `yaml:"EndWorkingHour"`
		MaxOvertimeHour      int     `yaml:"MaxOvertimeHour"`
		MultiplyOvertimeRate float64 `yaml:"MultiplyOvertimeRate"`
	}
)

// ReadModuleConfig ...
func ReadModuleConfig(cfg interface{}, filePath string) error {
	if filePath != "" {
		err := fileConfig.ReadConfig(cfg, filePath)
		if err != nil {
			log.Fatalf("failed to read config. %v", err)
		}
		return nil
	}
	err := fileConfig.ReadConfig(cfg, constant.DefaultConfigFile)
	if err != nil {
		log.Fatalf("failed to read config. %v", err)
	}
	return nil
}
