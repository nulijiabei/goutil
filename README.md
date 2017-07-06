----------------
Goutil

<a href="https://godoc.org/github.com/nulijiabei/goutil"><img src="https://godoc.org/github.com/nulijiabei/goutil?status.svg" alt="GoDoc"></a>

---

	const FORMAT_DATE string = "2006-01-02"
	const FORMAT_DATE_TIME string = "2006-01-02 15:04:05"
	const FORMAT_TIME string = "15:04:05"

	func CheckParents(aph string)
	func FCopy(src string, dst string) error
	func FWrite(path string, data []byte) error
	func Fexists(ph string) bool
	func FileA(ph string) *os.File
	func FileAF(ph string, callback func(*os.File))
	func FileO(ph string, flag int) *os.File
	func FileOF(ph string, flag int, callback func(*os.File))
	func FileR(ph string) *os.File
	func FileRF(ph string, callback func(*os.File))
	func FileW(ph string) *os.File
	func FileWF(ph string, callback func(*os.File))
	func GetEnv(v string) string
	func GetTime(layout string) string
	func HttpGetReq(u string, params map[string]string) ([]byte, error)
	func HttpMultipartPostReq(u string, ff string, params map[string]string) ([]byte, error)
	func HttpPostReq(u string, params map[string]string) ([]byte, error)
	func HttpPostReqPlus(u string, params map[string]interface{}, header map[string]string) ([]byte, error)
	func ImageDrawRGBA(img *image.RGBA, imgcode image.Image, x, y int)
	func ImageDrawRGBAOffSet(img *image.RGBA, imgcode image.Image, r image.Rectangle, x, y int)
	func ImageEncodeJPEG(ph string, img image.Image, option int) error
	func ImageEncodePNG(ph string, img image.Image) error
	func ImageJPEG(ph string) (image.Image, error)
	func ImagePNG(ph string) (image.Image, error)
	func ImageRGBA(width, height int) *image.RGBA
	func IsBlank(s string) bool
	func IsExist(name string) bool
	func IsSpace(c byte) bool
	func JsonMarshalFile(file string, v interface{}) error
	func JsonMarshalIndent(v interface{}) ([]byte, error)
	func JsonMarshalIndentFile(file string, v interface{}) error
	func JsonUnmarshalFile(file string, v interface{}) error
	func NoError(err error)
	func ParseTime(layout, value string) time.Time
	func String2Bool(_str string) bool
	func String2Float(data string, reData float64) float64
	func String2Int(s string, dft int) int
	func String2Int64(s string, dft int64) int64
	func String2Utf8(_str string) string
	func UnixMsSec(off int) int64
	func UnixNano() string
	type BeeMap
	func NewBeeMap() *BeeMap
	func (m *BeeMap) Check(k interface{}) bool
	func (m *BeeMap) Delete(k interface{})
	func (m *BeeMap) Get(k interface{}) interface{}
	func (m *BeeMap) Set(k interface{}, v interface{}) bool
	...
	
---