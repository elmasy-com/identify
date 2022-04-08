package identify

import (
	"net"
	"strconv"
)

type IPTypes interface {
	net.IP | *net.IP | net.IPAddr | *net.IPAddr | net.TCPAddr | *net.TCPAddr | net.UDPAddr | *net.UDPAddr | string
}

type PortTypes interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | string
}

// IsValidIPv4 whether ip is valid IPv4 address.
func IsValidIPv4[T IPTypes](ip T) bool {

	var i net.IP

	switch v := any(ip).(type) {
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
	default:
		return false
	}

	return i.To4() != nil
}

// IsValidIPv6 whether ip is valid IPv6 address.
func IsValidIPv6[T IPTypes](ip T) bool {

	var i net.IP

	switch v := any(ip).(type) {
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
	default:
		return false
	}

	return i.To16() != nil
}

func IsValidPort[T PortTypes](p T) bool {

	var (
		port int
		err  error
	)

	switch v := any(p).(type) {
	case int8:
		port = int(v)
	case uint8:
		port = int(v)
	case int16:
		port = int(v)
	case uint16:
		port = int(v)
	case int32:
		port = int(v)
	case uint32:
		port = int(v)
	case int64:
		port = int(v)
	case uint64:
		port = int(v)
	case int:
		port = int(v)
	case uint:
		port = int(v)
	case string:
		port, err = strconv.Atoi(v)
		if err != nil {
			return false
		}
	default:
		return false
	}

	return port > 0 && port < 65535
}
