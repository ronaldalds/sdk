package model

const TableNameMkCore = "mk_cores"

type MkCore struct {
	Codcor   int32  `gorm:"column:codcor;not null" json:"codcor,omitempty"`
	CorFundo string `gorm:"column:cor_fundo" json:"cor_fundo,omitempty"`
	CorTx    string `gorm:"column:cor_tx" json:"cor_tx,omitempty"`
	CorBorda string `gorm:"column:cor_borda" json:"cor_borda,omitempty"`
	CorLetra string `gorm:"column:cor_letra" json:"cor_letra,omitempty"`
}

func (*MkCore) TableName() string {
	return TableNameMkCore
}
