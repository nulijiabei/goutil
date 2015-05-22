package goutil

import (
	"fmt"
	"net"
	"strings"
)

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
	if IsBlank(Trim(HardwareAddr)) {
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
