package config

import (
	//"database/sql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	// gorm v1
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	db *gorm.DB
)

func Connect() {
	// option 1 from tutorial
	/*
		//db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
		d, err := gorm.Open("mssql", "sqlserver://uae:Demo1234@localhost:9010?database=TestDev") //8000
		if err != nil {
			panic(err)
		}
		db = d
	*/

	// second option from stack overflow
	/*
			dsn := "sqlserver://user_log:dhhdf127ihd@11.111.1.111?database=Database_Log"
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	*/
	dsn := "sqlserver://uae:Demo1234@localhost:1434?database=TestDev"
	var err error
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) //8000
	//d, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) //8000
	//d, err := gorm.Open("mssql", "sqlserver://uae:Demo1234@localhost:1433?database=TestDev") //8000 // 9930 //1433 // localhost/localhost:1433
	if err != nil {
		panic(err)
	}
	//db = d
	//GetDB()
	//defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
