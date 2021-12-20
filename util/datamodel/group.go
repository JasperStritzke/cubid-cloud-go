package datamodel

type ServerGroup struct {
	Name                  string `json:"name"`
	Memory                int    `json:"memory"`
	MinOnlineServices     int    `json:"min_online_services"`
	MaxOnlineServices     int    `json:"max_online_services"`
	AssociatedExecutorUid string `json:"associated_executor_uid"`
}
