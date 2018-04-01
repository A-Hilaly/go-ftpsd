package data

func tokenizedAuthentification(provider, email, token string) (*User, bool, error){
    user, err := selectUserBy("email", email)
    if err != nil {
        return nil, false, err
    }
    if user.AuthToken != "" && token == user.AuthToken {
        return user, true, nil
    }
    return nil, false, ErrorAuthentificationFailed
}

func defaultAuthentification(un, value, password string) (*User, bool, error) {
    user, err := selectUserBy(un, value)
    if err != nil {
        return nil, false, err
    }
    if user.Password != "" && Hash(password) == user.Password {
		return user, true, nil
	}
    return nil, false, ErrorAuthentificationFailed
}
