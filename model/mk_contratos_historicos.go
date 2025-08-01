package model

import (
	"time"
)

const TableNameMkContratosHistorico = "mk_contratos_historicos"

type MkContratosHistorico struct {
	// Relationsships Init
	Contrato MkContrato `gorm:"foreignKey:codcontrato;references:cd_contrato" json:"contrato,omitempty"`
	// Relationsships End
	Codcontratohist int32      `gorm:"column:codcontratohist;not null" json:"codcontratohist,omitempty"`
	CdContrato      int32      `gorm:"column:cd_contrato" json:"cd_contrato,omitempty"`
	Dt              *time.Time `gorm:"column:dt" json:"dt,omitempty"`
	DtHr            *time.Time `gorm:"column:dt_hr" json:"dt_hr,omitempty"`
	Operador        string     `gorm:"column:operador" json:"operador,omitempty"`
	CdOperacao      int32      `gorm:"column:cd_operacao" json:"cd_operacao,omitempty"`
	TxExtra         string     `gorm:"column:tx_extra" json:"tx_extra,omitempty"`
	CdPlanoVelho    int32      `gorm:"column:cd_plano_velho;comment:usado para eventos de upgrade e downgrade" json:"cd_plano_velho,omitempty"`                // usado para eventos de upgrade e downgrade
	CdPlanoNovo     int32      `gorm:"column:cd_plano_novo;comment:usado para eventos de upgrade e downgrade" json:"cd_plano_novo,omitempty"`                  // usado para eventos de upgrade e downgrade
	Vlr             float64    `gorm:"column:vlr;comment:usado para eventos de upgrade e downgrade tb pode ser usado para outras coisas" json:"vlr,omitempty"` // usado para eventos de upgrade e downgrade tb pode ser usado para outras coisas
	HashControle    string     `gorm:"column:hash_controle" json:"hash_controle,omitempty"`
	Efetivado       string     `gorm:"column:efetivado" json:"efetivado,omitempty"`
}

func (*MkContratosHistorico) TableName() string {
	return TableNameMkContratosHistorico
}
