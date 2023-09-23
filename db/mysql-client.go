package db

import (
	"fmt"
	"web/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// DB Client
var MysqlClient *gorm.DB

// Client to MySQL
func ClientToMySQL(v *viper.Viper) error {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			v.GetString("mysql1.user"),
			v.GetString("mysql1.pwd"),
			v.GetString("mysql1.ip"),
			v.GetString("mysql1.port"),
			v.GetString("mysql1.dbname")))
	if err != nil {
		return err
	}
	MysqlClient = db
	//  Auto Migrate
	MysqlClient.Debug().AutoMigrate(&model.User{})
	return nil
}
