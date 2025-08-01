package model

import (
	"time"
)

const TableNameMkCidade = "mk_cidades"

type MkCidade struct {
	// Relationsships Init
	Estado MkEstado `gorm:"foreignKey:codestado;references:codestado" json:"estado,omitempty"`
	// Relationsships End
	Codcidade     int32      `gorm:"column:codcidade;not null" json:"codcidade,omitempty"`
	Codestado     int32      `gorm:"column:codestado;not null" json:"codestado,omitempty"`
	Cidade        string     `gorm:"column:cidade;not null" json:"cidade,omitempty"`
	Ibge          string     `gorm:"column:ibge" json:"ibge,omitempty"`
	Latitude      string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude     string     `gorm:"column:longitude" json:"longitude,omitempty"`
	NaoEncontrato string     `gorm:"column:nao_encontrato" json:"nao_encontrato,omitempty"`
	IDOperador    int32      `gorm:"column:id_operador" json:"id_operador,omitempty"`
	DtCadastro    *time.Time `gorm:"column:dt_cadastro" json:"dt_cadastro,omitempty"`
	UserOperador  string     `gorm:"column:user_operador" json:"user_operador,omitempty"`
}

func (*MkCidade) TableName() string {
	return TableNameMkCidade
}
