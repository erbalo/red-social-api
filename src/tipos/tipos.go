package tipos

// TipoRegistro e um tipo pra exportar como constante
type TipoRegistro string

// REGISTRO e um tipo de cadastro
const (
	REGISTRO TipoRegistro = "cadastro"
	EDICION  TipoRegistro = "edicao"
)
