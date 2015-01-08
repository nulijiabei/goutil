package goutil

import (
	"fmt"
	"net"
	"strings"
)

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

func main() {
	fmt.Println(GetInterfaceByName("en0"))
}
