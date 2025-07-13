package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func newServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (s *Server) Handle(conn net.Conn) {
	//
	fmt.Println("link success")

}
func (s *Server) start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	// close listen socket

	defer listener.Close()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept fail", err)
			continue
		}

		//do handle
		go s.Handle(conn)
	}
}
