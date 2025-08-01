package model

import (
	"time"
)

const TableNameMkCompromissoPessoa = "mk_compromisso_pessoa"

type MkCompromissoPessoa struct {
	// Relationsships Init
	Colaborador MkColaborador `gorm:"foreignKey:codcolaborador;references:cdpessoa" json:"colaborador,omitempty"`
	// Relationsships End
	Codcompromisso int32      `gorm:"column:codcompromisso;not null" json:"codcompromisso,omitempty"`
	Cdpessoa       int32      `gorm:"column:cdpessoa;not null" json:"cdpessoa,omitempty"`
	Papel          string     `gorm:"column:papel;not null" json:"papel,omitempty"`
	Confirmou      string     `gorm:"column:confirmou;not null;default:P" json:"confirmou,omitempty"`
	DhInclusao     *time.Time `gorm:"column:dh_inclusao;not null;default:now()" json:"dh_inclusao,omitempty"`
	DhConfirmacao  *time.Time `gorm:"column:dh_confirmacao" json:"dh_confirmacao,omitempty"`
}

func (*MkCompromissoPessoa) TableName() string {
	return TableNameMkCompromissoPessoa
}
