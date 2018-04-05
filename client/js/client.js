//var request = require('request');
const url = require('url');

var ApiURL = "http://localhost:9000/";
var Token = "TOKEN";

var options = {
  json: {"token":"xyz"}
};

var srequest = require('sync-request');


console.log("h1");

var r = srequest("GET", "http://localhost:9000/user/create", options)
console.log(r.body)
console.log(JSON.parse(r.body))

console.log("h2");
/*
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
    console.log(bd, "kaka");
    return err, resp, bd;
}

var a, b, c = requestJson("dev/healthcheck", "GET");
*/
