package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
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
func IsExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 确保父目录的存在 ...
func CheckParents(aph string) {
	pph := path.Dir(aph)
	err := os.MkdirAll(pph, os.ModeDir|0755)
	if nil != err {
		panic(err)
	}
}

// 是否存在某一路径
func Fexists(ph string) bool {
	return IsExist(ph)
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

// -------------------------------------------------
// image

// 读取JPEG图片返回image.Image对象
func ImageJPEG(ph string) (image.Image, error) {
	// 打开图片文件
	f, fileErr := os.Open(ph)
	if fileErr != nil {
		return nil, fileErr
	}
	// 退出时关闭文件
	defer f.Close()
	// 解码
	j, jErr := jpeg.Decode(f)
	if jErr != nil {
		return nil, jErr
	}
	// 返回解码后的图片
	return j, nil
}

// 读取PNG图片返回image.Image对象
func ImagePNG(ph string) (image.Image, error) {
	// 打开图片文件
	f, fileErr := os.Open(ph)
	if fileErr != nil {
		return nil, fileErr
	}
	// 退出时关闭文件
	defer f.Close()
	// 解码
	p, pErr := png.Decode(f)
	if pErr != nil {
		return nil, pErr
	}
	// 返回解码后的图片
	return p, nil
}

// 按照分辨率创建一张空白图片对象
func ImageRGBA(width, height int) *image.RGBA {
	// 建立图像,image.Rect(最小X,最小Y,最大X,最小Y)
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// 将图片绘制到图片
func ImageDrawRGBA(img *image.RGBA, imgcode image.Image, x, y int) {
	// 绘制图像
	// image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
	// image.ZP ZP is the zero Point
	// image.Pt Pt is shorthand for Point{X, Y}
	draw.Draw(img, img.Bounds(), imgcode, image.Pt(x, y), draw.Over)
}

// 将图片绘制到图片
func ImageDrawRGBAOffSet(img *image.RGBA, imgcode image.Image, r image.Rectangle, x, y int) {
	// 绘制图像
	// image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
	// image.ZP ZP is the zero Point
	// image.Pt Pt is shorthand for Point{X, Y}
	// r image.Rectangle img.Bounds() or img.Bounds().Add(offset)
	draw.Draw(img, r, imgcode, image.Pt(x, y), draw.Over)
}

// JPEG将编码生成图片
// 选择编码参数,质量范围从1到100,更高的是更好 &jpeg.Options{90}
func ImageEncodeJPEG(ph string, img image.Image, option int) error {
	// 确保文件父目录存在
	CheckParents(ph)
	// 打开文件等待写入
	f := FileW(ph)
	// 保证文件正常关闭
	defer f.Close()
	// 写入文件
	return jpeg.Encode(f, img, &jpeg.Options{option})
}

// PNG将编码生成图片
func ImageEncodePNG(ph string, img image.Image) error {
	// 确保文件父目录存在
	CheckParents(ph)
	// 打开文件等待写入
	f := FileW(ph)
	// 保证文件正常关闭
	defer f.Close()
	// 写入文件
	return png.Encode(f, img)
}

// --------------------------------------------
// file

// 调用者将负责关闭文件
func FileA(ph string) *os.File {
	// 确定文件的父目录是存在的
	CheckParents(ph)
	// 打开文件，文件不存在则创建,追加方式
	f, err := os.OpenFile(ph, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if nil != err {
		panic(err)
	}
	return f
}

// 用回调的方式打文件以便追加内容，回调函数不需要关心文件关闭等问题
func FileAF(ph string, callback func(*os.File)) {
	f := FileA(ph)
	if nil != f {
		defer f.Close()
		callback(f)
	}
}

// 调用者将负责关闭文件
func FileW(ph string) *os.File {
	// 确定文件的父目录是存在的
	CheckParents(ph)
	// 打开文件
	f, err := os.OpenFile(ph, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if nil != err {
		panic(err)
	}
	return f
}

// 用回调的方式打文件以便复写内容，回调函数不需要关心文件关闭等问题
func FileWF(ph string, callback func(*os.File)) {
	f := FileW(ph)
	// 开始写入
	if nil != f {
		defer f.Close()
		if nil != callback {
			callback(f)
		}
	}
}

/*
	将从自己磁盘目录，只读的方式打开一个文件。
	如果文件不存在，或者打开错误，则返回 nil。
	调用者将负责关闭文件
*/
func FileR(ph string) *os.File {
	f, err := os.Open(ph)
	if nil != err {
		return nil
	}
	return f
}

// 用回调的方式打文件以便读取内容，回调函数不需要关心文件关闭等问题
func FileRF(ph string, callback func(*os.File)) {
	f := FileR(ph)
	if nil != f {
		defer f.Close()
		callback(f)
	}
}

// 自定义模式打开文件
// 调用者将负责关闭文件
func FileO(ph string, flag int) *os.File {
	// 确定文件的父目录是存在的
	CheckParents(ph)
	// 打开文件
	f, err := os.OpenFile(ph, flag, 0666)
	if nil != err {
		return nil
	}
	return f
}

// 用自定义的模式打文件以便替换内容，回调函数不需要关心文件关闭等问题
func FileOF(ph string, flag int, callback func(*os.File)) {
	f := FileO(ph, flag)
	// 开始写入
	if nil != f {
		defer f.Close()
		if nil != callback {
			callback(f)
		}
	}
}

// 强制覆盖写入文件
func FWrite(path string, data []byte) error {
	// 保证目录存在
	CheckParents(path)
	// 写入文件
	return ioutil.WriteFile(path, data, 0644)
}
