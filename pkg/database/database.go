package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//connect to database
func Database() (*gorm.DB, error) {
	var host = os.Getenv("DATABASE_HOST")
	var port = os.Getenv("DATABASE_PORT")
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var dbname = os.Getenv("DATABASE_NAME")

	//var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)
	var dbInfo = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	autoMigrate(db)

	return db, err
}
