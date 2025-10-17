package tencentcloudservices

// import (
// 	"common/conf"
// 	"crypto/md5"
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// type TencentCloudLive struct {
// 	Ini *conf.TencentCloudConfig
// }

// // 创建一个直播
// func NewTencentCloudLive() *TencentCloudLive {
// 	var tencentCloudLive TencentCloudLive
// 	tencentCloudLive.Ini = &conf.Conf.TencentCloud
// 	return &tencentCloudLive
// }

// // 创建一个流名称
// func (q *TencentCloudLive) CreateStreamKey(roomId string) (streamKey string) {
// 	streamKey = fmt.Sprintf("%s%s", q.Ini.Css.StreamKeyPrefix, roomId)
// 	return
// }

// // 根据一个流名称得到直播间ID
// func (q *TencentCloudLive) GetRoomIdForStreamKey(streamKey string) (roomId string, err error) {
// 	roomId = streamKey[len(q.Ini.Css.StreamKeyPrefix):]
// 	return
// }

// // 获取推流地址
// func (q *TencentCloudLive) GetPushUrl(streamName string, expireAt int64) (url string) {
// 	var extStr, key string
// 	key = conf.Conf.TencentCloud.Css.PushSecret
// 	if key != "" && expireAt != 0 {
// 		txTime := strings.ToUpper(strconv.FormatInt(expireAt, 16))
// 		txSecret := md5.Sum([]byte(key + streamName + txTime))
// 		txSecretStr := fmt.Sprintf("%x", txSecret)
// 		extStr = "?txSecret=" + txSecretStr + "&txTime=" + txTime
// 	}
// 	url = "地址:rtmp://" + conf.Conf.TencentCloud.Css.PushUrl + "/" + conf.Conf.TencentCloud.Css.AppName + "/推流码:" + streamName + extStr
// 	return
// }

// // 获取拉流地址
// // txTime过期时间
// func (q *TencentCloudLive) GetPullUrl(streamName string, format string, txTime int64) (url string) {
// 	txTime = time.Now().Unix() + txTime
// 	txSecret := fmt.Sprintf("%x", md5.Sum([]byte(conf.Conf.TencentCloud.Css.PullSecret+streamName+strings.ToUpper(strconv.FormatInt(txTime, 16)))))
// 	switch format {
// 	case "HLS":
// 		url = "http://" + conf.Conf.TencentCloud.Css.PullUrl + "/" + conf.Conf.TencentCloud.Css.AppName + "/" + streamName + ".m3u8" + "?txSecret=" + txSecret + "&txTime=" + strings.ToUpper(strconv.FormatInt(txTime, 16))
// 	case "RTMP":
// 		url = "rtmp://" + conf.Conf.TencentCloud.Css.PullUrl + "/" + conf.Conf.TencentCloud.Css.AppName + "/" + streamName + "?txSecret=" + txSecret + "&txTime=" + strings.ToUpper(strconv.FormatInt(txTime, 16))
// 	case "FLV":
// 		//url = "https://" + conf.Conf.TencentCloud.Css.PullUrl + "/" + conf.Conf.TencentCloud.Css.AppName + "/" + streamKey + ".flv" + "?txSecret=" + txSecret + "&txTime=" + strings.ToUpper(strconv.FormatInt(txTime, 16))
// 		url = "https://" + conf.Conf.TencentCloud.Css.PullUrl + "/" + conf.Conf.TencentCloud.Css.AppName + "/" + streamName + ".flv"
// 	case "WebRTC":
// 		url = "webrtc://" + conf.Conf.TencentCloud.Css.PullUrl + "/" + conf.Conf.TencentCloud.Css.AppName + "/" + streamName + "?txSecret=" + txSecret + "&txTime=" + strings.ToUpper(strconv.FormatInt(txTime, 16))
// 	default:
// 		return
// 	}
// 	return
// }

// // 回调鉴权
// func (q *TencentCloudLive) Authentication(sign string, t int64) (b bool) {
// 	newSign := md5.Sum([]byte(conf.Conf.TencentCloud.Css.CallbackKey + fmt.Sprintf("%d", t)))
// 	fmt.Println(fmt.Sprintf("%x", newSign))
// 	if sign == fmt.Sprintf("%x", newSign) {
// 		return true
// 	}
// 	return
// }
