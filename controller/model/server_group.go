package model

import "github.com/JasperStritzke/cubid-cloud/util/datamodel"

type ServerGroup struct {
	Servers []ServerInfo
	Info    datamodel.ServerGroup
}
