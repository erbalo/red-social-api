package modelos

import (
	"api/src/tipos"
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os metodos para valdar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa tipos.TipoRegistro) error {
	usuario.formatar()

	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa tipos.TipoRegistro) error {
	if usuario.Nome == "" {
		return errors.New("O nome e obrigatorio e nao pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick e obrigatorio e nao pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email e obrigatorio e nao pode estar em branco")
	}

	if etapa == tipos.REGISTRO && usuario.Senha == "" {
		return errors.New("A senha e obrigatoria e nao pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
