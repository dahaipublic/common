package util

import (
	"net"

	"github.com/ip2location/ip2location-go"
)

func GetLocalIP() (ipv4 string, err error) {
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	// 遍历每个网络接口
	for _, iface := range interfaces {
		// 排除无效接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		// 遍历每个地址
		for _, addr := range addrs {
			// 将地址转换为 IP 地址类型
			ip, ok := addr.(*net.IPNet)
			if !ok || ip.IP.IsLoopback() {
				continue
			}

			// 检查 IP 地址版本
			if ip.IP.To4() != nil {
				ipv4 = ip.IP.String()
				return ipv4, nil
			}
		}
	}
	return
}

func GetIPAttribution(ip string) (ipInfo ip2location.IP2Locationrecord, err error) {
	if ip == "::1" || ip == "127.0.0.1" {
		return
	}
	ipdb, err := ip2location.OpenDB("./IP-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ISP-DOMAIN-MOBILE.BIN")
	if err != nil {
		return
	}
	ipInfo, err = ipdb.Get_all(ip)
	if err != nil {
		ipInfo = ip2location.IP2Locationrecord{}
		return
	}
	return
}
