package system

import (

    "github.com/a-hilaly/supfile-api/core/models"
)

// Create supfile user (Storage + system user)
func CreateUser() {
    _, _ := models.NewUser(first, last, email, password, type_)
}

func DropUser(email string) {
    //_ := models.DropUser()
}
