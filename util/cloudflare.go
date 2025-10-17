package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const TurnstileSecret = "0x4AAAAAABm3UoHQ5D-JNNPQSmBZQ6UWRAs" // 替换成你的 Secret

func VerifyTurnstile(token, remoteIP string) (bool, error) {
	data := map[string]string{
		"secret":   TurnstileSecret,
		"response": token,
		"remoteip": remoteIP,
	}
	body, _ := json.Marshal(data)

	resp, err := http.Post("https://challenges.cloudflare.com/turnstile/v0/siteverify", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)
	var result struct {
		Success bool `json:"success"`
	}
	if err := json.Unmarshal(resBody, &result); err != nil {
		return false, err
	}

	return result.Success, nil
}
