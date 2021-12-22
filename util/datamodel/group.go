package datamodel

type ServerGroupInfo struct {
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	Memory            int      `json:"memory"`
	MinOnlineServices int      `json:"min_online_services"`
	MaxOnlineServices int      `json:"max_online_services"`
	ExecutorNames     []string `json:"executor_names"`
}
