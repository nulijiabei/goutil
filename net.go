package goutil

import (
	"fmt"
	"net"
	"strings"
)

// IP地址和MAC地址
type Interface struct {
	IPAddr       string
	HardwareAddr string
}

// 通过接口名称获取(IP地址和MAC地址)
func GetInterfaceByName(_interface string) (*Interface, error) {
	inter, err := net.InterfaceByName(_interface)
	if err != nil {
		return nil, err
	}
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var HardwareAddr string
	for _, v := range strings.Split(inter.HardwareAddr.String(), ":") {
		HardwareAddr += v
	}
	var IPAddr string
	for k, v := range addr {
		if k == inter.Index {
			IPAddr = v.String()
		}
	}
	return &Interface{strings.Split(IPAddr, "/")[0], strings.ToUpper(HardwareAddr)}, nil
}

func main() {
	fmt.Println(GetInterfaceByName("en0"))
}
