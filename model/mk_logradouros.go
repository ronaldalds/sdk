package model

import (
	"time"
)

const TableNameMkLogradouro = "mk_logradouros"

type MkLogradouro struct {
	// Relationsships Init
	Bairro MkBairro `gorm:"foreignKey:codbairro;references:codbairro" json:"bairro,omitempty"`
	// Relationsships End
	Codlogradouro int32      `gorm:"column:codlogradouro;not null" json:"codlogradouro,omitempty"`
	Logradouro    string     `gorm:"column:logradouro;not null" json:"logradouro,omitempty"`
	Codbairro     int32      `gorm:"column:codbairro;not null" json:"codbairro,omitempty"`
	Cep           string     `gorm:"column:cep" json:"cep,omitempty"`
	Codcidade     int32      `gorm:"column:codcidade;not null" json:"codcidade,omitempty"`
	NaoEncontrato string     `gorm:"column:nao_encontrato" json:"nao_encontrato,omitempty"`
	IDOperador    int32      `gorm:"column:id_operador" json:"id_operador,omitempty"`
	DtCadastro    *time.Time `gorm:"column:dt_cadastro" json:"dt_cadastro,omitempty"`
	UserOperador  string     `gorm:"column:user_operador" json:"user_operador,omitempty"`
}

func (*MkLogradouro) TableName() string {
	return TableNameMkLogradouro
}
