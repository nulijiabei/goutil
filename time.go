package goutil

import (
	"strconv"
	"time"
)

const FORMAT_DATE string = "2006-01-02"
const FORMAT_TIME string = "15:04:05"

// 获得当前日志
func GetDate() string {
	return time.Now().Format(FORMAT_DATE)
}

// 获得当前系统时间
func GetTime() string {
	return time.Now().Format(FORMAT_TIME)
}

// 获取本地时间戳纳秒,以字符串格式返回
func UnixNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
