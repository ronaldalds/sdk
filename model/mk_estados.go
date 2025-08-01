package model

const TableNameMkEstado = "mk_estados"

type MkEstado struct {
	Codestado     int32  `gorm:"column:codestado;not null" json:"codestado,omitempty"`
	Siglaestado   string `gorm:"column:siglaestado" json:"siglaestado,omitempty"`
	Nomeestado    string `gorm:"column:nomeestado;not null" json:"nomeestado,omitempty"`
	Ibge          int32  `gorm:"column:ibge" json:"ibge,omitempty"`
	NaoEncontrato string `gorm:"column:nao_encontrato" json:"nao_encontrato,omitempty"`
}

func (*MkEstado) TableName() string {
	return TableNameMkEstado
}
