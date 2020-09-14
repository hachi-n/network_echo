package tcp

import (
	"net"
	"syscall"
)

type TcpClient struct {
	inetV4 *syscall.SockaddrInet4
	fd     int
}

func NewTcpClient(ip string, port int) *TcpClient {
	var byteIpv4Addr [4]byte
	ipv4Addr := net.ParseIP(ip).To4()
	copy(byteIpv4Addr[:], ipv4Addr)

	inet4 := &syscall.SockaddrInet4{
		// Echo Protocol Wellknown Port.
		Port: port,
		Addr: byteIpv4Addr,
	}

	return &TcpClient{
		inetV4: inet4,
	}
}

func (t *TcpClient) Socket() (err error) {
	t.fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	return
}

func (t *TcpClient) Connect() (err error) {
	err = syscall.Connect(t.fd, t.inetV4)
	return
}

func (t *TcpClient) Send(message string) (err error) {
	err = syscall.Sendto(t.fd, []byte(message), 0, t.inetV4)
	return
}

func (t *TcpClient) Recv(buf []byte) (recvMessage []byte, err error) {
	_, _, err = syscall.Recvfrom(t.fd, buf, 0)
	recvMessage = buf
	return
}

func (t *TcpClient) Close() (err error) {
	return syscall.Close(t.fd)
}
