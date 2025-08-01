package model

const TableNameMkAgendaGrupoColaborador = "mk_agenda_grupo_colaborador"

type MkAgendaGrupoColaborador struct {
	// Relationsships Init
	AgendaGrupo []MkAgendaGrupo `gorm:"foreignKey:codagenda_grupo;references:codagenda_grupo" json:"agenda_grupo,omitempty"`
	Colaborador MkColaborador   `gorm:"foreignKey:codcolaborador;references:codcolaborador" json:"colaborador,omitempty"`
	// Relationsships End
	CodagendaGrupoUsuario int32  `gorm:"column:codagenda_grupo_usuario;not null" json:"codagenda_grupo_usuario,omitempty"`
	Codcolaborador        int32  `gorm:"column:codcolaborador;not null" json:"codcolaborador,omitempty"`
	CodagendaGrupo        int32  `gorm:"column:codagenda_grupo;not null" json:"codagenda_grupo,omitempty"`
	Permissoes            string `gorm:"column:permissoes;not null" json:"permissoes,omitempty"`
}

func (*MkAgendaGrupoColaborador) TableName() string {
	return TableNameMkAgendaGrupoColaborador
}
