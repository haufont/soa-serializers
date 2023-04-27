package udp

import (
	"net"
	"strconv"
	"strings"
)

const (
	BufferSize  = 8192
	DefaultAddr = "0.0.0.0"
)

func JoinAddrAndPort(addr string, port int) string {
	return strings.Join([]string{addr, strconv.Itoa(port)}, ":")
}

func InitUDPServer(port int) *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", JoinAddrAndPort(DefaultAddr, port))
	if err != nil {
		panic(err)
	}
	udpServer, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	return udpServer
}

func InitUDPMulticastServer(addr string, port int) *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", JoinAddrAndPort(addr, port))
	if err != nil {
		panic(err)
	}
	udpServer, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	return udpServer
}

func InitUDPClient(addr string, port int) *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", JoinAddrAndPort(addr, port))
	if err != nil {
		panic(err)
	}
	udpClient, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	return udpClient
}

type IHandler interface {
	Handle(req string) (resp string)
}

func Serve(udpServer *net.UDPConn, handler IHandler) {
	request := make([]byte, BufferSize)
	for {
		sz, addr, err := udpServer.ReadFrom(request)
		if err != nil {
			continue
		}
		resp := handler.Handle(string(request[:sz]))
		_, _ = udpServer.WriteTo([]byte(resp), addr)
	}
}
