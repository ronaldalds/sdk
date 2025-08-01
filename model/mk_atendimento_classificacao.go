package model

const TableNameMkAtendimentoClassificacao = "mk_atendimento_classificacao"

type MkAtendimentoClassificacao struct {
	Codatclass            int32  `gorm:"column:codatclass;not null" json:"codatclass,omitempty"`
	Descricao             string `gorm:"column:descricao" json:"descricao,omitempty"`
	Aplicacao             string `gorm:"column:aplicacao" json:"aplicacao,omitempty"`
	HorasFechamento       int32  `gorm:"column:horas_fechamento" json:"horas_fechamento,omitempty"`
	CorFundo              int32  `gorm:"column:cor_fundo" json:"cor_fundo,omitempty"`
	CorBorda              int32  `gorm:"column:cor_borda" json:"cor_borda,omitempty"`
	CorLetra              int32  `gorm:"column:cor_letra" json:"cor_letra,omitempty"`
	GerarIndiceReclamacao string `gorm:"column:gerar_indice_reclamacao" json:"gerar_indice_reclamacao,omitempty"`
	Encerramento          string `gorm:"column:encerramento" json:"encerramento,omitempty"`
	Inativar              string `gorm:"column:inativar;default:N" json:"inativar,omitempty"`
}

func (*MkAtendimentoClassificacao) TableName() string {
	return TableNameMkAtendimentoClassificacao
}
