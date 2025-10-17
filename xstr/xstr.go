package xstr

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInts format int64 slice like:n1,n2,n3.
func JoinInts(is []int64) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// SplitInts split string into int64 slice.
func SplitInts(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

// string 转byte[]
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// byte[] 转 string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwsyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 最短设备码
const CodeMinLength = 6

// 字符串反转
func reverse(str string) string {
	var result []byte
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	return string(result)
}

// transform GBK bytes to UTF-8 bytes
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// transform UTF-8 bytes to GBK bytes
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// transform GBK string to UTF-8 string and replace it, if transformed success, returned nil error, or died by error message
func StrToUtf8(str *string) error {
	b, err := GbkToUtf8([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}

// transform UTF-8 string to GBK string and replace it, if transformed success, returned nil error, or died by error message
func StrToGBK(str *string) error {
	b, err := Utf8ToGbk([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 字符串填补0
func StrPad(code string, length int) string {
	var addString = "0"
	if len(code) < length {
		code = addString + code
		code = StrPad(code, length)
	}
	return code
}

// 生成的code加密
func EncodeCode(code string, l int) string {
	var sourceString = "E5FCD1G3HQA4BNPJO2RSTUV67MWX8I9KLYZ"
	// 截取最后一位字符串
	length := len(code)
	if length < l {
		code = StrPad(code, 8)
		length = len(code)
	}
	lastChar := code[length-1:]

	// 最后一位字符串在源字符串出现的位置
	step := strings.Index(sourceString, lastChar) - (length - 3)
	codeByte := []byte(code)
	for i := 0; i < length-1; i++ {
		local := 0
		b := strings.Index(sourceString, string(codeByte[i]))
		if b == -1 {
			b = 0
		}
		if step%2 != 0 {
			local = b + step - i
		} else {
			local = b + step + i
		}

		if local < 0 {
			local = len(sourceString) + local
		}

		if local >= len(sourceString) {
			local = local - len(sourceString)
		}
		codeByte[i] = sourceString[local]
	}
	return string(codeByte)
}

// 获取excel文件的列，最大26列
func GetExcelHeader(data []string, K int, rows map[string]string) (map[string]string, error) {
	if len(data) > 26 {
		return nil, errors.New("列数超出，最大支持26列")
	}
	keys := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	res := make(map[string]string, 0)
	if rows != nil {
		res = rows
	}

	for k, v := range data {
		res[keys[k]+strconv.Itoa(K)] = v
	}
	return res, nil
}

// 将一个切片相同参数赋值到另一个切片
func ConvertStructData(src interface{}, dst ...interface{}) error {
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}

	for _, v := range dst {
		if err := json.Unmarshal(jsonStr, v); err != nil {
			return err
		}
	}

	return nil
}

// 判断切片中是否有指定元素
func SliceInValueInt64(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// 判断切片中是否有指定元素2
func SliceInValueString(slice []*string, value string) bool {
	for _, v := range slice {
		if *v == value {
			return true
		}
	}
	return false
}
