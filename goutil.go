package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

// string to int64
func String2Int64(s string, dft int64) int64 {
	var re, err = strconv.ParseInt(s, 10, 64)
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
const FORMAT_DATE_TIME string = "2006-01-02 15:04:05"

// 获取本地时间戳纳秒,以字符串格式返回
func UnixNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// 获取本地事件毫秒
func UnixMsSec(off int) int64 {
	return (time.Now().Unix() + int64(off)) * 1000
}

// 获得当前系统日期
//const FORMAT_DATE string = "2006-01-02"
//const FORMAT_TIME string = "15:04:05"
//const FORMAT_DATE_TIME string = "2006-01-02 15:04:05"
func GetTime(layout string) string {
	return time.Now().Format(layout)
}

// 解析日期
//const FORMAT_DATE string = "2006-01-02"
//const FORMAT_TIME string = "15:04:05"
//const FORMAT_DATE_TIME string = "2006-01-02 15:04:05"
func ParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	panic(err)
	return t
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
// http

// Http Get Request
func HttpGetReq(u string, params map[string]string) ([]byte, error) {
	p, _ := url.Parse(u)
	q := p.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	p.RawQuery = q.Encode()
	resp, err := http.Get(p.String())
	if err != nil {
		return nil, err
	}
	//	if resp.StatusCode != 200 {
	//		log.Println(u, resp.StatusCode)
	//	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Http Post Request
func HttpPostReq(u string, params map[string]string) ([]byte, error) {
	values := make(url.Values)
	for k, v := range params {
		values.Add(k, v)
	}
	resp, err := http.PostForm(u, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Http Post Request Plus
func HttpPostReqPlus(u string, params map[string]interface{}, header map[string]string) ([]byte, error) {
	client := &http.Client{}
	js, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u, strings.NewReader(string(js)))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// multipart/form-data 文件上传 ...
// <input type="file" name="file" />
func HttpMultipartPostReq(u string, ff string, params map[string]string) ([]byte, error) {
	f, err := os.Open(ff)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	file, err := w.CreateFormFile("file", filepath.Base(ff))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, f)
	content_type := w.FormDataContentType()
	for k, v := range params {
		w.WriteField(k, v)
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", content_type)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
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
