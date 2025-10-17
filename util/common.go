package util

import (
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"common/conf"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 获取订单编号
func GetOrderUuid() (orderNum string) {
	orderNum = uuid.NewString()
	if len(orderNum) > 3 && orderNum[:3] == "live" {
		return GetOrderUuid()
	}
	return
}

const earthRadius = 6371 // 地球半径，单位：千米

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// 将经纬度转换为弧度
	lat1Rad := degToRad(lat1)
	lon1Rad := degToRad(lon1)
	lat2Rad := degToRad(lat2)
	lon2Rad := degToRad(lon2)

	// 应用Haversine公式计算距离
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func CalculateAge(birthDate string) int {
	layout := "2006-01-02" // 定义日期格式
	t, err := time.Parse(layout, birthDate)
	if err != nil {
		fmt.Println("无效的日期格式")
		return -1
	}

	now := time.Now()
	age := now.Year() - t.Year()

	// 如果生日还未到，则减去一岁
	if now.YearDay() < t.YearDay() {
		age--
	}

	return age
}

// 截取文件路径
func ImgSliceFilePath(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	// 返回去掉前导斜杠的路径
	if len(u.Path) >= 2 {
		return u.Path[1:]
	}
	return u.Path
}

// 签名图片设置访问时间
func SignImgAccessTime(str string) string {
	if str == "" {
		return ""
	}

	// 创建 OSS 客户端
	client, err := oss.New(conf.Conf.AliyunCloud.Oss.Endpoint, conf.Conf.AliyunCloud.AccessKeyID, conf.Conf.AliyunCloud.AccessKeySecret)
	if err != nil {
		fmt.Println("Error creating OSS client:", err)
		return ""
	}

	// 获取存储空间
	bucket, err := client.Bucket(conf.Conf.AliyunCloud.Oss.BucketName)
	if err != nil {
		fmt.Println("Error getting bucket:", err)
		return ""
	}

	// 生成签名的访问链接，设置有效期为10分钟
	signedURL, err := bucket.SignURL(str, oss.HTTPGet, 10*60)
	if err != nil {
		fmt.Println("Error signing URL:", err)
		return ""
	}

	return signedURL
}

// 提取比赛id
func ExtractTwoNumbers(s string) (int8, uint64, bool) {
	re := regexp.MustCompile(`-(\d+)-(\d+)$`)
	match := re.FindStringSubmatch(s)
	if len(match) == 3 {
		// 转换第一个为 int8
		n1, err1 := strconv.ParseInt(match[1], 10, 8)
		// 转换第二个为 uint64
		n2, err2 := strconv.ParseUint(match[2], 10, 64)
		if err1 == nil && err2 == nil {
			return int8(n1), n2, true
		}
	}
	return 0, 0, false
}

func GetTimestamps(t uint64) (start, end int64) {
	// 获取当前时间
	now := time.Now()

	// 获取今天的 0 点
	year, month, day := now.Date()
	location := now.Location()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, location)

	// 根据类型调整时间
	var targetDate time.Time
	//switch t {
	//case 1: // 今天
	//	targetDate = startOfDay
	//case 2: // 明天
	//	targetDate = startOfDay.Add(24 * time.Hour)
	//case 3: // 后天
	//	targetDate = startOfDay.Add(48 * time.Hour)
	//case 4: // 第四天
	//	targetDate = startOfDay.Add(72 * time.Hour)
	//case 5: // 第五天
	//	targetDate = startOfDay.Add(96 * time.Hour)
	//case 6: // 第六天
	//	targetDate = startOfDay.Add(120 * time.Hour)
	//case 7: // 第七天
	//	targetDate = startOfDay.Add(144 * time.Hour)
	//default:
	//	// 默认是今天
	//	targetDate = startOfDay
	//}
	if t <= 1 || t > 31 {
		targetDate = startOfDay
	} else {
		//targetDate = startOfDay.Add(24 * (t - 1) * time.Hour)
		targetDate = startOfDay.Add(time.Duration(24*int64(t-1)) * time.Hour)

	}

	// 返回该天的 0 点和 24 点时间戳
	start = targetDate.Unix()
	end = targetDate.Add(24 * time.Hour).Unix()

	return
}

// 根据日期(yyyy-MM-dd)提取开始/结束 时间戳
func GetDayStartEndTimestamp(dateStr string) (startTs int64, endTs int64, err error) {
	const layout = "2006-01-02"
	start, err := time.ParseInLocation(layout, dateStr, time.Local)
	if err != nil {
		return
	}
	end := start.Add(24 * time.Hour).Add(-time.Second)

	startTs = start.Unix()
	endTs = end.Unix()
	return
}
