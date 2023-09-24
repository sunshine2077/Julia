package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"web/db"
	"web/route"

	"github.com/spf13/viper"
)

// 初始化配置
func InitCfg(filename string) (*viper.Viper, error) {
	// Load Config Package
	v := viper.New()
	// Load Redis Config
	v.SetConfigFile(filename)
	// Read the Config
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	// Over
	return v, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// Step1. Load the config file
	globalCfg, err := InitCfg("./conf.yml")
	if err != nil {
		log.Panicf("config file not found: %s", err.Error())
		return
	}
	// Step2. Load Tika
	err = db.ClientToTika(ctx, globalCfg)
	if err != nil {
		log.Panicf("load tika error: %s", err.Error())
		return
	}
	// Step3. Load Qdrant
	err = db.ClientToQdrant(ctx, globalCfg)
	if err != nil {
		log.Panicf("load qdrant error: %s", err.Error())
		return
	}
	// // Step3. Load Mysql
	// err = db.ClientToMySQL(globalCfg)
	// if err != nil {
	// 	log.Panicf("load mysql error: %s", err.Error())
	// 	return
	// }
	// // Step4. Load Redis
	// err = db.ClientToRedis(globalCfg)
	// if err != nil {
	// 	log.Panicf("load redis error: %s", err.Error())
	// 	return
	// }
	// Step5. Load Route
	engine := route.LoadRoute()
	// Step6. Build listen
	engine.Run(fmt.Sprintf(":%s", globalCfg.GetString("service.port")))
}
