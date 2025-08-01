package model

const TableNameMkProfilePgto = "mk_profile_pgto"

type MkProfilePgto struct {
	Codprofile                    int32   `gorm:"column:codprofile;not null" json:"codprofile,omitempty"`
	NomeProfile                   string  `gorm:"column:nome_profile;not null" json:"nome_profile,omitempty"`
	FormaPgto                     int32   `gorm:"column:forma_pgto;comment:codigo da forma de pgto associada." json:"forma_pgto,omitempty"` // codigo da forma de pgto associada.
	CodigoCedente                 string  `gorm:"column:codigo_cedente" json:"codigo_cedente,omitempty"`
	Agencia                       string  `gorm:"column:agencia" json:"agencia,omitempty"`
	ContaCorrente                 string  `gorm:"column:conta_corrente" json:"conta_corrente,omitempty"`
	InstrucoesBoleto              string  `gorm:"column:instrucoes_boleto;comment:instrucoes sugeridas para a impressao no boleto." json:"instrucoes_boleto,omitempty"` // instrucoes sugeridas para a impressao no boleto.
	ValorBoleto                   float64 `gorm:"column:valor_boleto;comment:valor de acrescimo no boleto." json:"valor_boleto,omitempty"`                              // valor de acrescimo no boleto.
	Posto                         string  `gorm:"column:posto;comment:posto bancario em caso de boleto." json:"posto,omitempty"`                                        // posto bancario em caso de boleto.
	Banco                         string  `gorm:"column:banco;comment:numero de referencia do banco." json:"banco,omitempty"`                                           // numero de referencia do banco.
	Ano                           string  `gorm:"column:ano;comment:ano de referencia boleto bancario." json:"ano,omitempty"`                                           // ano de referencia boleto bancario.
	Sequencial                    int32   `gorm:"column:sequencial;comment:numero do sequencial do boleto." json:"sequencial,omitempty"`                                // numero do sequencial do boleto.
	Entidade                      int32   `gorm:"column:entidade" json:"entidade,omitempty"`
	NomeCedente                   string  `gorm:"column:nome_cedente;comment:nome do emitente da cobranca." json:"nome_cedente,omitempty"` // nome do emitente da cobranca.
	DocCedente                    string  `gorm:"column:doc_cedente" json:"doc_cedente,omitempty"`
	LiberacaoAutomatica           string  `gorm:"column:liberacao_automatica;comment:se a liberacao e automatica no pgto." json:"liberacao_automatica,omitempty"` // se a liberacao e automatica no pgto.
	MsgLinha1                     string  `gorm:"column:msg_linha1" json:"msg_linha1,omitempty"`
	MsgLinha2                     string  `gorm:"column:msg_linha2" json:"msg_linha2,omitempty"`
	MsgLinha3                     string  `gorm:"column:msg_linha3" json:"msg_linha3,omitempty"`
	MsgLinha4                     string  `gorm:"column:msg_linha4" json:"msg_linha4,omitempty"`
	MsgLinha5                     string  `gorm:"column:msg_linha5" json:"msg_linha5,omitempty"`
	MsgLinha6                     string  `gorm:"column:msg_linha6" json:"msg_linha6,omitempty"`
	MsgLinha7                     string  `gorm:"column:msg_linha7" json:"msg_linha7,omitempty"`
	MsgLinha8                     string  `gorm:"column:msg_linha8" json:"msg_linha8,omitempty"`
	TxAux1                        string  `gorm:"column:tx_aux1" json:"tx_aux1,omitempty"`
	TxAux2                        string  `gorm:"column:tx_aux2" json:"tx_aux2,omitempty"`
	TxAux3                        string  `gorm:"column:tx_aux3" json:"tx_aux3,omitempty"`
	TxAux4                        string  `gorm:"column:tx_aux4" json:"tx_aux4,omitempty"`
	TxAux5                        string  `gorm:"column:tx_aux5" json:"tx_aux5,omitempty"`
	TxAux6                        string  `gorm:"column:tx_aux6" json:"tx_aux6,omitempty"`
	TxAux7                        string  `gorm:"column:tx_aux7" json:"tx_aux7,omitempty"`
	LocalPgto                     string  `gorm:"column:local_pgto" json:"local_pgto,omitempty"`
	Aceite                        string  `gorm:"column:aceite" json:"aceite,omitempty"`
	Carteira                      string  `gorm:"column:carteira" json:"carteira,omitempty"`
	CedenteBarra                  string  `gorm:"column:cedente_barra;comment:Obrigatoria para carteira especial Unibanco" json:"cedente_barra,omitempty"` // Obrigatoria para carteira especial Unibanco
	Modalidade                    string  `gorm:"column:modalidade;comment:Obrigatorio Banco Nossa Caixa" json:"modalidade,omitempty"`                     // Obrigatorio Banco Nossa Caixa
	NumConvenio                   string  `gorm:"column:num_convenio;comment:Obrigatorio Banco do Brasil." json:"num_convenio,omitempty"`                  // Obrigatorio Banco do Brasil.
	URLJsp                        string  `gorm:"column:url_jsp" json:"url_jsp,omitempty"`
	Empresa                       int32   `gorm:"column:empresa" json:"empresa,omitempty"`
	SeqRemessa                    int32   `gorm:"column:seq_remessa" json:"seq_remessa,omitempty"`
	DvConta                       string  `gorm:"column:dv_conta" json:"dv_conta,omitempty"`
	AtivarSms                     string  `gorm:"column:ativar_sms" json:"ativar_sms,omitempty"`
	AtivarEmail                   string  `gorm:"column:ativar_email" json:"ativar_email,omitempty"`
	TxSms                         string  `gorm:"column:tx_sms" json:"tx_sms,omitempty"`
	TxEmail                       string  `gorm:"column:tx_email" json:"tx_email,omitempty"`
	SMTP                          string  `gorm:"column:smtp" json:"smtp,omitempty"`
	LoginSMTP                     string  `gorm:"column:login_smtp" json:"login_smtp,omitempty"`
	PasswordSMTP                  string  `gorm:"column:password_smtp" json:"password_smtp,omitempty"`
	EmailAutenticao               string  `gorm:"column:email_autenticao" json:"email_autenticao,omitempty"`
	NomeDoBanco                   string  `gorm:"column:nome_do_banco" json:"nome_do_banco,omitempty"`
	RetRemoveDv                   string  `gorm:"column:ret_remove_dv" json:"ret_remove_dv,omitempty"`
	DiasReagendamento             int32   `gorm:"column:dias_reagendamento" json:"dias_reagendamento,omitempty"`
	RegraFormatacao               int32   `gorm:"column:regra_formatacao" json:"regra_formatacao,omitempty"`
	TipoCobranca                  string  `gorm:"column:tipo_cobranca" json:"tipo_cobranca,omitempty"`
	PostarTitulo                  string  `gorm:"column:postar_titulo" json:"postar_titulo,omitempty"`
	PercentualMultaAtraso         string  `gorm:"column:percentual_multa_atraso" json:"percentual_multa_atraso,omitempty"`
	Especie                       string  `gorm:"column:especie" json:"especie,omitempty"`
	DiasParaProtesto              int32   `gorm:"column:dias_para_protesto" json:"dias_para_protesto,omitempty"`
	PercentualJurosDia            string  `gorm:"column:percentual_juros_dia" json:"percentual_juros_dia,omitempty"`
	PercentualDescontoAntecipacao string  `gorm:"column:percentual_desconto_antecipacao" json:"percentual_desconto_antecipacao,omitempty"`
	Protestar                     string  `gorm:"column:protestar" json:"protestar,omitempty"`
	ControleAno                   int32   `gorm:"column:controle_ano" json:"controle_ano,omitempty"`
	AuxTipoLayout                 int32   `gorm:"column:aux_tipo_layout" json:"aux_tipo_layout,omitempty"`
	AgenciaDv                     string  `gorm:"column:agencia_dv" json:"agencia_dv,omitempty"`
	CdOperacao                    string  `gorm:"column:cd_operacao" json:"cd_operacao,omitempty"`
	LocalPgto2                    string  `gorm:"column:local_pgto2" json:"local_pgto2,omitempty"`
	CdFornecidoAgencia            string  `gorm:"column:cd_fornecido_agencia" json:"cd_fornecido_agencia,omitempty"`
	CdFornecidoAgenciaDv          string  `gorm:"column:cd_fornecido_agencia_dv" json:"cd_fornecido_agencia_dv,omitempty"`
	ModeloImpressao               string  `gorm:"column:modelo_impressao" json:"modelo_impressao,omitempty"`
	LinkImagem                    string  `gorm:"column:link_imagem" json:"link_imagem,omitempty"`
	MensagensCabecalho            string  `gorm:"column:mensagens_cabecalho" json:"mensagens_cabecalho,omitempty"`
	Sigcb                         string  `gorm:"column:sigcb" json:"sigcb,omitempty"`
	ImgCabecalho                  []uint8 `gorm:"column:img_cabecalho" json:"img_cabecalho,omitempty"`
	LayoutRetorno                 int32   `gorm:"column:layout_retorno" json:"layout_retorno,omitempty"`
	LayoutRemessa                 int32   `gorm:"column:layout_remessa" json:"layout_remessa,omitempty"`
	AuxFrentePreencherRecibo      string  `gorm:"column:aux_frente_preencher_recibo" json:"aux_frente_preencher_recibo,omitempty"`
	AuxFrenteLinha7               string  `gorm:"column:aux_frente_linha7" json:"aux_frente_linha7,omitempty"`
	AuxFrenteLinha8               string  `gorm:"column:aux_frente_linha8" json:"aux_frente_linha8,omitempty"`
	AuxFrenteLinha9               string  `gorm:"column:aux_frente_linha9" json:"aux_frente_linha9,omitempty"`
	GerarNossoNum                 string  `gorm:"column:gerar_nosso_num" json:"gerar_nosso_num,omitempty"`
	QuebrarPagina                 string  `gorm:"column:quebrar_pagina" json:"quebrar_pagina,omitempty"`
	CcAlternativa                 string  `gorm:"column:cc_alternativa;comment:numero alternativo de conta bancaria para processamento de retorno." json:"cc_alternativa,omitempty"` // numero alternativo de conta bancaria para processamento de retorno.
	RemGerarCarne                 string  `gorm:"column:rem_gerar_carne" json:"rem_gerar_carne,omitempty"`
	Beta                          string  `gorm:"column:beta" json:"beta,omitempty"`
	ImpNossoNumRetBanco           string  `gorm:"column:imp_nosso_num_ret_banco" json:"imp_nosso_num_ret_banco,omitempty"`
	OperacaoConta                 int32   `gorm:"column:operacao_conta" json:"operacao_conta,omitempty"`
	MargemSuperior                float64 `gorm:"column:margem_superior" json:"margem_superior,omitempty"`
	LytPersonalizado              int32   `gorm:"column:lyt_personalizado" json:"lyt_personalizado,omitempty"`
	Apelido                       string  `gorm:"column:apelido" json:"apelido,omitempty"`
	NomenLiquidacao               string  `gorm:"column:nomen_liquidacao" json:"nomen_liquidacao,omitempty"`
	ComportamentoTarifa           string  `gorm:"column:comportamento_tarifa" json:"comportamento_tarifa,omitempty"`
	ComportamentoAcresimos        string  `gorm:"column:comportamento_acresimos" json:"comportamento_acresimos,omitempty"`
	ComportamentoDescontos        string  `gorm:"column:comportamento_descontos" json:"comportamento_descontos,omitempty"`
	ComportamentoAMenos           string  `gorm:"column:comportamento_a_menos" json:"comportamento_a_menos,omitempty"`
	NomenAcrescimos               string  `gorm:"column:nomen_acrescimos" json:"nomen_acrescimos,omitempty"`
	NomenDescontos                string  `gorm:"column:nomen_descontos" json:"nomen_descontos,omitempty"`
	NomenTarifas                  string  `gorm:"column:nomen_tarifas" json:"nomen_tarifas,omitempty"`
	AvisoLiquidaAMenos            string  `gorm:"column:aviso_liquida_a_menos" json:"aviso_liquida_a_menos,omitempty"`
	Flags                         string  `gorm:"column:flags" json:"flags,omitempty"`
	CdPeculiaridade               int32   `gorm:"column:cd_peculiaridade" json:"cd_peculiaridade,omitempty"`
	Observacoes                   string  `gorm:"column:observacoes" json:"observacoes,omitempty"`
	Homologado                    string  `gorm:"column:homologado" json:"homologado,omitempty"`
	DvCedente                     string  `gorm:"column:dv_cedente" json:"dv_cedente,omitempty"`
	Flags2                        string  `gorm:"column:flags_2" json:"flags_2,omitempty"`
	CdFlash                       string  `gorm:"column:cd_flash" json:"cd_flash,omitempty"`
	CdUnificacaoProfile           int32   `gorm:"column:cd_unificacao_profile" json:"cd_unificacao_profile,omitempty"`
	NomenScm                      string  `gorm:"column:nomen_scm" json:"nomen_scm,omitempty"`
	NomenSva                      string  `gorm:"column:nomen_sva" json:"nomen_sva,omitempty"`
	PracaDevolucao                string  `gorm:"column:praca_devolucao" json:"praca_devolucao,omitempty"`
	OcultarDoc                    string  `gorm:"column:ocultar_doc" json:"ocultar_doc,omitempty"`
	TipoImpressao                 string  `gorm:"column:tipo_impressao" json:"tipo_impressao,omitempty"`
	ProfileConsEntConfirmada      int32   `gorm:"column:profile_cons_ent_confirmada" json:"profile_cons_ent_confirmada,omitempty"`
	CdModeloImpressao             int32   `gorm:"column:cd_modelo_impressao" json:"cd_modelo_impressao,omitempty"`
	FebrabanSegmento              string  `gorm:"column:febraban_segmento" json:"febraban_segmento,omitempty"`
	FebrabanIdentificacao         string  `gorm:"column:febraban_identificacao" json:"febraban_identificacao,omitempty"`
	ModeloImpressaoEnvio          int32   `gorm:"column:modelo_impressao_envio" json:"modelo_impressao_envio,omitempty"`
	DesativarProfile              string  `gorm:"column:desativar_profile" json:"desativar_profile,omitempty"`
	CdEmpresa                     int32   `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	AceitaBaixaAlteracao          string  `gorm:"column:aceita_baixa_alteracao" json:"aceita_baixa_alteracao,omitempty"`
	AceitaAlteracao               string  `gorm:"column:aceita_alteracao" json:"aceita_alteracao,omitempty"`
	FebrabanFormatoVenc           string  `gorm:"column:febraban_formato_venc" json:"febraban_formato_venc,omitempty"`
	StatusProfile                 int32   `gorm:"column:status_profile" json:"status_profile,omitempty"`
	BeneficiarioNome              string  `gorm:"column:beneficiario_nome" json:"beneficiario_nome,omitempty"`
	BeneficiarioDocumento         string  `gorm:"column:beneficiario_documento" json:"beneficiario_documento,omitempty"`
	BeneficiarioLogradouro        string  `gorm:"column:beneficiario_logradouro" json:"beneficiario_logradouro,omitempty"`
	BeneficiarioBairro            string  `gorm:"column:beneficiario_bairro" json:"beneficiario_bairro,omitempty"`
	BeneficiarioCep               string  `gorm:"column:beneficiario_cep" json:"beneficiario_cep,omitempty"`
	BeneficiarioUf                string  `gorm:"column:beneficiario_uf" json:"beneficiario_uf,omitempty"`
	BeneficiarioMunicipio         string  `gorm:"column:beneficiario_municipio" json:"beneficiario_municipio,omitempty"`
	BeneficiarioEmail             string  `gorm:"column:beneficiario_email" json:"beneficiario_email,omitempty"`
	BeneficiarioCodCliente        string  `gorm:"column:beneficiario_cod_cliente" json:"beneficiario_cod_cliente,omitempty"`
	BeneficiarioTelefone          string  `gorm:"column:beneficiario_telefone" json:"beneficiario_telefone,omitempty"`
	L2AgenciaNumero               string  `gorm:"column:l2_agencia_numero" json:"l2_agencia_numero,omitempty"`
	L2AgenciaDv                   string  `gorm:"column:l2_agencia_dv" json:"l2_agencia_dv,omitempty"`
	L2ContaCorrenteNumero         string  `gorm:"column:l2_conta_corrente_numero" json:"l2_conta_corrente_numero,omitempty"`
	L2ContaCorrenteDv             string  `gorm:"column:l2_conta_corrente_dv" json:"l2_conta_corrente_dv,omitempty"`
	L2Carteira                    string  `gorm:"column:l2_carteira" json:"l2_carteira,omitempty"`
	L2CarteiraModalidade          string  `gorm:"column:l2_carteira_modalidade" json:"l2_carteira_modalidade,omitempty"`
	L2CarteiraTipo                string  `gorm:"column:l2_carteira_tipo" json:"l2_carteira_tipo,omitempty"`
	L2CarteiraConvenio            string  `gorm:"column:l2_carteira_convenio" json:"l2_carteira_convenio,omitempty"`
	L2CarteiraConvenioDv          string  `gorm:"column:l2_carteira_convenio_dv" json:"l2_carteira_convenio_dv,omitempty"`
	L2Ios                         string  `gorm:"column:l2_ios" json:"l2_ios,omitempty"`
	L2ModeloDefault               string  `gorm:"column:l2_modelo_default" json:"l2_modelo_default,omitempty"`
	L2ModeloSac                   string  `gorm:"column:l2_modelo_sac" json:"l2_modelo_sac,omitempty"`
	L2ConvenioNumero              string  `gorm:"column:l2_convenio_numero" json:"l2_convenio_numero,omitempty"`
	L2MultaAtraso                 float64 `gorm:"column:l2_multa_atraso" json:"l2_multa_atraso,omitempty"`
	L2AcrescimosDiario            float64 `gorm:"column:l2_acrescimos_diario" json:"l2_acrescimos_diario,omitempty"`
	L2LayoutRemessa               int32   `gorm:"column:l2_layout_remessa" json:"l2_layout_remessa,omitempty"`
	GatewayID                     string  `gorm:"column:gateway_id" json:"gateway_id,omitempty"`
	GatewaySecret                 string  `gorm:"column:gateway_secret" json:"gateway_secret,omitempty"`
	GerencianetSandbox            string  `gorm:"column:gerencianet_sandbox" json:"gerencianet_sandbox,omitempty"`
	URLNotificacoesGw             string  `gorm:"column:url_notificacoes_gw" json:"url_notificacoes_gw,omitempty"`
	GwProducao                    string  `gorm:"column:gw_producao" json:"gw_producao,omitempty"`
	GwDiasFrenteMonitoramento     int32   `gorm:"column:gw_dias_frente_monitoramento" json:"gw_dias_frente_monitoramento,omitempty"`
	GwDiasAtrasMonitoramento      int32   `gorm:"column:gw_dias_atras_monitoramento" json:"gw_dias_atras_monitoramento,omitempty"`
	GwTempoRenovaToken            int32   `gorm:"column:gw_tempo_renova_token" json:"gw_tempo_renova_token,omitempty"`
	L2CodigoCedente               string  `gorm:"column:l2_codigo_cedente" json:"l2_codigo_cedente,omitempty"`
	CartaoLojaNumero              string  `gorm:"column:cartao_loja_numero" json:"cartao_loja_numero,omitempty"`
	CartaoLojaChave               string  `gorm:"column:cartao_loja_chave" json:"cartao_loja_chave,omitempty"`
	CartaoPadrao                  string  `gorm:"column:cartao_padrao" json:"cartao_padrao,omitempty"`
	CartaoProducao                string  `gorm:"column:cartao_producao" json:"cartao_producao,omitempty"`
	L2ControleImpressao           string  `gorm:"column:l2_controle_impressao" json:"l2_controle_impressao,omitempty"`
	OrdemPgtoAtiva                string  `gorm:"column:ordem_pgto_ativa" json:"ordem_pgto_ativa,omitempty"`
	CdLayoutOrdemPgto             int32   `gorm:"column:cd_layout_ordem_pgto" json:"cd_layout_ordem_pgto,omitempty"`
	PathLogo                      string  `gorm:"column:path_logo" json:"path_logo,omitempty"`
	CartaoHabTestesHomologacao    string  `gorm:"column:cartao_hab_testes_homologacao" json:"cartao_hab_testes_homologacao,omitempty"`
	L2Posto                       string  `gorm:"column:l2_posto" json:"l2_posto,omitempty"`
	CdPessoaCartaoHomologa        int32   `gorm:"column:cd_pessoa_cartao_homologa" json:"cd_pessoa_cartao_homologa,omitempty"`
	NumBancoExtra                 string  `gorm:"column:num_banco_extra" json:"num_banco_extra,omitempty"`
	Cooperativa                   string  `gorm:"column:cooperativa" json:"cooperativa,omitempty"`
	IgnorarNossoNumeroDv          string  `gorm:"column:ignorar_nosso_numero_dv;default:N" json:"ignorar_nosso_numero_dv,omitempty"`
	PixChave                      string  `gorm:"column:pix_chave" json:"pix_chave,omitempty"`
	CdCredencialPixBb             int32   `gorm:"column:cd_credencial_pix_bb" json:"cd_credencial_pix_bb,omitempty"`
	StatusProfilePix              int32   `gorm:"column:status_profile_pix" json:"status_profile_pix,omitempty"`
	PixHabTestesHomologacao       string  `gorm:"column:pix_hab_testes_homologacao" json:"pix_hab_testes_homologacao,omitempty"`
	CdPessoaPixHomologa           int32   `gorm:"column:cd_pessoa_pix_homologa" json:"cd_pessoa_pix_homologa,omitempty"`
	WebhookMknextConfigurado      string  `gorm:"column:webhook_mknext_configurado" json:"webhook_mknext_configurado,omitempty"`
	RegistrarBoletoAuto           string  `gorm:"column:registrar_boleto_auto" json:"registrar_boleto_auto,omitempty"`
	VersaoAPIPix                  string  `gorm:"column:versao_api_pix" json:"versao_api_pix,omitempty"`
	CdProfilePix                  int32   `gorm:"column:cd_profile_pix;comment:Coluna criada para especificar a profile que será utilizada para gerar o qr code pix quando a profile for de uma entidade MK Next." json:"cd_profile_pix,omitempty"` // Coluna criada para especificar a profile que será utilizada para gerar o qr code pix quando a profile for de uma entidade MK Next.
	CdContaPixConcilia            int32   `gorm:"column:cd_conta_pix_concilia" json:"cd_conta_pix_concilia,omitempty"`
}

func (*MkProfilePgto) TableName() string {
	return TableNameMkProfilePgto
}
