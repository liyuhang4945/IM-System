package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	s    *Server
}

func (u *User) Online() {
	u.s.mapLock.Lock()
	u.s.OnlineMap[u.Name] = u
	u.s.mapLock.Unlock()

	//广播消息
	u.s.BroadCast(u, "已上线")
}
func (u *User) Offline() {
	u.s.mapLock.Lock()
	delete(u.s.OnlineMap, u.Name)
	u.s.mapLock.Unlock()

	//广播消息
	u.s.BroadCast(u, "下线")
}
func (u *User) DoMessage(msg string) {
	u.s.BroadCast(u, msg)
}

// new User
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
		s:    server,
	}

	// 启动当前user channel
	go user.ListenMessage()
	return user
}

// 监听user channel 有消息发给客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
