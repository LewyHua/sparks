package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sparks/config"
	"sparks/model"
	"time"
)

var (
	DB *gorm.DB
)

func Init(appConfig *config.AppConfig) (err error) {
	var conf = appConfig.MySQLConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Address,
		conf.Port,
		conf.Database,
	)

	mysqlLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
			Colorful:                  false,
		})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLog, PrepareStmt: true})
	if err != nil {
		log.Println(dsn)
		log.Fatal("connect to mysql failed:", err)
	}

	err = DB.AutoMigrate(&model.User{})
	err = DB.AutoMigrate(&model.Video{})
	err = DB.AutoMigrate(&model.Relation{})
	err = DB.AutoMigrate(&model.Comment{})
	err = DB.AutoMigrate(&model.Favorite{})
	if err != nil {
		zap.L().Error("auto migrate table failed", zap.Error(err))
		return
	}
	return nil
}
