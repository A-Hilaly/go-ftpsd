# go-ftpsd

# Requirements

- MySQL database
- config.json file


config.json contain informations about :
- running port
- mode
- security token
- database engine type
- database name
- database credentials

```python
{
    "port" : 9000,
    "mode" : "DEFAULT",     
    "auth" : {
        "token" : "xyz",   
    },
    "logging" : {
        "type" : "stdout",
        "logdir" : "*",
        "logtable" : "*"
    },
    "database" : {
        "creds" : {
            "database" : "DB", 
                                     

            "host" : "localhost",
            "port" : 3036,
            "user" : "root",
            "password" : "hello"
        }
    }
}
```

# Running the Api

Running with released binaries (make sure config.json is on the working directory):

- `./bin/runserver-secure`

Running with go compiler:

- `go run main.go`

# Compile the Api

```shell
cd /path/to/go-ftpsd
env GOOS=target-OS GOARCH=target-architecture NAME=binary-name go build -o $NAME
#https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
```
