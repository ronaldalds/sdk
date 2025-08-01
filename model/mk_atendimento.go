package model

import (
	"time"
)

const TableNameMkAtendimento = "mk_atendimento"

type MkAtendimento struct {
	// Relationsships Init
	AtendimentoClassificacao MkAtendimentoClassificacao `gorm:"foreignKey:codatclass;references:classificacao_atendimento" json:"atendimento_classificacao,omitempty"`
	// Relationsships End
	Codatendimento             int32      `gorm:"column:codatendimento;not null" json:"codatendimento,omitempty"`
	ClassificacaoAtendimento   int32      `gorm:"column:classificacao_atendimento" json:"classificacao_atendimento,omitempty"`
	ClienteCadastrado          int32      `gorm:"column:cliente_cadastrado" json:"cliente_cadastrado,omitempty"`
	Conexao                    int32      `gorm:"column:conexao" json:"conexao,omitempty"`
	Protocolo                  string     `gorm:"column:protocolo" json:"protocolo,omitempty"`
	DtAbertura                 *time.Time `gorm:"column:dt_abertura" json:"dt_abertura,omitempty"`
	HrAbertura                 string     `gorm:"column:hr_abertura" json:"hr_abertura,omitempty"`
	JaCliente                  string     `gorm:"column:ja_cliente" json:"ja_cliente,omitempty"`
	Login                      string     `gorm:"column:login" json:"login,omitempty"`
	CdAssinante                int32      `gorm:"column:cd_assinante" json:"cd_assinante,omitempty"`
	Nome                       string     `gorm:"column:nome" json:"nome,omitempty"`
	Cep                        string     `gorm:"column:cep" json:"cep,omitempty"`
	Logradouro                 int32      `gorm:"column:logradouro" json:"logradouro,omitempty"`
	Bairro                     int32      `gorm:"column:bairro" json:"bairro,omitempty"`
	Cidade                     int32      `gorm:"column:cidade" json:"cidade,omitempty"`
	Estado                     int32      `gorm:"column:estado" json:"estado,omitempty"`
	Email                      string     `gorm:"column:email" json:"email,omitempty"`
	FoneCelular                string     `gorm:"column:fone_celular" json:"fone_celular,omitempty"`
	Fone                       string     `gorm:"column:fone" json:"fone,omitempty"`
	Observacao                 string     `gorm:"column:observacao" json:"observacao,omitempty"`
	Contato                    string     `gorm:"column:contato" json:"contato,omitempty"`
	ConsultaExpressa           string     `gorm:"column:consulta_expressa" json:"consulta_expressa,omitempty"`
	ConsultaExpressaAPI        string     `gorm:"column:consulta_expressa_api" json:"consulta_expressa_api,omitempty"`
	Log                        string     `gorm:"column:log" json:"log,omitempty"`
	OperadorAtendimento        string     `gorm:"column:operador_atendimento" json:"operador_atendimento,omitempty"`
	QuemContatou               string     `gorm:"column:quem_contatou" json:"quem_contatou,omitempty"`
	ComoFoiContato             int32      `gorm:"column:como_foi_contato" json:"como_foi_contato,omitempty"`
	InstrucoesRetorno          string     `gorm:"column:instrucoes_retorno;comment:como devera ser dado o retorno" json:"instrucoes_retorno,omitempty"`           // como devera ser dado o retorno
	InfoCliente                string     `gorm:"column:info_cliente;comment:defeito reclamado ou informacao passada pelo cliente" json:"info_cliente,omitempty"` // defeito reclamado ou informacao passada pelo cliente
	Indicacao                  string     `gorm:"column:indicacao;comment:indicacoes e informacoes extras." json:"indicacao,omitempty"`                           // indicacoes e informacoes extras.
	TextoEncerramento          string     `gorm:"column:texto_encerramento" json:"texto_encerramento,omitempty"`
	ParecerTecnico             string     `gorm:"column:parecer_tecnico" json:"parecer_tecnico,omitempty"`
	Encaminhamento             int32      `gorm:"column:encaminhamento" json:"encaminhamento,omitempty"`
	Finalizado                 string     `gorm:"column:finalizado" json:"finalizado,omitempty"`
	DtFinaliza                 *time.Time `gorm:"column:dt_finaliza" json:"dt_finaliza,omitempty"`
	HrFinaliza                 string     `gorm:"column:hr_finaliza" json:"hr_finaliza,omitempty"`
	OperadorEncerramento       string     `gorm:"column:operador_encerramento" json:"operador_encerramento,omitempty"`
	ClassificacaoApoio         int32      `gorm:"column:classificacao_apoio" json:"classificacao_apoio,omitempty"`
	OperadorAbertura           string     `gorm:"column:operador_abertura" json:"operador_abertura,omitempty"`
	FoneContato                string     `gorm:"column:fone_contato" json:"fone_contato,omitempty"`
	NomeContato                string     `gorm:"column:nome_contato" json:"nome_contato,omitempty"`
	EmailContato               string     `gorm:"column:email_contato" json:"email_contato,omitempty"`
	ProtocoloReferencia        string     `gorm:"column:protocolo_referencia" json:"protocolo_referencia,omitempty"`
	ObsAuxContato              string     `gorm:"column:obs_aux_contato" json:"obs_aux_contato,omitempty"`
	CdProcesso                 int32      `gorm:"column:cd_processo" json:"cd_processo,omitempty"`
	Talk                       string     `gorm:"column:talk" json:"talk,omitempty"`
	RetTalk                    string     `gorm:"column:ret_talk" json:"ret_talk,omitempty"`
	ControleHora               float64    `gorm:"column:controle_hora" json:"controle_hora,omitempty"`
	Situacao                   string     `gorm:"column:situacao" json:"situacao,omitempty"`
	CdTicket                   int32      `gorm:"column:cd_ticket" json:"cd_ticket,omitempty"`
	CdSubprocesso              int32      `gorm:"column:cd_subprocesso" json:"cd_subprocesso,omitempty"`
	CdControleSubprocesso      int32      `gorm:"column:cd_controle_subprocesso" json:"cd_controle_subprocesso,omitempty"`
	OperadorResgate            string     `gorm:"column:operador_resgate" json:"operador_resgate,omitempty"`
	ControleHoraResgate        float64    `gorm:"column:controle_hora_resgate" json:"controle_hora_resgate,omitempty"`
	DataResgate                *time.Time `gorm:"column:data_resgate" json:"data_resgate,omitempty"`
	HoraResgate                string     `gorm:"column:hora_resgate" json:"hora_resgate,omitempty"`
	ControleHoraLibResgate     float64    `gorm:"column:controle_hora_lib_resgate" json:"controle_hora_lib_resgate,omitempty"`
	VersaoAtendimento          string     `gorm:"column:versao_atendimento" json:"versao_atendimento,omitempty"`
	DtHrLimResgate             *time.Time `gorm:"column:dt_hr_lim_resgate" json:"dt_hr_lim_resgate,omitempty"`
	DtHrResgate                *time.Time `gorm:"column:dt_hr_resgate" json:"dt_hr_resgate,omitempty"`
	PermCancelResgate          string     `gorm:"column:perm_cancel_resgate" json:"perm_cancel_resgate,omitempty"`
	EmDeb                      string     `gorm:"column:em_deb" json:"em_deb,omitempty"`
	EmBloqueio                 string     `gorm:"column:em_bloqueio" json:"em_bloqueio,omitempty"`
	EmOs                       string     `gorm:"column:em_os" json:"em_os,omitempty"`
	EmNotificacao              string     `gorm:"column:em_notificacao" json:"em_notificacao,omitempty"`
	EmAviso                    string     `gorm:"column:em_aviso" json:"em_aviso,omitempty"`
	IDChamada                  string     `gorm:"column:id_chamada" json:"id_chamada,omitempty"`
	NaoCliente                 string     `gorm:"column:nao_cliente" json:"nao_cliente,omitempty"`
	Prioridade                 string     `gorm:"column:prioridade" json:"prioridade,omitempty"`
	GUIDTmp                    string     `gorm:"column:guid_tmp" json:"guid_tmp,omitempty"`
	SetorAtual                 int32      `gorm:"column:setor_atual" json:"setor_atual,omitempty"`
	GrupoSetorSupervisor       int32      `gorm:"column:grupo_setor_supervisor" json:"grupo_setor_supervisor,omitempty"`
	GrupoParticipante          int32      `gorm:"column:grupo_participante" json:"grupo_participante,omitempty"`
	EncaminhadoParaOs          string     `gorm:"column:encaminhado_para_os" json:"encaminhado_para_os,omitempty"`
	UserAdmSetor               int32      `gorm:"column:user_adm_setor" json:"user_adm_setor,omitempty"`
	CdEmpresa                  int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	DtHrLimiteFimProcesso      *time.Time `gorm:"column:dt_hr_limite_fim_processo" json:"dt_hr_limite_fim_processo,omitempty"`
	ControleHoraFim            float64    `gorm:"column:controle_hora_fim" json:"controle_hora_fim,omitempty"`
	MinutosDiferenca           int32      `gorm:"column:minutos_diferenca" json:"minutos_diferenca,omitempty"`
	TempoVisual                string     `gorm:"column:tempo_visual" json:"tempo_visual,omitempty"`
	Atrasado                   string     `gorm:"column:atrasado" json:"atrasado,omitempty"`
	TempoTotalAtendimento      int32      `gorm:"column:tempo_total_atendimento" json:"tempo_total_atendimento,omitempty"`
	TempoTotalVisual           string     `gorm:"column:tempo_total_visual" json:"tempo_total_visual,omitempty"`
	DtHrStart                  *time.Time `gorm:"column:dt_hr_start" json:"dt_hr_start,omitempty"`
	ControleHoraStart          float64    `gorm:"column:controle_hora_start" json:"controle_hora_start,omitempty"`
	ClassificacaoEncerramento  int32      `gorm:"column:classificacao_encerramento" json:"classificacao_encerramento,omitempty"`
	NotaCliente                int32      `gorm:"column:nota_cliente" json:"nota_cliente,omitempty"`
	LocalAvaliacao             string     `gorm:"column:local_avaliacao" json:"local_avaliacao,omitempty"`
	ComentarioAvaliacao        string     `gorm:"column:comentario_avaliacao" json:"comentario_avaliacao,omitempty"`
	AvaliacaoEnviada           string     `gorm:"column:avaliacao_enviada" json:"avaliacao_enviada,omitempty"`
	AvaliacaoEnviadaEm         *time.Time `gorm:"column:avaliacao_enviada_em" json:"avaliacao_enviada_em,omitempty"`
	RetornoAvaliacaoEm         *time.Time `gorm:"column:retorno_avaliacao_em" json:"retorno_avaliacao_em,omitempty"`
	TmpSLA                     int32      `gorm:"column:tmp_sla" json:"tmp_sla,omitempty"`
	PerfilAtendido             int32      `gorm:"column:perfil_atendido" json:"perfil_atendido,omitempty"`
	DtHrInsert                 *time.Time `gorm:"column:dt_hr_insert" json:"dt_hr_insert,omitempty"`
	CdContrato                 int32      `gorm:"column:cd_contrato" json:"cd_contrato,omitempty"`
	NivelSLA                   int32      `gorm:"column:nivel_sla" json:"nivel_sla,omitempty"`
	CdTecResponsavelContato    int32      `gorm:"column:cd_tec_responsavel_contato" json:"cd_tec_responsavel_contato,omitempty"`
	IDHashInsert               string     `gorm:"column:id_hash_insert" json:"id_hash_insert,omitempty"`
	CdOperadorResgate          int32      `gorm:"column:cd_operador_resgate" json:"cd_operador_resgate,omitempty"`
	CdControleProtocolo        int32      `gorm:"column:cd_controle_protocolo" json:"cd_controle_protocolo,omitempty"`
	CdWorkspace                int32      `gorm:"column:cd_workspace" json:"cd_workspace,omitempty"`
	Complemento                string     `gorm:"column:complemento" json:"complemento,omitempty"`
	Numero                     string     `gorm:"column:numero;comment:" json:"numero,omitempty"`
	TempoTotalAtraso           int32      `gorm:"column:tempo_total_atraso" json:"tempo_total_atraso,omitempty"`
	DhFim                      *time.Time `gorm:"column:dh_fim" json:"dh_fim,omitempty"`
	NivelNotificacaoEletronica int32      `gorm:"column:nivel_notificacao_eletronica;comment:1- sem notificacoes 2- abertura e fechamento 3- todas (processos, subprocessos, OS, etc.)" json:"nivel_notificacao_eletronica,omitempty"` // 1- sem notificacoes 2- abertura e fechamento 3- todas (processos, subprocessos, OS, etc.)
	CdLead                     int32      `gorm:"column:cd_lead" json:"cd_lead,omitempty"`
	CdDialogo                  int32      `gorm:"column:cd_dialogo" json:"cd_dialogo,omitempty"`
	DtHrFimPrevProcesso        *time.Time `gorm:"column:dt_hr_fim_prev_processo" json:"dt_hr_fim_prev_processo,omitempty"`
	IDCircuito                 string     `gorm:"column:id_circuito" json:"id_circuito,omitempty"`
	InfoClienteEspera          string     `gorm:"column:info_cliente_espera" json:"info_cliente_espera,omitempty"`
	HrChegada                  string     `gorm:"column:hr_chegada" json:"hr_chegada,omitempty"`
	ObsMigracao                string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
}

func (*MkAtendimento) TableName() string {
	return TableNameMkAtendimento
}
