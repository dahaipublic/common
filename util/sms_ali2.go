package util

import (
	"encoding/json"
	"fmt"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20180501 "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// Description:
//
// 使用凭据初始化账号Client
//
// @return Client
//
// @throws Exception
func CreateClient2() (_result *dysmsapi20180501.Client, _err error) {
	// 工程代码建议使用更安全的无AK方式，凭据配置方式请参见：https://help.aliyun.com/document_detail/378661.html。
	//credential, _err := credential.NewCredential(nil)
	//if _err != nil {
	//	return _result, _err
	//}

	config := &openapi.Config{
		AccessKeyId: tea.String(""),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(""),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.eu-central-1.aliyuncs.com")
	_result = &dysmsapi20180501.Client{}
	_result, _err = dysmsapi20180501.NewClient(config)
	return _result, _err
}

func SendALiSmS2(to, message string) (_err error) {
	client, _err := CreateClient2()
	if _err != nil {
		return _err
	}

	//batchSendMessageToGlobeRequest := &dysmsapi20180501.BatchSendMessageToGlobeRequest{
	//	To:      tea.String(to),
	//	Message: tea.String(message),
	//}
	sendMessageToGlobeRequest := &dysmsapi20180501.SendMessageToGlobeRequest{
		To:      tea.String(to),
		Message: tea.String(message),
		//TaskId:  tea.String(req.SenderId),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		response1, _err := client.SendMessageToGlobeWithOptions(sendMessageToGlobeRequest, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}
		fmt.Println(response1)

		return nil
	}()
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
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
