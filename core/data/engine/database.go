package engine

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/a-hilaly/supfile-api/core/config"
)

//
const (
    Type         string = "MYSQL"  // MySQL Database
    Protocol     string = "@tcp"   // TCP Protocol
    SettingsForm string = "?charset=utf8&parseTime=True&loc=Local"
)

var (
    DB      *gorm.DB  // Production database engine
    TestDB  *gorm.DB  // Tests      database engine
    Magicword string = ""          // Magic word ;p
)

func SetMagicWordFromConfig(conf config.DatabaseConfig) {
	c := conf.Creds
	s := c.User + ":" + c.Password + Protocol + "(" + c.Host + ")/" + c.Database + SettingsForm
	Magicword = s
}

// init() procs automatically when this file is imported
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
	db.DB().SetMaxIdleConns(0)
	//db.LogMode(true)
	DB = db
}

// This function will create a temporarily database for running testing cases
func TestDBInit(magic string) {
	db, err := gorm.Open("mysql", magic)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(0)
	db.LogMode(true)
	TestDB = db
}

// Using this function in order to get a connection db
// you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
