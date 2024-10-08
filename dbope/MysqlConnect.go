package dbope

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var Err error

func init() {
	Db, Err = gorm.Open("mysql", "root:12345678@tcp(192.168.202.101:3306)/sb_demo?charset=utf8&parseTime=true&loc=Local")
	if Err != nil {
		fmt.Printf("mysql connect error! err = %v\n", Err)
	}
}
