package data

func tokenizedAuthentification(provider, email, token string) (*User, bool, error){
    user, err := SelectUserBy("email", email)
    if err != nil {
        return nil, false, err
    }
    if user.AuthToken != "" && token == user.AuthToken {
        return user, true, nil
    }
    return nil, false, ErrorAuthentificationFailed
}

func defaultAuthentification(email, password string) (*User, bool, error) {
    user, err := SelectUserBy("email", email)
    if err != nil {
        return nil, false, err
    }
    if user.Password != "" && Hash(password) == user.Password {
		return user, true, nil
	}
    return nil, false, ErrorAuthentificationFailed
}

// Authentificate user
func AuthentificateUser(method, email, pass string) (*User, bool, error) {
	switch method {
    case "simple":
        return defaultAuthentification(email, pass)
    case "facebook":
        return tokenizedAuthentification("facebook", email, pass)
    case "google":
        return tokenizedAuthentification("google", email, pass)
    default:
        return nil, false, ErrorUnkownMethod
    }
}
