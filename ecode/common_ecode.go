package ecode

// type Errno struct {
// 	Code    int
// 	Message string `example:"\"OK\" or \"some error info\""`
// }

// func (err Errno) Error() string {
// 	return err.Message
// }

// var (
// 	//common
// 	OK          = &Errno{Code: 0, Message: "Ok"}
// 	InternalErr = &Errno{Code: 10001, Message: "内部错误"}

// 	ErrRoute            = &Errno{Code: 10003, Message: "非法url"}
// 	ErrInvalidReq       = &Errno{Code: 10004, Message: "请求参数非法"}
// 	ErrMarshalFailed    = &Errno{Code: 10005, Message: "Struct转换为JSON对象失败"}
// 	ErrUnmarshalFailed  = &Errno{Code: 10006, Message: "JSON对象转换为Struct失败"}
// 	ErrIdEmpty          = &Errno{Code: 10007, Message: "id不能为空"}
// 	ErrMobileEmpty      = &Errno{Code: 10008, Message: "mobile不能为空"}
// 	ErrUpdateDataFailed = &Errno{Code: 10009, Message: "update data failed"}
// 	ErrDeleteDataFailed = &Errno{Code: 10010, Message: "delete data failed"}
// 	ErrCreateDataFailed = &Errno{Code: 10011, Message: "create data failed"}
// 	ErrInitDataFailed   = &Errno{Code: 10012, Message: "init data failed"}
// 	ErrStructEmpty      = &Errno{Code: 10013, Message: "struct不能为空"}

// 	//user
// 	ErrNoAuth       = &Errno{Code: 20001, Message: "没有权限"}
// 	ErrTokenEmpty   = &Errno{Code: 20002, Message: "未登录"}
// 	ErrTokenInvalid = &Errno{Code: 20003, Message: "登录已过期"}

// 	//redis设置
// 	ErrRedisSetFail  = &Errno{Code: 30001, Message: "设置缓存失败"}
// 	ErrRedisHSetFail = &Errno{Code: 30001, Message: "设置哈希缓存失败"}
// 	ErrRedisZAddFail = &Errno{Code: 30001, Message: "设置有序集合缓存失败"}

// 	//全局错误设置
// 	ErrFailedAndRetry = &Errno{Code: 100, Message: "failed,retry."}
// )

// func DecodeErr(err error, code int) (int, string) {
// 	if err == nil {
// 		return OK.Code, OK.Message
// 	}
// 	switch err_type := err.(type) {
// 	case *Errno:
// 		return err_type.Code, err_type.Message
// 	}
// 	if err.Error() != "" {
// 		InternalErr.Message = err.Error()
// 	}
// 	return code, InternalErr.Message
// }
