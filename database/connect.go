package database

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func Connect(){
	db, err := gorm.Open("mysql", "root:12345678@/online_shop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Cannot connect to database")
	} else {
		fmt.Println("Successfully")
	}
	defer db.Close()
}