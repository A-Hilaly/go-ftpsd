var request = require('request');
const url = require('url');

var ApiURL = "http://www.amine.in/";
var Token = "TOKEN";


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
