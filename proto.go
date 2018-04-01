package main

import (
    "fmt"

    "github.com/a-hilaly/supfile-api/core/data"
    "github.com/a-hilaly/supfile-api/core/data/engine"
    "github.com/a-hilaly/supfile-api/core/config"
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

func test() {
    conf := config.LoadConfig()
    engine.SetMagicWordFromConfig(conf.Database)
    engine.Init()
    //data.Init()
    man := data.NewManager("first")
    config := man.GetConfig()
    fmt.Println(config)
    str, err := man.GetUserAccountID("username", "amine")
    fmt.Println(str, err)
    user, err := man.GetUser(str)
    fmt.Println(user, err)

    user, err = man.CreateUser("amiine", "huge@lol", "normal", "simple", "ueh")
    fmt.Println(user, err)
}

func main() {
    test()
}
