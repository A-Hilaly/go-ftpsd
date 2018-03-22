package engine

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/a-hilaly/supfile-api/config"
)

type Database struct {
	*gorm.DB
}

var Magicword string = ""

func SetMagicWordFromConfig(conf config.DatabaseConfig) {
	c := conf.Creds
	s := c.User + ":" + c.Password + "@tcp(" + c.Host + ")/" + c.Database + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(s)
	Magicword = s
}

var DB *gorm.DB
var TestDB *gorm.DB

func init() {
	//InitDB(Magicword)
}

func Init() {
	InitDB(Magicword)
}

// Opening a database and save the reference to `Database` struct.
func InitDB(magic string) {
	db, err := gorm.Open("mysql", magic)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(100)
	//db.LogMode(true)
	DB = db
}

// This function will create a temporarily database for running testing cases
func TestDBInit(magic string) {
	db, err := gorm.Open("mysql", magic)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(100)
	db.LogMode(true)
	TestDB = db
}

// Using this function in order to get a connection db
// you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
