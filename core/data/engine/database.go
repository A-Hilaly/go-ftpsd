package engine

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/a-hilaly/go-ftpsd/core/config"
)

//
const (
	Type         string = "MYSQL" // MySQL Database
	Protocol     string = "@tcp"  // TCP Protocol
	SettingsForm string = "?charset=utf8&parseTime=True&loc=Local"
)

var (
	DB        *gorm.DB      // Production database engine
	TestDB    *gorm.DB      // Tests      database engine
	Magicword string   = "" // Magic word ;p
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

func Init() error {
	return InitDB(Magicword)
}

// Opening a database and save the reference to `Database` struct.
func InitDB(magic string) error {
	db, err := gorm.Open("mysql", magic)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(0)
	//db.LogMode(true)
	DB = db
	return nil
}

// This function will create a temporarily database for running testing cases
func TestDBInit(magic string) error {
	db, err := gorm.Open("mysql", magic)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(0)
	db.LogMode(true)
	TestDB = db
	return nil
}

// Using this function in order to get a connection db
// you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
