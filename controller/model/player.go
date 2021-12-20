package model

type Player struct {
	Name       string
	UUID       string
	ServerName string
	ProxyName  string
}

func (p *Player) Connect(server string) {

}
