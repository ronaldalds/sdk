package model

import (
	"time"
)

const TableNameMkContratosGerado = "mk_contratos_gerados"

type MkContratosGerado struct {
	Codemissao       int32      `gorm:"column:codemissao;not null" json:"codemissao,omitempty"`
	TextoContrato    string     `gorm:"column:texto_contrato" json:"texto_contrato,omitempty"`
	Contatado        string     `gorm:"column:contatado" json:"contatado,omitempty"`
	Contratante      string     `gorm:"column:contratante" json:"contratante,omitempty"`
	Emissao          *time.Time `gorm:"column:emissao" json:"emissao,omitempty"`
	Contrato         int32      `gorm:"column:contrato" json:"contrato,omitempty"`
	TipoContrato     int32      `gorm:"column:tipo_contrato" json:"tipo_contrato,omitempty"`
	CdPerfilContrato int32      `gorm:"column:cd_perfil_contrato" json:"cd_perfil_contrato,omitempty"`
	Status           int32      `gorm:"column:status;comment:1 aceite pedente 2 aceito" json:"status,omitempty"` // 1 aceite pedente 2 aceito
	DhGeracao        *time.Time `gorm:"column:dh_geracao" json:"dh_geracao,omitempty"`
	Titulo           string     `gorm:"column:titulo" json:"titulo,omitempty"`
	CdRegistroAceite int32      `gorm:"column:cd_registro_aceite" json:"cd_registro_aceite,omitempty"`
	GerarAceite      string     `gorm:"column:gerar_aceite" json:"gerar_aceite,omitempty"`
}

func (*MkContratosGerado) TableName() string {
	return TableNameMkContratosGerado
}
