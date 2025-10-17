package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "AIzaSyAk2RgMjOz5m262cK6S5UHArJ1f8TesXL0" // ⚠️ 在这里替换成你申请的 API Key

// Translate 调用 Google Cloud Translation API 翻译文本
// text: 待翻译的文字
// targetLang: 目标语言代码，例如 "en"、"zh-CN"、"tr"
func Translate(text, targetLang string) (string, error) {
	url := "https://translation.googleapis.com/language/translate/v2?key=" + apiKey

	// 构造请求体
	body := map[string]interface{}{
		"q":      text,
		"target": targetLang,
	}
	jsonBody, _ := json.Marshal(body)

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应
	data, _ := ioutil.ReadAll(resp.Body)

	// 解析 JSON
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 提取翻译结果
	if trans, ok := result["data"].(map[string]interface{}); ok {
		if translations, ok := trans["translations"].([]interface{}); ok && len(translations) > 0 {
			if translatedText, ok := translations[0].(map[string]interface{})["translatedText"].(string); ok {
				return translatedText, nil
			}
		}
	}

	return "", fmt.Errorf("翻译失败: %s", string(data))
}
