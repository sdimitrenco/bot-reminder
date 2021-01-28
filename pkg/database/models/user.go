package models

type User struct {
	Id        int32 `gorm:"primaryKey"`
	Username  string
	FirstName string
	LastName  string
	ChatId    string
}

func (u *User) GetFullName() string {
	return u.Username
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetChatId() string {
	return u.ChatId
}
