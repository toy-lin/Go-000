package dao

import (
	"Week02/model"
	"database/sql"
	"github.com/pkg/errors"
)

// DAO层
type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (dao *UserDao) GetAllUsers() (users []model.User, err error) {
	data, err := dbGetUsers()
	// 对于查询不到数据的情况，一般在DAO层处理掉，不往上抛
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		err = errors.Wrap(err, "dao: failed query all users")
		return
	}
	// 填充数据到users切片中
	for range data {
		//todo
	}
	return
}

func dbGetUsers() ([]map[string]interface{}, error) {
	return nil, sql.ErrNoRows
}
