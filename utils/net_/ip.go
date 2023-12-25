package net_

import (
	"MyCodeArchive_Go/utils/fault"
	"errors"
	"net"
	"strings"
)

// CheckIps 对ip的数量上下限进行检测，能接收ip，网段和*。
func CheckIps(ips []string) *fault.Fault {
	if len(ips) == 0 {
		return fault.ParamEmpty("policy ip")
	}
	if len(ips) > 20 {
		return fault.ParamCount("policy ip", 20)
	}

	// ip检测
	for _, ip := range ips {
		if ip == "*" {
			continue
		}
		isCIDR := strings.Contains(ip, "/")
		if isCIDR {
			if _, _, err := net.ParseCIDR(ip); err != nil {
				return fault.ParseIp(ip, err)
			}
			continue
		}
		if net.ParseIP(ip) == nil {
			return fault.ParseIp(ip, errors.New(""))
		}
	}
	return nil
}
