package util

import (
	"fmt"
	"net/url"
	"strings"
	xtime "time"
)

// 分页参数
type PageParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// 获取分页参数
func GetPageParams(page int, size int) (params PageParams) {
	if page < 1 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	if size > 10000 {
		size = 10000
	}
	params.Offset = (page - 1) * size
	params.Limit = size
	return
}

// 替换字符窜
func HideStr(str string, start int, length int, replace string) string {
	strLen := len(str)
	if strLen < start {
		return str
	}
	var newStr strings.Builder
	newStr.WriteString(str[:start])
	if strLen < start+length {
		newStr.WriteString(strings.Repeat(replace, strLen-start))
	} else {
		newStr.WriteString(strings.Repeat(replace, length))
		newStr.WriteString(str[start+length:])
	}
	return newStr.String()
}

// url_encode
func EncodeURIComponent(str string) string {
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	return str
}

// IsSuccessive 是否满足连续lens数量的顺子号码
// lens 连续数字的数量
func IsSuccessive(n, lens int) bool {
	//统计正顺次数 12345
	z := 0
	//统计反顺次数  654321
	f := 0
	//判断3个数字是否是顺子，只需要判断2次
	lens = lens - 1
	for {
		// 个位数
		g := n % 10
		n = n / 10
		// 十位数
		s := n % 10

		if s-g == 1 {
			f = f + 1
		} else {
			f = 0
		}

		if g-s == 1 {
			z = z + 1
		} else {
			z = 0
		}

		if f == lens || z == lens {
			return true
		}

		if n == 0 {
			return false
		}
	}
}

// IsAlike 是否连续lens位相同数字
func IsAlike(n, lens int) bool {
	c := 0
	lens = lens - 1
	var g, s int
	for {
		g = n % 10
		n = n / 10
		s = n % 10

		if s == g {
			c = c + 1
		} else {
			c = 0
		}

		if c == lens {
			return true
		}

		if n == 0 {
			return false
		}
	}
}

// 时间戳格式化
func TimeUnix(timeUnix int64) string {
	if timeUnix == 0 {
		return ""
	}
	unix := xtime.Unix(timeUnix, 0)
	format := unix.Format("2006-01-02 15:04:05")
	return format
}

// 时间转换时间戳
func TimeForUnix(t string) int64 {
	if t == "" {
		return 0
	}
	loc, _ := xtime.LoadLocation("Local")
	tmp, _ := xtime.ParseInLocation("2006-01-02 15:04:05", t, loc)
	return tmp.Unix()
}

// 数字缩写
func NumberAbbreviation(num float64) (numStr string) {
	if num < 10000 {
		numStr = fmt.Sprintf("%d", int64(num))
	} else {
		numStr = fmt.Sprintf("%.2fk", num/10000)
	}
	return
}
