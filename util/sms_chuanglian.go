package util

import (
	"bytes"
	. "common"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type SmsSendRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Msg      string `json:"msg"`
}

func SendSmsCL(smsRequest SmsSendRequest) (_err EErrCode) {
	//password := smsRequest.Password
	//password := "j1HVgr313Of24f"
	nonce := fmt.Sprintf("%d", time.Now().UnixMilli())

	// 构造签名
	paramMap := structToSortedMap(smsRequest)
	paramMap["nonce"] = nonce
	sign := generateSign(smsRequest.Password, paramMap)

	// 构造请求体
	requestBody, err := json.Marshal(smsRequest)
	if err != nil {
		return Err_Param
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", "https://sgap.253.com/send/sms", bytes.NewBuffer(requestBody))
	if err != nil {
		return Err_RemoteCall
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("nonce", nonce)
	req.Header.Set("sign", sign)

	resp, err := client.Do(req)
	if err != nil {
		return Err_RemoteCall
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	fmt.Println("响应：", string(bodyBytes))

	if result["code"] == "0" {
		if data, ok := result["data"].(map[string]interface{}); ok {
			fmt.Println("messageId =", data["messageId"])
		}
	} else {
		return Err_RemoteCall
	}
	return
}

// 将结构体转为按 key 排序的 map
func structToSortedMap(req SmsSendRequest) map[string]string {
	m := make(map[string]string)
	m["account"] = req.Account
	m["mobile"] = req.Mobile
	m["msg"] = req.Msg
	return m
}

// 生成签名
func generateSign(password string, paramMap map[string]string) string {
	var keys []string
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		v := paramMap[k]
		if strings.TrimSpace(v) != "" {
			sb.WriteString(k)
			sb.WriteString(v)
		}
	}
	sb.WriteString(password)

	hash := md5.Sum([]byte(sb.String()))
	return hex.EncodeToString(hash[:])
}
