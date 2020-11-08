package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

type Auth struct {
	AuthId uint `gorm: "primary_key" json: "auth_id" gorm:"auto-increment"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Token string `json:"token"`
	StudentId string `json:"student_id" form:"student_id"`
	Mail string `json:"mail" form:"mail"`
	Major string `json:"major" form:"major"`
}

type Token struct {
	Token string `header:"Token"`
}


func (user *Auth) Create() error{
	err := db.Create(&user).Error
	return err
}

func (user *Auth) Query(name string) error {
	err := db.Where(&Auth{Username:name}).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *Auth) Update(attribute,value string) {
	db.Model(user).Update(attribute,value)
}

func (token *Token) Query() bool{
	var user Auth
	err := db.Where("token = ?",token.Token).First(&user).Error
	return gorm.IsRecordNotFoundError(err)
}