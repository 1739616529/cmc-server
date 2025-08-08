package resp

/*
*
AA 表示系统/服务模块
BB 表示子模块或错误子类（如数据库错误01，参数错误02）
CC 表示具体错误代码（如不存在04）
*/

var (
	StatusBadRequest           = 100001 // 请求参数错误
	StatusUnauthorized         = 100002 // 未授权
	StatusForbidden            = 100003 // 无权限访问
	StatusNotFound             = 100004 // 资源未找到
	StatusInternalServerError  = 100005 // 服务器内部错误
	StatusCaptchaValidateFiled = 100006 // 验证码错误
	StatusCaptchaExpiration    = 100007 // 验证码失效
	StatusCaptchaFrequent      = 100008 // 验证码请求频繁

	// 01 用户

	StatusUserNotFound     = 100104 // 用户不存在
	StatusUserExists       = 100105 // 用户已存在
	StatusUserLoginFailed  = 100106 // 用户登录失败
	StatusUserTokenExpired = 100107 // 用户登录信息过期

	// 02 权限

	StatusRolePrimissionNotFound = 100204 // 角色或权限不存在
	StatusRolePrimissionExists   = 100205 // 角色或权限已存在
	StatusRolePrimissionFailed   = 100206 // 角色或权限操作失败
	StatusRolePrimissionInvalid  = 100207 // 角色或权限无效

	// 99 数据库

	StatusDBQueryFailed  = 990001 // 数据库查询失败
	StatusDBInsertFailed = 990002 // 数据库插入失败
	StatusDBUpdateFailed = 990003 // 数据库更新失败
	StatusDBDeleteFailed = 990005 // 数据库删除失败
)
