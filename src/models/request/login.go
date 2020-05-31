package request

type LoginForm struct {
	Username string `json:"username" vd:"len($)>0;msg:'用户名不能为空'"`
	Password string `json:"password" vd:"len($)>0;msg:'密码不能为空'"`
}

type RegisterForm struct {
	Username string `json:"username" vd:"len($)>0;msg:'用户名不能为空'"`
	Password string `json:"password" vd:"len($)>0;msg:'密码不能为空'"`
}
