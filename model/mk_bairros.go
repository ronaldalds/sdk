package model

import (
	"time"
)

const TableNameMkBairro = "mk_bairros"

type MkBairro struct {
	// Relationsships Init
	Cidade MkCidade `gorm:"foreignKey:codcidade;references:codcidade" json:"cidade,omitempty"`
	// Relationsships End
	Codbairro     int32      `gorm:"column:codbairro;not null" json:"codbairro,omitempty"`
	Bairro        string     `gorm:"column:bairro;not null" json:"bairro,omitempty"`
	Codcidade     int32      `gorm:"column:codcidade;not null" json:"codcidade,omitempty"`
	CdMicroarea   int32      `gorm:"column:cd_microarea" json:"cd_microarea,omitempty"`
	Latitude      string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude     string     `gorm:"column:longitude" json:"longitude,omitempty"`
	NaoEncontrato string     `gorm:"column:nao_encontrato" json:"nao_encontrato,omitempty"`
	IDOperador    int32      `gorm:"column:id_operador" json:"id_operador,omitempty"`
	DtCadastro    *time.Time `gorm:"column:dt_cadastro" json:"dt_cadastro,omitempty"`
	UserOperador  string     `gorm:"column:user_operador" json:"user_operador,omitempty"`
}

func (*MkBairro) TableName() string {
	return TableNameMkBairro
}
