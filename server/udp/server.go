package udp

import (
	"fmt"
	"github.com/hachi-n/network_echo/lib/handler"
	"net"
	"syscall"
)

type UdpServer struct {
	QueueLimit int

	inetV4 *syscall.SockaddrInet4
	fd     int
}

func NewUdpServer(ip string, port int, queueLimit int) *UdpServer {
	var byteIpv4Addr [4]byte
	ipv4Addr := net.ParseIP(ip).To4()
	copy(byteIpv4Addr[:], ipv4Addr)

	inet4 := &syscall.SockaddrInet4{
		// Echo Protocol Wellknown Port.
		Port: port,
		Addr: byteIpv4Addr,
	}

	return &UdpServer{
		inetV4:     inet4,
		QueueLimit: queueLimit,
	}
}

func (t *UdpServer) Socket() (err error) {
	t.fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	return
}

func (t *UdpServer) Bind() (err error) {
	return syscall.Bind(t.fd, t.inetV4)
}

func (t *UdpServer) Serve() (err error) {
	err = syscall.Listen(t.fd, t.QueueLimit)
	if err != nil {
		return
	}

	fmt.Println("UDP Server Start!")
	for {
		handler.EchoHandling(t.fd, t.inetV4)
	}
}

func (t *UdpServer) Close() (err error) {
	return syscall.Close(t.fd)
}
