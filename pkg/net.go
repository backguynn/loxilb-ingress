package pkg

import (
	"errors"
	"net"
)

// GetLocalNonLoopBackIP - get a non-loopback IP of this pod
func GetLocalNonLoopBackIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", errors.New("no addresses found")
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("non-lo addresses not found")
}
