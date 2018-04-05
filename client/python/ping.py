import sys
import requests


class url:
    create = 'http://localhost:9000/user/create'
    auth = 'http://localhost:9000/user/auth'
    drop = 'http://localhost:9000/user/drop'
    info = 'http://localhost:9000/user/info'
    update = 'http://localhost:9000/user/update'


class params:
    """
        Json objects to request
    schema:
    {
        "token" : "xxx",
        "data"  : {
            "username" : "a",
            ...
        },
        "option" : {
            "account_id", "",
            "data" : true,
            "system" : true,
            "actions" : [""]
        }
    }
    """

    create = dict(
        token="xyz",
        data=dict(
            username="aminetest",
            password="lol",
            email="hilalyamine@gmail.com",
            auth_type="simple",
            account_type="admin"
        ),
        option=dict(
            data=True,
        ),
    )

    auth = dict(
        token="xyz",
        data=dict(
            #username="aminetest",
            password="lol",
            #email="hilalyamine@gmail.com",
            auth_type="simple",
        ),
        option=dict(
            account_id='d4af1ee98a4e80afdcc66fb27d76195a1b285d5fc52cdc3956f2fe5552c12aee',
        )
    )

    drop = dict(
        token="xyz",
        option=dict(
            data=True,
            account_id='d4af1ee98a4e80afdcc66fb27d76195a1b285d5fc52cdc3956f2fe5552c12aee',
        )
    )

    info = dict(
        token="xyz",
        option=dict(
            data=True,
            account_id='d4af1ee98a4e80afdcc66fb27d76195a1b285d5fc52cdc3956f2fe5552c12aee',
        )
    )

    update = dict(
        token="xyz",
        data=dict(
            username="aminenewname",
            email="hilaly1amine@gmail.com",
            max_storage=30,
        ),
        option=dict(
            account_id='d4af1ee98a4e80afdcc66fb27d76195a1b285d5fc52cdc3956f2fe5552c12aee',
        )
    )


if __name__ == "__main__":

    args = sys.argv

    if args[1] == "create":
        resp = requests.post(url=url.create, json=params.create)
        data = resp.json()
        print(data)

    elif args[1] == "auth":
        resp = requests.post(url=url.auth, json=params.auth)
        data = resp.json()
        print(data)

    elif args[1] == "drop":
        resp = requests.post(url=url.drop, json=params.drop)
        data = resp.json()
        print(data)

    elif args[1] == "info":
        resp = requests.post(url.info, json=params.info)
        data = resp.json()
        print(data)

    elif args[1] == "update":
        resp = requests.post(url.update, json=params.update)
        data = resp.json()
        print(data)
