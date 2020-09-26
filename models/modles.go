package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"harbor/pkg/toml"
	"log"
)

var db *gorm.DB

type Model struct {
	Uuid int `gorm:"primary_key;type:char(32);not null" json:"uuid"`
}

func init() {
	var err error
	mysql := toml.GetMysqlConfig()
	db, err = gorm.Open(mysql.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.User, mysql.Pwd, mysql.Host, mysql.Database))
	if err != nil {
		log.Fatal(err)
	}
	err = db.DB().Ping()
	if err != nil{
		log.Fatalln(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return mysql.TablePrefix + defaultTableName;
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}