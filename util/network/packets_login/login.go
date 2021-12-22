package packets

const LoginRequestAction = "REQUEST_LOGIN"

type PacketAuthLoginRequest struct {
	Key string `json:"key"`
}

const LoginResponseAction = "RESPONSE_LOGIN"

type PacketAuthLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
