package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	//"github.com/satori/go.uuid"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// SHA1
func SHA1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// hmac SHA1
func HmacSHA1(data, key string) string {
	m := hmac.New(sha1.New, []byte(key))
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

//float64 转 String工具类，保留 prec 位小数
func FloatToString(input_num float64, prec int) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', prec, 64)
}

// 时间字符串转换秒数，传入格式 HH:mm:ss.000
func String2Seconds(timeStr string) (uint32, error) {
	if timeStr == "" {
		return 0, nil
	}
	var seconds uint32
	t := strings.Split(timeStr, ":")
	hour, err := strconv.Atoi(t[0])
	if err != nil {
		return 0, err
	}
	seconds += uint32(hour) * 3600

	minute, err := strconv.Atoi(t[1])
	if err != nil {
		return 0, err
	}
	seconds += uint32(minute) * 60

	second, err := func(seconds string) (uint32, error) {
		s := strings.Split(seconds, ".")
		second, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, err
		}
		return uint32(second), nil
	}(t[2])
	if err != nil {
		return 0, err
	}
	return second + seconds, nil
}

//获取随机数 (min, max]
func GetRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	//logs.Error("rangeNum-max-min", max, min)
	if min == max {
		return max
	}
	randNum := rand.Intn(max-min) + min
	return randNum
}

// 判断obj是否在target中，target支持的类型arrary,slice,map
func Contains(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

/**
获取当前日期字符串
格式：20060102
*/
func GetNowDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	result := tm.Format("20060102")
	return result
}

/**
获取当前日期字符串
格式：2006-01-02
*/
func GetFormatNowDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	result := tm.Format("2006-01-02")
	return result
}

/**
获取指定日期时间戳
待转化为时间戳的字符串 datetime := "2015-01-01 00:00:00"
*/
func GetFormatDate(datetime string) int64 {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	return timestamp
}

/**
写入文本文件，追加模式，如果不存在自动创建文件
*/
func WriteFile(path string, data string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		logs.Error("WriteFile-OpenFile", err)
	}
	defer file.Close()
	bytes := []byte(data + "\n")
	_, err = file.Write(bytes)
	if err != nil {
		logs.Error("WriteFile-Write", err)
	}
}

/**
判断数组是否包含字符串
*/

func InArray(array []string, val string) bool {
	result := false
	num := len(array)
	for i := 0; i < num; i++ {
		if array[i] == val {
			result = true
		}
	}
	return result
}
func InArrayInt(array []int, val int) bool {
	result := false
	num := len(array)
	for i := 0; i < num; i++ {
		if array[i] == val {
			result = true
		}
	}
	return result
}

/**
获取客户端IP
*/
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	if !CheckIp(remoteAddr) {
		if ip := req.Header.Get("X-Real-IP"); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
		}
	}
	return remoteAddr
}

/**
获取当前13位时间戳字符串
*/
func GetNowUnixMillisecond() string {
	result := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	return result
}

/**
返回当前纳秒时间字符串 19位
*/
func GetNowUnixNa() string {
	result := strconv.Itoa(int(time.Now().UnixNano()))
	return result
}

/**
获取 UUID
*/
func GetUUID() string {
	result, err := uuid.NewV4()
	if err != nil {
		logs.Error("getUUID", err)
	}
	return result.String()
}

/**
字符串转float64
*/
func StrToFloat64(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		logs.Error("strToFloat64", err)
	}
	return result
}

//加上 0.5是为了四舍五入，想保留几位小数的话把2改掉即可。
func Decimal(value float64) float64 {
	//return math.Trunc(value*1e2+0.5) * 1e-2
	return math.Trunc(value*1e2) * 1e-2
}

/**
获取当前时间字符串
格式：2006-01-02 15:04:05
*/
func GetNowDateTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	result := tm.Format("2006-01-02 15:04:05")
	return result
}

/**
获取当前时间字符串
格式：2017-03-15 20:49:26.978 +0800 CST
*/
func MsToTime(ms string) (string, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		logs.Error("msToTime", err)
		return "", err
	}
	tm := time.Unix(0, msInt*int64(time.Millisecond))
	return tm.String(), nil
}

/**
计算两个字符串的差
str1、str2 为数字型字符串
返回差额字符串，毫秒级，0.000
*/
func TwoStringLack(str1, str2 string) string {
	str1ToI, err := strconv.Atoi(str1)
	str2ToI, err := strconv.Atoi(str2)
	if err != nil {
		logs.Error("twoStringLack", err)
	}
	lack := str2ToI - str1ToI
	lackStr := FloatToString(float64(lack)/1000, 3)
	return lackStr
}

/*
计算渠道请求时间毫秒
*/
func GetUnionRequestTime(UnionRequestTime string) int {
	second := StrToFloat64(UnionRequestTime)
	ms := second * 1000
	unionRequestTime := IntFromFloat64(ms)
	return unionRequestTime
}

/**
验证IP是否合法
*/
func CheckIp(ip string) bool {
	if ip == "" {
		return false
	}
	if ip == "127.0.0.1" {
		return false
	}
	ipArr := strings.Split(ip, ".")
	if len(ipArr) == 4 {
		if (ipArr[0] == "192" && ipArr[1] == "168") || ipArr[0] == "10" {
			return false
		}
		ipArr1, err := strconv.Atoi(ipArr[1])
		if err != nil {
			logs.Error("tools-checkIp", err)
			return false
		}
		if ipArr[0] == "172" && ipArr1 >= 16 && ipArr1 <= 31 {
			return false
		}
	}
	return true
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

//验证版本号
func CheckApiVersion(apiVersion string) bool {
	var result bool = false
	apiVersionArr := strings.Split(apiVersion, ".")
	if len(apiVersionArr) == 3 {
		result = true
	}
	return result
}

//判断是否增加点击宏的整数替换(版本号 >= 1.1.0)
func CheckClickMacroInt(apiVersion string) (bool, error) {
	var result bool = false
	apiVersionArr := strings.Split(apiVersion, ".")
	major := apiVersionArr[0]
	minor := apiVersionArr[1]
	micro := apiVersionArr[2]
	majorInt, err := strconv.Atoi(major)
	minorInt, err := strconv.Atoi(minor)
	microInt, err := strconv.Atoi(micro)
	if err != nil {
		return false, err
	}
	if majorInt > 1 || (majorInt == 1 && minorInt > 1) || (majorInt == 1 && minorInt == 1 && microInt >= 0) {
		result = true
	}
	return result, nil
}

// base编码
func BASE64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

// base解码
func BASE64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}

//base64加密
func Base64Encode(src []byte, base64Table string) []byte {
	var coder = base64.NewEncoding(base64Table)
	return []byte(coder.EncodeToString(src))
}
func Base64Decode(src []byte, base64Table string) ([]byte, error) {
	var coder = base64.NewEncoding(base64Table)
	return coder.DecodeString(string(src))
}

// 获取程序运行路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error("getCurrentDirectory", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取两个数组的交集
func GetIntersect(arr1, arr2 []string) []string {
	var result []string
	for _, val := range arr1 {
		if InArray(arr2, val) {
			result = append(result, val)
		}
	}
	return result
}

// 手机号验证
func VerifyMobileFormat(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// 获取当前时间与当日零点的时间差
func GetTodayEndSecond() time.Duration {
	now := time.Now().Unix()
	// 获取到23:59:59 的时间戳 +1 秒
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	endSecond := endTime.Unix()
	expire := endSecond - now + 1
	return time.Duration(expire)
}

// 生成订单ID
func GetOrderNumber(orderId int64, payType int8, userId string) string {
	var prefix string
	switch payType {
	case 1:
		prefix = "wx"
	}
	date := GetNowDate()
	orderNumber := prefix + strconv.Itoa(GetRangeNum(1000, 9999)) + date + userId
	return orderNumber
}

// 时间戳转日期, 转换后的格式 20191108
func UnixToDate(unix int64) string {
	tm := time.Unix(unix, 0)
	return tm.Format("20060102")
}

//struct转map例子
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// 日志文件添加hostname
func GetLogPath(path string) string {
	hostName, _ := os.Hostname()
	result := strings.Replace(path, "${HOSTNAME}", hostName, -1)
	return result
}

// HmacSha256 加密
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

//start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
//       负数 - 在从字符串结尾的指定位置开始
//       0 - 在字符串中的第一个字符处开始
//length:正数 - 从 start 参数所在的位置返回
//       负数 - 从字符串末端返回

func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}
func TimestrToTimestamp(time_str string, flag int) int64 {
	var t int64
	loc, _ := time.LoadLocation("Local")
	if flag == 1 {
		t1, _ := time.ParseInLocation("2006.01.02 15:04:05", time_str, loc)
		t = t1.Unix()
	} else if flag == 2 {
		t1, _ := time.ParseInLocation("2006-01-02 15:04", time_str, loc)
		t = t1.Unix()
	} else if flag == 3 {
		t1, _ := time.ParseInLocation("2006-01-02", time_str, loc)
		t = t1.Unix()
	} else if flag == 4 {
		t1, _ := time.ParseInLocation("2006.01.02", time_str, loc)
		t = t1.Unix()
	} else {
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)
		t = t1.Unix()
	}
	return t
}

// 判断版本号 v1 >= v2 true, 否则 false
func CheckVersion(v1, v2 string) bool {
	v1Int, _ := strconv.Atoi(v1)
	v2Int, _ := strconv.Atoi(v2)
	if v1Int >= v2Int {
		return true
	}
	return false
}
