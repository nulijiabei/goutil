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
	var HardwareAddr string
	for _, v := range strings.Split(inter.HardwareAddr.String(), ":") {
		HardwareAddr += v
	}
	addrs, err := inter.Addrs()
	if err != nil {
		return nil, err
	}
	for _, v := range addrs {
		addr := strings.Split(v.String(), "/")
		if len(strings.Split(addr[0], ".")) == 4 {
			return &Interface{addr[0], strings.ToUpper(HardwareAddr)}, nil
		}
	}
	return nil, fmt.Errorf("...")
}

func main() {
	fmt.Println(GetInterfaceByName("en0"))
}
