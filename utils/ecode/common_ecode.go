package ecode

var (
	OK           = add(0)    // 正确
	ServerErr    = add(-500) // 服务器错误
	NoLogin      = add(-101) // 账号未登录
	AccessDenied = add(-403) // 访问权限不足
)
