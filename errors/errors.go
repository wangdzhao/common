package errors

// 定义一个公共的错误结构体
type CommonError struct {
	Code    int    // 错误码
	English string // 错误的英文描述
	Chinese string // 错误的中文描述
}

// 创建一个新的CommonError
func NewCommonError(code int, english, chinese string) *CommonError {
	return &CommonError{
		Code:    code,
		English: english,
		Chinese: chinese,
	}
}

// 公共错误定义
var (
	ERR_BAD_REQUEST            = NewCommonError(400, "Bad Request", "错误的请求")
	ERR_UNAUTHORIZED           = NewCommonError(401, "Unauthorized", "未授权")
	ERR_FORBIDDEN              = NewCommonError(403, "Forbidden", "禁止访问")
	ERR_NOT_FOUND              = NewCommonError(404, "Resource Not Found", "资源未找到")
	ERR_METHOD_NOT_ALLOWED     = NewCommonError(405, "Method Not Allowed", "不允许的方法")
	ERR_CONFLICT               = NewCommonError(409, "Conflict", "冲突")
	ERR_UNSUPPORTED_MEDIA_TYPE = NewCommonError(415, "Unsupported Media Type", "不支持的媒体类型")
	ERR_INTERNAL_SERVER        = NewCommonError(500, "Internal Server Error", "内部服务器错误")
	ERR_NOT_IMPLEMENTED        = NewCommonError(501, "Not Implemented", "未实施")
	ERR_BAD_GATEWAY            = NewCommonError(502, "Bad Gateway", "错误的网关")
	ERR_SERVICE_UNAVAILABLE    = NewCommonError(503, "Service Unavailable", "服务不可用")
	ERR_GATEWAY_TIMEOUT        = NewCommonError(504, "Gateway Timeout", "网关超时")
)

// 网络请求错误
var (
	ERR_NETWORK_TIMEOUT        = NewCommonError(10001, "Network Timeout", "网络超时")
	ERR_CONNECTION_REFUSED     = NewCommonError(10002, "Connection Refused", "连接被拒绝")
	ERR_NETWORK_UNAVAILABLE    = NewCommonError(10003, "Network Unavailable", "网络不可用")
	ERR_NETWORK_PROTOCOL_ERROR = NewCommonError(10004, "Network Protocol Error", "网络协议错误")
	ERR_HOST_UNREACHABLE       = NewCommonError(10005, "Host Unreachable", "主机不可达")
	ERR_INVALID_URL            = NewCommonError(10006, "Invalid URL", "无效的URL")
	ERR_REQUEST_TOO_LARGE      = NewCommonError(10007, "Request Entity Too Large", "请求实体过大")
	ERR_RETURN_PARAMS          = NewCommonError(10008, "Failed to return the interface data. Procedure", "接口参数返回失败")
)

// 数据库操作错误
var (
	ERR_DB_CONNECTION_FAILED  = NewCommonError(20001, "Database Connection Failed", "数据库连接失败")
	ERR_DB_QUERY_FAILED       = NewCommonError(20002, "Database Query Failed", "数据库查询失败")
	ERR_DB_UPDATE_FAILED      = NewCommonError(20003, "Database Update Failed", "数据库更新失败")
	ERR_DB_RECORD_NOT_FOUND   = NewCommonError(20004, "Database Record Not Found", "数据库记录未找到")
	ERR_DB_TRANSACTION_FAILED = NewCommonError(20005, "Database Transaction Failed", "数据库事务失败")
	ERR_DB_MIGRATION_FAILED   = NewCommonError(20006, "Database Migration Failed", "数据库迁移失败")
	ERR_DB_CONFLICT           = NewCommonError(20007, "Database Conflict", "数据库冲突")
)

// 资源操作错误
var (
	ERR_FILE_NOT_FOUND          = NewCommonError(30001, "File Not Found", "文件未找到")
	ERR_FILE_ACCESS_DENIED      = NewCommonError(30002, "File Access Denied", "文件访问被拒绝")
	ERR_RESOURCE_LIMIT_EXCEEDED = NewCommonError(30003, "Resource Limit Exceeded", "资源限制超出")
	ERR_RESOURCE_ALREADY_EXISTS = NewCommonError(30004, "Resource Already Exists", "资源已存在")
	ERR_RESOURCE_NOT_AVAILABLE  = NewCommonError(30005, "Resource Not Available", "资源不可用")
	ERR_FILE_WRITE_FAILED       = NewCommonError(30006, "File Write Failed", "文件写入失败")
	ERR_RESOURCE_INITIALIZATION = NewCommonError(30007, "Resource Initialization Failed", "资源初始化失败")
	ERR_ALGORITHM_CALL_FAILURE  = NewCommonError(30008, "Algorithm call failure", "算法调用失败")
)

// 逻辑处理错误
var (
	ERR_INVALID_PARAMETER   = NewCommonError(40001, "Invalid Parameter", "无效的参数")
	ERR_UNAUTHORIZED_ACCESS = NewCommonError(40002, "Unauthorized Access", "未授权的访问")
	ERR_OPERATION_FAILED    = NewCommonError(40003, "Operation Failed", "操作失败")
	ERR_ITEM_NOT_FOUND      = NewCommonError(40004, "Item Not Found", "项目未找到")
	ERR_ITEM_ALREADY_EXISTS = NewCommonError(40005, "Item Already Exists", "项目已存在")
)

// Redis操作错误
var (
	ERR_REDIS_CONNECTION_FAILED     = NewCommonError(50001, "Redis Connection Failed", "Redis连接失败")
	ERR_REDIS_TIMEOUT               = NewCommonError(50002, "Redis Timeout", "Redis操作超时")
	ERR_REDIS_KEY_NOT_FOUND         = NewCommonError(50003, "Redis Key Not Found", "Redis键未找到")
	ERR_REDIS_DATA_SERIALIZATION    = NewCommonError(50004, "Data Serialization Failed", "数据序列化失败")
	ERR_REDIS_DATA_DESERIALIZATION  = NewCommonError(50005, "Data Deserialization Failed", "数据反序列化失败")
	ERR_REDIS_QUOTA_EXCEEDED        = NewCommonError(50006, "Redis Quota Exceeded", "Redis配额超出")
	ERR_REDIS_OPERATION_NOT_ALLOWED = NewCommonError(50007, "Operation Not Allowed", "不允许的操作")
	ERR_REDIS_CMD_EXECUTION_FAILED  = NewCommonError(50008, "Redis Command Execution Failed", "Redis命令执行失败")
	ERR_REDIS_SET_OPERATION_FAILED  = NewCommonError(50009, "Redis SET Operation Failed", "Redis SET操作失败")
	ERR_REDIS_GET_OPERATION_FAILED  = NewCommonError(50010, "Redis GET Operation Failed", "Redis GET操作失败")
	ERR_REDIS_DEL_OPERATION_FAILED  = NewCommonError(50011, "Redis DEL Operation Failed", "Redis DEL操作失败")
	ERR_REDIS_EXPIRE_FAILED         = NewCommonError(50012, "Redis EXPIRE Operation Failed", "Redis EXPIRE操作失败")
	ERR_REDIS_TX_BEGIN_FAILED       = NewCommonError(50013, "Redis Transaction Begin Failed", "Redis事务开始失败")
	ERR_REDIS_TX_COMMIT_FAILED      = NewCommonError(50014, "Redis Transaction Commit Failed", "Redis事务提交失败")
)
