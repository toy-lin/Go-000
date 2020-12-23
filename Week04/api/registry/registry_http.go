package registry

import (
	"Week04/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	req := LoginRequest{
		Username: username,
		Password: password,
	}

	ls := service.NewLoginService()
	c, _ := context.WithTimeout(context.Background(), time.Second*10)
	reply, err := ls.Login(c, req)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, reply)
}
