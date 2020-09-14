package tcp

import (
	"fmt"
	"github.com/hachi-n/network_echo/lib/handler"
	"net"
	"syscall"
)

type TcpServer struct {
	QueueLimit int

	inetV4 *syscall.SockaddrInet4
	fd     int
}

func NewTcpServer(ip string, port int, queueLimit int) *TcpServer {
	var byteIpv4Addr [4]byte
	ipv4Addr := net.ParseIP(ip).To4()
	copy(byteIpv4Addr[:], ipv4Addr)

	inet4 := &syscall.SockaddrInet4{
		// Echo Protocol Wellknown Port.
		Port: port,
		Addr: byteIpv4Addr,
	}

	return &TcpServer{
		inetV4:     inet4,
		QueueLimit: queueLimit,
	}
}

func (t *TcpServer) Socket() (err error) {
	t.fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	return
}

func (t *TcpServer) Bind() (err error) {
	return syscall.Bind(t.fd, t.inetV4)
}

func (t *TcpServer) Serve() (err error) {
	err = syscall.Listen(t.fd, t.QueueLimit)
	if err != nil {
		return
	}

	fmt.Println("TCP Server Start!")
	for {
		nfd, sockAddr, err := syscall.Accept(t.fd)
		if err != nil {
			return err
		}
		handler.EchoHandling(nfd, sockAddr)
	}
}

func (t *TcpServer) Close() (err error) {
	return syscall.Close(t.fd)
}
