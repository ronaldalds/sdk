package model

const TableNameMkNotificacaoParadaConexo = "mk_notificacao_parada_conexoes"

type MkNotificacaoParadaConexo struct {
	Codconexaoparada int32 `gorm:"column:codconexaoparada;not null" json:"codconexaoparada,omitempty"`
	CdParada         int32 `gorm:"column:cd_parada" json:"cd_parada,omitempty"`
	CdConexao        int32 `gorm:"column:cd_conexao" json:"cd_conexao,omitempty"`
}

func (*MkNotificacaoParadaConexo) TableName() string {
	return TableNameMkNotificacaoParadaConexo
}
