package database

import(
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


func Connect(){
	dsn := "root:12345678@tcp(127.0.0.1:3306)/online_shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	
	if err != nil {
		fmt.Println("ko connect")
        panic(err.Error())
    }
	defer db.Close()
}

