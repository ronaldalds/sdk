package model

import (
	"time"
)

const TableNameMkFatura = "mk_faturas"

type MkFatura struct {
	Codfatura                  int32      `gorm:"column:codfatura;not null" json:"codfatura,omitempty"`
	DataLancamento             *time.Time `gorm:"column:data_lancamento" json:"data_lancamento,omitempty"`
	UsuarioLancamento          string     `gorm:"column:usuario_lancamento" json:"usuario_lancamento,omitempty"`
	ValorTotal                 float64    `gorm:"column:valor_total" json:"valor_total,omitempty"`
	DataVencimento             *time.Time `gorm:"column:data_vencimento" json:"data_vencimento,omitempty"`
	Excluida                   string     `gorm:"column:excluida;default:N;comment:S/N" json:"excluida,omitempty"` // S/N
	UsuarioExclusao            string     `gorm:"column:usuario_exclusao" json:"usuario_exclusao,omitempty"`
	DataExclusao               *time.Time `gorm:"column:data_exclusao" json:"data_exclusao,omitempty"`
	CdPessoa                   int32      `gorm:"column:cd_pessoa" json:"cd_pessoa,omitempty"`
	TipoLancamento             string     `gorm:"column:tipo_lancamento" json:"tipo_lancamento,omitempty"`
	Liquidado                  string     `gorm:"column:liquidado;default:N" json:"liquidado,omitempty"`
	DataLiquidacao             *time.Time `gorm:"column:data_liquidacao" json:"data_liquidacao,omitempty"`
	HoraLiquidacao             string     `gorm:"column:hora_liquidacao" json:"hora_liquidacao,omitempty"`
	VlrLiquidacao              float64    `gorm:"column:vlr_liquidacao" json:"vlr_liquidacao,omitempty"`
	UsuarioLiquidacao          string     `gorm:"column:usuario_liquidacao" json:"usuario_liquidacao,omitempty"`
	HoraExclusao               string     `gorm:"column:hora_exclusao" json:"hora_exclusao,omitempty"`
	CdFaturamento              int32      `gorm:"column:cd_faturamento" json:"cd_faturamento,omitempty"`
	GUIDSelecao                string     `gorm:"column:guid_selecao" json:"guid_selecao,omitempty"`
	TxMultaAtraso              float64    `gorm:"column:tx_multa_atraso" json:"tx_multa_atraso,omitempty"`
	TxJuroAtraso               float64    `gorm:"column:tx_juro_atraso" json:"tx_juro_atraso,omitempty"`
	DtRefInicial               *time.Time `gorm:"column:dt_ref_inicial" json:"dt_ref_inicial,omitempty"`
	DtRefFinal                 *time.Time `gorm:"column:dt_ref_final" json:"dt_ref_final,omitempty"`
	Descricao                  string     `gorm:"column:descricao" json:"descricao,omitempty"`
	FormaPgtoLiquidacao        int32      `gorm:"column:forma_pgto_liquidacao;comment:1- Dinheiro 2- E-commerce 3- cartao 4- deposito 5- pag-seguro 6- profile bancaria (avulso) 7- liquidacao por conciliacao 8- conciliacao manual pelo fluxo de caixa 98- liquidacao a partir de remocao saldos 99- outro (sem integracao)" json:"forma_pgto_liquidacao,omitempty"` // 1- Dinheiro 2- E-commerce 3- cartao 4- deposito 5- pag-seguro 6- profile bancaria (avulso) 7- liquidacao por conciliacao 8- conciliacao manual pelo fluxo de caixa 98- liquidacao a partir de remocao saldos 99- outro (sem integracao)
	ObsLiquidacao              string     `gorm:"column:obs_liquidacao" json:"obs_liquidacao,omitempty"`
	DocumentoLiquidacao        string     `gorm:"column:documento_liquidacao" json:"documento_liquidacao,omitempty"`
	PontoRecebimentoLiquidacao int32      `gorm:"column:ponto_recebimento_liquidacao" json:"ponto_recebimento_liquidacao,omitempty"`
	DtEfetivacaoCredito        *time.Time `gorm:"column:dt_efetivacao_credito" json:"dt_efetivacao_credito,omitempty"`
	VlrAcrescimoAplicado       float64    `gorm:"column:vlr_acrescimo_aplicado" json:"vlr_acrescimo_aplicado,omitempty"`
	VlrDescontoAplicado        float64    `gorm:"column:vlr_desconto_aplicado" json:"vlr_desconto_aplicado,omitempty"`
	VlrTarifaAplicada          float64    `gorm:"column:vlr_tarifa_aplicada" json:"vlr_tarifa_aplicada,omitempty"`
	DescontoMaximoAplicavel    float64    `gorm:"column:desconto_maximo_aplicavel" json:"desconto_maximo_aplicavel,omitempty"`
	DataDescontoMaximo         *time.Time `gorm:"column:data_desconto_maximo" json:"data_desconto_maximo,omitempty"`
	Suspenso                   string     `gorm:"column:suspenso;default:N" json:"suspenso,omitempty"`
	DataVencimentoOriginal     *time.Time `gorm:"column:data_vencimento_original" json:"data_vencimento_original,omitempty"`
	ValorTotalOriginal         float64    `gorm:"column:valor_total_original" json:"valor_total_original,omitempty"`
	NumReprogramacao           int32      `gorm:"column:num_reprogramacao" json:"num_reprogramacao,omitempty"`
	InfoReprogramacao          string     `gorm:"column:info_reprogramacao" json:"info_reprogramacao,omitempty"`
	CdMotivoSuspensao          int32      `gorm:"column:cd_motivo_suspensao" json:"cd_motivo_suspensao,omitempty"`
	Exportacao                 string     `gorm:"column:exportacao" json:"exportacao,omitempty"`
	ExportacaoDataHora         *time.Time `gorm:"column:exportacao_data_hora" json:"exportacao_data_hora,omitempty"`
	ExportacaoArquivo          string     `gorm:"column:exportacao_arquivo" json:"exportacao_arquivo,omitempty"`
	Agrupamento                string     `gorm:"column:agrupamento" json:"agrupamento,omitempty"`
	UsuarioSuspensao           string     `gorm:"column:usuario_suspensao" json:"usuario_suspensao,omitempty"`
	DataSuspensao              *time.Time `gorm:"column:data_suspensao" json:"data_suspensao,omitempty"`
	ObsSuspensao               string     `gorm:"column:obs_suspensao" json:"obs_suspensao,omitempty"`
	ObsFatura                  string     `gorm:"column:obs_fatura" json:"obs_fatura,omitempty"`
	LiquidadoPorRetorno        string     `gorm:"column:liquidado_por_retorno" json:"liquidado_por_retorno,omitempty"`
	NumArqRetorno              int32      `gorm:"column:num_arq_retorno" json:"num_arq_retorno,omitempty"`
	CdEmpresa                  int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	Estornado                  string     `gorm:"column:estornado;default:N" json:"estornado,omitempty"`
	DataEstorno                *time.Time `gorm:"column:data_estorno" json:"data_estorno,omitempty"`
	HoraEstorno                string     `gorm:"column:hora_estorno" json:"hora_estorno,omitempty"`
	VlrEstorno                 float64    `gorm:"column:vlr_estorno" json:"vlr_estorno,omitempty"`
	UsuarioEstorno             string     `gorm:"column:usuario_estorno" json:"usuario_estorno,omitempty"`
	Tipo                       string     `gorm:"column:tipo;default:R;comment:P - Pagar; R - Receber" json:"tipo,omitempty"` // P - Pagar; R - Receber
	DocumentoCobranca          string     `gorm:"column:documento_cobranca" json:"documento_cobranca,omitempty"`
	DtHrEventoLiquidacao       *time.Time `gorm:"column:dt_hr_evento_liquidacao" json:"dt_hr_evento_liquidacao,omitempty"`
	DocFiscaisGerados          string     `gorm:"column:doc_fiscais_gerados;default:N" json:"doc_fiscais_gerados,omitempty"`
	NomenclaturaIntegracao     string     `gorm:"column:nomenclatura_integracao;comment:Identifica origem de criacao da fatura MOV - mk_estoque_master_parcelas (entradas e saidas de estoque)" json:"nomenclatura_integracao,omitempty"` // Identifica origem de criacao da fatura MOV - mk_estoque_master_parcelas (entradas e saidas de estoque)
	CdVinculado                int32      `gorm:"column:cd_vinculado;comment:codigo complementar da nomenclatura de integracao" json:"cd_vinculado,omitempty"`                                                                            // codigo complementar da nomenclatura de integracao
	BarrasCobranca             string     `gorm:"column:barras_cobranca" json:"barras_cobranca,omitempty"`
	LdCobranca                 string     `gorm:"column:ld_cobranca" json:"ld_cobranca,omitempty"`
	DhGeracao                  *time.Time `gorm:"column:dh_geracao" json:"dh_geracao,omitempty"`
	IDOperador                 int32      `gorm:"column:id_operador" json:"id_operador,omitempty"`
	TipoCobranca               int32      `gorm:"column:tipo_cobranca;comment:1- boleto 2- arrecadacao 3- boleto proprio 4- debito 5-cartao" json:"tipo_cobranca,omitempty"` // 1- boleto 2- arrecadacao 3- boleto proprio 4- debito 5-cartao
	DhRegistro                 *time.Time `gorm:"column:dh_registro;default:now()" json:"dh_registro,omitempty"`
	NumeroCobranca             string     `gorm:"column:numero_cobranca" json:"numero_cobranca,omitempty"`
	CdProfileCobranca          int32      `gorm:"column:cd_profile_cobranca" json:"cd_profile_cobranca,omitempty"`
	CdDocCobranca              int32      `gorm:"column:cd_doc_cobranca" json:"cd_doc_cobranca,omitempty"`
	OpcaoLcto                  int32      `gorm:"column:opcao_lcto;comment:1- sem rateio 2- com rateio" json:"opcao_lcto,omitempty"` // 1- sem rateio 2- com rateio
	PlanoContas                string     `gorm:"column:plano_contas" json:"plano_contas,omitempty"`
	CcDepartamentos            int32      `gorm:"column:cc_departamentos" json:"cc_departamentos,omitempty"`
	CcServicos                 int32      `gorm:"column:cc_servicos" json:"cc_servicos,omitempty"`
	CcProjetos                 int32      `gorm:"column:cc_projetos" json:"cc_projetos,omitempty"`
	CcAreaNegocios             int32      `gorm:"column:cc_area_negocios" json:"cc_area_negocios,omitempty"`
	DhInsert                   *time.Time `gorm:"column:dh_insert;default:now()" json:"dh_insert,omitempty"`
	ContaFixa                  string     `gorm:"column:conta_fixa;default:N" json:"conta_fixa,omitempty"`
	CdAgrupador                int32      `gorm:"column:cd_agrupador" json:"cd_agrupador,omitempty"`
	SaldoRestante              float64    `gorm:"column:saldo_restante" json:"saldo_restante,omitempty"`
	ValorLiquidado             float64    `gorm:"column:valor_liquidado" json:"valor_liquidado,omitempty"`
	VerGeracaoDocFiscal        int32      `gorm:"column:ver_geracao_doc_fiscal" json:"ver_geracao_doc_fiscal,omitempty"`
	TeveLiquidacaoParcial      string     `gorm:"column:teve_liquidacao_parcial;default:N" json:"teve_liquidacao_parcial,omitempty"`
	GUIDImpressao              string     `gorm:"column:guid_impressao" json:"guid_impressao,omitempty"`
	ComoPagar                  int32      `gorm:"column:como_pagar" json:"como_pagar,omitempty"`
	ComoPagarBkp               int32      `gorm:"column:como_pagar_bkp" json:"como_pagar_bkp,omitempty"`
	TipoCobrancaBkp            int32      `gorm:"column:tipo_cobranca_bkp" json:"tipo_cobranca_bkp,omitempty"`
	Migra                      int32      `gorm:"column:migra" json:"migra,omitempty"`
	CdProfile                  int32      `gorm:"column:cd_profile" json:"cd_profile,omitempty"`
	ObsMigracao                string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
	DataInativacao             *time.Time `gorm:"column:data_inativacao" json:"data_inativacao,omitempty"`
	IndAcaoMassa               string     `gorm:"column:ind_acao_massa;default:N" json:"ind_acao_massa,omitempty"`
	DhAcaoMassa                *time.Time `gorm:"column:dh_acao_massa" json:"dh_acao_massa,omitempty"`
}

func (*MkFatura) TableName() string {
	return TableNameMkFatura
}
