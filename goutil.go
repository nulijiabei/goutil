package goutil

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// --------------------------------- //
// a2b

// string to float
func String2Float(data string, reData float64) float64 {
	i, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return reData
	}
	return i
}

// string to int
func String2Int(s string, dft int) int {
	var re, err = strconv.Atoi(s)
	if err != nil {
		return dft
	}
	return re
}

// string to utf8
func String2Utf8(_str string) string {
	var utf string
	b := []byte(_str)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		utf += string(r)
		b = b[size:]
	}
	return utf
}

// --------------------------------- //
// string

// 是不是空字符
func IsSpace(c byte) bool {
	if c >= 0x00 && c <= 0x20 {
		return true
	}
	return false
}

// 判断一个字符串是不是空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func IsBlank(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if !IsSpace(b) {
			return false
		}
	}
	return true
}

// --------------------------------- //
// time

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

// --------------------------------- //
// os

// 判断一个路径是否存在
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// --------------------------------- //
// net

// 获取指定网络接口地址
func GetInterfaceByName() map[string]interface{} {
	inter := make(map[string]interface{})
	devs := [...]string{"eth0", "eth1", "ppp0", "tun0"}
	for _, dev := range devs {
		data := make(map[string]string)
		network, err := GetNetworkAddrByName(dev)
		if err == nil {
			data["network"] = network.(string)
		}
		hardware, err := GetHardwareAddrByName(dev)
		if err == nil {
			data["hardware"] = hardware.(string)
		}
		if len(data) > 0 {
			inter[dev] = data
		}
	}
	return inter
}

// 通过接口名称获取硬件地址
func GetHardwareAddrByName(_interface string) (interface{}, error) {
	inter, err := net.InterfaceByName(_interface)
	if err != nil {
		return nil, err
	}
	var HardwareAddr string
	for _, v := range strings.Split(inter.HardwareAddr.String(), ":") {
		HardwareAddr += v
	}
	if IsBlank(strings.TrimSpace(HardwareAddr)) {
		return nil, fmt.Errorf("...")
	}
	return strings.ToUpper(HardwareAddr), nil
}

// 通过接口名称获取网络地址
func GetNetworkAddrByName(_interface string) (interface{}, error) {
	inter, err := net.InterfaceByName(_interface)
	if err != nil {
		return nil, err
	}
	addrs, err := inter.Addrs()
	if err != nil {
		return nil, err
	}
	for _, v := range addrs {
		addr := strings.Split(v.String(), "/")
		if len(strings.Split(addr[0], ".")) == 4 {
			return addr[0], nil
		}
	}
	return nil, fmt.Errorf("...")
}
