package echo

import (
	"fmt"
	"github.com/hachi-n/network_echo/lib/handler"
	"net"
	"os"
	"syscall"
)

func ListenAndServe(ipString string, queueLimit int) {
	ipv4Addr := net.ParseIP(ipString).To4()

	var byteIpv4Addr [4]byte
	copy(byteIpv4Addr[:], ipv4Addr)

	inet4 := &syscall.SockaddrInet4{
		// Echo Protocol Welknown Port.
		Port: 7,
		Addr: byteIpv4Addr,
	}

	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sock)
		os.Exit(1)
	}

	err = syscall.Bind(sock, inet4)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sock)
		os.Exit(1)
	}

	err = syscall.Listen(sock, queueLimit)
	if err != nil {
		fmt.Println(err)
		syscall.Close(sock)
		os.Exit(1)
	}

	fmt.Println("Listen Start!")
	for {
		nfd, sockAddr, err := syscall.Accept(sock)
		if err != nil {
			fmt.Println(err)
			syscall.Close(sock)
			os.Exit(1)
		}

		handler.EchoHandling(nfd, sockAddr)
	}
}
