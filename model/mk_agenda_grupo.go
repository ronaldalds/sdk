package model

const TableNameMkAgendaGrupo = "mk_agenda_grupo"

type MkAgendaGrupo struct {
	// Relationsships Init
	AgendaGrupoColaborador []MkAgendaGrupoColaborador `gorm:"foreignKey:codagenda_grupo;references:codagenda_grupo" json:"agenda_grupo_colaborador,omitempty"`
	// Relationsships End
	CodagendaGrupo      int32  `gorm:"column:codagenda_grupo;not null" json:"codagenda_grupo,omitempty"`
	Nome                string `gorm:"column:nome;not null" json:"nome,omitempty"`
	HrProcessamentoAuto string `gorm:"column:hr_processamento_auto" json:"hr_processamento_auto,omitempty"`
	HrAprovacaoAuto     string `gorm:"column:hr_aprovacao_auto" json:"hr_aprovacao_auto,omitempty"`
	PrevisaoInpe        string `gorm:"column:previsao_inpe" json:"previsao_inpe,omitempty"`
	AcessoRestrito      string `gorm:"column:acesso_restrito;comment:S - o acesso fica limitado na agenda N - a agenda fica acessivel" json:"acesso_restrito,omitempty"` // S - o acesso fica limitado na agenda N - a agenda fica acessivel
}

func (*MkAgendaGrupo) TableName() string {
	return TableNameMkAgendaGrupo
}
