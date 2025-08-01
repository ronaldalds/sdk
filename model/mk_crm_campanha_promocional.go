package model

import (
	"time"
)

const TableNameMkCrmCampanhaPromocional = "mk_crm_campanha_promocional"

type MkCrmCampanhaPromocional struct {
	Codcamppromo int32      `gorm:"column:codcamppromo;not null" json:"codcamppromo,omitempty"`
	Descricao    string     `gorm:"column:descricao;not null" json:"descricao,omitempty"`
	Observacoes  string     `gorm:"column:observacoes" json:"observacoes,omitempty"`
	DataInicial  *time.Time `gorm:"column:data_inicial" json:"data_inicial,omitempty"`
	DataFinal    *time.Time `gorm:"column:data_final" json:"data_final,omitempty"`
}

func (*MkCrmCampanhaPromocional) TableName() string {
	return TableNameMkCrmCampanhaPromocional
}
