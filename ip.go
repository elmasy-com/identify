package identify

import (
	"net"
)

type IPVersion uint8

const (
	INVALID IPVersion = 0
	IPV4    IPVersion = 4
	IPV6    IPVersion = 6
)

// GetIPVersion returns the IP version of ip, or INVALID if not valid or not an IP address at all.
// The known types (and its pointers) that can hold an IP address: net.IP, net.IPAddr, net.TCPAddr, net.UDPAddr and string.
func GetIPVersion(ip any) IPVersion {

	var i net.IP

	switch v := ip.(type) {
	case net.IP:
		i = v
	case *net.IP:
		i = *v
	case net.IPAddr:
		i = v.IP
	case *net.IPAddr:
		i = v.IP
	case net.TCPAddr:
		i = v.IP
	case *net.TCPAddr:
		i = v.IP
	case net.UDPAddr:
		i = v.IP
	case *net.UDPAddr:
		i = v.IP
	case string:
		i = net.ParseIP(v)
	case *string:
		i = net.ParseIP(*v)
	default:
		return INVALID
	}

	switch {
	case i == nil:
		return INVALID
	case i.To4() != nil:
		return IPV4
	case i.To16() != nil:
		return IPV6
	default:
		return INVALID
	}

}
