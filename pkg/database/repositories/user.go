package repositories

import (
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/models"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (user User) Create(u *models.User) *models.User {
	user.db.Create(&u)
	return u
}

func (user User) Save(u *models.User) *models.User {
	user.db.Save(&u)
	return u
}

func (user User) FindByUserId(id int32) *models.User {
	var u models.User
	user.db.First(&u, id)
	return &u
}

func (user User) FindByChatId(s string) *models.User {
	var u models.User
	user.db.First(&u, "chat_id = ?", s)
	return &u
}

func (user User) GetAllRecords() []models.User {
	var u []models.User
	user.db.Find(&u)
	return u
}
