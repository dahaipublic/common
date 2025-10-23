package nami

import (
	"fmt"
	"github.com/dahaipublic/common/database"
	"github.com/dahaipublic/common/model"
	"strconv"
	"strings"
	"time"
)

func MatchLiveArrayToString(arr []interface{}) string {
	var strArr []string
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			strArr = append(strArr, strconv.Itoa(val))
		case float64:
			strArr = append(strArr, fmt.Sprintf("%d", int64(val)))
		case string:
			strArr = append(strArr, `"`+val+`"`)
		case []int:
			strArr = append(strArr, fmt.Sprintf("%v", val))
		case []interface{}:
			strArr = append(strArr, fmt.Sprintf("%v", val))
		}
	}
	return strings.Join(strArr, ",")
}
func GetLivePushUrl() string {
	pushUrl, err := database.Redis.GetRedis().Get(model.RedisKeyMatchVideoPushUrl).Result()
	if err != nil {
		pushUrl = "rtmp://livepush.migucloudeg.com"
		SetLivePushUrl(pushUrl)
	}
	return pushUrl
}
func SetLivePushUrl(pushUrl string) {
	database.Redis.GetRedis().Set(model.RedisKeyMatchVideoPushUrl, pushUrl, time.Second*model.RedisExpirationOneWeek*9999)
}
func GetLivePlayUrl() string {
	playUrl, err := database.Redis.GetRedis().Get(model.RedisKeyMatchVideoPlayUrl).Result()
	if err != nil {
		playUrl = "https://liveplay.migucloudeg.com"
		SetLivePlayUrl(playUrl)
	}
	return playUrl
}
func SetLivePlayUrl(playUrl string) {
	database.Redis.GetRedis().Set(model.RedisKeyMatchVideoPlayUrl, playUrl, time.Second*model.RedisExpirationOneWeek*9999)
}
