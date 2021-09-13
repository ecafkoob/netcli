package pkg

import (
	"net"
)

func GetNS(host string) ([]*net.NS, error) {
	ns, err := net.LookupNS(host)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func GetCNAME(host string) (string, error) {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return "", err
	}
	return cname, nil
}

func GetIP(host string) ([]net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
