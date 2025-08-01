package model

import (
	"time"
)

const TableNameMkCompromisso = "mk_compromissos"

type MkCompromisso struct {
	// Relationsships Init
	CompromissoPessoa []MkCompromissoPessoa `gorm:"foreignKey:codcompromisso;references:codcompromisso" json:"compromisso_pessoa,omitempty"`
	Core              MkCore                `gorm:"foreignKey:codcor;references:com_cor_de_fundo" json:"core,omitempty"`
	OS                *MkOS                 `gorm:"foreignKey:codos;references:cd_integracao" json:"os,omitempty"`
	// Relationsships End
	Codcompromisso         int32      `gorm:"column:codcompromisso;not null" json:"codcompromisso,omitempty"`
	ComTitulo              string     `gorm:"column:com_titulo;not null" json:"com_titulo,omitempty"`
	CdIntegracao           int32      `gorm:"column:cd_integracao" json:"cd_integracao,omitempty"`
	ComInicio              *time.Time `gorm:"column:com_inicio;not null" json:"com_inicio,omitempty"`
	ComFim                 *time.Time `gorm:"column:com_fim" json:"com_fim,omitempty"`
	ComEventoDiario        bool       `gorm:"column:com_evento_diario;not null" json:"com_evento_diario,omitempty"`
	ComEditavel            bool       `gorm:"column:com_editavel;not null" json:"com_editavel,omitempty"`
	ComDescricao           string     `gorm:"column:com_descricao" json:"com_descricao,omitempty"`
	ComRealizado           string     `gorm:"column:com_realizado;default:N" json:"com_realizado,omitempty"`
	ComTxRealizado         string     `gorm:"column:com_tx_realizado" json:"com_tx_realizado,omitempty"`
	ComDtRealizacao        *time.Time `gorm:"column:com_dt_realizacao" json:"com_dt_realizacao,omitempty"`
	ComOperadorAbertura    string     `gorm:"column:com_operador_abertura" json:"com_operador_abertura,omitempty"`
	ComOperadorRealizacao  string     `gorm:"column:com_operador_realizacao" json:"com_operador_realizacao,omitempty"`
	Cliente                int32      `gorm:"column:cliente" json:"cliente,omitempty"`
	ComCorDeFundo          int32      `gorm:"column:com_cor_de_fundo" json:"com_cor_de_fundo,omitempty"`
	EnviarEmailCliente     string     `gorm:"column:enviar_email_cliente" json:"enviar_email_cliente,omitempty"`
	EnviarEmailOperador    string     `gorm:"column:enviar_email_operador" json:"enviar_email_operador,omitempty"`
	ComCodigo              string     `gorm:"column:com_codigo" json:"com_codigo,omitempty"`
	CdFuncionario          int32      `gorm:"column:cd_funcionario" json:"cd_funcionario,omitempty"`
	FuncionarioLembrar     string     `gorm:"column:funcionario_lembrar" json:"funcionario_lembrar,omitempty"`
	FuncionarioLembrarDt   *time.Time `gorm:"column:funcionario_lembrar_dt" json:"funcionario_lembrar_dt,omitempty"`
	FuncionarioLembrarHt   string     `gorm:"column:funcionario_lembrar_ht" json:"funcionario_lembrar_ht,omitempty"`
	FuncionarioLembrarTipo string     `gorm:"column:funcionario_lembrar_tipo" json:"funcionario_lembrar_tipo,omitempty"`
	BloquearPorOperador    string     `gorm:"column:bloquear_por_operador" json:"bloquear_por_operador,omitempty"`
	ClienteLembrar         string     `gorm:"column:cliente_lembrar" json:"cliente_lembrar,omitempty"`
	ClienteLembrarDt       *time.Time `gorm:"column:cliente_lembrar_dt" json:"cliente_lembrar_dt,omitempty"`
	ClienteLembrarHr       string     `gorm:"column:cliente_lembrar_hr" json:"cliente_lembrar_hr,omitempty"`
	ClienteLembrarTipo     string     `gorm:"column:cliente_lembrar_tipo" json:"cliente_lembrar_tipo,omitempty"`
	SistemaIntegrado       string     `gorm:"column:sistema_integrado" json:"sistema_integrado,omitempty"`
	ComCorDaBorda          int32      `gorm:"column:com_cor_da_borda" json:"com_cor_da_borda,omitempty"`
	ComCorDoTexto          int32      `gorm:"column:com_cor_do_texto" json:"com_cor_do_texto,omitempty"`
	Lembrado               string     `gorm:"column:lembrado" json:"lembrado,omitempty"`
	LembradoCliente        string     `gorm:"column:lembrado_cliente" json:"lembrado_cliente,omitempty"`
	CdOperador             int32      `gorm:"column:cd_operador" json:"cd_operador,omitempty"`
	OperadorLembrar        string     `gorm:"column:operador_lembrar" json:"operador_lembrar,omitempty"`
	OperadorEnviarTodos    string     `gorm:"column:operador_enviar_todos" json:"operador_enviar_todos,omitempty"`
	TmpFlagExibicao        string     `gorm:"column:tmp_flag_exibicao" json:"tmp_flag_exibicao,omitempty"`
	DtHrAbertura           *time.Time `gorm:"column:dt_hr_abertura" json:"dt_hr_abertura,omitempty"`
	CdEmpresa              int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	TipoCompromisso        int32      `gorm:"column:tipo_compromisso" json:"tipo_compromisso,omitempty"`
	ComInicioReal          *time.Time `gorm:"column:com_inicio_real" json:"com_inicio_real,omitempty"`
	ComFimReal             *time.Time `gorm:"column:com_fim_real" json:"com_fim_real,omitempty"`
	Codmicroarea           int32      `gorm:"column:codmicroarea" json:"codmicroarea,omitempty"`
	Prioridade             int32      `gorm:"column:prioridade;not null;default:50" json:"prioridade,omitempty"`
	EquipeMinima           int32      `gorm:"column:equipe_minima;not null;default:1" json:"equipe_minima,omitempty"`
	CdagendaGrupo          int32      `gorm:"column:cdagenda_grupo" json:"cdagenda_grupo,omitempty"`
	BloqueioAgenda         string     `gorm:"column:bloqueio_agenda;default:N" json:"bloqueio_agenda,omitempty"`
}

func (*MkCompromisso) TableName() string {
	return TableNameMkCompromisso
}
