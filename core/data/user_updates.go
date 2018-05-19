package data

import (
	"github.com/a-hilaly/go-ftpsd/core/data/engine"
)

func updateUserBy(by, value, field string, nvalue interface{}) error {
	user, err := selectUserBy(by, value)
	if err != nil {
		return err
	}
	return engine.DB.Model(user).Update(field, value).Error
}

func updateUserMapBy(by, value string, values map[string]interface{}) error {
	user, err := selectUserBy(by, value)
	if err != nil {
		return err
	}
	return engine.DB.Model(user).Updates(values).Error
}

/*
switch field {
case "username":
    user.Username = nvalue.(string)
case "lastname":
    user.Username = nvlaue.(string)
case "email":
    user.Email = nvlaue.(string)
case "password":
    user.Password = nvlaue.(string)
case "auth_type":
    user.AuthType = nvlaue.(string)
case "account_id":
    user.AccountId = nvlaue.(string)
case "auth_type":
    user.AuthType = nvlaue.(string)
case "auth_token":
    user.AuthToken = nvlaue.(string)
case "state":
    user.State = nvlaue.(string)
case "max_storage":
    user.MaxStorage = nvalue.(int)
}
*/
