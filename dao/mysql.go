package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	//dst := "root:ty82885279@(127.0.0.1:3306)/ielts?charset=utf8mb4&parseTime=True&loc=Local"
	//dst := "root:root123@(127.0.0.1:3306)/ielts?charset=utf8mb4&parseTime=True&loc=Local"
	dst := "root:Mx123123.@(127.0.0.1:3306)/ielts?charset=utf8mb4&parseTime=True&loc=Local"
	//dst := "root:root123@(127.0.0.1:3306)/ielts?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dst)
	if err != nil {

		return err
	}
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
