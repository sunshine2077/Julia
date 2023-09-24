package db

import (
	"context"
	"fmt"
	"web/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// DB Client
var MysqlClient *gorm.DB

// Client to MySQL
func ClientToMySQL(ctx context.Context, v *viper.Viper) error {
	var err error
	MysqlClient, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			v.GetString("mysql1.user"),
			v.GetString("mysql1.pwd"),
			v.GetString("mysql1.ip"),
			v.GetInt("mysql1.port"),
			v.GetString("mysql1.dbname")))
	if err != nil {
		return err
	}
	MysqlClient.DB().SetMaxOpenConns(v.GetInt("mysql1.maxOpenConns"))
	MysqlClient.DB().SetMaxIdleConns(v.GetInt("mysql1.maxIdleConns"))
	err = MysqlClient.DB().PingContext(ctx)
	if err != nil {
		return err
	}
	// Auto Migrate
	MysqlClient.Debug().AutoMigrate(&model.User{})
	return nil
}
