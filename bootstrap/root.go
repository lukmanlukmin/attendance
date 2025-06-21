// Package bootstrap ...
package bootstrap

import (
	"attendance/bootstrap/repository"
	"attendance/bootstrap/service"
	"attendance/config"

	connDB "github.com/lukmanlukmin/go-lib/database/connection"
	"github.com/lukmanlukmin/go-lib/kafka"
)

// Bootstrap ...
type Bootstrap struct {
	Repository *repository.Repository
	Service    *service.Service
}

// NewBootstrap ...
func NewBootstrap(conf *config.Config) *Bootstrap {
	connectionDB := connDB.New(connDB.DBConfig{
		MasterDSN:     conf.PostgreSQLConfig.DSN,
		EnableSlave:   false,
		RetryInterval: conf.PostgreSQLConfig.RetryInterval,
		MaxIdleConn:   conf.PostgreSQLConfig.MaxIdleConn,
		MaxConn:       conf.PostgreSQLConfig.MaxConn,
	}, connDB.DriverPostgres)

	kafkaProducer := kafka.NewProducer(&conf.KafkaConfig)

	repo := repository.LoadRepository(connectionDB, kafkaProducer)
	svc := service.LoadServices(repo, conf)
	bs := &Bootstrap{
		Repository: repo,
		Service:    svc,
	}
	LoadDefaultData(bs)
	return bs
}
