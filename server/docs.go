/*
# Api docs

##### Create user

```python
#
URL = "/user/create"
Method = "POST"

# Request body
request = {
    "level"    : 2,
    "username" : "alex123",
    "email"    : "alexdu123@edu.fr",
    "authtype" : "simple",
    "pass"     : "password",
}

# Response body
response = {
    "success" : true,
    "errors"  : ["no error"],
    "data"    : {userid : "9x3780019"}
}
```

##### Authentificate user

Simple authentification : email/username + password

```python
#
URL = "/user/auth"
Method = "POST"

# Request body
request = {
    "email"    : "alexdu123@edu.fr",
    "authtype" : "simple",
    "pass"     : "password",
}

# Response body
response = {
    "success" : true,
    "errors"  : ["no error"],
    "data"    : {userid : "9x3780019"}
}
```

Facebook / Google / Twitter authentification

```python
#
URL = "/user/auth"
Method = "POST"

# Request body
request = {
    "email"    : "alexdu123@edu.fr",
    "authtype" : "facebook",
    "pass"     : "token"
}

# Response body
response = {
    "success" : true,
    "errors"  : ["no error"],
    "data"    : {userid : "9x3780019"}
}
```

##### Update user

```python
#
URL = "/user/update"
Method = "POST"

# Request body
request = {
    "userid"  : "12345678",
    "updates" : {
        "firstname" : "alexandre",
        "lastname"  : "dupont",
        "state"     : "inactive"
    }
}

# Response body
response = {
    "success" : true,
    "errors"  : ["no error"],
}
```

##### Actions on user

```python
#
URL = "/user/actions"
Method = "POST"

# Request body
request = {
    "userid"  : "12345678",
    "email"   : "zinjai@tai.china",
    "actions" : ["activate"], # drop, ban, limit, remove, stats
    "args"    : []


# Response body
response = {
    "success" : true,
    "errors"  : ["no error"],
    "data"    : {userid : "9x3780019"},
}
```
*/
package server
