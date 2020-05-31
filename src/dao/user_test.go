package dao

import (
	"com.phpstu/webapp/src/config"
	"com.phpstu/webapp/src/models"
	"com.phpstu/webapp/src/models/tables"
	"reflect"
	"testing"
)

func init() {
	//初始化配置文件 日志
	if err := config.Init("../../config.yaml", "WEBAPP"); err != nil {
		panic(err)
	}
	models.Db = &models.Databases{}
	//初始化数据库
	if err := models.Db.InitMySql(); err != nil {
		panic(err)
	}
	//defer models.Db.CloseMysql()
}

func TestUserDao_GetUserInfoById(t *testing.T) {
	type args struct {
		uid int
	}
	tests := []struct {
		name     string
		args     args
		wantUser *tables.User
	}{
		{
			name: "获取用户",
			args: args{
				1000,
			},
			wantUser: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ud := &UserDao{}
			if gotUser := ud.GetUserInfoById(tt.args.uid); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserInfoById() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserDao_Register(t *testing.T) {
	type args struct {
		uid      int
		userName string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "插入",
			args: args{
				uid:      1000,
				userName: "zhangsan",
				password: "adfasdfasdfasdfasdfasd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ud := &UserDao{}
			if err := ud.Register(tt.args.uid, tt.args.userName, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
