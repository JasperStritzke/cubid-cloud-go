package server_group

import (
	"github.com/JasperStritzke/cubid-cloud/controller/model"
	"github.com/JasperStritzke/cubid-cloud/util/datamodel"
)

type ServerGroup struct {
	Info    datamodel.ServerGroupInfo
	Servers []model.ServerInfo
}
