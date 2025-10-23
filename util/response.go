package util

import (
	"github.com/dahaipublic/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// type Gin struct {
// 	Ctx *gin.Context
// }

// type Response struct {
// 	Code int         `json:"code" mapstructure:"code" msgpack:"code"`
// 	Data interface{} `json:"data" mapstructure:"data" msgpack:"data"`
// 	Msg  string      `json:"msg" mapstructure:"msg" msgpack:"msg"`
// }

// type ResponseData struct {
// 	Current int64       `json:"current"`
// 	Size    int64       `json:"size"`
// 	Records interface{} `json:"records"`
// 	Total   int64       `json:"total"`
// }

type Response = common.THttpResp

// json
func Success(c *gin.Context, data interface{}) {
	accLan := c.Request.Header.Get("Accept-Language")
	c.JSON(http.StatusOK, common.THttpResp{
		Code: common.No_Error,
		Msg:  common.GetErrStr(accLan, common.EErrCode(0)),
		Data: data,
	})
	return
}

// json分页带总数

func SuccessCount[T any](c *gin.Context, data []T, count, page, limit int64) {
	c.JSON(http.StatusOK, &common.THttpResp{
		Code: common.No_Error,
		Msg:  "",
		Data: &common.TPageXX[T]{
			Total: int32(count),
			List:  data,
		},
	})
	return
}

// type TPage struct {
// 	List  []interface{} `json:"list"`
// 	Total int32         `json:"total"`
// }

// type common.TPageXX[T any] struct {
// 	List  []*T  `json:"list"`
// 	Total int32 `json:"total"`
// }

// error
// func Error(c *gin.Context, code int32, msg ...string) { // msg 可选参数

// 	// c.Writer.WriteHeader(200)
// 	// c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
// 	// c.Writer.WriteString(string(common.JsonResult.Get(common.EErrCode(code))))
// 	accLan := c.Request.Header.Get("Accept-Language")
// 	msgStr := strings.Join(msg, " ")
// 	// 使用 JSON 响应； 参数数据的话返回那个参数错误了 fullMsg
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":      code,
// 		"msg":       common.GetErrStr(accLan, common.EErrCode(code)),
// 		"extraInfo": msgStr,
// 	})
// 	return
// }

func RespCodeResult(c *gin.Context, errCode common.EErrCode, msg ...string) {

	accLan := c.Request.Header.Get("Accept-Language")
	if len(msg) == 0 {
		c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		c.Writer.Write(common.JsonResult.Get(accLan, errCode))
	} else {
		msgStr := strings.Join(msg, " ")
		c.JSON(http.StatusOK, gin.H{
			"code":      errCode,
			"msg":       common.GetErrStr(accLan, common.EErrCode(errCode)),
			"extraInfo": msgStr,
		})
	}
}

func MakeCodeResp(accLan string, errCode common.EErrCode, msg ...string) (rsp common.THttpResp) {
	rsp.Code = errCode
	rsp.Msg = common.GetErrStr(accLan, common.EErrCode(errCode))
	if len(msg) != 0 {
		rsp.ExtraInfo = strings.Join(msg, " ")
	}
	return
}

func RespDataResult(c *gin.Context, data interface{}) {
	accLan := c.Request.Header.Get("Accept-Language")
	c.JSON(http.StatusOK, gin.H{
		"code": common.No_Error,
		"msg":  common.GetErrStr(accLan, common.EErrCode(common.No_Error)),
		"data": data,
	})
}

func RespCommonResult(c *gin.Context, data interface{}, errCode common.EErrCode) {
	if errCode != common.No_Error {
		RespCodeResult(c, errCode)
	} else {
		RespDataResult(c, data)
	}
}

func RespCommonPageResult[T any](c *gin.Context, data []T, count int64, errCode common.EErrCode) {
	if errCode != common.No_Error {
		RespCodeResult(c, errCode)
	} else {
		page := &common.TPageXX[T]{
			Total: int32(count),
			List:  data,
		}
		RespDataResult(c, page)
	}
}

func GetLanIDByHttpHead(c *gin.Context) common.ELanDef {
	accLan := c.Request.Header.Get("Accept-Language")
	return common.GetLanID(accLan)
}

// 获取userID
func GetUserID(context *gin.Context) (userID uint64) {
	existUserID, ok := context.Get("userID")
	if !ok {
		return
	}
	userID = existUserID.(uint64)
	return
}
