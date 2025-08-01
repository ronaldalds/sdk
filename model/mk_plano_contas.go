package model

import (
	"time"
)

const TableNameMkPlanoConta = "mk_plano_contas"

type MkPlanoConta struct {
	// Relationsships Init
	Contrato MkContrato `gorm:"foreignKey:codcontrato;references:codvinculado" json:"contratos,omitempty"`
	// Relationsships End
	Codconta     int32 `gorm:"column:codconta;not null" json:"codconta,omitempty"`
	Codvinculado int32 `gorm:"column:codvinculado;comment:codigo de vinculo com outras tabelas." json:"codvinculado,omitempty"` // codigo de vinculo com outras tabelas.

	TipoConta                string     `gorm:"column:tipo_conta;not null;comment:P/R" json:"tipo_conta,omitempty"` // P/R
	DescricaoConta           string     `gorm:"column:descricao_conta;not null" json:"descricao_conta,omitempty"`
	ValorLancamento          float64    `gorm:"column:valor_lancamento;not null" json:"valor_lancamento,omitempty"`
	CodclienteCodfornecedor  int32      `gorm:"column:codcliente_codfornecedor;not null" json:"codcliente_codfornecedor,omitempty"`
	NomenclaturaIntegracao   string     `gorm:"column:nomenclatura_integracao;comment:nomenclatura da integracao do sistema, ex. OS ou CNT" json:"nomenclatura_integracao,omitempty"` // nomenclatura da integracao do sistema, ex. OS ou CNT
	ValorTotal               float64    `gorm:"column:valor_total;comment:valor total dos lancamento, o somatorio das parcelas da conta." json:"valor_total,omitempty"`               // valor total dos lancamento, o somatorio das parcelas da conta.
	DataLancamento           *time.Time `gorm:"column:data_lancamento;not null" json:"data_lancamento,omitempty"`
	DataVencimento           *time.Time `gorm:"column:data_vencimento;not null" json:"data_vencimento,omitempty"`
	Acrescimo                float64    `gorm:"column:acrescimo" json:"acrescimo,omitempty"`
	Desconto                 float64    `gorm:"column:desconto" json:"desconto,omitempty"`
	FormaPgto                int32      `gorm:"column:forma_pgto;comment:campo que define a forma de pgto adotada para a conta." json:"forma_pgto,omitempty"` // campo que define a forma de pgto adotada para a conta.
	NumeroParcelas           int32      `gorm:"column:numero_parcelas;comment:numero de parcelas" json:"numero_parcelas,omitempty"`                           // numero de parcelas
	ParcelaAtual             int32      `gorm:"column:parcela_atual;comment:parcela atual do lancamento" json:"parcela_atual,omitempty"`                      // parcela atual do lancamento
	CodigoBarras             string     `gorm:"column:codigo_barras" json:"codigo_barras,omitempty"`
	IDConta                  string     `gorm:"column:id_conta;comment:id de identificacao da conta" json:"id_conta,omitempty"` // id de identificacao da conta
	CodigoBarrasDigitavel    string     `gorm:"column:codigo_barras_digitavel" json:"codigo_barras_digitavel,omitempty"`
	Liquidado                string     `gorm:"column:liquidado;comment:se a conta esta liquidada." json:"liquidado,omitempty"` // se a conta esta liquidada.
	DataLiquidacao           *time.Time `gorm:"column:data_liquidacao" json:"data_liquidacao,omitempty"`
	Observacoes              string     `gorm:"column:observacoes" json:"observacoes,omitempty"`
	Instrucoes               string     `gorm:"column:instrucoes;comment:instrucoes para impressao no boleto." json:"instrucoes,omitempty"` // instrucoes para impressao no boleto.
	Documento                string     `gorm:"column:documento" json:"documento,omitempty"`
	Status                   string     `gorm:"column:status;comment:status da conexao." json:"status,omitempty"`                                 // status da conexao.
	Emitido                  string     `gorm:"column:emitido;comment:se a parcela foi emitida ao banco ou ao cliente." json:"emitido,omitempty"` // se a parcela foi emitida ao banco ou ao cliente.
	ProfileEmissao           int32      `gorm:"column:profile_emissao;comment:profile de emissao vinculada." json:"profile_emissao,omitempty"`    // profile de emissao vinculada.
	Grupo                    int32      `gorm:"column:grupo" json:"grupo,omitempty"`
	Subgrupo                 int32      `gorm:"column:subgrupo" json:"subgrupo,omitempty"`
	LiquidadoSac             string     `gorm:"column:liquidado_sac" json:"liquidado_sac,omitempty"`
	PontoRecebimento         int32      `gorm:"column:ponto_recebimento" json:"ponto_recebimento,omitempty"`
	MotivoEstorno            int32      `gorm:"column:motivo_estorno" json:"motivo_estorno,omitempty"`
	ObsPgto                  string     `gorm:"column:obs_pgto" json:"obs_pgto,omitempty"`
	NfGerada                 string     `gorm:"column:nf_gerada" json:"nf_gerada,omitempty"`
	TmpNf                    string     `gorm:"column:tmp_nf" json:"tmp_nf,omitempty"`
	LiquidadoPorRetorno      string     `gorm:"column:liquidado_por_retorno" json:"liquidado_por_retorno,omitempty"`
	NfPreGerada              string     `gorm:"column:nf_pre_gerada" json:"nf_pre_gerada,omitempty"`
	CodPgtoSac               int32      `gorm:"column:cod_pgto_sac" json:"cod_pgto_sac,omitempty"`
	VlrLiquidacao            float64    `gorm:"column:vlr_liquidacao" json:"vlr_liquidacao,omitempty"`
	InfoExtra                string     `gorm:"column:info_extra" json:"info_extra,omitempty"`
	NumReprogramacao         int32      `gorm:"column:num_reprogramacao" json:"num_reprogramacao,omitempty"`
	InfoReprogramacao        string     `gorm:"column:info_reprogramacao" json:"info_reprogramacao,omitempty"`
	IDContaRetBanco          string     `gorm:"column:id_conta_ret_banco" json:"id_conta_ret_banco,omitempty"`
	TmpGrupo                 int32      `gorm:"column:tmp_grupo" json:"tmp_grupo,omitempty"`
	TmpSubgrupo              int32      `gorm:"column:tmp_subgrupo" json:"tmp_subgrupo,omitempty"`
	DataVencimentoOriginal   *time.Time `gorm:"column:data_vencimento_original" json:"data_vencimento_original,omitempty"`
	ValorOriginal            float64    `gorm:"column:valor_original" json:"valor_original,omitempty"`
	MaxExcecaoSac            *time.Time `gorm:"column:max_excecao_sac" json:"max_excecao_sac,omitempty"`
	DtBaseExcecaoSac         *time.Time `gorm:"column:dt_base_excecao_sac" json:"dt_base_excecao_sac,omitempty"`
	ObsPgtoParcela           string     `gorm:"column:obs_pgto_parcela" json:"obs_pgto_parcela,omitempty"`
	IgnoraNf21               string     `gorm:"column:ignora_nf21" json:"ignora_nf21,omitempty"`
	LoteNf                   int32      `gorm:"column:lote_nf" json:"lote_nf,omitempty"`
	AuxVlrPago               float64    `gorm:"column:aux_vlr_pago" json:"aux_vlr_pago,omitempty"`
	AuxValorTroco            float64    `gorm:"column:aux_valor_troco" json:"aux_valor_troco,omitempty"`
	IgnoraValidaID           string     `gorm:"column:ignora_valida_id" json:"ignora_valida_id,omitempty"`
	IDConta2                 string     `gorm:"column:id_conta2" json:"id_conta2,omitempty"`
	TxMulta                  float64    `gorm:"column:tx_multa" json:"tx_multa,omitempty"`
	TxAcrescimoMes           float64    `gorm:"column:tx_acrescimo_mes" json:"tx_acrescimo_mes,omitempty"`
	TmpVlrDescAbatimento     float64    `gorm:"column:tmp_vlr_desc_abatimento" json:"tmp_vlr_desc_abatimento,omitempty"`
	TmpIgnoraDecAbatimento   string     `gorm:"column:tmp_ignora_dec_abatimento" json:"tmp_ignora_dec_abatimento,omitempty"`
	TmpDiasDescAbatimento    int32      `gorm:"column:tmp_dias_desc_abatimento" json:"tmp_dias_desc_abatimento,omitempty"`
	DtEfetivacaoCredito      *time.Time `gorm:"column:dt_efetivacao_credito" json:"dt_efetivacao_credito,omitempty"`
	TarifaLiquidacao         float64    `gorm:"column:tarifa_liquidacao" json:"tarifa_liquidacao,omitempty"`
	VersaoLiquidacao         string     `gorm:"column:versao_liquidacao" json:"versao_liquidacao,omitempty"`
	NumArquivoRetorno        int32      `gorm:"column:num_arquivo_retorno" json:"num_arquivo_retorno,omitempty"`
	UnidadeFinanceira        string     `gorm:"column:unidade_financeira" json:"unidade_financeira,omitempty"`
	NossoNumero              string     `gorm:"column:nosso_numero" json:"nosso_numero,omitempty"`
	ContaMaster              int32      `gorm:"column:conta_master" json:"conta_master,omitempty"`
	Agendar                  string     `gorm:"column:agendar" json:"agendar,omitempty"`
	SugestaoCor              int32      `gorm:"column:sugestao_cor" json:"sugestao_cor,omitempty"`
	TmpVencido               string     `gorm:"column:tmp_vencido" json:"tmp_vencido,omitempty"`
	TmpVenceHj               string     `gorm:"column:tmp_vence_hj" json:"tmp_vence_hj,omitempty"`
	RashNf                   string     `gorm:"column:rash_nf" json:"rash_nf,omitempty"`
	CdBarrasNotaFatura       string     `gorm:"column:cd_barras_nota_fatura" json:"cd_barras_nota_fatura,omitempty"`
	GUIDTmp                  string     `gorm:"column:guid_tmp" json:"guid_tmp,omitempty"`
	VersaoNf21               string     `gorm:"column:versao_nf21" json:"versao_nf21,omitempty"`
	SituacaoNf               string     `gorm:"column:situacao_nf" json:"situacao_nf,omitempty"`
	GUIDReferenciaNf         string     `gorm:"column:guid_referencia_nf" json:"guid_referencia_nf,omitempty"`
	GUIDTmpGeracaoMassa      string     `gorm:"column:guid_tmp_geracao_massa" json:"guid_tmp_geracao_massa,omitempty"`
	MeioPgto                 int32      `gorm:"column:meio_pgto" json:"meio_pgto,omitempty"`
	OperadorLiquidacaoManual string     `gorm:"column:operador_liquidacao_manual" json:"operador_liquidacao_manual,omitempty"`
	LiquidacaoRetroativa     string     `gorm:"column:liquidacao_retroativa;default:N" json:"liquidacao_retroativa,omitempty"`
	DtLctoRetroativo         *time.Time `gorm:"column:dt_lcto_retroativo" json:"dt_lcto_retroativo,omitempty"`
	Codprofile               int32      `gorm:"column:codprofile" json:"codprofile,omitempty"`
	DescontoProcessado       string     `gorm:"column:desconto_processado" json:"desconto_processado,omitempty"`
	IDUnificacao             int32      `gorm:"column:id_unificacao" json:"id_unificacao,omitempty"`
	GUIDTmpGeracaoMassaRec   string     `gorm:"column:guid_tmp_geracao_massa_rec" json:"guid_tmp_geracao_massa_rec,omitempty"`
	NfGeradaPreLyt02         string     `gorm:"column:nf_gerada_pre_lyt02" json:"nf_gerada_pre_lyt02,omitempty"`
	GUIDTmpGeracaoMassaPre   string     `gorm:"column:guid_tmp_geracao_massa_pre" json:"guid_tmp_geracao_massa_pre,omitempty"`
	NfseGerada               string     `gorm:"column:nfse_gerada" json:"nfse_gerada,omitempty"`
	DtRefInicial             *time.Time `gorm:"column:dt_ref_inicial" json:"dt_ref_inicial,omitempty"`
	DtRefFinal               *time.Time `gorm:"column:dt_ref_final" json:"dt_ref_final,omitempty"`
	XDiasReferencia          int32      `gorm:"column:x_dias_referencia" json:"x_dias_referencia,omitempty"`
	Suspenso                 string     `gorm:"column:suspenso;default:N" json:"suspenso,omitempty"`
	GUIDSelecao              string     `gorm:"column:guid_selecao" json:"guid_selecao,omitempty"`
	ProfileRedefinida        string     `gorm:"column:profile_redefinida" json:"profile_redefinida,omitempty"`
	ProfileOriginal          int32      `gorm:"column:profile_original" json:"profile_original,omitempty"`
	DtVctoAntesExcSac        *time.Time `gorm:"column:dt_vcto_antes_exc_sac" json:"dt_vcto_antes_exc_sac,omitempty"`
	ContaFixa                string     `gorm:"column:conta_fixa;default:N" json:"conta_fixa,omitempty"`
	ContaFixaInicio          *time.Time `gorm:"column:conta_fixa_inicio" json:"conta_fixa_inicio,omitempty"`
	ContaFixaTermino         *time.Time `gorm:"column:conta_fixa_termino" json:"conta_fixa_termino,omitempty"`
	ContaFixaCodcontaRef     int32      `gorm:"column:conta_fixa_codconta_ref" json:"conta_fixa_codconta_ref,omitempty"`
	VersaoContasFaturamento  string     `gorm:"column:versao_contas_faturamento" json:"versao_contas_faturamento,omitempty"`
	ReciboNegDebito          string     `gorm:"column:recibo_neg_debito" json:"recibo_neg_debito,omitempty"`
	CdMotivoSuspensao        int32      `gorm:"column:cd_motivo_suspensao" json:"cd_motivo_suspensao,omitempty"`
	CdUnidOrcamentaria       int32      `gorm:"column:cd_unid_orcamentaria" json:"cd_unid_orcamentaria,omitempty"`
	DescontoInatividadeTmp   string     `gorm:"column:desconto_inatividade_tmp" json:"desconto_inatividade_tmp,omitempty"`
	TokenInsercaoExterna     string     `gorm:"column:token_insercao_externa" json:"token_insercao_externa,omitempty"`
	ReciboIntegrGerado       string     `gorm:"column:recibo_integr_gerado" json:"recibo_integr_gerado,omitempty"`
	ControleHora             int64      `gorm:"column:controle_hora" json:"controle_hora,omitempty"`
	ArquivoImportacao        string     `gorm:"column:arquivo_importacao" json:"arquivo_importacao,omitempty"`
	UsuarioSuspensao         string     `gorm:"column:usuario_suspensao" json:"usuario_suspensao,omitempty"`
	DataSuspensao            *time.Time `gorm:"column:data_suspensao" json:"data_suspensao,omitempty"`
	ObsSuspensao             string     `gorm:"column:obs_suspensao" json:"obs_suspensao,omitempty"`
	CompDtLanMes             int32      `gorm:"column:comp_dt_lan_mes" json:"comp_dt_lan_mes,omitempty"`
	CompDtLanAno             int32      `gorm:"column:comp_dt_lan_ano" json:"comp_dt_lan_ano,omitempty"`
	CompDtLanMesAno          string     `gorm:"column:comp_dt_lan_mes_ano" json:"comp_dt_lan_mes_ano,omitempty"`
	CompDtVencMes            int32      `gorm:"column:comp_dt_venc_mes" json:"comp_dt_venc_mes,omitempty"`
	CompDtVencAno            int32      `gorm:"column:comp_dt_venc_ano" json:"comp_dt_venc_ano,omitempty"`
	CompDtVencMesAno         string     `gorm:"column:comp_dt_venc_mes_ano" json:"comp_dt_venc_mes_ano,omitempty"`
	CompDtVencOrigMes        int32      `gorm:"column:comp_dt_venc_orig_mes" json:"comp_dt_venc_orig_mes,omitempty"`
	CompDtVencOrigAno        int32      `gorm:"column:comp_dt_venc_orig_ano" json:"comp_dt_venc_orig_ano,omitempty"`
	CompDtVencOrigMesAno     string     `gorm:"column:comp_dt_venc_orig_mes_ano" json:"comp_dt_venc_orig_mes_ano,omitempty"`
	CompDtLiqMes             int32      `gorm:"column:comp_dt_liq_mes" json:"comp_dt_liq_mes,omitempty"`
	CompDtLiqAno             int32      `gorm:"column:comp_dt_liq_ano" json:"comp_dt_liq_ano,omitempty"`
	CompDtLiqMesAno          string     `gorm:"column:comp_dt_liq_mes_ano" json:"comp_dt_liq_mes_ano,omitempty"`
	CdEmpresa                int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	HoraLiquidado            string     `gorm:"column:hora_liquidado" json:"hora_liquidado,omitempty"`
	Estornado                string     `gorm:"column:estornado" json:"estornado,omitempty"`
	DataEstorno              *time.Time `gorm:"column:data_estorno" json:"data_estorno,omitempty"`
	HoraEstorno              string     `gorm:"column:hora_estorno" json:"hora_estorno,omitempty"`
	VlrEstorno               float64    `gorm:"column:vlr_estorno" json:"vlr_estorno,omitempty"`
	UsuarioEstorno           string     `gorm:"column:usuario_estorno" json:"usuario_estorno,omitempty"`
	DtHrEventoLiquidacao     *time.Time `gorm:"column:dt_hr_evento_liquidacao" json:"dt_hr_evento_liquidacao,omitempty"`
	DtHrEventoInsercao       *time.Time `gorm:"column:dt_hr_evento_insercao" json:"dt_hr_evento_insercao,omitempty"`
	BloqAlteraRemove         string     `gorm:"column:bloq_altera_remove;default:N" json:"bloq_altera_remove,omitempty"`
	BloqTxMotivo             string     `gorm:"column:bloq_tx_motivo" json:"bloq_tx_motivo,omitempty"`
	TmpControleLoopGeracao   string     `gorm:"column:tmp_controle_loop_geracao" json:"tmp_controle_loop_geracao,omitempty"`
	DhRegistro               *time.Time `gorm:"column:dh_registro;default:now()" json:"dh_registro,omitempty"`
	DerivaMassaErro          string     `gorm:"column:deriva_massa_erro" json:"deriva_massa_erro,omitempty"`
	DerivaMassaDesc          string     `gorm:"column:deriva_massa_desc" json:"deriva_massa_desc,omitempty"`
	HashReplicacao           string     `gorm:"column:hash_replicacao" json:"hash_replicacao,omitempty"`
	ClassificacaoBilling     string     `gorm:"column:classificacao_billing;comment:001- billing puro" json:"classificacao_billing,omitempty"` // 001- billing puro
	LiqCdLeitura             int32      `gorm:"column:liq_cd_leitura" json:"liq_cd_leitura,omitempty"`
	LiqCdExtrato             int32      `gorm:"column:liq_cd_extrato" json:"liq_cd_extrato,omitempty"`
	LiquidadoViaConciliacao  string     `gorm:"column:liquidado_via_conciliacao" json:"liquidado_via_conciliacao,omitempty"`
	CdPlanoDerivacao         int32      `gorm:"column:cd_plano_derivacao" json:"cd_plano_derivacao,omitempty"`
	VlrPago                  float64    `gorm:"column:vlr_pago" json:"vlr_pago,omitempty"`
	VlrSaldo                 float64    `gorm:"column:vlr_saldo" json:"vlr_saldo,omitempty"`
	TeveLiquidacaoParcial    string     `gorm:"column:teve_liquidacao_parcial;default:N" json:"teve_liquidacao_parcial,omitempty"`
	CcDepartamentos          int32      `gorm:"column:cc_departamentos" json:"cc_departamentos,omitempty"`
	CcServicos               int32      `gorm:"column:cc_servicos" json:"cc_servicos,omitempty"`
	CcProjetos               int32      `gorm:"column:cc_projetos" json:"cc_projetos,omitempty"`
	CcAreaNegocios           int32      `gorm:"column:cc_area_negocios" json:"cc_area_negocios,omitempty"`
	Migra                    int32      `gorm:"column:migra" json:"migra,omitempty"`
	ObsMigracao              string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
}

func (*MkPlanoConta) TableName() string {
	return TableNameMkPlanoConta
}
