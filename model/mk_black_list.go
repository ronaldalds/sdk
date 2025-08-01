package model

import (
	"time"
)

const TableNameMkBlackList = "mk_black_list"

type MkBlackList struct {
	// Relationsships Init
	Logradouro MkLogradouro `gorm:"foreignKey:codlogradouro;references:cd_logradouro" json:"logradouro,omitempty"`
	// Relationsships End
	NomeCliente           string     `gorm:"column:nome_cliente" json:"nome_cliente,omitempty"`
	Cpf                   string     `gorm:"column:cpf" json:"cpf,omitempty"`
	Cnpj                  string     `gorm:"column:cnpj" json:"cnpj,omitempty"`
	CdLogradouro          int32      `gorm:"column:cd_logradouro" json:"cd_logradouro,omitempty"`
	Codblacklist          int32      `gorm:"column:codblacklist;not null" json:"codblacklist,omitempty"`
	DtEntrada             *time.Time `gorm:"column:dt_entrada" json:"dt_entrada,omitempty"`
	HrEntrada             string     `gorm:"column:hr_entrada" json:"hr_entrada,omitempty"`
	DtSaida               *time.Time `gorm:"column:dt_saida" json:"dt_saida,omitempty"`
	HrSaida               string     `gorm:"column:hr_saida" json:"hr_saida,omitempty"`
	CdCliente             int32      `gorm:"column:cd_cliente" json:"cd_cliente,omitempty"`
	CdUf                  int32      `gorm:"column:cd_uf" json:"cd_uf,omitempty"`
	CdCidade              int32      `gorm:"column:cd_cidade" json:"cd_cidade,omitempty"`
	CdBairro              int32      `gorm:"column:cd_bairro" json:"cd_bairro,omitempty"`
	Numero                int32      `gorm:"column:numero" json:"numero,omitempty"`
	Complemento           string     `gorm:"column:complemento" json:"complemento,omitempty"`
	ConsiderarComplemento string     `gorm:"column:considerar_complemento" json:"considerar_complemento,omitempty"`
	Cep                   string     `gorm:"column:cep" json:"cep,omitempty"`
	CdUraRet              int32      `gorm:"column:cd_ura_ret" json:"cd_ura_ret,omitempty"`
	Descrito              string     `gorm:"column:descrito" json:"descrito,omitempty"`
	Removido              string     `gorm:"column:removido" json:"removido,omitempty"`
	Acao                  string     `gorm:"column:acao" json:"acao,omitempty"`
	Numeroap              int32      `gorm:"column:numeroap" json:"numeroap,omitempty"`
	Bloco                 string     `gorm:"column:bloco" json:"bloco,omitempty"`
	Referencia            string     `gorm:"column:referencia" json:"referencia,omitempty"`
}

func (*MkBlackList) TableName() string {
	return TableNameMkBlackList
}
