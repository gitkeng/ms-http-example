package main

import (
	"github.com/gitkeng/microservice"
	"github.com/gitkeng/microservice/log"
)

func main() {
	cfgLocation := "./conf/app_setting.yaml"
	conf, err := microservice.NewConfig(cfgLocation)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	logCfg, isLogCfgFound := conf.GetLogConfig()
	if !isLogCfgFound {
		log.Fatalf("Error: log config not found")
	}
	apiCfg, isAPICfgFound := conf.GetAPIConfig()
	if !isAPICfgFound {
		log.Fatalf("Error: log apit not found")
	}
	dbCfgs, isDBCfgFound := conf.GetDBConfigs()
	if !isDBCfgFound {
		log.Fatalf("Error: db config not found")
	}
	redisCfgs, isRedisCfgFound := conf.GetRedisConfigs()
	if !isRedisCfgFound {
		log.Fatalf("Error: redis config not found")
	}

	ms, err := microservice.New(
		microservice.WithAPIConfig(apiCfg),
		microservice.WithLogConfig(logCfg),
		microservice.WithDBConfigs(dbCfgs...),
		microservice.WithRedisConfigs(redisCfgs...),
	)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	// register handler
	registerEmployeeHandler(ms)

	if err := ms.Start(); err != nil {
		log.Warnf("Error: %s", err.Error())
	}
}
