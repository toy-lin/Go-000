package service

import (
    "Week02/dao"
    "Week02/model"
)

// 服务层
var defaultUserDao *dao.UserDao

func init() {
    defaultUserDao = dao.NewUserDao()
}

type UserService struct {
}

func NewUserService() *UserService {
    return &UserService{}
}

func (service *UserService) GetAllUsers() ([]model.User, error) {
    return defaultUserDao.GetAllUsers()
}
