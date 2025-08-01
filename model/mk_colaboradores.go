package model

const TableNameMkColaboradore = "mk_colaboradores"

type MkColaborador struct {
	// Relationsships Init
	Usuario                FrUsuario                 `gorm:"foreignKey:usr_codigo;references:usr_codigo" json:"usuario,omitempty"`
	AgendaGrupoColaborador *MkAgendaGrupoColaborador `gorm:"foreignKey:codcolaborador;references:codcolaborador" json:"agenda_grupo_colaborador,omitempty"`
	// Relationsships End
	Codcolaborador        int32  `gorm:"column:codcolaborador;not null" json:"codcolaborador,omitempty"`
	Codempresa            int32  `gorm:"column:codempresa;default:1" json:"codempresa,omitempty"`
	UsrCodigo             int32  `gorm:"column:usr_codigo;not null" json:"usr_codigo,omitempty"`
	Nivel                 int32  `gorm:"column:nivel;not null;default:1" json:"nivel,omitempty"`
	Terceirizado          string `gorm:"column:terceirizado;default:N" json:"terceirizado,omitempty"`
	PrioridadeAgendamento int32  `gorm:"column:prioridade_agendamento;default:1;comment:1 - Agedamento primeiro\r\n\r\n5 - Agendamento por ultimo" json:"prioridade_agendamento,omitempty"` // 1 - Agedamento primeiro\r\n\r\n5 - Agendamento por ultimo
	LiderEquipe           string `gorm:"column:lider_equipe;default:S" json:"lider_equipe,omitempty"`
	Codbaseoperacional    int32  `gorm:"column:codbaseoperacional" json:"codbaseoperacional,omitempty"`
	ChatUsuario           string `gorm:"column:chat_usuario" json:"chat_usuario,omitempty"`
	ChatSenha             string `gorm:"column:chat_senha" json:"chat_senha,omitempty"`
	LimiteDiario          int32  `gorm:"column:limite_diario" json:"limite_diario,omitempty"`
	TurnoTrabalho         int32  `gorm:"column:turno_trabalho" json:"turno_trabalho,omitempty"`
}

func (*MkColaborador) TableName() string {
	return TableNameMkColaboradore
}
