/*
Provides support functionality used by Eureka client.
*/
package util

import "net"

/*
Gets IP address of current machine as a string, also returns a bool to indicate if
the operation was successful.
*/
func GetIP() (string, bool) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", false
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), true
			}
		}
	}
	return "", false
}
