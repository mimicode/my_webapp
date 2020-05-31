package errno

// 10 - 系统错误 00  - 01
// 20 - 用户输入 00  - 01

var (
	OK                  = Errno{Code: 0, Message: "ok"}
	InternalServerError = Errno{Code: 100001, Message: "Internal server error."}
	BindJsonErr         = &Errno{Code: 10002, Message: "参数无法解析."}
	RegisterFailErr     = Errno{Code: 100101, Message: "注册失败，请稍后重试!"}
	LoginFailErr        = Errno{Code: 100102, Message: "登录失败，请稍后重试!"}

	RequiredLoginError = Errno{Code: 200101, Message: "请登录"}
	MisUserError       = Errno{Code: 200102, Message: "用户名不能为空"}
	UsernameExistsErr  = Errno{Code: 200103, Message: "用户名已存在,请修改"}
	UserPwdErr         = Errno{Code: 200105, Message: "用户名或密码错误"}
)
