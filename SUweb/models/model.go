package models

import (
	"SUweb/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db,err = gorm.Open(conf.DBType,conf.DbUrl)
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&Auth{},&Project{},&Member{},&Tutor{},
	               &ProjectComment{},&Alumni{},&Blog{})
}
