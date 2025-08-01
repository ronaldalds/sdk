package model

import (
	"time"
)

const TableNameMkContrato = "mk_contratos"

type MkContrato struct {
	// Relationsships Init
	Pessoa           MkPessoa            `gorm:"foreignKey:codpessoa;references:cliente;" json:"pessoa,omitempty"`
	Conexoes         []MkConexao         `gorm:"foreignKey:contrato;references:codcontrato;" json:"conexoes,omitempty"`
	ContratosGerados []MkContratosGerado `gorm:"foreignKey:contrato;references:codcontrato;" json:"contratos_gerados,omitempty"`
	CrmLeads         MkCrmLead           `gorm:"foreignKey:codlead;references:cd_lead;" json:"crm_lead,omitempty"`
	CrmProduto       MkCrmProduto        `gorm:"foreignKey:cd_plano_principal;references:plano_acesso;" json:"crm_produto,omitempty"`
	PlanosAcesso     []MkPlanosAcesso    `gorm:"foreignKey:codplano;references:plano_acesso;" json:"planos_acesso,omitempty"`
	// Relationsships End
	Codcontrato                 int32      `gorm:"column:codcontrato;not null" json:"codcontrato,omitempty"`
	Cliente                     int32      `gorm:"column:cliente;not null;comment:cliente associado ao contrato" json:"cliente,omitempty"`
	CdLead                      int32      `gorm:"column:cd_lead" json:"cd_lead,omitempty"`
	Cancelado                   string     `gorm:"column:cancelado;default:N;comment:S ou N" json:"cancelado,omitempty"`
	Adesao                      *time.Time `gorm:"column:adesao;not null" json:"adesao,omitempty"`
	PrevisaoVencimento          *time.Time `gorm:"column:previsao_vencimento;comment:data prevista para o vencimento." json:"previsao_vencimento,omitempty"`                                                                                                                           // data prevista para o vencimento.
	PlanoAcesso                 int32      `gorm:"column:plano_acesso;not null;comment:plano de acesso." json:"plano_acesso,omitempty"`                                                                                                                                                // plano de acesso.
	BloquearAutomaticamente     int32      `gorm:"column:bloquear_automaticamente;comment:0 para nao bloquear automaticamente, ou o numero de dias para o bloqueio." json:"bloquear_automaticamente,omitempty"`                                                                        // 0 para nao bloquear automaticamente, ou o numero de dias para o bloqueio.
	FormaPgto                   int32      `gorm:"column:forma_pgto;comment:1 - Cartao de Credito, 2 - Debito em Conta, 3 - Duplicata Bancaria" json:"forma_pgto,omitempty"`                                                                                                           // 1 - Cartao de Credito, 2 - Debito em Conta, 3 - Duplicata Bancaria
	ProfilePgto                 int32      `gorm:"column:profile_pgto;comment:profile de configuracao da forma de pgto." json:"profile_pgto,omitempty"`                                                                                                                                // profile de configuracao da forma de pgto.
	Vendedor                    int32      `gorm:"column:vendedor;comment:vendedor responsavel pelo contrato." json:"vendedor,omitempty"`                                                                                                                                              // vendedor responsavel pelo contrato.
	RenovadoAte                 *time.Time `gorm:"column:renovado_ate;comment:em caso de renovacao de contrato." json:"renovado_ate,omitempty"`                                                                                                                                        // em caso de renovacao de contrato.
	HistoricoEventos            string     `gorm:"column:historico_eventos;comment:historico do contrato, em caso de renovacoes, trocas de planos, etc." json:"historico_eventos,omitempty"`                                                                                           // historico do contrato, em caso de renovacoes, trocas de planos, etc.
	AceitaEmails                string     `gorm:"column:aceita_emails;comment:se aceita emails de cobranca e notificacao." json:"aceita_emails,omitempty"`                                                                                                                            // se aceita emails de cobranca e notificacao.
	BloqAutoContrato            int32      `gorm:"column:bloq_auto_contrato;comment:bloquear automaticamente caso o cliente nao aceite o contrato eletronico. 0 para nao bloquear, ou o informe o numero de dias que o cliente tem para aceitar." json:"bloq_auto_contrato,omitempty"` // bloquear automaticamente caso o cliente nao aceite o contrato eletronico. 0 para nao bloquear, ou o informe o numero de dias que o cliente tem para aceitar.
	DiaVencimento               int32      `gorm:"column:dia_vencimento;comment:dia de vencimento das duplicatas ou parcelas." json:"dia_vencimento,omitempty"`                                                                                                                        // dia de vencimento das duplicatas ou parcelas.
	ContratoEletronico          string     `gorm:"column:contrato_eletronico;comment:se vai utilizar o contrato eletronico." json:"contrato_eletronico,omitempty"`                                                                                                                     // se vai utilizar o contrato eletronico.
	VlrAdesao                   float64    `gorm:"column:vlr_adesao;comment:caso seja cobrado valor de adesao" json:"vlr_adesao,omitempty"`                                                                                                                                            // caso seja cobrado valor de adesao
	ParcelamentoAdesao          int32      `gorm:"column:parcelamento_adesao;comment:numero de vezes q o valor de adesao pode ser parcelado." json:"parcelamento_adesao,omitempty"`                                                                                                    // numero de vezes q o valor de adesao pode ser parcelado.
	CodigoOs                    int32      `gorm:"column:codigo_os;comment:caso a adesao esteja associada a uma OS de instalacao." json:"codigo_os,omitempty"`                                                                                                                         // caso a adesao esteja associada a uma OS de instalacao.
	NumeroIps                   int32      `gorm:"column:numero_ips;comment:numero de ips permitidos para esse contrato." json:"numero_ips,omitempty"`                                                                                                                                 // numero de ips permitidos para esse contrato.
	PrimeiroVencimento          *time.Time `gorm:"column:primeiro_vencimento" json:"primeiro_vencimento,omitempty"`
	PercentualMulta             float64    `gorm:"column:percentual_multa" json:"percentual_multa,omitempty"`
	PercentualAcrescimo         float64    `gorm:"column:percentual_acrescimo" json:"percentual_acrescimo,omitempty"`
	AceitaPgtoSac               string     `gorm:"column:aceita_pgto_sac" json:"aceita_pgto_sac,omitempty"`
	PercentualMultaSac          float64    `gorm:"column:percentual_multa_sac" json:"percentual_multa_sac,omitempty"`
	PrecisaRenovar              string     `gorm:"column:precisa_renovar" json:"precisa_renovar,omitempty"`
	PrePago                     string     `gorm:"column:pre_pago" json:"pre_pago,omitempty"`
	DescontoAplicado            int32      `gorm:"column:desconto_aplicado" json:"desconto_aplicado,omitempty"`
	DtCancelamento              *time.Time `gorm:"column:dt_cancelamento" json:"dt_cancelamento,omitempty"`
	UserCancelamento            string     `gorm:"column:user_cancelamento" json:"user_cancelamento,omitempty"`
	MotivoCancelamento          string     `gorm:"column:motivo_cancelamento" json:"motivo_cancelamento,omitempty"`
	RemoverParcelasRestantes    string     `gorm:"column:remover_parcelas_restantes" json:"remover_parcelas_restantes,omitempty"`
	ParcelasARemover            string     `gorm:"column:parcelas_a_remover" json:"parcelas_a_remover,omitempty"`
	RemoverConexao              string     `gorm:"column:remover_conexao" json:"remover_conexao,omitempty"`
	DtAgendadaRemocao           *time.Time `gorm:"column:dt_agendada_remocao" json:"dt_agendada_remocao,omitempty"`
	HistoricoAlteracao          string     `gorm:"column:historico_alteracao" json:"historico_alteracao,omitempty"`
	NovoDiaVencimento           int32      `gorm:"column:novo_dia_vencimento" json:"novo_dia_vencimento,omitempty"`
	NovoPlano                   int32      `gorm:"column:novo_plano" json:"novo_plano,omitempty"`
	ParcelasAAlterar            string     `gorm:"column:parcelas_a_alterar" json:"parcelas_a_alterar,omitempty"`
	AgenciaDebito               int32      `gorm:"column:agencia_debito" json:"agencia_debito,omitempty"`
	IDClienteBanco              string     `gorm:"column:id_cliente_banco" json:"id_cliente_banco,omitempty"`
	Nf21                        string     `gorm:"column:nf_21" json:"nf_21,omitempty"`
	NfPre                       string     `gorm:"column:nf_pre" json:"nf_pre,omitempty"`
	AlteraValorCancel           string     `gorm:"column:altera_valor_cancel" json:"altera_valor_cancel,omitempty"`
	VlrParcialCancelamento      float64    `gorm:"column:vlr_parcial_cancelamento" json:"vlr_parcial_cancelamento,omitempty"`
	AlteraValorAlter            string     `gorm:"column:altera_valor_alter" json:"altera_valor_alter,omitempty"`
	VlrParcialAlteracao         float64    `gorm:"column:vlr_parcial_alteracao" json:"vlr_parcial_alteracao,omitempty"`
	RegraComissao               int32      `gorm:"column:regra_comissao" json:"regra_comissao,omitempty"`
	PreCadastro                 int32      `gorm:"column:pre_cadastro" json:"pre_cadastro,omitempty"`
	ContratoAssinado            string     `gorm:"column:contrato_assinado" json:"contrato_assinado,omitempty"`
	DtAssinatura                *time.Time `gorm:"column:dt_assinatura" json:"dt_assinatura,omitempty"`
	Grupo                       int32      `gorm:"column:grupo" json:"grupo,omitempty"`
	Subgrupo                    int32      `gorm:"column:subgrupo" json:"subgrupo,omitempty"`
	BloqueiaPorOs               string     `gorm:"column:bloqueia_por_os" json:"bloqueia_por_os,omitempty"`
	RenovarAuto                 string     `gorm:"column:renovar_auto" json:"renovar_auto,omitempty"`
	IgnoraRenovacaoAuto         string     `gorm:"column:ignora_renovacao_auto;default:N" json:"ignora_renovacao_auto,omitempty"`
	RenovarPorValorParcela      string     `gorm:"column:renovar_por_valor_parcela" json:"renovar_por_valor_parcela,omitempty"`
	TmpComMensalidade           string     `gorm:"column:tmp_com_mensalidade" json:"tmp_com_mensalidade,omitempty"`
	TmpDtVencimento             *time.Time `gorm:"column:tmp_dt_vencimento" json:"tmp_dt_vencimento,omitempty"`
	TmpVlrParcela               float64    `gorm:"column:tmp_vlr_parcela" json:"tmp_vlr_parcela,omitempty"`
	InativarConexao             string     `gorm:"column:inativar_conexao;default:N" json:"inativar_conexao,omitempty"`
	InativarCliente             string     `gorm:"column:inativar_cliente;default:N" json:"inativar_cliente,omitempty"`
	Comodato                    int32      `gorm:"column:comodato" json:"comodato,omitempty"`
	IDClienteEmpresa            string     `gorm:"column:id_cliente_empresa" json:"id_cliente_empresa,omitempty"`
	ReciboNumerado              string     `gorm:"column:recibo_numerado" json:"recibo_numerado,omitempty"`
	NovaFormaPgto               int32      `gorm:"column:nova_forma_pgto" json:"nova_forma_pgto,omitempty"`
	NovaProfile                 int32      `gorm:"column:nova_profile" json:"nova_profile,omitempty"`
	ComodatoSimples             string     `gorm:"column:comodato_simples" json:"comodato_simples,omitempty"`
	DescricaoComodato           string     `gorm:"column:descricao_comodato" json:"descricao_comodato,omitempty"`
	DtComodato                  *time.Time `gorm:"column:dt_comodato" json:"dt_comodato,omitempty"`
	ObsComodato                 string     `gorm:"column:obs_comodato" json:"obs_comodato,omitempty"`
	TmpLogin                    string     `gorm:"column:tmp_login" json:"tmp_login,omitempty"`
	ConcederDescontoInatividade string     `gorm:"column:conceder_desconto_inatividade" json:"conceder_desconto_inatividade,omitempty"`
	UnidadeFinanceira           string     `gorm:"column:unidade_financeira" json:"unidade_financeira,omitempty"`
	Operador                    string     `gorm:"column:operador" json:"operador,omitempty"`
	ComRegRetorno               string     `gorm:"column:com_reg_retorno" json:"com_reg_retorno,omitempty"`
	ComDtRet                    *time.Time `gorm:"column:com_dt_ret" json:"com_dt_ret,omitempty"`
	ComHrRet                    string     `gorm:"column:com_hr_ret" json:"com_hr_ret,omitempty"`
	ComDescEquipamentoRet       string     `gorm:"column:com_desc_equipamento_ret" json:"com_desc_equipamento_ret,omitempty"`
	ComObsRet                   string     `gorm:"column:com_obs_ret" json:"com_obs_ret,omitempty"`
	ComOperadorRet              string     `gorm:"column:com_operador_ret" json:"com_operador_ret,omitempty"`
	ComOsRet                    int32      `gorm:"column:com_os_ret" json:"com_os_ret,omitempty"`
	VlrRenovacao                float64    `gorm:"column:vlr_renovacao" json:"vlr_renovacao,omitempty"`
	ModeloNf                    string     `gorm:"column:modelo_nf" json:"modelo_nf,omitempty"`
	SupPlanoDifere              string     `gorm:"column:sup_plano_difere" json:"sup_plano_difere,omitempty"`
	SupContratosVencidos        string     `gorm:"column:sup_contratos_vencidos" json:"sup_contratos_vencidos,omitempty"`
	SupParcelasPorVencer        int32      `gorm:"column:sup_parcelas_por_vencer" json:"sup_parcelas_por_vencer,omitempty"`
	SupParcelasVencidas         int32      `gorm:"column:sup_parcelas_vencidas" json:"sup_parcelas_vencidas,omitempty"`
	SupCancComCon               string     `gorm:"column:sup_canc_com_con" json:"sup_canc_com_con,omitempty"`
	SupContSemCon               string     `gorm:"column:sup_cont_sem_con" json:"sup_cont_sem_con,omitempty"`
	SupDiasParcAtrasadas        int32      `gorm:"column:sup_dias_parc_atrasadas" json:"sup_dias_parc_atrasadas,omitempty"`
	ComodatoDtRet               *time.Time `gorm:"column:comodato_dt_ret" json:"comodato_dt_ret,omitempty"`
	ComodatoHrRet               string     `gorm:"column:comodato_hr_ret" json:"comodato_hr_ret,omitempty"`
	ComodatoRet                 string     `gorm:"column:comodato_ret" json:"comodato_ret,omitempty"`
	ComodatoInfoRet             string     `gorm:"column:comodato_info_ret" json:"comodato_info_ret,omitempty"`
	ComodatoRetNumOs            int32      `gorm:"column:comodato_ret_num_os" json:"comodato_ret_num_os,omitempty"`
	SupProbRetCom               string     `gorm:"column:sup_prob_ret_com" json:"sup_prob_ret_com,omitempty"`
	SupCancComParc              int32      `gorm:"column:sup_canc_com_parc" json:"sup_canc_com_parc,omitempty"`
	SupVlrUltParcela            float64    `gorm:"column:sup_vlr_ult_parcela" json:"sup_vlr_ult_parcela,omitempty"`
	SupVlrParcDfPlano           string     `gorm:"column:sup_vlr_parc_df_plano" json:"sup_vlr_parc_df_plano,omitempty"`
	SupTxComoRenovar            string     `gorm:"column:sup_tx_como_renovar" json:"sup_tx_como_renovar,omitempty"`
	Nfse                        string     `gorm:"column:nfse" json:"nfse,omitempty"`
	MotivoCancelamento2         int32      `gorm:"column:motivo_cancelamento_2" json:"motivo_cancelamento_2,omitempty"`
	IgnoraUpdateConexao         string     `gorm:"column:ignora_update_conexao" json:"ignora_update_conexao,omitempty"`
	DiasReferencia              int32      `gorm:"column:dias_referencia" json:"dias_referencia,omitempty"`
	TxRefDiasCancelamento       string     `gorm:"column:tx_ref_dias_cancelamento" json:"tx_ref_dias_cancelamento,omitempty"`
	TxRefDiasAlteracao          string     `gorm:"column:tx_ref_dias_alteracao" json:"tx_ref_dias_alteracao,omitempty"`
	Suspenso                    string     `gorm:"column:suspenso;default:N" json:"suspenso,omitempty"`
	SuspensoDt                  *time.Time `gorm:"column:suspenso_dt" json:"suspenso_dt,omitempty"`
	SuspensoHr                  string     `gorm:"column:suspenso_hr" json:"suspenso_hr,omitempty"`
	SuspensoOperador            string     `gorm:"column:suspenso_operador" json:"suspenso_operador,omitempty"`
	SuspensoMotivo              int32      `gorm:"column:suspenso_motivo" json:"suspenso_motivo,omitempty"`
	CdContaTmpAlteracao         int32      `gorm:"column:cd_conta_tmp_alteracao" json:"cd_conta_tmp_alteracao,omitempty"`
	MetodoFaturamento           string     `gorm:"column:metodo_faturamento" json:"metodo_faturamento,omitempty"`
	Versao                      string     `gorm:"column:versao" json:"versao,omitempty"`
	DiasRedBanda                int32      `gorm:"column:dias_red_banda" json:"dias_red_banda,omitempty"`
	FatAuto                     string     `gorm:"column:fat_auto" json:"fat_auto,omitempty"`
	CdRegraFaturamento          int32      `gorm:"column:cd_regra_faturamento" json:"cd_regra_faturamento,omitempty"`
	CdEnderecoNf                int64      `gorm:"column:cd_endereco_nf" json:"cd_endereco_nf,omitempty"`
	AguardaAtivacao             string     `gorm:"column:aguarda_ativacao" json:"aguarda_ativacao,omitempty"`
	DataHoraAtivacao            *time.Time `gorm:"column:data_hora_ativacao" json:"data_hora_ativacao,omitempty"`
	UserAtivacao                string     `gorm:"column:user_ativacao" json:"user_ativacao,omitempty"`
	CdContratoVinculado         int32      `gorm:"column:cd_contrato_vinculado" json:"cd_contrato_vinculado,omitempty"`
	CdEnderecoNfse              int32      `gorm:"column:cd_endereco_nfse" json:"cd_endereco_nfse,omitempty"`
	IptvConteudoAdulto          string     `gorm:"column:iptv_conteudo_adulto" json:"iptv_conteudo_adulto,omitempty"`
	DtVenda                     *time.Time `gorm:"column:dt_venda" json:"dt_venda,omitempty"`
	DtAtivacao                  *time.Time `gorm:"column:dt_ativacao" json:"dt_ativacao,omitempty"`
	ObsCancelamentoSist         string     `gorm:"column:obs_cancelamento_sist" json:"obs_cancelamento_sist,omitempty"`
	ObsAlteracaoSist            string     `gorm:"column:obs_alteracao_sist" json:"obs_alteracao_sist,omitempty"`
	ObsNf                       string     `gorm:"column:obs_nf" json:"obs_nf,omitempty"`
	CdEmpresa                   int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	DtAgendadaInativConex       *time.Time `gorm:"column:dt_agendada_inativ_conex" json:"dt_agendada_inativ_conex,omitempty"`
	DtAgendadaInativCli         *time.Time `gorm:"column:dt_agendada_inativ_cli" json:"dt_agendada_inativ_cli,omitempty"`
	CdPerfilCartaoCredito       int32      `gorm:"column:cd_perfil_cartao_credito" json:"cd_perfil_cartao_credito,omitempty"`
	CdReguaConfiguracao         int32      `gorm:"column:cd_regua_configuracao" json:"cd_regua_configuracao,omitempty"`
	DtHrInsert                  *time.Time `gorm:"column:dt_hr_insert" json:"dt_hr_insert,omitempty"`
	GUIDIDGerecao               string     `gorm:"column:guid_id_gerecao" json:"guid_id_gerecao,omitempty"`
	AbrirOsRetirada             string     `gorm:"column:abrir_os_retirada" json:"abrir_os_retirada,omitempty"`
	TvCodAssinante              string     `gorm:"column:tv_cod_assinante" json:"tv_cod_assinante,omitempty"`
	TvIdorder                   string     `gorm:"column:tv_idorder" json:"tv_idorder,omitempty"`
	CdTabelaSLA                 int32      `gorm:"column:cd_tabela_sla" json:"cd_tabela_sla,omitempty"`
	ReverContasAntesFaturar     string     `gorm:"column:rever_contas_antes_faturar" json:"rever_contas_antes_faturar,omitempty"`
	TipoPlano                   int32      `gorm:"column:tipo_plano" json:"tipo_plano,omitempty"`
	IgnoraFaturaNaGeracaoConta  string     `gorm:"column:ignora_fatura_na_geracao_conta;default:N" json:"ignora_fatura_na_geracao_conta,omitempty"`
	CriarContaPeloVlrSugerido   string     `gorm:"column:criar_conta_pelo_vlr_sugerido;default:N" json:"criar_conta_pelo_vlr_sugerido,omitempty"`
	ErroAoAtivarAuto            string     `gorm:"column:erro_ao_ativar_auto;default:N" json:"erro_ao_ativar_auto,omitempty"`
	CdAgrupador                 int32      `gorm:"column:cd_agrupador" json:"cd_agrupador,omitempty"`
	PlanoContasAdesao           string     `gorm:"column:plano_contas_adesao" json:"plano_contas_adesao,omitempty"`
	ObsAdesao                   string     `gorm:"column:obs_adesao" json:"obs_adesao,omitempty"`
	VlrAdesaoCondicao           float64    `gorm:"column:vlr_adesao_condicao" json:"vlr_adesao_condicao,omitempty"`
	ParcelamentoAdesaoCondicao  int32      `gorm:"column:parcelamento_adesao_condicao" json:"parcelamento_adesao_condicao,omitempty"`
	ConsultorCrm                int32      `gorm:"column:consultor_crm" json:"consultor_crm,omitempty"`
	EnviarBv                    string     `gorm:"column:enviar_bv" json:"enviar_bv,omitempty"`
	EmailDstBv                  string     `gorm:"column:email_dst_bv" json:"email_dst_bv,omitempty"`
	AssuntoBv                   string     `gorm:"column:assunto_bv" json:"assunto_bv,omitempty"`
	TxBv                        string     `gorm:"column:tx_bv" json:"tx_bv,omitempty"`
	GravarModeloBv              string     `gorm:"column:gravar_modelo_bv" json:"gravar_modelo_bv,omitempty"`
	GravarModeloBvPlano         string     `gorm:"column:gravar_modelo_bv_plano" json:"gravar_modelo_bv_plano,omitempty"`
	GravarBvDefault             string     `gorm:"column:gravar_bv_default" json:"gravar_bv_default,omitempty"`
	CdBv                        int32      `gorm:"column:cd_bv" json:"cd_bv,omitempty"`
	FormaPgtoAde                int32      `gorm:"column:forma_pgto_ade" json:"forma_pgto_ade,omitempty"`
	GerarAdesaoImediata         string     `gorm:"column:gerar_adesao_imediata" json:"gerar_adesao_imediata,omitempty"`
	AdesaoGerada                string     `gorm:"column:adesao_gerada;default:N" json:"adesao_gerada,omitempty"`
	GerarAdeAgora               string     `gorm:"column:gerar_ade_agora;default:N" json:"gerar_ade_agora,omitempty"`
	ObsExtraNf                  string     `gorm:"column:obs_extra_nf" json:"obs_extra_nf,omitempty"`
	TvUsuario                   string     `gorm:"column:tv_usuario" json:"tv_usuario,omitempty"`
	TvSenha                     string     `gorm:"column:tv_senha" json:"tv_senha,omitempty"`
	ContaEmailBv                int32      `gorm:"column:conta_email_bv" json:"conta_email_bv,omitempty"`
	TvSituacao                  string     `gorm:"column:tv_situacao;not null;default:N;comment:N - Normal P - Bloqueio parcial T - Bloqueio total" json:"tv_situacao,omitempty"` // N - Normal P - Bloqueio parcial T - Bloqueio total
	CdComposicaoDerivacao       int32      `gorm:"column:cd_composicao_derivacao" json:"cd_composicao_derivacao,omitempty"`
	NumTerminalTelefonico       string     `gorm:"column:num_terminal_telefonico" json:"num_terminal_telefonico,omitempty"`
	LicencaCodigoIntegracao     string     `gorm:"column:licenca_codigo_integracao" json:"licenca_codigo_integracao,omitempty"`
	CdNatOperacaoExtra          int32      `gorm:"column:cd_nat_operacao_extra" json:"cd_nat_operacao_extra,omitempty"`
	PlanoAcesso1                int32      `gorm:"column:plano_acesso1" json:"plano_acesso1,omitempty"`
	Mig1                        int32      `gorm:"column:mig1" json:"mig1,omitempty"`
	PrimeiraParcAdesao          *time.Time `gorm:"column:primeira_parc_adesao" json:"primeira_parc_adesao,omitempty"`
	Migra                       int32      `gorm:"column:migra" json:"migra,omitempty"`
	DataFidelidade              *time.Time `gorm:"column:data_fidelidade" json:"data_fidelidade,omitempty"`
	DestacarIrNf2122            string     `gorm:"column:destacar_ir_nf_21_22;comment:Task SC-6711" json:"destacar_ir_nf_21_22,omitempty"` // Task SC-6711
	LeadExterno                 int32      `gorm:"column:lead_externo;comment:Task SC-6912" json:"lead_externo,omitempty"`                 // Task SC-6912
	DestacarPisNf2122           string     `gorm:"column:destacar_pis_nf_21_22" json:"destacar_pis_nf_21_22,omitempty"`
	DestacarCofinsNf2122        string     `gorm:"column:destacar_cofins_nf_21_22" json:"destacar_cofins_nf_21_22,omitempty"`
	DestacarCsslNf2122          string     `gorm:"column:destacar_cssl_nf_21_22" json:"destacar_cssl_nf_21_22,omitempty"`
	EnderecoCobranca            string     `gorm:"column:endereco_cobranca" json:"endereco_cobranca,omitempty"`
	GerarDescontoIcms           string     `gorm:"column:gerar_desconto_icms;comment:Task SC-7382" json:"gerar_desconto_icms,omitempty"` // Task SC-7382
	DestacarIrpjNf2122          string     `gorm:"column:destacar_irpj_nf_21_22" json:"destacar_irpj_nf_21_22,omitempty"`
	ObsMigracao                 string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
	ObsFidelidade               string     `gorm:"column:obs_fidelidade" json:"obs_fidelidade,omitempty"`
	LicencaBloqueioStatus       string     `gorm:"column:licenca_bloqueio_status" json:"licenca_bloqueio_status,omitempty"`
}

func (*MkContrato) TableName() string {
	return TableNameMkContrato
}
