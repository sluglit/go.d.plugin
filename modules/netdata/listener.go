package netdata

import "net"

type TcpServer struct {
	Addr string

	server net.Listener
}

func (t *TcpServer) Listen() (err error) {
	t.server, err = net.Listen("udp", t.Addr)
	if err != nil {
		return err
	}

	return nil
}

func (t *TcpServer) Close() {
	if t.server != nil {
		t.server.Close()
	}
}
