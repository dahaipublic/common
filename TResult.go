/*
go**************************************************************************

	File            : TResult.go
	Subsystem       : commmon
	Author          :
	Date&Time       : 2015-8-5 18:17
	Description     :

	Revision        :

	History
	-------


	Copyright (c) .

**************************************************************************go
*/
package common

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type EErrCode = int32

// common err code
// 代码区间 [0-199], 该区间为预留字段
const (
	No_Error       EErrCode = 0 // 操作成功
	Err_Unknown    EErrCode = 1 // 未知错误
	Err_System     EErrCode = 2 // 系统错误，通用系统逻辑错误，指代码执行到非预计的地方和非正常的逻辑
	Err_Param      EErrCode = 3 // 参数错误，通用参数错误
	Err_Cookies    EErrCode = 4 // 参数错误，读写cookies错误
	Err_Redis      EErrCode = 5 // 服务异常，通用redis错误
	Err_DB         EErrCode = 6 // 服务异常，通用数据库错误
	Err_Marshal    EErrCode = 7 // 数据异常，通用序列化错误
	Err_UnMarshal  EErrCode = 8 // 数据异常，通用反序列化错误
	Err_RemoteCall EErrCode = 9 // 通信失败，通用服务器之间请求错误，指网络错误

	Err_HttpReadBody  EErrCode = 10 // Http读body错误
	Err_RemoteReturn  EErrCode = 11 // 对方返回错误
	Err_VerifySignRet EErrCode = 12 // 验证对方数据签名错误，用于需要请求回包签名校验的接口
	Err_HexDecode     EErrCode = 13 // 二进制转码错误
	Err_XmlDecode     EErrCode = 14 // xml解析错误
	Err_JsonDecode    EErrCode = 15 // json解析错误
	Err_ParamSign     EErrCode = 16 // 参数签名验证错误，用于需要参数签名校验的接口
	Err_NotSupport    EErrCode = 17 // 不支持的支付类型、方法调用
	Err_Timeout       EErrCode = 18 // 请求超时
	Err_SQLNoRows     EErrCode = 19 // SQL单行查询无数据

	Err_DelRefData      EErrCode = 20 // 不能删除还被使用的数据
	Err_RedisEmptyValue EErrCode = 21 // redis，key对应的值不存在
	Err_AsyncDoing      EErrCode = 22 // 请求处理中，请稍后查询
	Err_SQLInject       EErrCode = 23 // SQL注入
	Err_XSSInject       EErrCode = 24 // XSS注入
	Err_IllParam        EErrCode = 25 // 参数非法请求，利用接口非法请求不相关的数据。正常的页面不提供对应api功能
	Err_IllRequest      EErrCode = 26 // 代表利用接口越权操纵，比如利用自己的token，请求不能操作的数据。正常的页面不提供对应api功能
	Err_Perission       EErrCode = 27 // 用于页面按钮可点，但是点击会权限不足的报错的场景
	Err_Role            EErrCode = 28 // 角色不合法
	Err_HttpStatus      EErrCode = 29 // http状态异常

	Err_AuditState  EErrCode = 30 // 未审核
	Err_NoEntry     EErrCode = 31 // 服务未找到或不可用
	Err_FindService EErrCode = 32 // 未找到服务
	Err_VerService  EErrCode = 33 // 服务版本不支持
	Err_AccountLock EErrCode = 34 // 特指内部账号没锁
	Err_EntLock     EErrCode = 35 // 特指接入的三方接入账号被锁
	Err_UserLock    EErrCode = 36 // 特指接入的三方接入账号的一个用户被锁
	Err_NameIsExist EErrCode = 37 // 名称已存在，用于需要名称唯一的场景
	Err_Sensitivie  EErrCode = 38 // 含有敏感词

	//
	//  登录、验证码、短信
	//
	Err_VerifyCaptcha         EErrCode = 40  // 图形验证码错误
	Err_NotFoundUser          EErrCode = 41  // 没找到用户
	Err_CheckSMS              EErrCode = 42  // 校验短信失败
	Err_SMSVerifyCodeErrTimes EErrCode = 43  // 短信验证码错误次数太多，请重新发送短信验证码
	Err_CheckPhone            EErrCode = 44  // 手机号验证错误
	Err_CheckPsw              EErrCode = 45  // 登陆失败
	Err_OldPsw                EErrCode = 46  // 旧密码错误
	Err_PSWExpire             EErrCode = 47  // 密码过期
	Err_NewOldPswNotSame      EErrCode = 48  // 新密码不可以跟旧密码一样
	Err_SMSCoolDown           EErrCode = 49  // 短信验证码间隔时间未到
	Err_SMSLimitTimes         EErrCode = 50  // 短信验证码当天发送次数超限
	Err_SMSCodeTimeOut        EErrCode = 51  //验证码过期
	Err_LoginTooManyFail      EErrCode = 108 //登录错误次数太多
	Err_AccountHasBeenLocked  EErrCode = 109 //账号已被临时锁定

	//
	// token及请求校验
	//
	Err_ReqTimeNull   EErrCode = 52 // Head传参X-Unix-Time为空
	Err_ReqTimeFormat EErrCode = 53 // Head传参X-Unix-Time值，数据格式不合法
	Err_ReqTimeExpire EErrCode = 54 // Head传参X-Unix-Time值，过期
	Err_TokenNull     EErrCode = 55 // Head传参Authorization值, Token为空
	Err_TokenFormat   EErrCode = 56 // Token格式错误
	Err_TokenSM3      EErrCode = 57 // Token哈希校验
	Err_TokenRedis    EErrCode = 58 // redis不存在
	Err_TokenChange   EErrCode = 59 // token有新值
	Err_IPWhitelist   EErrCode = 60 // ip白名单
	Err_RandNull      EErrCode = 61 // Head传参Gz-Rand为空

	//
	// SM234加解密
	//
	Err_GenerateSM2Key  EErrCode = 63 // 生成SM2密钥失败
	Err_SM2KeyNull      EErrCode = 64 // SM2，key为空",
	Err_SM2Decrypt      EErrCode = 65 // SM2解密错误
	Err_SM2DataTooLarge EErrCode = 66 // 需要SM2解密的数据太大
	Err_SM3SignNull     EErrCode = 67 // Head传参X-Request-Head-Sign为空
	Err_SM3KeyNull      EErrCode = 68 // 系统设置的SM3 key为空
	Err_SM3SignCheck    EErrCode = 69 // SM3数据签名不一致
	Err_SM4KeyNull      EErrCode = 70 // SM4，key为空
	Err_SM4Decrypt      EErrCode = 71 // SM4解密错误
	Err_SM4Expire       EErrCode = 72 // SM4密钥过期

	//
	// 用户、企业管理(例如用户、企业添加，编辑）
	//
	Err_IDTypeOrNumber EErrCode = 75 // 证件类型或证件ID与数据库存储的不一致
	Err_UserIsExist    EErrCode = 76 // 用户名或用户已存在
	Err_UserNotExist   EErrCode = 77 // 用户名或用户不存在
	Err_EntIsExist     EErrCode = 78 // 企业已经存在
	Err_EntNotExist    EErrCode = 79 // 企业不存在

	Err_PhoneIsExist    EErrCode = 82 // 该手机号已经被使用
	Err_PhoneNotExist   EErrCode = 83 // 手机号不存在
	Err_RoleNameIsExist EErrCode = 84 // 该手机号已经被使用

	//
	// 文件功能
	//
	Err_UploadFile EErrCode = 85 // 上传文件失败
	Err_FileTooBig EErrCode = 86 // 文件太大
	Err_FileType   EErrCode = 87 // 文件类型错误
	Err_FileExpire EErrCode = 88 // 文件链接已过期

	//
	// 聊天功能
	//
	Err_UserIsBeBan      EErrCode = 100 //账户已被禁言过
	Err_InviteLinkNotCfg EErrCode = 101 //未配置邀请链接信息
	Err_MatchIsOver      EErrCode = 102
	//Error_SignedInToDay EErrCode = 100 //今日已签到过
	//Err_
	Err_IP_Locked   EErrCode = 103
	Err_USER_Locked EErrCode = 104

	ERR_Cloudflare_Lack_Token  EErrCode = 105
	ERR_Cloudflare_Verify_Err  EErrCode = 106
	ERR_Cloudflare_Verify_Fail EErrCode = 107
)

var Err_Oracle = Err_DB
var Err_PgDB = Err_DB
var Err_MySQL = Err_DB

var ErrMap = map[EErrCode][]string{

	No_Error:       {"操作成功", "ok", "", ""}, // 0
	Err_Unknown:    {"未知错误", "unknown error", "unknown error", "unknown error"},
	Err_System:     {"系统错误", "system error", "system error", "system error"},
	Err_Param:      {"参数错误", "param error", "param error", "param error"}, // 通用参数错误
	Err_Cookies:    {"参数错误", "param error", "param error", "param error"},
	Err_Redis:      {"服务异常", "service error", "service error", "service error"},
	Err_DB:         {"服务异常", "service error", "service error", "service error"},
	Err_Marshal:    {"数据异常", "data error", "data error", "data error"},
	Err_UnMarshal:  {"数据异常", "data error", "data error", "data error"},
	Err_RemoteCall: {"通信失败", "communication failure", "communication failure", "communication failure"},

	Err_HttpReadBody:  {"数据异常", "data error", "data error", "data error"},                                  // 10
	Err_RemoteReturn:  {"通信失败", "communication failure", "communication failure", "communication failure"}, // 11
	Err_VerifySignRet: {"通信失败", "communication failure", "communication failure", "communication failure"},
	Err_HexDecode:     {"数据异常", "data error", "data error", "data error"},
	Err_XmlDecode:     {"数据异常", "data error", "data error", "data error"},
	Err_JsonDecode:    {"数据异常", "data error", "data error", "data error"},
	Err_ParamSign:     {"签名错误", "signature error", "signature error", "signature error"},
	Err_NotSupport:    {"不支持的类型", "unsupported type", "unsupported type", "unsupported type"},
	Err_Timeout:       {"请求超时", "request timeout", "request timeout", "request timeout"},
	Err_SQLNoRows:     {"SQL单行查询无数据", "sql single row query no data", "sql single row query no data", "sql single row query no data"},

	Err_DelRefData: {"该数据被其他模块在使用。请先删除其他模块对该数据的引用",
		"this data is used by other modules. please delete the references to this data from other modules first.",
		"this data is used by other modules. please delete the references to this data from other modules first.",
		"this data is used by other modules. please delete the references to this data from other modules first."}, // 20
	Err_RedisEmptyValue: {"redis key对应的值不存在", "the value corresponding to the redis key does not exist",
		"the value corresponding to the redis key does not exist",
		"the value corresponding to the redis key does not exist"}, // 21
	Err_AsyncDoing: {"请求处理中，请稍后查询", "the request is being processed. please check again later.", "", ""},
	Err_SQLInject:  {"数据含有非法字符", "the data contains illegal characters", "", ""},                                // 内部使用，不直接对客户端使用，客户端应该发返回Err_Param
	Err_XSSInject:  {"数据含有非法字符", "the data contains illegal characters", "", ""},                                // 内部使用，不直接对客户端使用，客户端应该发返回Err_Param
	Err_IllParam:   {"参数错误", "param error", "", ""},                                                             // 代表其他的非法的参数，比如ID类的不存在，邮件，手机号正则不匹配
	Err_IllRequest: {"数据请求异常，请联系管理员", "the data request is abnormal, please contact the administrator", "", ""}, //  代表利用接口越权操纵，比如利用自己的token，请求不能操作的数据。实际页面不会限制
	Err_Perission:  {"权限不足", "insufficient permissions", "", ""},                                                // 用于页面按钮可点，但是点击会权限不足的报错的场景
	Err_Role:       {"角色不合法", "invalid role", "", ""},
	Err_HttpStatus: {"HTTP状态异常", "http status abnormal", "", ""},

	Err_AuditState:  {"未审核", "not audited", "", ""},                         // 30
	Err_NoEntry:     {"服务暂时不可用", "service temporarily unavailable", "", ""}, // 31
	Err_FindService: {"未找到服务", "service not found", "", ""},
	Err_VerService:  {"服务版本不支持", "service version not supported", "", ""},
	Err_AccountLock: {"账号锁定", "account locked", "", ""},                                                // 特指内部账号没锁
	Err_EntLock:     {"接入账号被锁定，请联系管理员", "access account locked, please contact administrator", "", ""}, // 特指接入的三方接入账号被锁
	Err_UserLock:    {"登录账号被锁定，请联系管理员", "login account locked, please contact administrator", "", ""},  // 特指接入的三方接入账号的一个用户被锁
	Err_NameIsExist: {"名称已存在", "name already exists", "", ""},                                          //  37 用于需要名称唯一的场景
	Err_Sensitivie:  {"您的发言违规", "Your speech contains illegal words", "Your speech contains illegal words", ""},
	//
	//  登录、验证码、短信
	//
	Err_VerifyCaptcha:         {"验证码错误", "incorrect captcha", "incorrect captcha", ""},                                                              // 40
	Err_NotFoundUser:          {"账号或手机号或密码错误", "incorrect account or phone number or password", "ncorrect account or phone number or password", ""}, // 特指登录时候的报错，用户名不存在或者密码错误
	Err_CheckSMS:              {"校验短信失败", "sms verification failed", "sms verification failed", ""},
	Err_SMSVerifyCodeErrTimes: {"短信验证码错误次数太多，请重新发送短信验证码", "too many sms verification code errors, please resend sms verification code", "too many sms verification code errors, please resend sms verification code", ""},
	Err_CheckPhone:            {"手机号验证错误", "phone number verification error", "phone number verification error", ""},
	Err_CheckPsw:              {"账号或手机号或密码错误", "incorrect account or phone number or password", "incorrect account or phone number or password", ""}, // 用于非登录场景的密码校验，比如关键操作需要二次密码验证
	Err_OldPsw:                {"原始密码错误", "original password error", "original password error", ""},
	Err_PSWExpire:             {"密码已过期", "password expired", "password expired", ""},
	Err_NewOldPswNotSame:      {"新密码不可以跟旧密码一样", "new password cannot be the same as old password", "new password cannot be the same as old password", ""},
	Err_SMSCoolDown:           {"短信发送间隔时间未到", "sms sending interval not reached", "sms sending interval not reached", ""},                    // 49
	Err_SMSLimitTimes:         {"今天累计发送次数超限", "today's cumulative sending times exceeded", "today's cumulative sending times exceeded", ""},  // 50
	Err_SMSCodeTimeOut:        {"发送验证码超时", "The Sms code Send Timeout ", "", ""},                                                             // 50
	Err_LoginTooManyFail:      {"该IP登录失败次数过多", "The IP has too many failed login attempts", "The IP has too many failed login attempts", ""}, // 50
	Err_AccountHasBeenLocked:  {"账号已被临时锁定", "The account has been temporarily locked", "The account has been temporarily locked", ""},
	//
	// token及请求校验
	//
	Err_ReqTimeNull:   {"Head传参X-Unix-Time为空", "head parameter x-unix-time is empty", "", ""}, // 52
	Err_ReqTimeFormat: {"Head传参X-Unix-Time值，数据格式不合法", "head parameter x-unix-time value, data format is invalid", "", ""},
	Err_ReqTimeExpire: {"Head传参X-Unix-Time值，过期", "head parameter x-unix-time value expired", "", ""},
	Err_TokenNull:     {"Head传参Authorization值, Token为空", "head parameter authorization value, token is empty", "", ""},
	Err_TokenFormat:   {"Token格式错误", "token format error", "", ""}, // 56
	Err_TokenSM3:      {"Token哈希校验", "token hash verification", "", ""},
	Err_TokenRedis:    {"Token失效", "token invalid", "", ""},
	Err_TokenChange:   {"token有新值", "token has new value", "", ""},
	Err_IPWhitelist:   {"你没有访问权限", "You do not have access permission ", "You do not have access permission", ""}, // 60
	Err_RandNull:      {"Head传参X-Rand为空", "head parameter x-rand is empty", "", ""},

	//
	// SM234加解密
	//
	Err_GenerateSM2Key:  {"生成SM2密钥失败", "failed to generate sm2 key", "", ""}, // 63
	Err_SM2KeyNull:      {"SM2，key为空", "sm2 key is empty", "", ""},
	Err_SM2Decrypt:      {"数据错误", "data error", "", ""}, // SM2解密错误
	Err_SM2DataTooLarge: {"数据错误", "data error", "", ""}, // 需要SM2解密的数据太大
	Err_SM3SignNull:     {"Head传参X-Request-Head-Sign为空", "head parameter x-request-head-sign is empty", "", ""},
	Err_SM3KeyNull:      {"系统设置的SM3 key为空", "system set sm3 key is empty", "", ""},
	Err_SM3SignCheck:    {"SM3数据签名不一致", "sm3 data signature inconsistency", "", ""},
	Err_SM4KeyNull:      {"SM4，key为空", "sm4 key is empty", "", ""}, // 70
	Err_SM4Decrypt:      {"SM4解密错误", "sm4 decryption error", "", ""},
	Err_SM4Expire:       {"SM4密钥过期", "sm4 key expired", "", ""}, // 72

	//
	// 用户、企业管理(例如用户、企业添加，编辑）
	//
	Err_IDTypeOrNumber: {"证件类型或证件ID与数据库存储的不一致", "id type or id number is inconsistent with database storage", "", ""}, // 75
	Err_UserIsExist:    {"用户名或用户已存在", "username or user already exists", "username or user already exists", ""},       // 76
	Err_UserNotExist:   {"用户名或用户不存在", "username or user does not exist", "username or user already exists", ""},
	Err_EntIsExist:     {"企业已经存在", "enterprise already exists", "", ""},
	Err_EntNotExist:    {"企业不存在", "enterprise does not exist", "", ""},

	Err_PhoneIsExist:    {"该手机号已经被使用", "this phone number is already in use", "this phone number is already in use", ""},
	Err_PhoneNotExist:   {"手机号不存在", "phone number does not exist", "phone number does not exist", ""}, // 83
	Err_RoleNameIsExist: {"该角色名称已经被使用", "this role name is already in use", "this role name is already in use", ""},

	//
	// 文件功能
	//
	Err_UploadFile: {"上传文件失败", "failed to upload file", "", ""}, // 85
	Err_FileTooBig: {"文件太大", "file too large", "", ""},
	Err_FileType:   {"文件类型错误", "file type error", "", ""},
	Err_FileExpire: {"文件链接已过期，请重新刷新再页面", "file link has expired, please refresh the page", "", ""}, // 88

	Err_UserIsBeBan:      {"用户被禁言", "user banned", "user banned", ""},                        //账户已被禁言过
	Err_InviteLinkNotCfg: {"未配置邀请链接信息", "invite link information is not configured", "", ""}, //未配置邀请链接信息
	Err_MatchIsOver:      {"比赛已结束", "match is over", "match is over", ""},                    //比赛已结束
	Err_IP_Locked:        {"登录错误次数太多，ip已经被锁定", "Ip is locked", "Ip is locke", ""},
	Err_USER_Locked:      {"登录错误次数太多，用户已经被锁定", "User is locked", "User is locked", ""},

	ERR_Cloudflare_Lack_Token:  {"Missing Turnstile token", "Missing Turnstile token", "Missing Turnstile token", ""},
	ERR_Cloudflare_Verify_Err:  {"Turnstile verification error", "Turnstile verification error", "Turnstile verification error", ""},
	ERR_Cloudflare_Verify_Fail: {"Turnstile verification failed", "Turnstile verification failed", "Turnstile verification failed", ""},
}

// func AddErrDef(errMap map[EErrCode]string) {
// 	for k, v := range errMap {
// 		ErrMap[k] = v
// 	}
// }

// Accept-Language
type ELanDef = int

const (
	ELan_zh_CN = 0 //  zh-CN 中国
	ELan_en_US = 1 //  en-US 美国
	ELan_tr_TR = 2 //  tr-TR 土耳其
	ELan_ar_EG = 3 //	ar-EG 埃及
	ELan_Max   = ELan_ar_EG
)

type CXXResult struct {
	errMap []map[EErrCode][]byte
}

type THttpResp struct {
	Code      EErrCode    `json:"code" mapstructure:"code" msgpack:"code"`
	Msg       string      `json:"msg" mapstructure:"msg" msgpack:"msg"`
	ExtraInfo string      `json:"extraInfo,omitempty" mapstructure:"extraInfo,omitempty" msgpack:"extraInfo,omitempty"`
	Data      interface{} `json:"data,omitempty" mapstructure:"data,omitempty" msgpack:"data,omitempty"`
}

//type THttpResp = THttpResponse

type TJsonResult struct {
	Code      EErrCode `json:"code"`
	Msg       string   `json:"msg"`
	ExtraInfo string   `json:"extraInfo,omitempty"`
}

type TPage struct {
	List  []interface{} `json:"list"`
	Total int32         `json:"total"`
}

type TPageXX[T any] struct {
	List  []T   `json:"list"`
	Total int32 `json:"total"`
}

func NewXXResult(
	errTable map[EErrCode][]string,
	initFun func(EErrCode, string) interface{},
	convFun func(interface{}) ([]byte, error)) *CXXResult {

	tmp := make([]map[EErrCode][]byte, ELan_Max+1, ELan_Max+1)
	tmp[ELan_zh_CN] = make(map[EErrCode][]byte)
	tmp[ELan_en_US] = make(map[EErrCode][]byte)
	tmp[ELan_tr_TR] = make(map[EErrCode][]byte)
	tmp[ELan_ar_EG] = make(map[EErrCode][]byte)
	for k, vv := range ErrMap {
		for i, v := range vv {
			res, err := convFun(initFun(k, v))
			if nil != err {
				panic("CXXResult Init failed!")
			}

			tmp[i][k] = res
		}
	}

	return &CXXResult{errMap: tmp}
}

func GetLanID(accLan string) (lanID ELanDef) {
	accLan = strings.Split(accLan, ",")[0] // 只取第一个
	switch accLan {
	case "zh-CN":
		lanID = ELan_zh_CN
	case "en-US":
		lanID = ELan_en_US
	case "tr-TR":
		lanID = ELan_tr_TR
	case "ar-EG", "ar-AR":
		lanID = ELan_ar_EG
	default:
		lanID = ELan_en_US
	}
	return
}
func GetErrStr(accLan string, ec EErrCode) string {
	lanID := GetLanID(accLan)
	reply, ok := ErrMap[ec]
	if ok {
		return reply[lanID]
	} else {
		return ErrMap[Err_Unknown][lanID]
	}
}

func (this *CXXResult) Get(accLan string, errCode EErrCode) []byte {
	lanID := GetLanID(accLan)
	errMap := this.errMap[lanID]
	reply, ok := errMap[errCode]
	if ok {
		return reply
	} else {
		//Warning("CXXResult:Get(%d)", errCode)
		//return this.Get(Err_Unknown)
		return errMap[Err_Unknown]
	}
}

// 代码扫描有误报，先注释把
func (this *CXXResult) ConvertJsonFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	file.WriteString("//自动生成，不要手动修改！\n")
	file.WriteString("var Result = {\n")
	for code, msg := range ErrMap {
		file.WriteString(fmt.Sprintf("\t%d: %q,\n", code, msg))
	}
	file.WriteString("};")
	file.Close()
	return nil
}

var JsonResult *CXXResult = nil

func InitJsonResult() {
	JsonResult = NewXXResult(
		ErrMap,
		func(errCode EErrCode, errStr string) interface{} { return &TJsonResult{Code: errCode, Msg: errStr} },
		json.Marshal)
}

func JsonMarshal(data interface{}) []byte {
	jsonData, _ := json.Marshal(data)
	ret := fmt.Sprintf(`{"code":0,"msg":"操作成功","data":%s}`, string(jsonData))
	return []byte(ret)
}

func NewJsonResult(accLan string, errCode EErrCode, extraInfo string) []byte {
	errMsg := GetErrStr(accLan, errCode)
	extra, _ := json.Marshal(extraInfo)

	ret := fmt.Sprintf(
		`{"code":%d,"msg":"%s","extraInfo":%s}`,
		errCode, errMsg, string(extra))
	return []byte(ret)
}

func JsonErrTest(err error, errCode *EErrCode) {
	if err != nil {
		*errCode = Err_JsonDecode
		// Todo: Log
	}
	return
}

func DBErrTest(err error, errCode *EErrCode) {
	if err != nil {
		*errCode = Err_DB
		// Todo: Log
	}
	return
}

func RedisErrTest(err error, errCode *EErrCode) {
	if err != nil {
		*errCode = Err_Redis
		// Todo: Log
	}
	return
}

func init() {
	InitJsonResult()
}
