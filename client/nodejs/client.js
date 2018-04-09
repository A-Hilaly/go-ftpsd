//!/bin/node

var request = require('request');
var url     = require('url');

var APIToken = "xyz";
var port     = 9000;

var mainUrl  = "http://localhost:9000/";
var userUrl = mainUrl + "user/";
var devUrl  = mainUrl + "dev/";


// return request option (URI, method, json body)
var createUserCall = function (
    create_data,   //bool
    create_system, //bool
    username,      //string
    password,      //string
    email,         //string
    auth_type,     //string
    account_type,  //string
    max_storage,   //int
) {
    return {
        uri    : userUrl + "create",
        method : "POST",
        json   : {
            token : APIToken,
            data  : {
                username     : username,     //"stricker23",
                email        : email,        //"hey@stricking.com",
                auth_type    : auth_type,    //"simple",
                account_type : account_type, //"admin" | "simple"
                password     : password,     //"password"
                max_storage  : max_storage,  //30 (Gb)
            },
            option : {
                system : create_system, //create user on OS
                data   : create_data,   //create user in database
            }
        }
    }
};

//request example using callCreateUser
var createUserRequest = function() {
    var options = createUserCall(
        true,
        true,
        "testusername",
        "pass",
        "email@k.o",
        "simple",
        "admin",
        30,
    );
    request(options, function (error, response, body) {
        //console.log(error, response, body);
    if (!error && response.statusCode == 200) {
        console.log(body);
    } else {
        console.log("error");
    }
    });
}

createUserRequest();


/*
var createUserOpt = {
  json: {"token":"xyz"}
};

var srequest = require('sync-request');

*/
/*

console.log("h1");

var r = srequest("GET", "http://localhost:9000/user/create", options)
console.log(r.body)
console.log(JSON.parse(r.body))

console.log("h2");


var bod;

var sendJsonRequest = function(method, exurl, data) {
    var options = {
      uri: url.resolve(ApiURL, exurl),
      method: 'GET',
      json: data
    };
    var a = request(options, function (error, response, body) {
      if (!error && response.statusCode == 200) {
        //console.log(body); // Print the shortened url.
        bod = body
        console.log(body);
      }
    });

    return a;
}


console.log("hello");
console.log(sendJsonRequest("GET", "/user/create", {"token":"xyz"}));
console.log(sendJsonRequest("GET", "/user/create", {"token":"xyz"}));
console.log("lol");


*/

/*
var xhr = new XMLHttpRequest();
var uurl = url.resolve(ApiURL, "/user/create");
xhr.open("POST", uurl, true);
xhr.setRequestHeader("Content-type", "application/json");
xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
        var json = JSON.parse(xhr.responseText);
        console.log(json.email + ", " + json.password);
    }
};
var data = JSON.stringify({"email": "hey@mail.com",
                           "password": "101010"});
xhr.send(data);

*/
/*
var requestJson = function(uri, method, onsuccess) {
    var err, resp, bd;
    if (method == "GET") {
        bd = request.get(url.resolve(ApiURL, uri), (error, response, body) => {console.log(body); return body})
    }
    else if (method == "POST") {
        err, resp, bd = request.get(url.resolve(ApiURL, uri), (error, response, body) => {return error, response, body});
    };
    console.log(bd, "k");
    return err, resp, bd;
}

var a, b, c = requestJson("dev/healthcheck", "GET");
*/
