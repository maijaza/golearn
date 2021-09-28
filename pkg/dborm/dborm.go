package dborm

import (
    "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
    // "github.com/jinzhu/gorm/dialects/mssql"
)


type Product struct {
	gorm.Model
	Code  string
	Price uint
  }

func Test(){
	//dsn := "sqlserver://sa:P@ssw0rd@localhost:1433?database=GmFitsRepoUAT_THB_OLD"
	dsn := "sqlserver://gmadm:P@ssw0rd@100.127.195.85:1433?database=test"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("A")
		fmt.Println(err)
	}
 
//	defer db.Close()
	db.AutoMigrate(&Product{})
	//db.Create(&Product{Code: "D43", Price: 100})

	var prd Product
	//db.Where("Code = 'D42'").Find(&prd)
    db.Table('products')
	fmt.Println(prd)
    
}

 