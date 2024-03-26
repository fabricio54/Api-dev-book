package usermodels

// Password representa um formato de uma estrutura para atualizar senha
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
