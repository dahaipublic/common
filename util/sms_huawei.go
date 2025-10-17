package util

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HuaWeiSmsRequest struct {
	Phone      string `json:"-"`
	Content    string `json:"content"`
	SenderId   string `json:"-" `
	ApiKey     string `json:"-"`
	ApiPwd     string `json:"-"`
	ApiUrl     string `json:"-"`
	TemplateId string `json:"-"`
	Signature  string `json:"-"`
}

// HuaWeiSmsResponse 定义结构体
type HuaWeiSmsResponse struct {
	Result      []ResultItem `json:"result"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
}

type ResultItem struct {
	OriginTo   string `json:"originTo"`
	CreateTime string `json:"createTime"`
	From       string `json:"from"`
	SmsMsgID   string `json:"smsMsgId"`
	Status     string `json:"status"`
	CountryID  string `json:"countryId"`
	Total      int    `json:"total"`
}

// WSSE_HEADER_FORMAT 无需修改,用于格式化鉴权头域,给"X-WSSE"参数赋值
const WSSE_HEADER_FORMAT = "UsernameToken Username=\"%s\",PasswordDigest=\"%s\",Nonce=\"%s\",Created=\"%s\""

// AUTH_HEADER_VALUE 无需修改,用于格式化鉴权头域,给"Authorization"参数赋值
const AUTH_HEADER_VALUE = "WSSE realm=\"SDP\",profile=\"UsernameToken\",type=\"Appkey\""

func SendSmsByHuaWei(request HuaWeiSmsRequest) (*HuaWeiSmsResponse, error) {
	//必填,请参考"开发准备"获取如下数据,替换为实际值
	apiAddress := fmt.Sprintf("%v/sms/batchSendSms/v1", request.ApiUrl) //"https://smsapi.ap-southeast-1.myhuaweicloud.com:443/sms/batchSendSms/v1" //APP接入地址(在控制台"应用管理"页面获取)+接口访问URI
	// 认证用的appKey和appSecret硬编码到代码中或者明文存储都有很大的安全风险，建议在配置文件或者环境变量中密文存放，使用时解密，确保安全；
	appKey := request.ApiKey         //APP_Key
	appSecret := request.ApiPwd      //APP_Secret
	sender := request.SenderId       //国内短信签名通道号或国际/港澳台短信通道号
	templateId := request.TemplateId //模板ID

	//条件必填,国内短信关注,当templateId指定的模板类型为通用模板时生效且必填,必须是已审核通过的,与模板类型一致的签名名称
	//国际/港澳台短信不用关注该参数
	signature := request.Signature //签名名称

	//必填,全局号码格式(包含国家码),示例:+8615123456789,多个号码之间用英文逗号分隔
	receiver := request.Phone //短信接收人号码

	//选填,短信状态报告接收地址,推荐使用域名,为空或者不填表示不接收状态报告
	statusCallBack := ""

	/*
	 * 选填,使用无变量模板时请赋空值 string templateParas = "";
	 * 单变量模板示例:模板内容为"您的验证码是${1}"时,templateParas可填写为"[\"369751\"]"
	 * 双变量模板示例:模板内容为"您有${1}件快递请到${2}领取"时,templateParas可填写为"[\"3\",\"人民公园正门\"]"
	 * 模板中的每个变量都必须赋值，且取值不能为空
	 * 查看更多模板和变量规范
	 */
	templateParas := fmt.Sprintf("[\"%v\"]", request.Content) //模板变量，此处以单变量验证码短信为例，请客户自行生成6位验证码，并定义为字符串类型，以杜绝首位0丢失的问题（例如：002569变成了2569）。

	body := buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Authorization"] = AUTH_HEADER_VALUE
	headers["X-WSSE"] = buildWsseHeader(appKey, appSecret)

	resp, err := post(apiAddress, []byte(body), headers)
	if err != nil {
		return nil, err
	}
	result := &HuaWeiSmsResponse{}
	if err := json.Unmarshal([]byte(resp), result); err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * sender,receiver,templateId不能为空
 */
func buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature string) string {
	param := "from=" + url.QueryEscape(sender) + "&to=" + url.QueryEscape(receiver) + "&templateId=" + url.QueryEscape(templateId)
	if templateParas != "" {
		param += "&templateParas=" + url.QueryEscape(templateParas)
	}
	if statusCallBack != "" {
		param += "&statusCallback=" + url.QueryEscape(statusCallBack)
	}
	if signature != "" {
		param += "&signature=" + url.QueryEscape(signature)
	}
	return param
}

func post(url string, param []byte, headers map[string]string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func buildWsseHeader(appKey, appSecret string) string {
	var cTime = time.Now().Format("2006-01-02T15:04:05Z")
	var nonce = uuid.New().String()
	nonce = strings.ReplaceAll(nonce, "-", "")

	h := sha256.New()
	h.Write([]byte(nonce + cTime + appSecret))
	passwordDigestBase64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf(WSSE_HEADER_FORMAT, appKey, passwordDigestBase64Str, nonce, cTime)
}
