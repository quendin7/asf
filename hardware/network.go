package hardware

import (
	"fmt"
	"net"
)

func GetNetworkInfo() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "unknown"
	}

	for _, iface := range interfaces {
		if iface.Name == "eth0" || iface.Name == "wlan0" || iface.Name == "enp5s0" || iface.Name == "wlp3s0" {
			addrs, err := iface.Addrs()
			if err == nil && len(addrs) > 0 {
				return fmt.Sprintf("%s (%s)", iface.Name, addrs[0].String())
			}
		}
	}

	return "No active connection"
}
