package xtime

import (
	"time"
)

// import (
//
//	"context"
//	"database/sql/driver"
//	"fmt"
//	"strconv"
//	xtime "time"
//
// )
//
// // Time be used to MySql timestamp converting.
// type Time int64
//
// // Scan scan time.
//
//	func (jt *Time) Scan(src interface{}) (err error) {
//		switch sc := src.(type) {
//		case xtime.Time:
//			*jt = Time(sc.Unix())
//		case string:
//			var i int64
//			i, err = strconv.ParseInt(sc, 10, 64)
//			*jt = Time(i)
//		}
//		return
//	}
//
// // Value get time value.
//
//	func (jt Time) Value() (driver.Value, error) {
//		return xtime.Unix(int64(jt), 0), nil
//	}
//
// // Time get time.
//
//	func (jt Time) Time() xtime.Time {
//		return xtime.Unix(int64(jt), 0)
//	}
//
// // Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration time.Duration

// // UnmarshalText unmarshal text to duration.
//
//	func (d *Duration) UnmarshalText(text []byte) error {
//		tmp, err := xtime.ParseDuration(string(text))
//		if err == nil {
//			*d = Duration(tmp)
//		}
//		return err
//	}
//
// // Shrink will decrease the duration by comparing with context's timeout duration
// // and return new timeout\context\CancelFunc.
//
//	func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
//		if deadline, ok := c.Deadline(); ok {
//			if ctimeout := xtime.Until(deadline); ctimeout < xtime.Duration(d) {
//				// deliver small timeout
//				return Duration(ctimeout), c, func() {}
//			}
//		}
//		ctx, cancel := context.WithTimeout(c, xtime.Duration(d))
//		return d, ctx, cancel
//	}
//
// // 补全时间格式
//
//	func (d Duration) DateFormat(oldStr string) (newStr string) {
//		fmt.Println(len(oldStr))
//		switch len(oldStr) {
//		case 4:
//			newStr = oldStr + "-01-01 00:00:00"
//		case 7:
//			newStr = oldStr + "-01 00:00:00"
//		case 10:
//			newStr = oldStr + " 00:00:00"
//		case 13:
//			newStr = oldStr + ":00:00"
//		case 16:
//			newStr = oldStr + ":00"
//		case 19:
//			newStr = oldStr
//		default:
//			newStr = ""
//		}
//		return
//	}
//
// // 今日零点时间戳
func TodayZeroTimestamp() (timestamp int64) {
	now := time.Now()
	year, month, day := now.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	timestamp = start.Unix()
	return
}

// 11位时间戳
func ElevenBitTimestamp() (millisecondTimestamp int64) {
	// 获取当前时间
	currentTime := time.Now()

	// 生成纳秒级的时间戳
	nanoTimestamp := currentTime.UnixNano()

	// 将纳秒转换为11位长度的时间戳（毫秒级）
	millisecondTimestamp = nanoTimestamp / 1e6
	return
}
func Get30DayZeroTimestamp() (timestamp int64) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	// 计算第90天的结束时间：开始时间 + 90天 - 1秒
	endOfDay := startOfDay.AddDate(0, 0, 30).Add(-time.Second)
	timestamp = endOfDay.Unix()
	return

}

func Get7DayZeroTimestamp() (timestamp int64) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	// 计算第90天的结束时间：开始时间 + 90天 - 1秒
	endOfDay := startOfDay.AddDate(0, 0, 7).Add(-time.Second)
	timestamp = endOfDay.Unix()
	return

}
func Get3DayZeroTimestamp() (timestamp int64) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	// 计算第90天的结束时间：开始时间 + 90天 - 1秒
	endOfDay := startOfDay.AddDate(0, 0, 3).Add(-time.Second)
	timestamp = endOfDay.Unix()
	return

}
func GetNext7Days() []string {
	var dates []string
	today := time.Now()
	for i := 0; i < 7; i++ {
		day := today.AddDate(0, 0, i)
		dates = append(dates, day.Format("2006-01-02"))
	}
	return dates
}
