package data

import "Week04/internal/biz"

func NewUserRepo() biz.UserRepo {
    return &userRepo{}
}

type userRepo struct {
    //redisClient redis
    //memCacheClient
    //mysqlClient
}

func (r *userRepo) CheckExists(user *biz.User) (bool, error) {
    return false, nil
}
