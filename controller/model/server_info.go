package model

const (
	StateLobby   = "LOBBY"
	StateInGame  = "INGAME"
	StateOffline = "OFFLINE"
)

type ServerId struct {
	Name string
	Id   int
}

type ServerInfo struct {
	Id          ServerId
	ServerGroup string
	MOTD        string
	MaxPlayers  int
	Players     []Player
	State       string
}

func (s *ServerInfo) IsLobby() bool {
	return s.State == StateLobby
}

func (s *ServerInfo) IsInGame() bool {
	return s.State == StateInGame
}

func (s *ServerInfo) IsOffline() bool {
	return s.State == StateOffline
}

func (s *ServerInfo) IsCustomState() bool {
	return !(s.IsLobby() && s.IsInGame() && s.IsOffline())
}
