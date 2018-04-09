# Supfile api

Supfile API

# Requirements

- MySQL database
- config.json file

# Config

```python
{
    "name" : "supfile-api",
    "port" : 9000,          # Using port
    "mode" : "DEFAULT",     
    "auth" : {
        "token" : "xyz",    # Communication token
        "dev"   : "DEV"
    },
    "logging" : {
        "type" : "stdout",
        "logdir" : "*",
        "logtable" : "*"
    },
    "database" : {
        "type" : "mysql",
        "creds" : {
            "database" : "SUPFILEDB", #NOTE database must be created manually
                                      #     tables can be created automatically

            "host" : "localhost",
            "port" : 3036,
            "user" : "amine",
            "password" : "a"
        }
    }
}

```

# Running the Api

`./bin/runserver-secure`

make sure config.json is on the working directory

With go compiler
`go run main.go`

# Compile the Api

```shell
cd /path/to/supfile-api
env GOOS=target-OS GOARCH=target-architecture NAME=binary-name go build -o $NAME

#https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
```

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
}

# Response body
response = {
    "success" : true,
    "errors"  : ["no error"],
    "data"    : {userid : "9x3780019"},
}
```
