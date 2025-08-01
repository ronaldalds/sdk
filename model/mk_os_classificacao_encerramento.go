package model

const TableNameMkOsClassificacaoEncerramento = "mk_os_classificacao_encerramento"

type MkOsClassificacaoEncerramento struct {
	Codclassifenc int32  `gorm:"column:codclassifenc;not null" json:"codclassifenc,omitempty"`
	Classificacao string `gorm:"column:classificacao" json:"classificacao,omitempty"`
	Anotacoes     string `gorm:"column:anotacoes" json:"anotacoes,omitempty"`
	Inativar      string `gorm:"column:inativar" json:"inativar,omitempty"`
}

func (*MkOsClassificacaoEncerramento) TableName() string {
	return TableNameMkOsClassificacaoEncerramento
}
