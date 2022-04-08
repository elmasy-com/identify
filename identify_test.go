package identify

import (
	"net"
	"testing"
)

func TestGetIPVersion(t *testing.T) {

	ip4 := "127.0.0.1"
	ip6 := "fe80::1"

	/*
		net.IP
	*/
	{
		i := net.ParseIP(ip4)
		if v := GetIPVersion(i); v != IPV4 {
			t.Fatalf("Failed to get the IP version with %s (net.IP), got %d", ip4, v)
		}

		i = net.ParseIP(ip6)
		if v := GetIPVersion(&i); v != IPV6 {
			t.Fatalf("Failed to get the IP version with %s (*net.IP), got %d", ip6, v)
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
		if v := GetIPVersion(*i); v != IPV4 {
			t.Fatalf("Failed to get the IP version with %s (net.IPAddr), got %d", ip4, v)
		}

		i, err = net.ResolveIPAddr("ip6", ip6)
		if err != nil {
			t.Fatalf("Failed to resolve IPAddr with %s: %s", ip6, err)
		}
		if v := GetIPVersion(i); v != IPV6 {
			t.Fatalf("Failed to get the IP version with %s (*net.IPAddr), got %d", ip6, v)
		}
	}

	/*
		net.TCPAddr
	*/
	{
		i, err := net.ResolveTCPAddr("tcp", ip4+":80")
		if err != nil {
			t.Fatalf("Failed to resolve TCPAddr with %s:80: %s", ip4, err)
		}
		if v := GetIPVersion(*i); v != IPV4 {
			t.Fatalf("Failed to get the IP version with %s (net.TCPAddr), got %d", ip4, v)
		}

		i, err = net.ResolveTCPAddr("tcp", "["+ip6+"]:80")
		if err != nil {
			t.Fatalf("Failed to resolve TCPAddr with [%s]:80 : %s", ip6, err)
		}
		if v := GetIPVersion(i); v != IPV6 {
			t.Fatalf("Failed to get the IP version with [%s]:80 (*net.TCPAddr), got %d", ip6, v)
		}
	}

	/*
		net.UDPAddr
	*/
	{
		i, err := net.ResolveTCPAddr("tcp", ip4+":80")
		if err != nil {
			t.Fatalf("Failed to resolve UDPAddr with %s:80: %s", ip4, err)
		}
		if v := GetIPVersion(*i); v != IPV4 {
			t.Fatalf("Failed to get the IP version with %s (net.UDPAddr), got %d", ip4, v)
		}

		i, err = net.ResolveTCPAddr("tcp", "["+ip6+"]:80")
		if err != nil {
			t.Fatalf("Failed to resolve UDPAddr with [%s]:80 : %s", ip6, err)
		}
		if v := GetIPVersion(i); v != IPV6 {
			t.Fatalf("Failed to get the IP version with [%s]:80 (*net.UDPAddr), got %d", ip6, v)
		}
	}

	/*
		string
	*/
	if v := GetIPVersion(ip4); v != IPV4 {
		t.Fatalf("Failed to get the IP version with%s, got %d", ip4, v)
	}

	if v := GetIPVersion(ip6); v != IPV6 {
		t.Fatalf("Failed to get the IP version with %s, got %d", ip6, v)
	}

}
