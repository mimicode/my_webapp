package dao

import (
	"com.phpstu/webapp/src/models"
	"com.phpstu/webapp/src/models/tables"
	"database/sql"
	"go.uber.org/zap"
)

type UserDao struct {
}

//获取用户信息
func (ud *UserDao) GetUserInfoById(uid int) *tables.User {
	var user tables.User
	if err := models.Db.MysqlDb.Get(&user, "select * from user where user_id = ?", uid); err != nil {
		if err != sql.ErrNoRows {
			zap.L().Error("query_error", zap.Error(err), zap.Int("uid", uid))
		}
		return nil
	}
	return &user
}

//获取用户信息
func (ud *UserDao) GetUserInfoByUsername(username string) *tables.User {
	var user tables.User
	if err := models.Db.MysqlDb.Get(&user, "select * from user where username = ?", username); err != nil {
		if err != sql.ErrNoRows {
			zap.L().Error("query_error", zap.Error(err), zap.String("username", username))
		}
		return nil
	}
	return &user
}

//注册账号
func (ud *UserDao) Register(uid int, userName, password string) error {

	execSql := "insert into user (user_id,username,password) value (?,?,?)"
	if _, err := models.Db.MysqlDb.Exec(execSql, uid, userName, password); err != nil {
		zap.L().Error("insert_error", zap.Error(err), zap.Int("user_id", uid), zap.String("username", userName), zap.String("password", password))
		return err
	}
	return nil
}

//检测用户名是否存在
func (ud UserDao) CheckUsernameExists(username string) bool {
	querySql := "select count(user_id) from user where username=?"
	var count int
	if err := models.Db.MysqlDb.Get(&count, querySql, username); err != nil {
		zap.L().Error("query_error", zap.Error(err), zap.String("sql", querySql), zap.String("username", username))
		return true
	}
	return count > 0
}
