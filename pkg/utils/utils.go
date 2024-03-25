package utils

import (
	"net"
	"time"

	"github.com/digineo/go-ping"
)

func PingHost(destinationIp string) (time.Duration, error) {
	bind := "0.0.0.0"

	var remoteAddr *net.IPAddr
	var pinger *ping.Pinger
	var rtt time.Duration

	remoteAddr = &net.IPAddr{
		IP: net.ParseIP(destinationIp),
	}

	pinger, err := ping.New(bind, "")
	if err != nil {
		return rtt, err
	}
	defer pinger.Close()

	timeout, _ := time.ParseDuration("30s")
	rtt, err = pinger.PingAttempts(remoteAddr, timeout, 3)
	if err != nil {
		return rtt, err
	}

	return rtt, nil
}
