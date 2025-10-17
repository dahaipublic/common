package util

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

func ConvertStructData(src interface{}, dst ...interface{}) error {
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}
	for _, v := range dst {
		json.Unmarshal(jsonStr, v)
	}

	return nil
}

func StructToMap(src interface{}, m *map[string]interface{}) (err error) {
	data, err := json.Marshal(&src)
	if err != nil {
		return
	}

	newMap := make(map[string]interface{})
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return
	}
	if len(newMap) > 0 {
		*m = make(map[string]interface{}, len(newMap))
		for k, v := range newMap {
			(*m)[CamelToUnderscore(k)] = v
		}
	}
	return
}

func CamelToUnderscore(key string) string {
	// 使用正则表达式匹配大写字母，并在前面加上下划线
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	underscoreKey := re.ReplaceAllString(key, "${1}_${2}")
	// 将所有大写字母转换为小写
	underscoreKey = strings.ToLower(underscoreKey)
	return underscoreKey
}

func StringToInt64Arr(input string) (intArr []int64) {
	if input == "" {
		return
	}
	// 使用逗号分隔字符串
	strArr := strings.Split(input, ",")

	// 创建 int64 数组
	intArr = make([]int64, len(strArr))

	// 将字符串转换为 int64 并存储到数组中
	for i, str := range strArr {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		intArr[i] = num
	}
	return
}
