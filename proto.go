package main

import (
    "fmt"

    //"github.com/a-hilaly/supfile-api/core/data"
    //"github.com/a-hilaly/supfile-api/core/data/engine"
    //"github.com/a-hilaly/supfile-api/core/config"
    //"github.com/a-hilaly/supfile-api/core/system"
)

func testData() {
    //conf := config.LoadConfig()
    //engine.SetMagicWordFromConfig(conf.Database)
    //engine.Init()
    //data.Init()
    //fmt.Println("hello")
    //_, err := data.NewUser("kappa12", "kappa@g2.com", "123", "test-account", "simple", "")
    //fmt.Println(user, err)

    //user, err := data.SelectUserBy("email", "kappa@g.com")
    //users, err := data.SelectUsersBy("account_type", "test-account")
    //e, err := data.UserExistBy("username", "kappa")
    //err2 := data.DropUserBy("username", "kappa1")
    //users, err2 := data.SelectAllUsers()
    //fmt.Println(users, err, err2)
    //user, t , err:= data.AuthentificateUser("facebook", "kappa@g3.com", "XXX")
    //fmt.Println(user, t, err)
}

func test() interface{} {
    return struct{A int; B int}{}
}

type SupfileSystemInterface interface {
    Hello()
}

type SupfileSystem struct {A int}

func (s *SupfileSystem) Hello() {
    fmt.Println("Hello", s.A)
}

type SupfileDataInterface interface {
}

type SupfileData struct {}

type SupfileCoreInterface interface {
    // Users
    System() (SupfileSystemInterface)
    Data() (SupfileDataInterface)
}

type SupfileCoreApi struct{
    sys  SupfileSystemInterface
    data SupfileDataInterface
}

func (sfc *SupfileCoreApi) System() (SupfileSystemInterface) {
    return sfc.sys
}

func (sfc *SupfileCoreApi) Data() (SupfileDataInterface) {
    return nil
}

func CoreApi() SupfileCoreInterface {
    return &SupfileCoreApi{sys: &SupfileSystem{}, data: &SupfileData{}}
}


func main() {
    a := CoreApi()
    fmt.Println(a)
    a.System().Hello()
    fmt.Println(a.Data())
}
