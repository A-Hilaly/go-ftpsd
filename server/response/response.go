package response


type Reponser interface {
    OnContext(c *gin.Context) error
}


type Response struct {
    success int         `json:"success"`
    Errors  []string    `json:"error"`
    Data    interface{} `json:"data"`
}
