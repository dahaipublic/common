package util

import (
	"common"
	"common/util"
	"time"
)

// 多语言词库，key 使用 int
var translations = map[string]map[int]map[int]string{
	"weekday": {
		common.ELan_en_US: {
			1: "Monday",
			2: "Tuesday",
			3: "Wednesday",
			4: "Thursday",
			5: "Friday",
			6: "Saturday",
			7: "Sunday",
		},
		common.ELan_zh_CN: {
			1: "周一",
			2: "周二",
			3: "周三",
			4: "周四",
			5: "周五",
			6: "周六",
			7: "周日",
		},
		common.ELan_tr_TR: {
			1: "Pazartesi",
			2: "Salı",
			3: "Çarşamba",
			4: "Perşembe",
			5: "Cuma",
			6: "Cumartesi",
			7: "Pazar",
		},
		common.ELan_ar_EG: {
			1: "الاثنين",
			2: "الثلاثاء",
			3: "الأربعاء",
			4: "الخميس",
			5: "الجمعة",
			6: "السبت",
			7: "الأحد",
		},
	},
	// 可以继续扩展其他类型，比如 month
	"month": {
		common.ELan_en_US: {
			1:  "January",
			2:  "February",
			3:  "March",
			4:  "April",
			5:  "May",
			6:  "June",
			7:  "July",
			8:  "August",
			9:  "September",
			10: "October",
			11: "November",
			12: "December",
		},
		common.ELan_zh_CN: {
			1:  "一月",
			2:  "二月",
			3:  "三月",
			4:  "四月",
			5:  "五月",
			6:  "六月",
			7:  "七月",
			8:  "八月",
			9:  "九月",
			10: "十月",
			11: "十一月",
			12: "十二月",
		},
		common.ELan_tr_TR: {
			1:  "Ocak",
			2:  "Şubat",
			3:  "Mart",
			4:  "Nisan",
			5:  "Mayıs",
			6:  "Haziran",
			7:  "Temmuz",
			8:  "Ağustos",
			9:  "Eylül",
			10: "Ekim",
			11: "Kasım",
			12: "Aralık",
		},
		common.ELan_ar_EG: {
			1:  "يناير",
			2:  "فبراير",
			3:  "مارس",
			4:  "أبريل",
			5:  "مايو",
			6:  "يونيو",
			7:  "يوليو",
			8:  "أغسطس",
			9:  "سبتمبر",
			10: "أكتوبر",
			11: "نوفمبر",
			12: "ديسمبر",
		},
	},
}

// common.GetTranslation 根据类型、语言、数字 key 获取对应文字
func GetTranslation(content string, lanId, key int) string {
	// 获取类型
	typeMap, ok := translations[content]
	if !ok {
		return ""
	}

	// 获取语言
	langMap, ok := typeMap[lanId]
	if !ok {
		langMap = typeMap[lanId] // 默认土语
	}

	// 查找 key
	if val, ok := langMap[key]; ok {
		return val
	}

	// 如果该语言没有，尝试英文
	if val, ok := typeMap[common.ELan_en_US][key]; ok {
		return val
	}

	return ""
}

// 根据时间戳返回对应的文字
func GetWeekTypeByTimestamp(ts int64, lanID common.ELanDef) string {
	t := time.Unix(ts, 0).In(time.Local)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // 周日改成 7
	}
	return util.GetTranslation("weekday", lanID, weekday)
}
