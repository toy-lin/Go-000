package registry

import (
    "context"
    "github.com/gin-gonic/gin"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginReply struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
}

type LoginHTTPServer interface {
    Login(context context.Context, request *LoginRequest)
}

func RegisterLoginHTTPServer(e *gin.Engine) {
    e.GET("/apis/login", Login)
}

func Login(ctx *gin.Context) {
    username, _ := ctx.GetQuery("username")
    password, _ := ctx.GetQuery("password")


}
