package models

import "time"

type User struct {
	Id          int32 `gorm:"primary_key"`
	UserId      string
	Password    string
	DateOfBirth time.Time
	FirstName   string
	LastName    string
}

func (User) TableName() string {
	return "user_profile"
}

func (User) FindAllUsers(limit int) (*[]User, error) {
	users := &[]User{}
	result := DB.Limit(limit).Find(users)
	return users, result.Error
}

func (User) FindUser(id string) (*User, error) {
	return User{}.FindUserBy(map[string]any{
		"id": id,
	})
}

func (User) FindUserBy(params map[string]any) (*User, error) {
	user := &User{}
	result := DB.First(user, params)
	return user, result.Error
}
