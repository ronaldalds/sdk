package model

const TableNameMkContratosOperaco = "mk_contratos_operacoes"

type MkContratosOperaco struct {
	Codcontratooperacao int32  `gorm:"column:codcontratooperacao;not null" json:"codcontratooperacao,omitempty"`
	DescricaoOperacao   string `gorm:"column:descricao_operacao" json:"descricao_operacao,omitempty"`
}

func (*MkContratosOperaco) TableName() string {
	return TableNameMkContratosOperaco
}
