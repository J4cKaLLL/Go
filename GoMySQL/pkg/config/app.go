package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "root:Kaitoz0407@tcp(localhost:3306)/mysql?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf("Tipo de d %T\n y valor de %v", d, &d)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
