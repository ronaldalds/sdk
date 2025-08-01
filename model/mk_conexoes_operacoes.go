package model

const TableNameMkConexoesOperaco = "mk_conexoes_operacoes"

type MkConexoesOperaco struct {
	Codconexaohistorico int32  `gorm:"column:codconexaohistorico;not null" json:"codconexaohistorico,omitempty"`
	Descricao           string `gorm:"column:descricao" json:"descricao,omitempty"`
}

func (*MkConexoesOperaco) TableName() string {
	return TableNameMkConexoesOperaco
}
