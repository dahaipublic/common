package util

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20180501 "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

func CreateClient(req ALiSmsRequest) (_result *dysmsapi20180501.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(req.AccessKeyID),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(req.AccessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.eu-central-1.aliyuncs.com")
	_result = &dysmsapi20180501.Client{}
	_result, _err = dysmsapi20180501.NewClient(config)
	return _result, _err
}

type ALiSmsRequest struct {
	Phone           string `json:"-"`
	Content         string `json:"content"`
	SenderId        string `json:"-" `
	AccessKeyID     string `json:"-"`
	AccessKeySecret string `json:"-"`
}

func SendMessageByALi(req ALiSmsRequest) (_err error) {
	client, _err := CreateClient(req)
	if _err != nil {
		return _err
	}

	sendMessageToGlobeRequest := &dysmsapi20180501.SendMessageToGlobeRequest{
		To:      tea.String(req.Phone),
		Message: tea.String(req.Content),
		TaskId:  tea.String(req.SenderId),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// Copy the code to run, please print the return value of the API by yourself.
		_, _err = client.SendMessageToGlobeWithOptions(sendMessageToGlobeRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// Only a printing example. Please be careful about exception handling and do not ignore exceptions directly in engineering projects.
		// print error message
		fmt.Println(tea.StringValue(error.Message))
		// Please click on the link below for diagnosis.
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
