package aliyunservices

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dahaipublic/common/conf"
	"path"

	"mime/multipart"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/live"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 增加直播录制转点播配置，将录制内容保存到点播媒资库
func AddLiveRecordVodConfig(streamName string) (err error) {
	liveClient, err := live.NewClientWithAccessKey("eu-central-1", conf.Conf.AliyunCloud.AccessKeyID, conf.Conf.AliyunCloud.AccessKeySecret)
	if err != nil {
		return
	}

	request := live.CreateAddLiveRecordVodConfigRequest()

	request.Scheme = "https"

	request.StreamName = streamName
	request.DomainName = conf.Conf.AliyunCloud.Live.DomainName
	request.AppName = conf.Conf.AliyunCloud.Live.AppName
	request.VodTranscodeGroupId = conf.Conf.AliyunCloud.Live.VodTranscodeGroupID
	request.StorageLocation = conf.Conf.AliyunCloud.Live.StorageLocation
	request.CycleDuration = requests.NewInteger(conf.Conf.AliyunCloud.Live.CycleDuration)

	_, err = liveClient.AddLiveRecordVodConfig(request)
	if err != nil {
		return
	}
	return
}

// 删除直播录制转点播配置
func DeleteLiveRecordVodConfig(streamName string) (err error) {
	liveClient, err := live.NewClientWithAccessKey("eu-central-1", conf.Conf.AliyunCloud.AccessKeyID, conf.Conf.AliyunCloud.AccessKeySecret)
	if err != nil {
		return
	}

	request := live.CreateDeleteLiveRecordVodConfigRequest()

	request.Scheme = "https"
	request.DomainName = conf.Conf.AliyunCloud.Live.DomainName
	request.AppName = conf.Conf.AliyunCloud.Live.AppName
	request.StreamName = streamName

	_, err = liveClient.DeleteLiveRecordVodConfig(request)
	if err != nil {
		return
	}
	return
}

// 生成推流地址
func PushUrl(streamName string) (pushUrl string) {
	timeStamp := time.Now().Unix() + 86400*3
	pushUrl = BuildPushURL(
		conf.Conf.AliyunCloud.Live.PushUrl,
		conf.Conf.AliyunCloud.Live.AnchorName,
		streamName,
		conf.Conf.AliyunCloud.Live.PushSecret,
		timeStamp)
	return
}

// 生成播放地址
func PlayUrl(streamName string) (playUrl string) {
	//未开启鉴权Key的情况下

	//playUrl = "https://" + conf.Conf.AliyunCloud.Live.PullUrl + "/" + conf.Conf.AliyunCloud.Live.AnchorName + "/" + streamName + ".flv"
	//开启鉴权key
	timeStamp := time.Now().Unix() + 86400*3
	playUrl = BuildPlayURL(
		conf.Conf.AliyunCloud.Live.PullUrl,
		conf.Conf.AliyunCloud.Live.AnchorName,
		streamName,
		conf.Conf.AliyunCloud.Live.PlaySecret,
		timeStamp,
		"m3u8")
	return

}

const PreHead = "upload/"

func GenerateAuthKey(app, stream, key string, expireUnix int64, rand, uid int, suffix string) string {
	// path 需要以 / 开头，播放/推流时 path 的形式通常为 "/{app}/{stream}" 或 "/{app}/{stream}.flv"
	path := fmt.Sprintf("/%s/%s%s", app, stream, suffix)
	plain := fmt.Sprintf("%s-%d-%d-%d-%s", path, expireUnix, rand, uid, key)

	sum := md5.Sum([]byte(plain))
	md5hex := hex.EncodeToString(sum[:])

	return fmt.Sprintf("%d-%d-%d-%s", expireUnix, rand, uid, md5hex)
}

// BuildPushURL 返回带鉴权的 RTMP 推流地址
// pushDomain: 推流域名，不带协议，例如 push.example.com
func BuildPushURL(pushDomain, app, stream, key string, expireUnix int64) string {
	auth := GenerateAuthKey(app, stream, key, expireUnix, 0, 0, "") // 推流一般 suffix 空
	return fmt.Sprintf("rtmp://%s/%s/%s?auth_key=%s", pushDomain, app, stream, auth)
}
func GetAuthUrl2(playUrl string, deviceType string) string {
	filename := path.Base(playUrl)
	streamName := strings.TrimSuffix(filename, path.Ext(filename))
	timeStamp := time.Now().Unix() + 43200*2

	if deviceType == "PC" {
		return GetAuthUrl(playUrl, streamName, timeStamp, ".flv")

	} else {
		playUrl = ReplaceFLVToM3U8(playUrl)
		return GetAuthUrl(playUrl, streamName, timeStamp, ".m3u8")

	}

}
func ReplaceFLVToM3U8(url string) string {
	if strings.HasSuffix(url, ".flv") {
		return strings.TrimSuffix(url, ".flv") + ".m3u8"
	}
	return url // 如果不是 .flv 结尾，就直接返回原始 URL
}

func GetAuthUrl(playUrl, streamName string, expireUnix int64, suffixType string) string {
	auth := GenerateAuthKey(
		"live",
		streamName,
		conf.Conf.AliyunCloud.Live.PlaySecret,
		expireUnix,
		0,
		0,
		suffixType)
	return playUrl + "?auth_key=" + auth
}

// BuildPlayURL 返回带鉴权的播放地址（支持 rtmp, flv, m3u8）
// playDomain: 播放域名，不带协议，例如 play.example.com
func BuildPlayURL(playDomain, app, stream, key string, expireUnix int64, mode string) string {
	var suffix string
	var proto string
	proto = "https://"

	suffix = ".m3u8"
	auth := GenerateAuthKey(app, stream, key, expireUnix, 0, 0, suffix)
	if mode == "rtmp" {
		return fmt.Sprintf("%s%s/%s/%s?auth_key=%s", proto, playDomain, app, stream, auth)
	}
	// HTTP 播放一般路径： http://domain/{app}/{stream}.flv  或 .m3u8
	return fmt.Sprintf("%s%s/%s/%s%s?auth_key=%s", proto, playDomain, app, stream, suffix, auth)
}
func OssUpload(filename string, s1 *multipart.FileHeader) (signedURL string, err error) {
	filename = strings.ReplaceAll(filename, "..", "")

	// 创建 OSS 客户端
	client, err := oss.New(conf.Conf.AliyunCloud.Oss.Endpoint, conf.Conf.AliyunCloud.AccessKeyID, conf.Conf.AliyunCloud.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(conf.Conf.AliyunCloud.Oss.BucketName)
	if err != nil {
		return
	}

	// 构建年月文件夹路径
	today := time.Now().Format("2006-01") // 格式化为年月
	objectKey := PreHead + today + "/" + filename

	// 检查文件夹是否存在，不存在则创建
	isExist, _ := bucket.IsObjectExist(today + "/")
	if !isExist {
		err = bucket.PutObject(today+"/", strings.NewReader(""))
		if err != nil {
			return
		}
	}

	imginfor, err := s1.Open()
	if err != nil {
		return
	}

	// 上传图片
	err = bucket.PutObject(objectKey, imginfor)
	if err != nil {
		return
	}

	// 生成永久有效的访问链接
	//signedURL, err = bucket.SignURL(objectKey, oss.HTTPGet, 10*60) // 签名图片有效 10分钟
	//if err != nil {
	//	return
	//}
	signedURL = "https://" + conf.Conf.AliyunCloud.Oss.BucketName + "." + conf.Conf.AliyunCloud.Oss.Endpoint + "/" + objectKey

	return
}
