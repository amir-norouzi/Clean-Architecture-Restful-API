package mysql

import (
	"RestApi/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db_name     = config.LoadConfig("DB_DATABASE")
	db_host     = config.LoadConfig("DB_HOST")
	db_port     = config.LoadConfig("DB_PORT")
	db_username = config.LoadConfig("DB_USERNAME")
	db_password = config.LoadConfig("DB_PASSWORD")
	Db          *gorm.DB
)

func Init() *gorm.DB {
	Db = connect()
	return Db
}

func connect() *gorm.DB {
	var err error
	dsn := db_username + ":" + db_password + "@tcp" + "(" + db_host + ":" + db_port + ")/" + db_name + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error occurred : error=%v", err)
		return nil
	}

	return db
}
