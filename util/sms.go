package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 聚合短信验证码
func JuHeSmsSend(phone, code string) (err error) {
	// 接口请求URL
	apiUrl := "http://v.juhe.cn/sms/send"

	// 初始化参数
	param := url.Values{}

	// 接口请求参数
	param.Set("mobile", phone)                           // 接收短信的手机号码
	param.Set("tpl_id", "257159")                        // 短信模板ID，请参考个人中心短信模板设置
	param.Set("tpl_value", "#code#="+code)               // 模板变量，如无则不用填写
	param.Set("key", "50379e98a88008bbd516553afa0b5238") // 接口请求Key

	// 发送请求
	data, err := Post(apiUrl, param)
	if err != nil {
		// 请求异常，根据自身业务逻辑进行调整修改
		err = errors.New(fmt.Sprintf("请求异常:\r\n%v", err))
		return
	}
	var netReturn map[string]interface{}
	jsonerr := json.Unmarshal(data, &netReturn)
	if jsonerr != nil {
		// 解析JSON异常，根据自身业务逻辑进行调整修改
		err = errors.New(fmt.Sprintf("请求异常:%v", jsonerr))
		return
	}
	errorCode := netReturn["error_code"]
	reason := netReturn["reason"]
	if errorCode.(float64) != 0 {
		// 查询失败，根据自身业务逻辑进行调整修改
		err = errors.New(fmt.Sprintf("请求失败:%v_%v", errorCode.(float64), reason))
		return
	}
	return
}

// post 方式发起网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 聚合国际短信验证码
func JuHeInternationalSmsSend(areaNum, phone, code string) (err error) {
	// 接口请求URL
	apiUrl := "http://v.juhe.cn/smsInternational/send"

	// 初始化参数
	param := url.Values{}

	// 接口请求参数
	param.Set("areaNum", areaNum)                        // 模板变量，如无则不用填写
	param.Set("mobile", phone)                           // 接收短信的手机号码
	param.Set("tplId", "12864")                          // 短信模板ID，请参考个人中心短信模板设置
	param.Set("tplValue", "#code#="+code)                // 模板变量，如无则不用填写
	param.Set("key", "50379e98a88008bbd516553afa0b5238") // 接口请求Key

	// 发送请求
	data, err := Post(apiUrl, param)
	if err != nil {
		// 请求异常，根据自身业务逻辑进行调整修改
		err = errors.New(fmt.Sprintf("请求异常:\r\n%v", err))
		return
	}
	var netReturn map[string]interface{}
	jsonerr := json.Unmarshal(data, &netReturn)
	if jsonerr != nil {
		// 解析JSON异常，根据自身业务逻辑进行调整修改
		err = errors.New(fmt.Sprintf("请求异常:%v", jsonerr))
		return
	}
	errorCode := netReturn["error_code"]
	reason := netReturn["reason"]
	if errorCode.(float64) != 0 {
		// 查询失败，根据自身业务逻辑进行调整修改
		err = errors.New(fmt.Sprintf("请求失败:%v_%v", errorCode.(float64), reason))
		return
	}
	return
}

// // 梦云短信验证码
//
//	func MonYunSmsSend(phone, code string) (err error) {
//		// 接口请求URL
//		apiUrl := "http://api01.monyun.cn:7901/sms/v2/std/single_send"
//
//		// 要转换和编码的 UTF-8 字符串
//		utf8Str := "您的验证码是" + code + "，在5分钟内有效。如非本人操作请忽略本短信。" // 以UTF-8编码的字符串
//
//		// 将 UTF-8 字符串转换为 GBK 编码
//		gbkStr, err := utf8ToGBK(utf8Str)
//		if err != nil {
//			log.Fatalf("Error converting UTF-8 to GBK: %v", err)
//		}
//
//		// 对 GBK 编码的字符串进行 URL 编码
//		encodedStr := url.QueryEscape(gbkStr)
//		// 初始化参数
//		var param struct {
//			Apikey  string `json:"apikey"`
//			Mobile  string `json:"mobile"`
//			Content string `json:"content"`
//		}
//
//		// 接口请求参数
//		param.Apikey = "25d82aa430340bb8bf99164984fe7e6f"
//		param.Mobile = phone
//		param.Content = encodedStr
//		var netReturn map[string]interface{}
//		//客户端句柄
//		client := xhttp.NewClient(
//			&xhttp.ClientConfig{
//				Dial:      60 * xtime.Duration(time.Second),
//				Timeout:   60 * xtime.Duration(time.Second),
//				KeepAlive: 120 * xtime.Duration(time.Second),
//			})
//		if err = client.PostJson(context.Background(), apiUrl, param, map[string]interface{}{}, &netReturn); err != nil {
//			return
//		}
//		log.Println("MonYunSmsSend:", netReturn)
//		errorCode := netReturn["result"]
//		reason := netReturn["desc"]
//		if errorCode.(float64) != 0 {
//			// 查询失败，根据自身业务逻辑进行调整修改
//			err = errors.New(fmt.Sprintf("请求失败:%v_%v", errorCode.(float64), reason))
//			return
//		}
//		return
//	}
//type HuaWeiSmsRequest struct {
//	Phone      string `json:"-"`
//	Content    string `json:"content"`
//	SenderId   string `json:"-" `
//	ApiKey     string `json:"-"`
//	ApiPwd     string `json:"-"`
//	ApiUrl     string `json:"-"`
//	TemplateId string `json:"-"`
//	Signature  string `json:"-"`
//}
//
//// HuaWeiSmsResponse 定义结构体
//type HuaWeiSmsResponse struct {
//	Result      []ResultItem `json:"result"`
//	Code        string       `json:"code"`
//	Description string       `json:"description"`
//}
//
//type ResultItem struct {
//	OriginTo   string `json:"originTo"`
//	CreateTime string `json:"createTime"`
//	From       string `json:"from"`
//	SmsMsgID   string `json:"smsMsgId"`
//	Status     string `json:"status"`
//	CountryID  string `json:"countryId"`
//	Total      int    `json:"total"`
//}
//
//// WSSE_HEADER_FORMAT 无需修改,用于格式化鉴权头域,给"X-WSSE"参数赋值
//const WSSE_HEADER_FORMAT = "UsernameToken Username=\"%s\",PasswordDigest=\"%s\",Nonce=\"%s\",Created=\"%s\""
//
//// AUTH_HEADER_VALUE 无需修改,用于格式化鉴权头域,给"Authorization"参数赋值
//const AUTH_HEADER_VALUE = "WSSE realm=\"SDP\",profile=\"UsernameToken\",type=\"Appkey\""
//
//func HuaweiYunSmsSend(request HuaWeiSmsRequest) (*HuaWeiSmsResponse, error) {
//	apiAddress := fmt.Sprintf("%v/sms/batchSendSms/v1", request.ApiUrl) //"https://smsapi.ap-southeast-1.myhuaweicloud.com:443/sms/batchSendSms/v1" //APP接入地址(在控制台"应用管理"页面获取)+接口访问URI
//	// 认证用的appKey和appSecret硬编码到代码中或者明文存储都有很大的安全风险，建议在配置文件或者环境变量中密文存放，使用时解密，确保安全；
//	appKey := request.ApiKey         //APP_Key
//	appSecret := request.ApiPwd      //APP_Secret
//	sender := request.SenderId       //国内短信签名通道号或国际/港澳台短信通道号
//	templateId := request.TemplateId //模板ID
//
//	//条件必填,国内短信关注,当templateId指定的模板类型为通用模板时生效且必填,必须是已审核通过的,与模板类型一致的签名名称
//	//国际/港澳台短信不用关注该参数
//	signature := request.Signature //签名名称
//
//	//必填,全局号码格式(包含国家码),示例:+8615123456789,多个号码之间用英文逗号分隔
//	receiver := request.Phone //短信接收人号码
//
//	//选填,短信状态报告接收地址,推荐使用域名,为空或者不填表示不接收状态报告
//	statusCallBack := ""
//	/*
//	 * 选填,使用无变量模板时请赋空值 string templateParas = "";
//	 * 单变量模板示例:模板内容为"您的验证码是${1}"时,templateParas可填写为"[\"369751\"]"
//	 * 双变量模板示例:模板内容为"您有${1}件快递请到${2}领取"时,templateParas可填写为"[\"3\",\"人民公园正门\"]"
//	 * 模板中的每个变量都必须赋值，且取值不能为空
//	 * 查看更多模板和变量规范
//	 */
//	templateParas := fmt.Sprintf("[\"%v\"]", request.Content) //模板变量，此处以单变量验证码短信为例，请客户自行生成6位验证码，并定义为字符串类型，以杜绝首位0丢失的问题（例如：002569变成了2569）。
//
//	body := buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature)
//	headers := make(map[string]string)
//	headers["Content-Type"] = "application/x-www-form-urlencoded"
//	headers["Authorization"] = AUTH_HEADER_VALUE
//	headers["X-WSSE"] = buildWsseHeader(appKey, appSecret)
//
//	resp, err := post(apiAddress, []byte(body), headers)
//	if err != nil {
//		//logx.Errorf("SendSmsByHuaWei post phone:%v,err:%v", request.Phone, err.Error())
//		return nil, err
//	}
//	result := &HuaWeiSmsResponse{}
//	if err := json.Unmarshal([]byte(resp), result); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
///**
// * sender,receiver,templateId不能为空
// */
//func buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature string) string {
//	param := "from=" + url.QueryEscape(sender) + "&to=" + url.QueryEscape(receiver) + "&templateId=" + url.QueryEscape(templateId)
//	if templateParas != "" {
//		param += "&templateParas=" + url.QueryEscape(templateParas)
//	}
//	if statusCallBack != "" {
//		param += "&statusCallback=" + url.QueryEscape(statusCallBack)
//	}
//	if signature != "" {
//		param += "&signature=" + url.QueryEscape(signature)
//	}
//	return param
//}
//func post(url string, param []byte, headers map[string]string) (string, error) {
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	client := &http.Client{Transport: tr}
//
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
//	if err != nil {
//		return "", err
//	}
//	for key, header := range headers {
//		req.Header.Set(key, header)
//	}
//
//	resp, err := client.Do(req)
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//	return string(body), nil
//}
//
//func buildWsseHeader(appKey, appSecret string) string {
//	var cTime = time.Now().Format("2006-01-02T15:04:05Z")
//	var nonce = uuid.New().String()
//	nonce = strings.ReplaceAll(nonce, "-", "")
//
//	h := sha256.New()
//	h.Write([]byte(nonce + cTime + appSecret))
//	passwordDigestBase64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))
//
//	return fmt.Sprintf(WSSE_HEADER_FORMAT, appKey, passwordDigestBase64Str, nonce, cTime)
//}
//
//func utf8ToGBK(utf8Str string) (string, error) {
//	utf8Bytes := []byte(utf8Str)
//	reader := transform.NewReader(strings.NewReader(string(utf8Bytes)), simplifiedchinese.GBK.NewEncoder())
//	gbkBytes, err := ioutil.ReadAll(reader)
//	if err != nil {
//		return "", err
//	}
//	return string(gbkBytes), nil
//}
