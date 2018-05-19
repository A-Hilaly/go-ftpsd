#!/bin/bash

function validate-string() {
    if [ "$1" = "" ]; then
        echo "false"
    else
        echo "true"
    fi
}

function validate-args() {
    for var in "$@"; do
        d=$(validate-string "$var")
        if [ "$d" = "false" ] || [ "$d" = "" ]; then
            echo "false"
            return 1
        fi
    done
    echo "true"
}

function user-exist() {
    valid=$(validate-args $@)
    if [ "valid" = "false" ]; then
        echo "false"
        return 1
    fi

    if id "$1" >/dev/null 2>&1; then
        echo "true"
    else
        echo "false"
    fi
}

function create-user() {
    username="$1"
    password="$2"

    valid=$(validate-args $username $password)
    if [ "$valid" = "false" ]; then
        echo "ERROR: empty args"
        return 1
    fi

    uexist=$(user-exist $username)
    if [ $uexist = "true" ]; then
        echo "ERROR: user already exist"
        return 1
    fi

    sudo mkdir -p "/home/ftp/$username"
    sudo useradd -s /bin/false -d /home/ftp/$username -g sftp $username
    sudo usermod --password $password $username
    sudo chown root:root /home/ftp/$username

    return 0
}


function drop-user() {
    username="$1"

    valid=$(validate-args $username)
    if [ "$valid" = "false" ]; then
        echo "ERROR: empty args"
        return 1
    fi

    uexist=$(user-exist $username)
    if [ $uexist = "false" ]; then
        echo "ERROR: user doesn't exist"
        return 1
    fi

    #NOTE:
    # sudo userdel -Z -r -f $username
    # userdel: -Z requires SELinux enabled kernel
    # userdel: -r mail spool (/var/mail/goki) not found

    sudo userdel -f $username

    return 0
}

function _help() {
    echo "commands: new | drop | exist"
}

function handle-args() {
    command="$1"
    username="$2"
    shift
    if [ "$command" = "new" ]; then
        password="$1"
        create-user $username $password
    elif [ "$command" = "drop" ]; then
        drop-user $username
    elif [ "$command" = "exist" ]; then
        user-exist $username
    else
        _help
    fi
}

handle-args $@