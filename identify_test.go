package identify

import (
	"net"
	"testing"
)

func TestIsValidIPv4(t *testing.T) {

	ip := "127.0.0.1"

	/*
		net.IP
	*/
	{
		i := net.ParseIP(ip)
		if !IsValidIPv4(i) {
			t.Fatalf("IsValidIPv4() failed with net.IP")
		}
	}

	/*
		net.IPAddr
	*/
	{
		i, err := net.ResolveIPAddr("ip4", ip)
		if err != nil {
			t.Fatalf("Failed to resolve IPAddr with %s: %s", ip, err)
		}
		if !IsValidIPv4(*i) {
			t.Fatalf("IsValidIPv4() failed with net.IPAddr")
		}
		if !IsValidIPv4(i) {
			t.Fatalf("IsValidIPv4() failed with *net.IPAddr")
		}
	}

	/*
		net.TCPAddr
	*/
	{
		i, err := net.ResolveTCPAddr("tcp", ip+":80")
		if err != nil {
			t.Fatalf("Failed to resolve TCPAddr with %s:80 : %s", ip, err)
		}
		if !IsValidIPv4(*i) {
			t.Fatalf("IsValidIPv4() failed with net.TCPAddr")
		}
		if !IsValidIPv4(i) {
			t.Fatalf("IsValidIPv4() failed with *net.TCPAddr")
		}
	}

	/*
		net.UDPAddr
	*/
	{
		i, err := net.ResolveUDPAddr("udp", ip+":80")
		if err != nil {
			t.Fatalf("Failed to resolve UDPAddr with %s:80 : %s", ip, err)
		}
		if !IsValidIPv4(*i) {
			t.Fatalf("IsValidIPv4() failed with net.UDPAddr")
		}
		if !IsValidIPv4(i) {
			t.Fatalf("IsValidIPv4() failed with *net.UDPAddr")
		}
	}

	/*
		string
	*/
	{
		if !IsValidIPv4(ip) {
			t.Fatalf("IsValidIPv4() failed with string")
		}

		if IsValidIPv4("invalid") {
			t.Fatalf("IsValidIPv4() failed with string")
		}
	}
}

func TestIsValidIPv6(t *testing.T) {

	ip := "fe80::1"

	/*
		net.IP
	*/
	{
		i := net.ParseIP(ip)
		if !IsValidIPv6(i) {
			t.Fatalf("IsValidIPv6() failed with net.IP")
		}
	}

	/*
		net.IPAddr
	*/
	{
		i, err := net.ResolveIPAddr("ip6", ip)
		if err != nil {
			t.Fatalf("Failed to resolve IPAddr with [%s]: %s", ip, err)
		}
		if !IsValidIPv6(*i) {
			t.Fatalf("IsValidIPv6() failed with net.IPAddr")
		}
		if !IsValidIPv6(i) {
			t.Fatalf("IsValidIPv6() failed with *net.IPAddr")
		}
	}

	/*
		net.TCPAddr
	*/
	{
		i, err := net.ResolveTCPAddr("tcp", "["+ip+"]:80")
		if err != nil {
			t.Fatalf("Failed to resolve TCPAddr with [%s]:80 : %s", ip, err)
		}
		if !IsValidIPv6(*i) {
			t.Fatalf("IsValidIPv6() failed with net.TCPAddr")
		}
		if !IsValidIPv6(i) {
			t.Fatalf("IsValidIPv6() failed with *net.TCPAddr")
		}
	}

	/*
		net.UDPAddr
	*/
	{
		i, err := net.ResolveUDPAddr("udp", "["+ip+"]:80")
		if err != nil {
			t.Fatalf("Failed to resolve UDPAddr with [%s]:80 : %s", ip, err)
		}
		if !IsValidIPv6(*i) {
			t.Fatalf("IsValidIPv6() failed with net.UDPAddr")
		}
		if !IsValidIPv6(i) {
			t.Fatalf("IsValidIPv6() failed with *net.UDPAddr")
		}
	}

	/*
		string
	*/
	{
		if !IsValidIPv6(ip) {
			t.Fatalf("IsValidIPv6() failed with string")
		}

		if IsValidIPv6("invalid") {
			t.Fatalf("IsValidIPv6() failed with string")
		}
	}
}

func TestIsValidIP(t *testing.T) {

	ip4 := "127.0.0.1"
	ip6 := "fe80::1"

	/*
		net.IP
	*/
	{
		i := net.ParseIP(ip4)
		if !IsValidIP(i) {
			t.Fatalf("IsValidIP() failed with net.IP")
		}
	}

	/*
		net.IPAddr
	*/
	{
		i, err := net.ResolveIPAddr("ip4", ip4)
		if err != nil {
			t.Fatalf("Failed to resolve IPAddr with %s: %s", ip4, err)
		}
		if !IsValidIP(*i) {
			t.Fatalf("IsValidIP() failed with net.IPAddr")
		}
		if !IsValidIP(i) {
			t.Fatalf("IsValidIP() failed with *net.IPAddr")
		}
	}

	/*
		net.TCPAddr
	*/
	{
		i, err := net.ResolveTCPAddr("tcp", "["+ip6+"]:80")
		if err != nil {
			t.Fatalf("Failed to resolve TCPAddr with [%s]:80 : %s", ip6, err)
		}
		if !IsValidIP(*i) {
			t.Fatalf("IsValidIP() failed with net.TCPAddr")
		}
		if !IsValidIP(i) {
			t.Fatalf("IsValidIP() failed with *net.TCPAddr")
		}
	}

	/*
		net.UDPAddr
	*/
	{
		i, err := net.ResolveUDPAddr("udp", ip4+":80")
		if err != nil {
			t.Fatalf("Failed to resolve UDPAddr with %s:80 : %s", ip4, err)
		}
		if !IsValidIP(*i) {
			t.Fatalf("IsValidIP() failed with net.UDPAddr")
		}
		if !IsValidIP(i) {
			t.Fatalf("IsValidIP() failed with *net.UDPAddr")
		}
	}

	/*
		string
	*/
	{
		if !IsValidIP(ip6) {
			t.Fatalf("IsValidIPv6() failed with string")
		}

		if IsValidIP("invalid") {
			t.Fatalf("IsValidIPv6() failed with string")
		}
	}
}
func TestIsValidPort(t *testing.T) {
	if !IsValidPort(80) {
		t.Fatalf("80 reported as invalid port!")
	}

	if !IsValidPort(uint8(80)) {
		t.Fatalf("80 reported as invalid port!")
	}

	if !IsValidPort(uint8(100)) {
		t.Fatalf("uint8(100) reported as an invalid port!")
	}

	if !IsValidPort(uint16(40000)) {
		t.Fatalf("uint16(40000) reported as an ivalid port!")
	}

	if IsValidPort(int32(400000)) {
		t.Fatalf("int32(400000) reported as a valid port!")
	}

	if IsValidPort(uint64(4000000)) {
		t.Fatalf("uint64(4000000) reported as a valid port!")
	}

	if IsValidPort(uint(40000000)) {
		t.Fatalf("uint(40000000) reported as a valid port!")
	}
}
