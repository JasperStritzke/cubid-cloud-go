package network

import (
	"encoding/json"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"net"
)

type Server struct {
	service     string
	tcpListener *net.TCPListener
	listeners   Listener
}

func NewServer(service string, listener Listener) *Server {
	return &Server{
		service:   service,
		listeners: listener,
	}
}

func (server *Server) Stop() error {
	err := server.tcpListener.Close()
	return err
}

func (server *Server) Start() {
	logger.Info("Starting server...")

	tcpAddr, err := net.ResolveTCPAddr("tcp", server.service)

	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	server.tcpListener = listener

	if err != nil {
		panic(err)
	}

	logger.Info("Listening on " + server.service + ".")

	for {
		conn, connErr := listener.Accept()

		if connErr != nil {
			if conn != nil {
				_ = conn.Close()
			}
			continue
		}

		go server.handleClient(conn)
	}
}

func (server *Server) handleClient(conn net.Conn) {
	logger.Info("Client from [" + conn.RemoteAddr().String() + "] connected.")

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	endpoint := Endpoint{
		conn:    conn,
		encoder: encoder,
	}

	server.listeners.ConnectListener(&endpoint)

	for {
		var pkg Packet
		err := decoder.Decode(&pkg)

		if err != nil {
			_ = conn.Close()
			break
		}

		server.listeners.PacketListener(&endpoint, pkg)
	}

	server.listeners.DisconnectListener(&endpoint)
	logger.Info("Client from [" + conn.RemoteAddr().String() + "] disconnected.")
}
