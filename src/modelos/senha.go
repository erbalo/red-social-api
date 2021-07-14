package modelos

// Senha representa um formato de requisicao de alteraco de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
