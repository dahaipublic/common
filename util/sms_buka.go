package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dahaipublic/common"
	"io"
	"net/http"

	"time"
)

type BukaSmsRequest struct {
	ApiKey  string `json:"-"`
	ApiPwd  string `json:"-"`
	AppId   string `json:"-" `
	Numbers string `json:"-"`
	Content string `json:"content"`
}

func SendSmsBuka(smsRequest BukaSmsRequest) (errCode common.EErrCode) {
	baseUrl := "https://api.laaffic.com/v3"

	url := baseUrl + "/sendSms"
	currentTime := time.Now().Unix()
	sign := generateMd5Sign(smsRequest.ApiKey, smsRequest.ApiPwd, currentTime)

	// 请求体参数
	body := map[string]interface{}{
		"appId":   smsRequest.AppId,
		"numbers": smsRequest.Numbers,
		"content": smsRequest.Content,
	}

	// 请求头
	headers := map[string]string{
		"Connection":   "Keep-Alive",
		"Content-Type": "application/json;charset=UTF-8",
		"Sign":         sign,
		"Timestamp":    fmt.Sprintf("%d", currentTime),
		"Api-Key":      smsRequest.ApiKey,
	}

	respBody, statusCode, err := doPostJson(url, headers, body)
	if err != nil {
		errCode = common.Err_Param
		return
	}

	if statusCode == http.StatusOK {
		fmt.Println(string(respBody))
	} else {
		errCode = common.Err_RemoteCall
	}
	return
}

func doPostJson(url string, headers map[string]string, body interface{}) ([]byte, int, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, 0, fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("error reading response: %w", err)
	}

	return respBody, resp.StatusCode, nil
}

func generateMd5Sign(apiKey, apiPwd string, datetime int64) string {
	data := apiKey + apiPwd + fmt.Sprintf("%d", datetime)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
