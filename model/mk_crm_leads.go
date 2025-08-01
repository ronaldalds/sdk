package model

import (
	"time"
)

const TableNameMkCrmLead = "mk_crm_leads"

type MkCrmLead struct {
	Codlead                   int32      `gorm:"column:codlead;not null" json:"codlead,omitempty"`
	DataCriacao               *time.Time `gorm:"column:data_criacao" json:"data_criacao,omitempty"`
	HoraCriacao               string     `gorm:"column:hora_criacao" json:"hora_criacao,omitempty"`
	OperadorCriacao           string     `gorm:"column:operador_criacao" json:"operador_criacao,omitempty"`
	LocalCriacao              string     `gorm:"column:local_criacao" json:"local_criacao,omitempty"`
	CdAtendimento             int32      `gorm:"column:cd_atendimento" json:"cd_atendimento,omitempty"`
	OperadorResponsavel       int32      `gorm:"column:operador_responsavel" json:"operador_responsavel,omitempty"`
	CdPessoa                  int32      `gorm:"column:cd_pessoa" json:"cd_pessoa,omitempty"`
	Documento                 string     `gorm:"column:documento" json:"documento,omitempty"`
	TipoCliente               int32      `gorm:"column:tipo_cliente;default:1;comment:1- residencial 2- corporativo" json:"tipo_cliente,omitempty"` // 1- residencial 2- corporativo
	NomeCliente               string     `gorm:"column:nome_cliente" json:"nome_cliente,omitempty"`
	DescritivoLocal           string     `gorm:"column:descritivo_local" json:"descritivo_local,omitempty"`
	ClassificacaoLocal        string     `gorm:"column:classificacao_local" json:"classificacao_local,omitempty"`
	InteresseEmInternetW      string     `gorm:"column:interesse_em_internet_w" json:"interesse_em_internet_w,omitempty"`
	InteresseEmTv             string     `gorm:"column:interesse_em_tv" json:"interesse_em_tv,omitempty"`
	InteresseEmTelefone       string     `gorm:"column:interesse_em_telefone" json:"interesse_em_telefone,omitempty"`
	InteresseEmOutros         string     `gorm:"column:interesse_em_outros" json:"interesse_em_outros,omitempty"`
	Informacoes               string     `gorm:"column:informacoes" json:"informacoes,omitempty"`
	StatusLead                int32      `gorm:"column:status_lead" json:"status_lead,omitempty"`
	Lat                       string     `gorm:"column:lat" json:"lat,omitempty"`
	Lon                       string     `gorm:"column:lon" json:"lon,omitempty"`
	InteresseEmInternetF      string     `gorm:"column:interesse_em_internet_f" json:"interesse_em_internet_f,omitempty"`
	EmInviabilidade           string     `gorm:"column:em_inviabilidade" json:"em_inviabilidade,omitempty"`
	DataHoraInviabilidade     *time.Time `gorm:"column:data_hora_inviabilidade" json:"data_hora_inviabilidade,omitempty"`
	Finalizado                string     `gorm:"column:finalizado" json:"finalizado,omitempty"`
	DataHoraFinalizacao       *time.Time `gorm:"column:data_hora_finalizacao" json:"data_hora_finalizacao,omitempty"`
	CdProdutoEscolhido        int32      `gorm:"column:cd_produto_escolhido" json:"cd_produto_escolhido,omitempty"`
	CdPedidoTransCartao       int32      `gorm:"column:cd_pedido_trans_cartao" json:"cd_pedido_trans_cartao,omitempty"`
	QualificacaoLead          int32      `gorm:"column:qualificacao_lead;comment:nivel de importancia da lead no negocio 1 - menor qualificacao 3 - maior qualificacao" json:"qualificacao_lead,omitempty"` // nivel de importancia da lead no negocio 1 - menor qualificacao 3 - maior qualificacao
	AgendaPendente            string     `gorm:"column:agenda_pendente" json:"agenda_pendente,omitempty"`
	DtHrEntradaEtapa          *time.Time `gorm:"column:dt_hr_entrada_etapa" json:"dt_hr_entrada_etapa,omitempty"`
	DtHrSaidaEtapa            *time.Time `gorm:"column:dt_hr_saida_etapa" json:"dt_hr_saida_etapa,omitempty"`
	DtHrPrevSaida             *time.Time `gorm:"column:dt_hr_prev_saida" json:"dt_hr_prev_saida,omitempty"`
	CdInviabilidadeTipo       int32      `gorm:"column:cd_inviabilidade_tipo" json:"cd_inviabilidade_tipo,omitempty"`
	CdTabelaSLA               int32      `gorm:"column:cd_tabela_sla" json:"cd_tabela_sla,omitempty"`
	VendaCrm                  int32      `gorm:"column:venda_crm;default:1;comment:1- Ativa 2- Reativo" json:"venda_crm,omitempty"`                                                                                  // 1- Ativa 2- Reativo
	MeioEntradaCrm            int32      `gorm:"column:meio_entrada_crm;default:9;comment:1-Outdoor 2-Facebook 3-Instagram 4-Site 5-TV 6-Radio 7-Jornal 8-Outros 9-Nao informado" json:"meio_entrada_crm,omitempty"` // 1-Outdoor 2-Facebook 3-Instagram 4-Site 5-TV 6-Radio 7-Jornal 8-Outros 9-Nao informado
	Protocolo                 string     `gorm:"column:protocolo" json:"protocolo,omitempty"`
	FormaAdesaoEscolhido      int32      `gorm:"column:forma_adesao_escolhido" json:"forma_adesao_escolhido,omitempty"`
	FormaMensalidadeEscolhido int32      `gorm:"column:forma_mensalidade_escolhido" json:"forma_mensalidade_escolhido,omitempty"`
	CodVencimento             int32      `gorm:"column:cod_vencimento" json:"cod_vencimento,omitempty"`
	CdLeadOrigem              int32      `gorm:"column:cd_lead_origem;comment:lead inviabilizada que deu origem a nova lead" json:"cd_lead_origem,omitempty"`                             // lead inviabilizada que deu origem a nova lead
	CodMotivoFechamento       int32      `gorm:"column:cod_motivo_fechamento;comment:Motivo de ganho - Por qual motivo o cliente fehou o negocio" json:"cod_motivo_fechamento,omitempty"` // Motivo de ganho - Por qual motivo o cliente fehou o negocio
	TipoContato               int32      `gorm:"column:tipo_contato;comment:0 - nao informado  1 - ativo 2 - receptivo" json:"tipo_contato,omitempty"`                                    // 0 - nao informado  1 - ativo 2 - receptivo
	CdEmpresa                 int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	ComissionadoInd           string     `gorm:"column:comissionado_ind;default:N;comment:S- Sim N- Nao" json:"comissionado_ind,omitempty"` // S- Sim N- Nao
	DataHoraComissaoInd       *time.Time `gorm:"column:data_hora_comissao_ind" json:"data_hora_comissao_ind,omitempty"`
	ComissionadoGrupo         string     `gorm:"column:comissionado_grupo;default:N;comment:S- Sim N- Nao" json:"comissionado_grupo,omitempty"` // S- Sim N- Nao
	DataHoraComissaoGrupo     *time.Time `gorm:"column:data_hora_comissao_grupo" json:"data_hora_comissao_grupo,omitempty"`
	AdicionarProdutosAvulso   string     `gorm:"column:adicionar_produtos_avulso" json:"adicionar_produtos_avulso,omitempty"`
	ItensCarregados           string     `gorm:"column:itens_carregados;default:N" json:"itens_carregados,omitempty"`
	CodPerfilativacao         int32      `gorm:"column:cod_perfilativacao" json:"cod_perfilativacao,omitempty"`
	DtLimite                  *time.Time `gorm:"column:dt_limite" json:"dt_limite,omitempty"`
	ImgDocViabilidade         []uint8    `gorm:"column:img_doc_viabilidade" json:"img_doc_viabilidade,omitempty"`
	JSONRouteViabilidade      string     `gorm:"column:json_route_viabilidade" json:"json_route_viabilidade,omitempty"`
	CdPlanoAtual              int32      `gorm:"column:cd_plano_atual;comment:plano antes do upgrade/downgrade" json:"cd_plano_atual,omitempty"` // plano antes do upgrade/downgrade
	CdPlanoNovo               int32      `gorm:"column:cd_plano_novo;comment:novo plano apos upgrade/downgrade" json:"cd_plano_novo,omitempty"`  // novo plano apos upgrade/downgrade
	VlrMensalidadeAtual       float64    `gorm:"column:vlr_mensalidade_atual" json:"vlr_mensalidade_atual,omitempty"`
	VlrMensalidadeNovo        float64    `gorm:"column:vlr_mensalidade_novo" json:"vlr_mensalidade_novo,omitempty"`
	CdContratoUpgrade         int32      `gorm:"column:cd_contrato_upgrade;comment:contrato que vai sofrer upgrade/downgrade" json:"cd_contrato_upgrade,omitempty"` // contrato que vai sofrer upgrade/downgrade
	ControleVisitas           string     `gorm:"column:controle_visitas" json:"controle_visitas,omitempty"`
	DataUltimaVisita          *time.Time `gorm:"column:data_ultima_visita" json:"data_ultima_visita,omitempty"`
	CdCampanhaPromo           int32      `gorm:"column:cd_campanha_promo" json:"cd_campanha_promo,omitempty"`
	TpCombo                   string     `gorm:"column:tp_combo" json:"tp_combo,omitempty"`
	DataRevisita              *time.Time `gorm:"column:data_revisita" json:"data_revisita,omitempty"`
	LeadQualificada           bool       `gorm:"column:lead_qualificada" json:"lead_qualificada,omitempty"`
	NomeMae                   string     `gorm:"column:nome_mae" json:"nome_mae,omitempty"`
	NomePai                   string     `gorm:"column:nome_pai" json:"nome_pai,omitempty"`
}

func (*MkCrmLead) TableName() string {
	return TableNameMkCrmLead
}
