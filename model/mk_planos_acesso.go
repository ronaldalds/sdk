package model

const TableNameMkPlanosAcesso = "mk_planos_acesso"

type MkPlanosAcesso struct {
	Codplano                 int32   `gorm:"column:codplano;not null" json:"codplano,omitempty"`
	Codservidor              int32   `gorm:"column:codservidor;comment:codigo de integracao com o servidor" json:"codservidor,omitempty"` // codigo de integracao com o servidor
	Descricao                string  `gorm:"column:descricao;not null" json:"descricao,omitempty"`
	MesesContrato            int32   `gorm:"column:meses_contrato;not null;comment:meses do contrato de adesao" json:"meses_contrato,omitempty"` // meses do contrato de adesao
	VelUp                    string  `gorm:"column:vel_up" json:"vel_up,omitempty"`
	VelDown                  string  `gorm:"column:vel_down;comment:velocidade de download" json:"vel_down,omitempty"`                // velocidade de download
	BrtThreUp                string  `gorm:"column:brt_thre_up;comment:velocidade burst threshold" json:"brt_thre_up,omitempty"`      // velocidade burst threshold
	BrtThreDown              string  `gorm:"column:brt_thre_down;comment:thresold down" json:"brt_thre_down,omitempty"`               // thresold down
	BurstUp                  string  `gorm:"column:burst_up;comment:vel de burst para upload" json:"burst_up,omitempty"`              // vel de burst para upload
	BurstDown                string  `gorm:"column:burst_down;comment:velocidade de burst para download" json:"burst_down,omitempty"` // velocidade de burst para download
	BurstTime                int32   `gorm:"column:burst_time;comment:tempo de burst em segundos" json:"burst_time,omitempty"`        // tempo de burst em segundos
	AtivarBurst              string  `gorm:"column:ativar_burst" json:"ativar_burst,omitempty"`
	VlrMensalidade           float64 `gorm:"column:vlr_mensalidade" json:"vlr_mensalidade,omitempty"`
	CirUp                    string  `gorm:"column:cir_up" json:"cir_up,omitempty"`
	CirDown                  string  `gorm:"column:cir_down" json:"cir_down,omitempty"`
	PerfilContrato           int32   `gorm:"column:perfil_contrato" json:"perfil_contrato,omitempty"`
	AtivarPlanoPersonalizado string  `gorm:"column:ativar_plano_personalizado" json:"ativar_plano_personalizado,omitempty"`
	AtivarPersonal           string  `gorm:"column:ativar_personal" json:"ativar_personal,omitempty"`
	AtivarDiario             string  `gorm:"column:ativar_diario" json:"ativar_diario,omitempty"`
	AtivarSemanal            string  `gorm:"column:ativar_semanal" json:"ativar_semanal,omitempty"`
	AtivarMensal             string  `gorm:"column:ativar_mensal" json:"ativar_mensal,omitempty"`
	RegraDiaria              int32   `gorm:"column:regra_diaria" json:"regra_diaria,omitempty"`
	RegraSemanal             int32   `gorm:"column:regra_semanal" json:"regra_semanal,omitempty"`
	RegraMensal              int32   `gorm:"column:regra_mensal" json:"regra_mensal,omitempty"`
	ComporValor              string  `gorm:"column:compor_valor" json:"compor_valor,omitempty"`
	Cfop                     int32   `gorm:"column:cfop" json:"cfop,omitempty"`
	TxIcms                   float64 `gorm:"column:tx_icms" json:"tx_icms,omitempty"`
	TxFust                   float64 `gorm:"column:tx_fust" json:"tx_fust,omitempty"`
	TxFuntel                 float64 `gorm:"column:tx_funtel" json:"tx_funtel,omitempty"`
	BaseCalculoIcms          float64 `gorm:"column:base_calculo_icms" json:"base_calculo_icms,omitempty"`
	ValorBase                float64 `gorm:"column:valor_base" json:"valor_base,omitempty"`
	DescricaoNf              string  `gorm:"column:descricao_nf" json:"descricao_nf,omitempty"`
	PlanoLivreServidor       string  `gorm:"column:plano_livre_servidor" json:"plano_livre_servidor,omitempty"`
	OutroPlanos              string  `gorm:"column:outro_planos" json:"outro_planos,omitempty"`
	Aplicacao                string  `gorm:"column:aplicacao" json:"aplicacao,omitempty"`
	ComporValorNfPre         string  `gorm:"column:compor_valor_nf_pre" json:"compor_valor_nf_pre,omitempty"`
	DescricaoNfPre           string  `gorm:"column:descricao_nf_pre" json:"descricao_nf_pre,omitempty"`
	CfopNfPre                int32   `gorm:"column:cfop_nf_pre" json:"cfop_nf_pre,omitempty"`
	VlrNfPre                 float64 `gorm:"column:vlr_nf_pre" json:"vlr_nf_pre,omitempty"`
	AliquotaNfPre            float64 `gorm:"column:aliquota_nf_pre" json:"aliquota_nf_pre,omitempty"`
	ValorTotalNf             float64 `gorm:"column:valor_total_nf" json:"valor_total_nf,omitempty"`
	ConsiderarVlrConta       string  `gorm:"column:considerar_vlr_conta" json:"considerar_vlr_conta,omitempty"`
	Pis                      float64 `gorm:"column:pis" json:"pis,omitempty"`
	Cofins                   float64 `gorm:"column:cofins" json:"cofins,omitempty"`
	Iss                      float64 `gorm:"column:iss" json:"iss,omitempty"`
	Pnbl                     string  `gorm:"column:pnbl" json:"pnbl,omitempty"`
	Dedicado                 string  `gorm:"column:dedicado" json:"dedicado,omitempty"`
	VlrVelocidade            int32   `gorm:"column:vlr_velocidade" json:"vlr_velocidade,omitempty"`
	VlrVelocidadeGarantido   int32   `gorm:"column:vlr_velocidade_garantido" json:"vlr_velocidade_garantido,omitempty"`
	MesLimiteParcelas        int32   `gorm:"column:mes_limite_parcelas" json:"mes_limite_parcelas,omitempty"`
	AddressList              int32   `gorm:"column:address_list" json:"address_list,omitempty"`
	Inativo                  string  `gorm:"column:inativo" json:"inativo,omitempty"`
	ComporValorAvancado      string  `gorm:"column:compor_valor_avancado" json:"compor_valor_avancado,omitempty"`
	AtivaReciboNumerado      string  `gorm:"column:ativa_recibo_numerado" json:"ativa_recibo_numerado,omitempty"`
	FracaoReciboNumerado     float64 `gorm:"column:fracao_recibo_numerado" json:"fracao_recibo_numerado,omitempty"`
	DescricaoReciboNumerado  string  `gorm:"column:descricao_recibo_numerado" json:"descricao_recibo_numerado,omitempty"`
	ObsReciboNumerado        string  `gorm:"column:obs_recibo_numerado" json:"obs_recibo_numerado,omitempty"`
	TabelaDescontos          int32   `gorm:"column:tabela_descontos" json:"tabela_descontos,omitempty"`
	Prioridade               int32   `gorm:"column:prioridade" json:"prioridade,omitempty"`
	StrMt                    string  `gorm:"column:str_mt" json:"str_mt,omitempty"`
	UndFinSugerida           string  `gorm:"column:und_fin_sugerida" json:"und_fin_sugerida,omitempty"`
	Ncm                      string  `gorm:"column:ncm" json:"ncm,omitempty"`
	FilterID                 string  `gorm:"column:filter_id" json:"filter_id,omitempty"`
	IngressPolicyIpv6        string  `gorm:"column:ingress_policy_ipv6" json:"ingress_policy_ipv6,omitempty"`
	EgressPolicyIpv6         string  `gorm:"column:egress_policy_ipv6" json:"egress_policy_ipv6,omitempty"`
	IngressPolicyIpv4        string  `gorm:"column:ingress_policy_ipv4" json:"ingress_policy_ipv4,omitempty"`
	EgressPolicyIpv4         string  `gorm:"column:egress_policy_ipv4" json:"egress_policy_ipv4,omitempty"`
	VoipEmpresa              int32   `gorm:"column:voip_empresa" json:"voip_empresa,omitempty"`
	TxUpRedVelocidade        int32   `gorm:"column:tx_up_red_velocidade" json:"tx_up_red_velocidade,omitempty"`
	TxDownRedVelocidade      int32   `gorm:"column:tx_down_red_velocidade" json:"tx_down_red_velocidade,omitempty"`
	TxUpRedVelocidadeMed     string  `gorm:"column:tx_up_red_velocidade_med" json:"tx_up_red_velocidade_med,omitempty"`
	TxDownRedVelocidadeMed   string  `gorm:"column:tx_down_red_velocidade_med" json:"tx_down_red_velocidade_med,omitempty"`
	DefBloqCnt               int32   `gorm:"column:def_bloq_cnt" json:"def_bloq_cnt,omitempty"`
	DefBloqRedVel            int32   `gorm:"column:def_bloq_red_vel" json:"def_bloq_red_vel,omitempty"`
	DefDiasBloq              int32   `gorm:"column:def_dias_bloq" json:"def_dias_bloq,omitempty"`
	DefIps                   int32   `gorm:"column:def_ips" json:"def_ips,omitempty"`
	DefGrupo                 int32   `gorm:"column:def_grupo" json:"def_grupo,omitempty"`
	DefSubgrupo              int32   `gorm:"column:def_subgrupo" json:"def_subgrupo,omitempty"`
	DefUndFin                string  `gorm:"column:def_und_fin" json:"def_und_fin,omitempty"`
	DefModeloNf              int32   `gorm:"column:def_modelo_nf" json:"def_modelo_nf,omitempty"`
	DefDesconto              int32   `gorm:"column:def_desconto" json:"def_desconto,omitempty"`
	Tipo                     int32   `gorm:"column:tipo" json:"tipo,omitempty"`
	CdBatch                  int32   `gorm:"column:cd_batch" json:"cd_batch,omitempty"`
	VlrVelocidadeUp          int32   `gorm:"column:vlr_velocidade_up" json:"vlr_velocidade_up,omitempty"`
	CdPlanoIptv              int32   `gorm:"column:cd_plano_iptv" json:"cd_plano_iptv,omitempty"`
	RateID                   int32   `gorm:"column:rate_id" json:"rate_id,omitempty"`
	TelefoniaBlqACobrar      int32   `gorm:"column:telefonia_blq_a_cobrar" json:"telefonia_blq_a_cobrar,omitempty"`
	TelefoniaLdn             int32   `gorm:"column:telefonia_ldn" json:"telefonia_ldn,omitempty"`
	TelefoniaLdi             int32   `gorm:"column:telefonia_ldi" json:"telefonia_ldi,omitempty"`
	TelefoniaVc1             int32   `gorm:"column:telefonia_vc1" json:"telefonia_vc1,omitempty"`
	ComposicaoPlanoVlr       string  `gorm:"column:composicao_plano_vlr" json:"composicao_plano_vlr,omitempty"`
	Grupo                    int32   `gorm:"column:grupo" json:"grupo,omitempty"`
	Limite                   int32   `gorm:"column:limite" json:"limite,omitempty"`
	HabilitarMkHotspot       string  `gorm:"column:habilitar_mk_hotspot" json:"habilitar_mk_hotspot,omitempty"`
	CdEmpresa                int32   `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	CdTelefoniaPlataforma    int32   `gorm:"column:cd_telefonia_plataforma" json:"cd_telefonia_plataforma,omitempty"`
	StrVelocidadeCmts        string  `gorm:"column:str_velocidade_cmts" json:"str_velocidade_cmts,omitempty"`
	IngressPolicyIpv6Reducao string  `gorm:"column:ingress_policy_ipv6_reducao" json:"ingress_policy_ipv6_reducao,omitempty"`
	IngressPolicyIpv4Reducao string  `gorm:"column:ingress_policy_ipv4_reducao" json:"ingress_policy_ipv4_reducao,omitempty"`
	EgressPolicyIpv4Reducao  string  `gorm:"column:egress_policy_ipv4_reducao" json:"egress_policy_ipv4_reducao,omitempty"`
	EgressPolicyIpv6Reducao  string  `gorm:"column:egress_policy_ipv6_reducao" json:"egress_policy_ipv6_reducao,omitempty"`
	CdReguaConfiguracao      int32   `gorm:"column:cd_regua_configuracao;comment:" json:"cd_regua_configuracao,omitempty"` //
	TvOperadora              string  `gorm:"column:tv_operadora" json:"tv_operadora,omitempty"`
	TvCodigoOferta           string  `gorm:"column:tv_codigo_oferta" json:"tv_codigo_oferta,omitempty"`
	StatusTestePlano         int32   `gorm:"column:status_teste_plano;comment:0 - plano ok 1 - problema de derivacao 2- sem valor 3- sem tipo definido 4-Tributação = Outros valores não pode ter base de cálculo e ICMS 5-Tributação = Conforme alíquota está com preenchimento de base de cálculo e/ou ICMS incorretos 6-Tributação = Isentar integralmente não pode ter valor de ICMS e base de cálculo" json:"status_teste_plano,omitempty"` // 0 - plano ok 1 - problema de derivacao 2- sem valor 3- sem tipo definido 4-Tributação = Outros valores não pode ter base de cálculo e ICMS 5-Tributação = Conforme alíquota está com preenchimento de base de cálculo e/ou ICMS incorretos 6-Tributação = Isentar integralmente não pode ter valor de ICMS e base de cálculo
	CdGroupReply             int32   `gorm:"column:cd_group_reply" json:"cd_group_reply,omitempty"`
	VelocidadesFormatadas    string  `gorm:"column:velocidades_formatadas" json:"velocidades_formatadas,omitempty"`
	InCiscoV4                string  `gorm:"column:in_cisco_v4" json:"in_cisco_v4,omitempty"`
	InCiscoV4R               string  `gorm:"column:in_cisco_v4_r" json:"in_cisco_v4_r,omitempty"`
	ECiscoV4                 string  `gorm:"column:e_cisco_v4" json:"e_cisco_v4,omitempty"`
	ECiscoV4R                string  `gorm:"column:e_cisco_v4_r" json:"e_cisco_v4_r,omitempty"`
	InCiscoV6                string  `gorm:"column:in_cisco_v6" json:"in_cisco_v6,omitempty"`
	InCiscoV6R               string  `gorm:"column:in_cisco_v6_r" json:"in_cisco_v6_r,omitempty"`
	ECiscoV6                 string  `gorm:"column:e_cisco_v6" json:"e_cisco_v6,omitempty"`
	ECiscoV6R                string  `gorm:"column:e_cisco_v6_r" json:"e_cisco_v6_r,omitempty"`
	InHuaweiV4               string  `gorm:"column:in_huawei_v4" json:"in_huawei_v4,omitempty"`
	InHuaweiV4R              string  `gorm:"column:in_huawei_v4_r" json:"in_huawei_v4_r,omitempty"`
	EHuaweiV4                string  `gorm:"column:e_huawei_v4" json:"e_huawei_v4,omitempty"`
	EHuaweiV4R               string  `gorm:"column:e_huawei_v4_r" json:"e_huawei_v4_r,omitempty"`
	InHuaweiV6               string  `gorm:"column:in_huawei_v6" json:"in_huawei_v6,omitempty"`
	InHuaweiV6R              string  `gorm:"column:in_huawei_v6_r" json:"in_huawei_v6_r,omitempty"`
	EHuaweiV6                string  `gorm:"column:e_huawei_v6" json:"e_huawei_v6,omitempty"`
	EHuaweiV6R               string  `gorm:"column:e_huawei_v6_r" json:"e_huawei_v6_r,omitempty"`
	InIpoeV4                 string  `gorm:"column:in_ipoe_v4" json:"in_ipoe_v4,omitempty"`
	InIpoeV4R                string  `gorm:"column:in_ipoe_v4_r" json:"in_ipoe_v4_r,omitempty"`
	EIpoeV4                  string  `gorm:"column:e_ipoe_v4" json:"e_ipoe_v4,omitempty"`
	EIpoeV4R                 string  `gorm:"column:e_ipoe_v4_r" json:"e_ipoe_v4_r,omitempty"`
	InIpoeV6                 string  `gorm:"column:in_ipoe_v6" json:"in_ipoe_v6,omitempty"`
	InIpoeV6R                string  `gorm:"column:in_ipoe_v6_r" json:"in_ipoe_v6_r,omitempty"`
	EIpoeV6                  string  `gorm:"column:e_ipoe_v6" json:"e_ipoe_v6,omitempty"`
	EIpoeV6R                 string  `gorm:"column:e_ipoe_v6_r" json:"e_ipoe_v6_r,omitempty"`
	VelocidadesFormatadasR   string  `gorm:"column:velocidades_formatadas_r" json:"velocidades_formatadas_r,omitempty"`
	ExclusivoCompoCrm        string  `gorm:"column:exclusivo_compo_crm;default:N;comment:se S pode somente ser usado dentro de CRM" json:"exclusivo_compo_crm,omitempty"` // se S pode somente ser usado dentro de CRM
	GUIDTmp                  string  `gorm:"column:guid_tmp" json:"guid_tmp,omitempty"`
	TipoPlanoPersonalizado   int32   `gorm:"column:tipo_plano_personalizado" json:"tipo_plano_personalizado,omitempty"`
	VendaEmAreasPre          string  `gorm:"column:venda_em_areas_pre" json:"venda_em_areas_pre,omitempty"`
	Plano                    int32   `gorm:"column:plano" json:"plano,omitempty"`
	QtdConexoes              int32   `gorm:"column:qtd_conexoes" json:"qtd_conexoes,omitempty"`
	InteresseFibra           string  `gorm:"column:interesse_fibra" json:"interesse_fibra,omitempty"`
	InteresseTelefone        string  `gorm:"column:interesse_telefone" json:"interesse_telefone,omitempty"`
	InteresseWireless        string  `gorm:"column:interesse_wireless" json:"interesse_wireless,omitempty"`
	InteresseTv              string  `gorm:"column:interesse_tv" json:"interesse_tv,omitempty"`
	InteresseOutros          string  `gorm:"column:interesse_outros" json:"interesse_outros,omitempty"`
	Tecnologia               string  `gorm:"column:tecnologia" json:"tecnologia,omitempty"`
	VelBitsDici              string  `gorm:"column:vel_bits_dici" json:"vel_bits_dici,omitempty"`
	ObsMigracao              string  `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
	SippulseVinculoBcore     int32   `gorm:"column:sippulse_vinculo_bcore" json:"sippulse_vinculo_bcore,omitempty"`
	PlanoDependencia         int32   `gorm:"column:plano_dependencia" json:"plano_dependencia,omitempty"`
}

func (*MkPlanosAcesso) TableName() string {
	return TableNameMkPlanosAcesso
}
