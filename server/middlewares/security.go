package middlewares

import (
    //"fmt"
    "errors"
    //"net/http"

    "github.com/gin-gonic/gin"

    //"github.com/a-hilaly/supfile-api/server/json"
)

var (
    ErrorUnknownToken = errors.New("Unkown token")
)


var hiddenToken string

func SetToken(str string) {hiddenToken = str}

func ValidateToken(token string) bool {return hiddenToken == token}


func TokenValidationMW() gin.HandlerFunc {
    return func(c *gin.Context) {
        //fmt.Println(c)
        /*
        fmt.Println(c)
        d := c.Copy()
        r, err := json.FromContext(*d, struct{}{})
        fmt.Println("mdl1", r, err)
        if err != nil {
            c.JSON(http.StatusOK, json.Response{
                Success : true,
                Errors  : []string{"Internal", err.Error()},
                Data    : nil,
            })
        }
        if !validateToken(r.Token) {
            c.JSON(http.StatusOK, json.Response{
                Success : true,
                Errors  : []string{"TOKEN: invalid token"},
                Data    : nil,
            })
        }
        fmt.Println("ALL OK")
        fmt.Println(c)
        */
        c.Next()
    }
}
