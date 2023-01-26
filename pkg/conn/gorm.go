package conn

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenGorm will open gorm connection and ping database to make sure connection
func OpenGorm(username, password, address, dbname string) (*gorm.DB, error) {
	dsnFormat := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsnFormat,
		username,
		password,
		address,
		dbname,
	)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// OpenGormFromViper will open gorm connection from viper config data
func OpenGormFromViper() (*gorm.DB, error) {
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "secret")
	viper.SetDefault("mysql.addr", "localhost:3306")
	viper.SetDefault("mysql.db", "deciduous")

	return OpenGorm(
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.db"),
	)
}
