package executor

import (
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/network"
)

type Executor struct {
	Info     Info
	endpoint *network.Endpoint
}

type Info struct {
	Name      string `json:"name"`
	SecretKey string `json:"secret_key"`
}

func (e *Executor) IsConnected() bool {
	return e.endpoint != nil
}

func (e *Executor) Connect(endpoint *network.Endpoint) {
	e.endpoint = endpoint
	logger.Infof("Executor %s[%s] connected.", e.Info.Name, endpoint.IP())
}

func (e *Executor) Disconnect() {
	e.endpoint = nil
	logger.Infof("Executor %s connected.", e.Info.Name)
}

func (e *Executor) PrintIllegalConnectionAttempt() {
	logger.Warnf("Executor %s tried to connect with illegal token.", e.Info.Name)
}

func (e *Executor) SendPacket(packet *network.Packet) error {
	if e.endpoint != nil {
		return e.endpoint.SendPacket(packet)
	}

	return nil
}

func (e *Executor) SendPackets(packets ...*network.Packet) error {
	if e.endpoint != nil {
		return e.endpoint.SendPackets(packets...)
	}

	return nil
}
