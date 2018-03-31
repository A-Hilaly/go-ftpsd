package handlers

type RequestModel interface {
    FromContext()
}


type ResponseError struct {
    Id          int    `json:"id"`
    Error       string `json:"error"`
    Description string `json:"description"`
}


type RequestJson struct {
    Token string
}
