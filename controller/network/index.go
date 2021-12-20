package network

import (
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"github.com/JasperStritzke/cubid-cloud/util/network"
)

func Start() {
	_ = config.InitConfigIfNotExists("net.json", fileutil.CodeTypeJSON, func() interface{} {
		return Config{
			Host: "127.0.0.1:4050",
		}
	})

	var cfg Config
	_ = config.LoadConfig("net.json", fileutil.CodeTypeJSON, &cfg)

	network.NewServer(cfg.Host, *network.NewListener(
		onConnect,
		onPacket,
		onDisconnect,
	))
}

func onConnect(endpoint *network.Endpoint) {

}

func onDisconnect(endpoint *network.Endpoint) {

}

func onPacket(endpoint *network.Endpoint, packet network.Packet) {

}
