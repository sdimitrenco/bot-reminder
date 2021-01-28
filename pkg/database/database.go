package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//connect to database
func Database() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	autoMigrate(db)

	return db, err
}
