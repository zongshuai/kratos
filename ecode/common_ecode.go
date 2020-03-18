package ecode

// All common ecode
var (
	OK = add(0) // 正确

	NotModified        = add(-304) // 木有改动
	TemporaryRedirect  = add(-307) // 撞车跳转
	RequestErr         = add(-400) // 请求错误
	Unauthorized       = add(-401) // 未认证
	AccessDenied       = add(-403) // 访问权限不足
	NothingFound       = add(-404) // 啥都木有
	MethodNotAllowed   = add(-405) // 不支持该方法
	Conflict           = add(-409) // 冲突
	Canceled           = add(-498) // 客户端取消请求
	ServerErr          = add(-500) // 服务器错误
	ServiceUnavailable = add(-503) // 过载保护,服务暂不可用
	Deadline           = add(-504) // 服务调用超时
	LimitExceed        = add(-509) // 超出限制

	RequestParamErr = add(4001) // 请求参数错误

	//SignParamErr = add(4011)  //签名参数错误
	//SignTimestampExpire = add(4012) //签名过期
	//SignNonceDuplicate = add(4013) //签名nonce重复
	//SignCheckErr = add(4014) //签名验证失败
)
