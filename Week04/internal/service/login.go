package service

import (
    "Week04/api/registry"
    "Week04/internal/biz"
    "Week04/internal/data"
    "context"
)

type User struct {
    Username string
    Password string
}

type LoginService interface {
    Login(ctx context.Context, req registry.LoginRequest) (registry.LoginReply, error)
}

func NewLoginService() LoginService {
    return &loginService{}
}

type loginService struct {
}

func (s *loginService) Login(ctx context.Context, req registry.LoginRequest) (reply registry.LoginReply, err error) {
    var user biz.User
    user.Username = req.Username
    user.Password = req.Password

    var userUseCase = biz.NewUserUseCase(data.NewUserRepo())
    userExists, err := userUseCase.Exists(&user)
    if err != nil {
        return
    }
    if !userExists {
        reply.Code = 1
        reply.Msg = "You should sign up first"
        return
    }

    reply.Code = 0
    reply.Msg = "Login success"
    return
}
