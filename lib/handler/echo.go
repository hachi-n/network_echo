package handler

import (
	"fmt"
	"os"
	"syscall"
)

func EchoHandling(nfd int, sockaddr syscall.Sockaddr) {

	buf := make([]byte, 5)
	recvMessageSize, _, err := syscall.Recvfrom(nfd, buf, 0)
	if err != nil {
		fmt.Println("recvfrom() error", err)
		syscall.Close(nfd)
		os.Exit(1)
	}

	for recvMessageSize > 0 {
		err = syscall.Sendto(nfd, buf, 0, sockaddr)
		if err != nil {
			fmt.Println("sendto() error", err)
			syscall.Close(nfd)
			os.Exit(1)
		}

		recvMessageSize, _, err = syscall.Recvfrom(nfd, buf, 0)
		if err != nil {
			fmt.Println("loop recvfrom() error", err)
			syscall.Close(nfd)
			os.Exit(1)
		}
	}

	syscall.Close(nfd)
}
