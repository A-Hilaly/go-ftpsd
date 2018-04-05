#/bin/bash


function build-binaries() {
    go build -o "$1"
}

if [ "$1" = "" ]; then
    echo "Help: build [name]"
else
    build-binaries $1
fi
