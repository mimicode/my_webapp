package service

import (
	"com.phpstu/webapp/src/dao"
	"com.phpstu/webapp/src/models/request"
	"com.phpstu/webapp/src/models/tables"
	"com.phpstu/webapp/src/pkg/auth"
	"com.phpstu/webapp/src/pkg/errno"
	"com.phpstu/webapp/src/pkg/jwt"
	"com.phpstu/webapp/src/pkg/snowflake"
	vd "github.com/bytedance/go-tagexpr/validator"
	"github.com/spf13/viper"
)

type UserServe struct {
	userDao *dao.UserDao
}

func (us *UserServe) getUserDao() *dao.UserDao {
	if us.userDao == nil {
		us.userDao = &dao.UserDao{}
	}
	return us.userDao
}
func (us *UserServe) GetUserInfoById(uid int) *tables.User {
	return us.getUserDao().GetUserInfoById(uid)
}

//注册账号
func (us *UserServe) Register(registerForm request.RegisterForm) error {
	if err := vd.Validate(registerForm); err != nil {
		return errno.NewErrNo(errno.MisUserError.Code, err.Error())
	}

	if us.getUserDao().CheckUsernameExists(registerForm.Username) {
		return errno.UsernameExistsErr
	}
	userId, err := snowflake.GenID()
	if err != nil {
		return errno.RegisterFailErr
	}

	err = us.userDao.Register(int(userId), registerForm.Username, auth.Md5(registerForm.Password))
	if err != nil {
		return errno.RegisterFailErr
	}
	return nil
}

//用户登录
func (us UserServe) Login(loginForm request.LoginForm) (error, string) {

	err := vd.Validate(loginForm)
	if err != nil {
		return errno.NewErrNo(errno.MisUserError.Code, err.Error()), ""
	}

	userInfo := us.getUserDao().GetUserInfoByUsername(loginForm.Username)
	if userInfo == nil {
		return errno.UserPwdErr, ""
	}

	if userInfo.Password != auth.Md5(loginForm.Password) {
		return errno.UserPwdErr, ""
	}
	if tokenString, err := jwt.Sign(jwt.SignContext{ID: int64(userInfo.UserId)}, viper.GetString("jwt"), 0); err != nil {
		return errno.LoginFailErr, ""
	} else {
		return nil, tokenString
	}

}
