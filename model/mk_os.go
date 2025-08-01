package model

import (
	"time"
)

const TableNameMkO = "mk_os"

type MkOS struct {
	// Relationsships Init
	Compromisso                 MkCompromisso                 `gorm:"foreignKey:cd_integracao;references:codos" json:"compromisso,omitempty"`
	Logradouro                  MkLogradouro                  `gorm:"foreignKey:codlogradouro;references:cd_logradouro" json:"logradouro,omitempty"`
	OSClassificacaoEncerramento MkOsClassificacaoEncerramento `gorm:"foreignKey:codclassifenc;references:classificacao_encerramento" json:"os_classificacao_encerramento,omitempty"`
	OSTipo                      MkOsTipo                      `gorm:"foreignKey:codostipo;references:tipo_os" json:"os_tipo,omitempty"`
	Pessoa                      MkPessoa                      `gorm:"foreignKey:codpessoa;references:cliente;" json:"pessoa,omitempty"`
	UsuarioTecnico              FrUsuario                     `gorm:"foreignKey:usr_codigo;references:operador_fech_tecnico" json:"usuario_tecnico,omitempty"`
	// Relationsships End
	Codos                         int32      `gorm:"column:codos;not null" json:"codos,omitempty"`
	Cliente                       int32      `gorm:"column:cliente" json:"cliente,omitempty"`
	CdLogradouro                  int32      `gorm:"column:cd_logradouro" json:"cd_logradouro,omitempty"`
	OperadorFechTecnico           string     `gorm:"column:operador_fech_tecnico" json:"operador_fech_tecnico,omitempty"`
	ClassificacaoEncerramento     int32      `gorm:"column:classificacao_encerramento" json:"classificacao_encerramento,omitempty"`
	TipoOsID                      int32      `gorm:"column:tipo_os" json:"tipo_os,omitempty"`
	DataAbertura                  *time.Time `gorm:"column:data_abertura;not null" json:"data_abertura,omitempty"`
	HoraAbertura                  string     `gorm:"column:hora_abertura;not null" json:"hora_abertura,omitempty"`
	DefeitoAssociado              int32      `gorm:"column:defeito_associado" json:"defeito_associado,omitempty"`
	IDEquipamento                 string     `gorm:"column:id_equipamento" json:"id_equipamento,omitempty"`
	InfoExtras                    string     `gorm:"column:info_extras" json:"info_extras,omitempty"`
	Status                        int32      `gorm:"column:status;not null;comment:1- aguardando agendamento, 2- aguardando conclusao, 3- Encerrado." json:"status,omitempty"` // 1- aguardando agendamento, 2- aguardando conclusao, 3- Encerrado.
	Observacoes                   string     `gorm:"column:observacoes" json:"observacoes,omitempty"`
	DataFechamento                *time.Time `gorm:"column:data_fechamento" json:"data_fechamento,omitempty"`
	HoraFechamento                string     `gorm:"column:hora_fechamento" json:"hora_fechamento,omitempty"`
	TotalServicos                 float64    `gorm:"column:total_servicos" json:"total_servicos,omitempty"`
	TotalPecas                    float64    `gorm:"column:total_pecas" json:"total_pecas,omitempty"`
	TecnicoResponsavel            int32      `gorm:"column:tecnico_responsavel" json:"tecnico_responsavel,omitempty"`
	TecnicoAtendimento            int32      `gorm:"column:tecnico_atendimento" json:"tecnico_atendimento,omitempty"`
	Indicacoes                    string     `gorm:"column:indicacoes" json:"indicacoes,omitempty"`
	DefeitoReclamado              string     `gorm:"column:defeito_reclamado" json:"defeito_reclamado,omitempty"`
	DefeitoConstatado             string     `gorm:"column:defeito_constatado" json:"defeito_constatado,omitempty"`
	GarantiaServicos              int32      `gorm:"column:garantia_servicos" json:"garantia_servicos,omitempty"`
	GarantiaPecas                 int32      `gorm:"column:garantia_pecas" json:"garantia_pecas,omitempty"`
	Faturado                      string     `gorm:"column:faturado;not null;default:N;comment:S/N" json:"faturado,omitempty"` // S/N
	ConexaoAssociada              int32      `gorm:"column:conexao_associada" json:"conexao_associada,omitempty"`
	ProtocoloOs                   string     `gorm:"column:protocolo_os" json:"protocolo_os,omitempty"`
	EmLaboratorio                 string     `gorm:"column:em_laboratorio;comment:S/N" json:"em_laboratorio,omitempty"` // S/N
	DataEntLab                    *time.Time `gorm:"column:data_ent_lab" json:"data_ent_lab,omitempty"`
	HoraEntLab                    string     `gorm:"column:hora_ent_lab" json:"hora_ent_lab,omitempty"`
	Rat                           string     `gorm:"column:rat" json:"rat,omitempty"`
	CodAgenda                     int32      `gorm:"column:cod_agenda" json:"cod_agenda,omitempty"`
	ServicoPrestado               string     `gorm:"column:servico_prestado" json:"servico_prestado,omitempty"`
	DataPrevEntrega               *time.Time `gorm:"column:data_prev_entrega" json:"data_prev_entrega,omitempty"`
	HoraPrevEntrega               string     `gorm:"column:hora_prev_entrega" json:"hora_prev_entrega,omitempty"`
	IndicacaoEntrega              string     `gorm:"column:indicacao_entrega" json:"indicacao_entrega,omitempty"`
	Encerrado                     string     `gorm:"column:encerrado" json:"encerrado,omitempty"`
	ClassificacaoCobranca         string     `gorm:"column:classificacao_cobranca" json:"classificacao_cobranca,omitempty"`
	EquipamentoDescritivo         string     `gorm:"column:equipamento_descritivo" json:"equipamento_descritivo,omitempty"`
	CabRelDescr                   string     `gorm:"column:cab_rel_descr" json:"cab_rel_descr,omitempty"`
	RodRelDescr                   string     `gorm:"column:rod_rel_descr" json:"rod_rel_descr,omitempty"`
	TxExtra                       string     `gorm:"column:tx_extra" json:"tx_extra,omitempty"`
	PrecoServico                  int32      `gorm:"column:preco_servico" json:"preco_servico,omitempty"`
	Operador                      string     `gorm:"column:operador" json:"operador,omitempty"`
	Grupo                         int32      `gorm:"column:grupo" json:"grupo,omitempty"`
	Subgrupo                      int32      `gorm:"column:subgrupo" json:"subgrupo,omitempty"`
	PontoRecebimento              int32      `gorm:"column:ponto_recebimento" json:"ponto_recebimento,omitempty"`
	GUIDTemp                      string     `gorm:"column:guid_temp" json:"guid_temp,omitempty"`
	VinculosIntegrador            string     `gorm:"column:vinculos_integrador" json:"vinculos_integrador,omitempty"`
	VlrDebAbertura                float64    `gorm:"column:vlr_deb_abertura" json:"vlr_deb_abertura,omitempty"`
	DebDtPesquisa                 *time.Time `gorm:"column:deb_dt_pesquisa" json:"deb_dt_pesquisa,omitempty"`
	DebTx                         string     `gorm:"column:deb_tx" json:"deb_tx,omitempty"`
	VlrBaixaImediata              float64    `gorm:"column:vlr_baixa_imediata" json:"vlr_baixa_imediata,omitempty"`
	UndFinSugerida                string     `gorm:"column:und_fin_sugerida" json:"und_fin_sugerida,omitempty"`
	TotalEncerrado                float64    `gorm:"column:total_encerrado" json:"total_encerrado,omitempty"`
	VersaoEncerramento            string     `gorm:"column:versao_encerramento" json:"versao_encerramento,omitempty"`
	Sincronizada                  string     `gorm:"column:sincronizada" json:"sincronizada,omitempty"`
	CdUf                          int32      `gorm:"column:cd_uf" json:"cd_uf,omitempty"`
	CdCidade                      int32      `gorm:"column:cd_cidade" json:"cd_cidade,omitempty"`
	CdBairro                      int32      `gorm:"column:cd_bairro" json:"cd_bairro,omitempty"`
	NumEndereco                   int32      `gorm:"column:num_endereco" json:"num_endereco,omitempty"`
	Cep                           string     `gorm:"column:cep" json:"cep,omitempty"`
	Complemento                   string     `gorm:"column:complemento" json:"complemento,omitempty"`
	TmpGUIDPendencias             string     `gorm:"column:tmp_guid_pendencias" json:"tmp_guid_pendencias,omitempty"`
	CdAtendimento                 int32      `gorm:"column:cd_atendimento" json:"cd_atendimento,omitempty"`
	FechamentoTecnico             string     `gorm:"column:fechamento_tecnico" json:"fechamento_tecnico,omitempty"`
	DtHrFechamentoTec             *time.Time `gorm:"column:dt_hr_fechamento_tec" json:"dt_hr_fechamento_tec,omitempty"`
	ModoFechTecnico               int32      `gorm:"column:modo_fech_tecnico" json:"modo_fech_tecnico,omitempty"`
	CdCondominio                  int32      `gorm:"column:cd_condominio" json:"cd_condominio,omitempty"`
	SincronizarAppMobile          string     `gorm:"column:sincronizar_app_mobile" json:"sincronizar_app_mobile,omitempty"`
	ResponsavelClienteVisita      string     `gorm:"column:responsavel_cliente_visita" json:"responsavel_cliente_visita,omitempty"`
	CdEndAdic                     int32      `gorm:"column:cd_end_adic" json:"cd_end_adic,omitempty"`
	LatitudeFechamento            string     `gorm:"column:latitude_fechamento" json:"latitude_fechamento,omitempty"`
	LongitudeFechamento           string     `gorm:"column:longitude_fechamento" json:"longitude_fechamento,omitempty"`
	UltimoStatusAppMk             string     `gorm:"column:ultimo_status_app_mk" json:"ultimo_status_app_mk,omitempty"`
	UltimoStatusAppMkTx           string     `gorm:"column:ultimo_status_app_mk_tx" json:"ultimo_status_app_mk_tx,omitempty"`
	CdEmpresa                     int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	UltimoStatusAppOp             int32      `gorm:"column:ultimo_status_app_op" json:"ultimo_status_app_op,omitempty"`
	NotaCliente                   int32      `gorm:"column:nota_cliente" json:"nota_cliente,omitempty"`
	LocalAvaliacao                string     `gorm:"column:local_avaliacao" json:"local_avaliacao,omitempty"`
	ComentarioAvaliacao           string     `gorm:"column:comentario_avaliacao" json:"comentario_avaliacao,omitempty"`
	AvaliacaoEnviada              string     `gorm:"column:avaliacao_enviada" json:"avaliacao_enviada,omitempty"`
	AvaliacaoEnviadaEm            *time.Time `gorm:"column:avaliacao_enviada_em" json:"avaliacao_enviada_em,omitempty"`
	RetornoAvaliacaoEm            *time.Time `gorm:"column:retorno_avaliacao_em" json:"retorno_avaliacao_em,omitempty"`
	FechamentoAtraso              string     `gorm:"column:fechamento_atraso" json:"fechamento_atraso,omitempty"`
	DtHrLimiteFechamento          *time.Time `gorm:"column:dt_hr_limite_fechamento" json:"dt_hr_limite_fechamento,omitempty"`
	TempoTotalMin                 int32      `gorm:"column:tempo_total_min" json:"tempo_total_min,omitempty"`
	TempoTotalTx                  string     `gorm:"column:tempo_total_tx" json:"tempo_total_tx,omitempty"`
	TempoAtrasoMin                int32      `gorm:"column:tempo_atraso_min" json:"tempo_atraso_min,omitempty"`
	TempoAtrasoTx                 string     `gorm:"column:tempo_atraso_tx" json:"tempo_atraso_tx,omitempty"`
	Reinscidencia                 string     `gorm:"column:reinscidencia" json:"reinscidencia,omitempty"`
	CdOsBaseReinscidencia         int32      `gorm:"column:cd_os_base_reinscidencia" json:"cd_os_base_reinscidencia,omitempty"`
	TmpSLA                        int32      `gorm:"column:tmp_sla" json:"tmp_sla,omitempty"`
	Latitude                      string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude                     string     `gorm:"column:longitude" json:"longitude,omitempty"`
	Cdagendagrupo                 int32      `gorm:"column:cdagendagrupo" json:"cdagendagrupo,omitempty"`
	DtHrFechamento                *time.Time `gorm:"column:dt_hr_fechamento" json:"dt_hr_fechamento,omitempty"`
	Categoria                     int32      `gorm:"column:categoria;default:1;comment:1- cliente 2- os provedor" json:"categoria,omitempty"` // 1- cliente 2- os provedor
	CdContrato                    int32      `gorm:"column:cd_contrato" json:"cd_contrato,omitempty"`
	NivelSLA                      int32      `gorm:"column:nivel_sla" json:"nivel_sla,omitempty"`
	IDHashInsercao                string     `gorm:"column:id_hash_insercao" json:"id_hash_insercao,omitempty"`
	AgendaDom                     string     `gorm:"column:agenda_dom" json:"agenda_dom,omitempty"`
	AgendaSeg                     string     `gorm:"column:agenda_seg" json:"agenda_seg,omitempty"`
	AgendaTer                     string     `gorm:"column:agenda_ter" json:"agenda_ter,omitempty"`
	AgendaQua                     string     `gorm:"column:agenda_qua" json:"agenda_qua,omitempty"`
	AgendaQui                     string     `gorm:"column:agenda_qui" json:"agenda_qui,omitempty"`
	AgendaSex                     string     `gorm:"column:agenda_sex" json:"agenda_sex,omitempty"`
	AgendaSab                     string     `gorm:"column:agenda_sab" json:"agenda_sab,omitempty"`
	AgendaDomHrInicio             string     `gorm:"column:agenda_dom_hr_inicio" json:"agenda_dom_hr_inicio,omitempty"`
	AgendaDomHrFim                string     `gorm:"column:agenda_dom_hr_fim" json:"agenda_dom_hr_fim,omitempty"`
	AgendaSegHrInicio             string     `gorm:"column:agenda_seg_hr_inicio" json:"agenda_seg_hr_inicio,omitempty"`
	AgendaTerHrInicio             string     `gorm:"column:agenda_ter_hr_inicio" json:"agenda_ter_hr_inicio,omitempty"`
	AgendaQuaHrInicio             string     `gorm:"column:agenda_qua_hr_inicio" json:"agenda_qua_hr_inicio,omitempty"`
	AgendaQuiHrInicio             string     `gorm:"column:agenda_qui_hr_inicio" json:"agenda_qui_hr_inicio,omitempty"`
	AgendaSexHrInicio             string     `gorm:"column:agenda_sex_hr_inicio" json:"agenda_sex_hr_inicio,omitempty"`
	AgendaSabHrInicio             string     `gorm:"column:agenda_sab_hr_inicio" json:"agenda_sab_hr_inicio,omitempty"`
	AgendaSegHrFim                string     `gorm:"column:agenda_seg_hr_fim" json:"agenda_seg_hr_fim,omitempty"`
	AgendaTerHrFim                string     `gorm:"column:agenda_ter_hr_fim" json:"agenda_ter_hr_fim,omitempty"`
	AgendaQuaHrFim                string     `gorm:"column:agenda_qua_hr_fim" json:"agenda_qua_hr_fim,omitempty"`
	AgendaQuiHrFim                string     `gorm:"column:agenda_qui_hr_fim" json:"agenda_qui_hr_fim,omitempty"`
	AgendaSexHrFim                string     `gorm:"column:agenda_sex_hr_fim" json:"agenda_sex_hr_fim,omitempty"`
	AgendaSabHrFim                string     `gorm:"column:agenda_sab_hr_fim" json:"agenda_sab_hr_fim,omitempty"`
	TxResumo                      string     `gorm:"column:tx_resumo" json:"tx_resumo,omitempty"`
	NotificacaoParadaAssociada    int32      `gorm:"column:notificacao_parada_associada" json:"notificacao_parada_associada,omitempty"`
	LocalManutencao               int32      `gorm:"column:local_manutencao" json:"local_manutencao,omitempty"`
	LocalDescritivo               string     `gorm:"column:local_descritivo" json:"local_descritivo,omitempty"`
	CdQuipamentoOsProvedor        int32      `gorm:"column:cd_quipamento_os_provedor" json:"cd_quipamento_os_provedor,omitempty"`
	OperadorFechamento            string     `gorm:"column:operador_fechamento" json:"operador_fechamento,omitempty"`
	DhInsert                      *time.Time `gorm:"column:dh_insert" json:"dh_insert,omitempty"`
	StatusAppMk                   int32      `gorm:"column:status_app_mk;comment:Coluna destinada ao status da O.S perante o aplicativo agentes da MK: 1 = Aguardando 2 = Deslocamento iniciado 3 = Servico iniciado 4 = O.S encerrada" json:"status_app_mk,omitempty"` // Coluna destinada ao status da O.S perante o aplicativo agentes da MK: 1 = Aguardando 2 = Deslocamento iniciado 3 = Servico iniciado 4 = O.S encerrada
	Emitida                       string     `gorm:"column:emitida" json:"emitida,omitempty"`
	ProvedorLocal                 int32      `gorm:"column:provedor_local" json:"provedor_local,omitempty"`
	ProvedorEquipamento           int32      `gorm:"column:provedor_equipamento" json:"provedor_equipamento,omitempty"`
	CdOsImpressa                  int32      `gorm:"column:cd_os_impressa" json:"cd_os_impressa,omitempty"`
	ImgAssinatura                 []uint8    `gorm:"column:img_assinatura" json:"img_assinatura,omitempty"`
	FecharAtendimento             string     `gorm:"column:fechar_atendimento" json:"fechar_atendimento,omitempty"`
	SolicitarAvaliacao            string     `gorm:"column:solicitar_avaliacao" json:"solicitar_avaliacao,omitempty"`
	TempoFechamentoOs             float64    `gorm:"column:tempo_fechamento_os;comment:tempo para fechamento de OS em horas." json:"tempo_fechamento_os,omitempty"` // tempo para fechamento de OS em horas.
	CdControleSLA                 int32      `gorm:"column:cd_controle_sla" json:"cd_controle_sla,omitempty"`
	CdWorkspace                   int32      `gorm:"column:cd_workspace" json:"cd_workspace,omitempty"`
	DhInicioAtividade             *time.Time `gorm:"column:dh_inicio_atividade" json:"dh_inicio_atividade,omitempty"`
	DhFimAtividade                *time.Time `gorm:"column:dh_fim_atividade" json:"dh_fim_atividade,omitempty"`
	TempoAtividade                int32      `gorm:"column:tempo_atividade" json:"tempo_atividade,omitempty"`
	TempoAtrasoAgenda             int32      `gorm:"column:tempo_atraso_agenda;comment:tempo em minutos para o atraso para o inicio da atividade. se o vlr for negativo, a atividade foi executada de forma antecipada." json:"tempo_atraso_agenda,omitempty"` // tempo em minutos para o atraso para o inicio da atividade. se o vlr for negativo, a atividade foi executada de forma antecipada.
	CdMovimentacao                int32      `gorm:"column:cd_movimentacao;comment:Codigo da movimentacao de venda" json:"cd_movimentacao,omitempty"`                                                                                                          // Codigo da movimentacao de venda
	CdLead                        int32      `gorm:"column:cd_lead" json:"cd_lead,omitempty"`
	Tecnologia                    string     `gorm:"column:tecnologia" json:"tecnologia,omitempty"`
	TipoIP                        int32      `gorm:"column:tipo_ip" json:"tipo_ip,omitempty"`
	CodenderecoIP                 int32      `gorm:"column:codendereco_ip;comment:codigo do endereco de ip" json:"codendereco_ip,omitempty"` // codigo do endereco de ip
	MacAddressConsiderado         string     `gorm:"column:mac_address_considerado" json:"mac_address_considerado,omitempty"`
	UsarIpv6                      string     `gorm:"column:usar_ipv6" json:"usar_ipv6,omitempty"`
	PoolMk                        string     `gorm:"column:pool_mk" json:"pool_mk,omitempty"`
	PoolAssociado                 int32      `gorm:"column:pool_associado" json:"pool_associado,omitempty"`
	RangePrefixoIpv6              int32      `gorm:"column:range_prefixo_ipv6;comment:Este e o range de prefixo que esta associado a tabela de ranges." json:"range_prefixo_ipv6,omitempty"`                       // Este e o range de prefixo que esta associado a tabela de ranges.
	RangeDelegadoIpv6             int32      `gorm:"column:range_delegado_ipv6;comment:Este e o range de prefixo delegado que o cliente esta vinculado na tabela de ranges." json:"range_delegado_ipv6,omitempty"` // Este e o range de prefixo delegado que o cliente esta vinculado na tabela de ranges.
	Wpa2Key                       string     `gorm:"column:wpa2_key" json:"wpa2_key,omitempty"`
	UserEap                       string     `gorm:"column:user_eap" json:"user_eap,omitempty"`
	SenhaEquipamento              string     `gorm:"column:senha_equipamento" json:"senha_equipamento,omitempty"`
	PrefixoIpv6                   int32      `gorm:"column:prefixo_ipv6;comment:Este e o prefixo que o cliente recebe no tunel pppoe e esta vinculado a tabela de IPs" json:"prefixo_ipv6,omitempty"`           // Este e o prefixo que o cliente recebe no tunel pppoe e esta vinculado a tabela de IPs
	PrefixoDelegadoIpv6           int32      `gorm:"column:prefixo_delegado_ipv6;comment:Este e o prefixo que a LAN do cliente distribui e esta ligado a tabela de ips" json:"prefixo_delegado_ipv6,omitempty"` // Este e o prefixo que a LAN do cliente distribui e esta ligado a tabela de ips
	MacAddress                    string     `gorm:"column:mac_address" json:"mac_address,omitempty"`
	CdSetorOrigem                 int32      `gorm:"column:cd_setor_origem;comment:codigo do setor de origen dos produtos para venda/comodato/imobilizdo" json:"cd_setor_origem,omitempty"` // codigo do setor de origen dos produtos para venda/comodato/imobilizdo
	CdComodato                    int32      `gorm:"column:cd_comodato;comment:codigo do comodato lancado pela os" json:"cd_comodato,omitempty"`                                            // codigo do comodato lancado pela os
	CdImobilizado                 int32      `gorm:"column:cd_imobilizado;comment:codigo da imobilizacao que recebeu os produtos da os" json:"cd_imobilizado,omitempty"`                    // codigo da imobilizacao que recebeu os produtos da os
	CdPontoImob                   int32      `gorm:"column:cd_ponto_imob" json:"cd_ponto_imob,omitempty"`
	CodplanoAcesso                int32      `gorm:"column:codplano_acesso" json:"codplano_acesso,omitempty"`
	Codservidor                   int32      `gorm:"column:codservidor" json:"codservidor,omitempty"`
	TempoExecucao                 int32      `gorm:"column:tempo_execucao;comment:- tempo total da execucao da OS contado em minutos" json:"tempo_execucao,omitempty"` // - tempo total da execucao da OS contado em minutos
	LeadsCarregarItens            string     `gorm:"column:leads_carregar_itens;default:N" json:"leads_carregar_itens,omitempty"`
	LeadItensCarregados           string     `gorm:"column:lead_itens_carregados;default:N" json:"lead_itens_carregados,omitempty"`
	VersaoFechamento              int32      `gorm:"column:versao_fechamento" json:"versao_fechamento,omitempty"`
	MovEstoqueSaida               string     `gorm:"column:mov_estoque_saida;default:N" json:"mov_estoque_saida,omitempty"`
	MovEstoqueImobilizado         string     `gorm:"column:mov_estoque_imobilizado;default:N" json:"mov_estoque_imobilizado,omitempty"`
	MovEstoqueComodato            string     `gorm:"column:mov_estoque_comodato;default:N" json:"mov_estoque_comodato,omitempty"`
	CdCancelamento                int32      `gorm:"column:cd_cancelamento" json:"cd_cancelamento,omitempty"`
	MovEstoquePendente            string     `gorm:"column:mov_estoque_pendente;default:N" json:"mov_estoque_pendente,omitempty"`
	MovEstEfetivado               string     `gorm:"column:mov_est_efetivado;default:N" json:"mov_est_efetivado,omitempty"`
	DefeitoConstatadoEncerramento int32      `gorm:"column:defeito_constatado_encerramento" json:"defeito_constatado_encerramento,omitempty"`
	IDCircuito                    string     `gorm:"column:id_circuito" json:"id_circuito,omitempty"`
	Processando                   string     `gorm:"column:processando;default:N" json:"processando,omitempty"`
	Migra                         int32      `gorm:"column:migra" json:"migra,omitempty"`
	ObsMigracao                   string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
	OsRegua                       string     `gorm:"column:os_regua;default:N" json:"os_regua,omitempty"`
	Cancelado                     string     `gorm:"column:cancelado;default:N" json:"cancelado,omitempty"`
	CdMotivoInviabilidade         int32      `gorm:"column:cd_motivo_inviabilidade" json:"cd_motivo_inviabilidade,omitempty"`
	DtHrCancelado                 *time.Time `gorm:"column:dt_hr_cancelado" json:"dt_hr_cancelado,omitempty"`
	DescricaoInviabilidade        string     `gorm:"column:descricao_inviabilidade" json:"descricao_inviabilidade,omitempty"`
	CdOperadorCancelamento        int32      `gorm:"column:cd_operador_cancelamento" json:"cd_operador_cancelamento,omitempty"`
	CriticasAgdAuto               string     `gorm:"column:criticas_agd_auto" json:"criticas_agd_auto,omitempty"`
	NaoAgendada                   string     `gorm:"column:nao_agendada" json:"nao_agendada,omitempty"`
}

// TableName MkO's table name
func (*MkOS) TableName() string {
	return TableNameMkO
}
