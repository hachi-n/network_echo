package echo

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func SendEchoMessage(ipString, message string) {
	ipv4Addr := net.ParseIP(ipString).To4()

	var byteIpv4Addr [4]byte
	copy(byteIpv4Addr[:], ipv4Addr)

	inet4 := &syscall.SockaddrInet4{
		// Echo Protocol Welknown Port.
		Port: 7,
		Addr: byteIpv4Addr,
	}

	sd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sd)
		os.Exit(1)
	}

	err = syscall.Connect(sd, inet4)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sd)
		os.Exit(1)
	}

	err = syscall.Sendto(sd, []byte(message), 0, inet4)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sd)
		os.Exit(1)
	}

	buf := make([]byte, len(message))

	_, _, err = syscall.Recvfrom(sd, buf, 0)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sd)
		os.Exit(1)
	}

	fmt.Println(string(buf))

	syscall.Close(sd)
}
