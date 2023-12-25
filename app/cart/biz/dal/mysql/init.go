package mysql

import (
	"os"

	"github.com/baiyutang/gomall/app/cart/biz/model"
	"github.com/baiyutang/gomall/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		DB.AutoMigrate(
			&model.CartItem{},
		)
	}
}
