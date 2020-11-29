package main

import (
    "Week02/service"
    "fmt"
)

// 业务层
func main() {
    userService := service.NewUserService()
    users, err := userService.GetAllUsers()
    if err != nil {
        fmt.Printf("biz: cannot get all users: %+v \n", err)
        return
    }
    if len(users) < 1{
        fmt.Println("Do not have any user")
    }
}
