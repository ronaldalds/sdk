package model

import (
	"time"
)

const TableNameMkConexo = "mk_conexoes"

type MkConexao struct {
	// Relationsships Init
	Contrato                  MkContrato                  `gorm:"foreignKey:codcontrato;references:contrato" json:"contrato,omitempty"`
	Pessoa                    MkPessoa                    `gorm:"foreignKey:codpessoa;references:codcliente;" json:"pessoa,omitempty"`
	NotificacaoParadaConexoes []MkNotificacaoParadaConexo `gorm:"foreignKey:cd_conexao;references:codconexao" json:"notificacao_parada_conexoes,omitempty"`
	OS                        []MkOS                      `gorm:"foreignKey:conexao_associada;references:codconexao" json:"os,omitempty"`
	FiberSplitter             MkFiberSplitter             `gorm:"foreignKey:codsplitter;references:cd_splitter" json:"fiber_splitter,omitempty"`
	FiberSplitterPorta        MkFiberSplitterPorta        `gorm:"foreignKey:codsplitterporta;references:id_porta_splitter" json:"fiber_splitter_porta,omitempty"`
	Logradouro                MkLogradouro                `gorm:"foreignKey:codlogradouro;references:logradouro" json:"logradouro,omitempty"`
	PontosAcesso              MkPontosAcesso              `gorm:"foreignKey:codpontoacesso;references:codponto_acesso" json:"pontos_acesso,omitempty"`
	Servidore                 MkServidore                 `gorm:"foreignKey:codmovimento;references:codservidor" json:"servidore,omitempty"`
	TabelaIPs                 MkTabelaIP                  `gorm:"foreignKey:codip;references:codendereco_ip" json:"tabela_ips,omitempty"`
	// Relationsships End
	Codconexao                 int32      `gorm:"column:codconexao;not null" json:"codconexao,omitempty"`
	Codcliente                 int32      `gorm:"column:codcliente;not null" json:"codcliente,omitempty"`
	CodContrato                int32      `gorm:"column:contrato;comment:numero do contrato associado." json:"codcontrato,omitempty"` // numero do contrato associado.
	CodEstado                  int32      `gorm:"column:uf" json:"uf,omitempty"`
	IDPortaSplitter            int32      `gorm:"column:id_porta_splitter" json:"id_porta_splitter,omitempty"`
	CdSplitter                 int32      `gorm:"column:cd_splitter" json:"cd_splitter,omitempty"`
	CodpontoAcesso             int32      `gorm:"column:codponto_acesso;comment:codigo do ponto de acesso" json:"codponto_acesso,omitempty"` // codigo do ponto de acesso
	CodLogradouro              int32      `gorm:"column:logradouro" json:"codlogradouro,omitempty"`
	Codservidor                int32      `gorm:"column:codservidor;not null;comment:codigo do servidor vinculado" json:"codservidor,omitempty"` // codigo do servidor vinculado
	CodenderecoIP              int32      `gorm:"column:codendereco_ip;comment:codigo do endereco de ip" json:"codendereco_ip,omitempty"`        // codigo do endereco de ip
	CodCidade                  int32      `gorm:"column:cidade" json:"cidade,omitempty"`
	CodBairro                  int32      `gorm:"column:bairro" json:"bairro,omitempty"`
	TipoConexao                int32      `gorm:"column:tipo_conexao;not null;comment:cobranca, cortesia, demonstracao, POP, suporte" json:"tipo_conexao,omitempty"` // cobranca, cortesia, demonstracao, POP, suporte
	Autenticacao               int32      `gorm:"column:autenticacao;not null;comment:nenhum, hotspot ou pppoe" json:"autenticacao,omitempty"`                       // nenhum, hotspot ou pppoe
	Username                   string     `gorm:"column:username" json:"username,omitempty"`
	Password                   string     `gorm:"column:password" json:"password,omitempty"`
	Instalador                 int32      `gorm:"column:instalador;comment:pessoa responsavel pela instalacao." json:"instalador,omitempty"`                    // pessoa responsavel pela instalacao.
	Cadastrado                 *time.Time `gorm:"column:cadastrado;not null;comment:data e hora da habilitacao" json:"cadastrado,omitempty"`                    // data e hora da habilitacao
	AdvertiseURL               string     `gorm:"column:advertise_url;comment:popup" json:"advertise_url,omitempty"`                                            // popup
	AdvertiseTime              int32      `gorm:"column:advertise_time;comment:tempo de replicacao do advertise" json:"advertise_time,omitempty"`               // tempo de replicacao do advertise
	SiteInicial                string     `gorm:"column:site_inicial;comment:pg inicial para o cliente" json:"site_inicial,omitempty"`                          // pg inicial para o cliente
	CodplanoAcesso             int32      `gorm:"column:codplano_acesso;not null;comment:codigo do plano de acesso escolhido" json:"codplano_acesso,omitempty"` // codigo do plano de acesso escolhido
	MacAddress                 string     `gorm:"column:mac_address;not null;default:00:00:00:00:00:00" json:"mac_address,omitempty"`
	MacAddressAux              string     `gorm:"column:mac_address_aux" json:"mac_address_aux,omitempty"`
	MacAddressConsiderado      string     `gorm:"column:mac_address_considerado;not null;default:00:00:00:00:00:00" json:"mac_address_considerado,omitempty"`
	ConexaoBloqueada           string     `gorm:"column:conexao_bloqueada;default:N;comment:se a conexao esta ou naum bloqueada." json:"conexao_bloqueada,omitempty"`                                             // se a conexao esta ou naum bloqueada.
	PathMonitorGrafico         string     `gorm:"column:path_monitor_grafico;comment:path do monitor grafico do mikrotik" json:"path_monitor_grafico,omitempty"`                                                  // path do monitor grafico do mikrotik
	URLInadimplencia           string     `gorm:"column:url_inadimplencia;comment:caso ele esteja entrando em inadimplencia end da url de direcionamento inicial." json:"url_inadimplencia,omitempty"`            // caso ele esteja entrando em inadimplencia end da url de direcionamento inicial.
	PontoAcesso                int32      `gorm:"column:ponto_acesso;comment:em caso da conexao ser do tipo ponto de acesso, aqui ira ser informado qual eh esse ponto de acesso." json:"ponto_acesso,omitempty"` // em caso da conexao ser do tipo ponto de acesso, aqui ira ser informado qual eh esse ponto de acesso.
	MotivoBloqueio             int32      `gorm:"column:motivo_bloqueio;comment:motivo pelo qual a conexao esta bloqueada" json:"motivo_bloqueio,omitempty"`                                                      // motivo pelo qual a conexao esta bloqueada
	Descricao                  string     `gorm:"column:descricao" json:"descricao,omitempty"`
	ExpiraEm                   *time.Time `gorm:"column:expira_em;comment:caso seje uma conexao de demonstracao, a dt de expiracao dessa conexao." json:"expira_em,omitempty"`     // caso seje uma conexao de demonstracao, a dt de expiracao dessa conexao.
	VelocidadesFormatadas      string     `gorm:"column:velocidades_formatadas;comment:relacao das velocidades formatadas para o radius." json:"velocidades_formatadas,omitempty"` // relacao das velocidades formatadas para o radius.
	SucessoInt01               string     `gorm:"column:sucesso_int_01;comment:se houve sucesso na integracao do radius local." json:"sucesso_int_01,omitempty"`                   // se houve sucesso na integracao do radius local.
	SucessoInt02               string     `gorm:"column:sucesso_int_02;comment:se houve sucesso na integracao do radius server slave 01" json:"sucesso_int_02,omitempty"`          // se houve sucesso na integracao do radius server slave 01
	SucessoInt03               string     `gorm:"column:sucesso_int_03;comment:se houve sucesso na integracao do radius server slave 02" json:"sucesso_int_03,omitempty"`          // se houve sucesso na integracao do radius server slave 02
	Wpa2Key                    string     `gorm:"column:wpa2_key" json:"wpa2_key,omitempty"`
	SenhaEquipamento           string     `gorm:"column:senha_equipamento" json:"senha_equipamento,omitempty"`
	RangeRota                  int32      `gorm:"column:range_rota;comment:em caso de utilizar roteamento, qual a rota associada." json:"range_rota,omitempty"`                // em caso de utilizar roteamento, qual a rota associada.
	CadastrarPortal            string     `gorm:"column:cadastrar_portal;comment:se deseja cadastrar esta conexao para acesso ao portal MK" json:"cadastrar_portal,omitempty"` // se deseja cadastrar esta conexao para acesso ao portal MK
	EmExcecao                  string     `gorm:"column:em_excecao;default:N" json:"em_excecao,omitempty"`
	MotivoExcecao              string     `gorm:"column:motivo_excecao" json:"motivo_excecao,omitempty"`
	PrazoExcecao               *time.Time `gorm:"column:prazo_excecao" json:"prazo_excecao,omitempty"`
	TmpExibicao                string     `gorm:"column:tmp_exibicao" json:"tmp_exibicao,omitempty"`
	TmpSeriaBloqueado          *time.Time `gorm:"column:tmp_seria_bloqueado" json:"tmp_seria_bloqueado,omitempty"`
	PermiteAcessoSac           string     `gorm:"column:permite_acesso_sac" json:"permite_acesso_sac,omitempty"`
	LicenseAp                  string     `gorm:"column:license_ap" json:"license_ap,omitempty"`
	LicenseServer              string     `gorm:"column:license_server" json:"license_server,omitempty"`
	QueueMbIn                  float64    `gorm:"column:queue_mb_in" json:"queue_mb_in,omitempty"`
	QueueMbOut                 float64    `gorm:"column:queue_mb_out" json:"queue_mb_out,omitempty"`
	QueuePacketsIn             float64    `gorm:"column:queue_packets_in" json:"queue_packets_in,omitempty"`
	QueuePacketsOut            float64    `gorm:"column:queue_packets_out" json:"queue_packets_out,omitempty"`
	QueueData                  *time.Time `gorm:"column:queue_data" json:"queue_data,omitempty"`
	QueueHora                  string     `gorm:"column:queue_hora" json:"queue_hora,omitempty"`
	QueueEncerrada             string     `gorm:"column:queue_encerrada" json:"queue_encerrada,omitempty"`
	WireDesconectado           string     `gorm:"column:wire_desconectado" json:"wire_desconectado,omitempty"`
	WireID                     string     `gorm:"column:wire_id" json:"wire_id,omitempty"`
	WireTxRate                 float64    `gorm:"column:wire_tx_rate" json:"wire_tx_rate,omitempty"`
	WireRxRate                 float64    `gorm:"column:wire_rx_rate" json:"wire_rx_rate,omitempty"`
	WireSignal                 string     `gorm:"column:wire_signal" json:"wire_signal,omitempty"`
	WireTxPackets              float64    `gorm:"column:wire_tx_packets" json:"wire_tx_packets,omitempty"`
	WireRxPackets              float64    `gorm:"column:wire_rx_packets" json:"wire_rx_packets,omitempty"`
	WireDownload               float64    `gorm:"column:wire_download" json:"wire_download,omitempty"`
	WireUpload                 float64    `gorm:"column:wire_upload" json:"wire_upload,omitempty"`
	DataLogin                  *time.Time `gorm:"column:data_login" json:"data_login,omitempty"`
	HoraLogin                  string     `gorm:"column:hora_login" json:"hora_login,omitempty"`
	LoginEncerrado             string     `gorm:"column:login_encerrado" json:"login_encerrado,omitempty"`
	IPAddServer                string     `gorm:"column:ip_add_server" json:"ip_add_server,omitempty"`
	IPAddPop                   string     `gorm:"column:ip_add_pop" json:"ip_add_pop,omitempty"`
	ActiveRevisado             string     `gorm:"column:active_revisado" json:"active_revisado,omitempty"`
	DataLoginEnc               *time.Time `gorm:"column:data_login_enc" json:"data_login_enc,omitempty"`
	HoraLoginEnc               string     `gorm:"column:hora_login_enc" json:"hora_login_enc,omitempty"`
	QueueRevisado              string     `gorm:"column:queue_revisado" json:"queue_revisado,omitempty"`
	WireRevisado               string     `gorm:"column:wire_revisado" json:"wire_revisado,omitempty"`
	QueueControl               string     `gorm:"column:queue_control" json:"queue_control,omitempty"`
	RotaIP                     int32      `gorm:"column:rota_ip" json:"rota_ip,omitempty"`
	RotaAtiva                  string     `gorm:"column:rota_ativa" json:"rota_ativa,omitempty"`
	TipoIP                     int32      `gorm:"column:tipo_ip" json:"tipo_ip,omitempty"`
	SeraBloqueado              string     `gorm:"column:sera_bloqueado;default:N" json:"sera_bloqueado,omitempty"`
	PoolAssociado              int32      `gorm:"column:pool_associado" json:"pool_associado,omitempty"`
	LacreEquipamento           string     `gorm:"column:lacre_equipamento" json:"lacre_equipamento,omitempty"`
	AddressList                int32      `gorm:"column:address_list" json:"address_list,omitempty"`
	FranquiaAtiva              string     `gorm:"column:franquia_ativa;default:N" json:"franquia_ativa,omitempty"`
	CdAssinante                string     `gorm:"column:cd_assinante" json:"cd_assinante,omitempty"`
	TipoCriptografia           int32      `gorm:"column:tipo_criptografia" json:"tipo_criptografia,omitempty"`
	Latitude                   string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude                  string     `gorm:"column:longitude" json:"longitude,omitempty"`
	AutoDesbloqueio            string     `gorm:"column:auto_desbloqueio;default:S" json:"auto_desbloqueio,omitempty"`
	PreCadastro                int32      `gorm:"column:pre_cadastro" json:"pre_cadastro,omitempty"`
	URLInadimplenciaDias       int32      `gorm:"column:url_inadimplencia_dias" json:"url_inadimplencia_dias,omitempty"`
	GUIDEstatisticas           string     `gorm:"column:guid_estatisticas" json:"guid_estatisticas,omitempty"`
	Condominio                 int32      `gorm:"column:condominio" json:"condominio,omitempty"`
	Nasipaddress               string     `gorm:"column:nasipaddress" json:"nasipaddress,omitempty"`
	Framedipaddress            string     `gorm:"column:framedipaddress" json:"framedipaddress,omitempty"`
	InputDia                   int64      `gorm:"column:input_dia" json:"input_dia,omitempty"`
	OutputDia                  int64      `gorm:"column:output_dia" json:"output_dia,omitempty"`
	InputMes                   int64      `gorm:"column:input_mes" json:"input_mes,omitempty"`
	OutputMes                  int64      `gorm:"column:output_mes" json:"output_mes,omitempty"`
	EmVelocidadeTemporaria     string     `gorm:"column:em_velocidade_temporaria;default:N" json:"em_velocidade_temporaria,omitempty"`
	VelocidadeTemporaria       string     `gorm:"column:velocidade_temporaria" json:"velocidade_temporaria,omitempty"`
	VelTempEntrada             string     `gorm:"column:vel_temp_entrada" json:"vel_temp_entrada,omitempty"`
	VelTempSaida               string     `gorm:"column:vel_temp_saida" json:"vel_temp_saida,omitempty"`
	UltimoBloqueioAuto         *time.Time `gorm:"column:ultimo_bloqueio_auto" json:"ultimo_bloqueio_auto,omitempty"`
	VencimentoDupBloqueio      *time.Time `gorm:"column:vencimento_dup_bloqueio" json:"vencimento_dup_bloqueio,omitempty"`
	VlrDuplicataBloqueio       float64    `gorm:"column:vlr_duplicata_bloqueio" json:"vlr_duplicata_bloqueio,omitempty"`
	DiaVencimentoBloqueio      int32      `gorm:"column:dia_vencimento_bloqueio" json:"dia_vencimento_bloqueio,omitempty"`
	LimitacaoUser              string     `gorm:"column:limitacao_user" json:"limitacao_user,omitempty"`
	LimitacaoDt                *time.Time `gorm:"column:limitacao_dt" json:"limitacao_dt,omitempty"`
	LimitacaoHr                string     `gorm:"column:limitacao_hr" json:"limitacao_hr,omitempty"`
	Acctsessionid              string     `gorm:"column:acctsessionid" json:"acctsessionid,omitempty"`
	FranquiaExcecao            string     `gorm:"column:franquia_excecao" json:"franquia_excecao,omitempty"`
	FranquiaDtExcecao          *time.Time `gorm:"column:franquia_dt_excecao" json:"franquia_dt_excecao,omitempty"`
	FranquiaTotalExcecao       float64    `gorm:"column:franquia_total_excecao" json:"franquia_total_excecao,omitempty"`
	FranquiaTotalExcecaoMedida string     `gorm:"column:franquia_total_excecao_medida" json:"franquia_total_excecao_medida,omitempty"`
	FranquiaExcecaoVitalicia   string     `gorm:"column:franquia_excecao_vitalicia" json:"franquia_excecao_vitalicia,omitempty"`
	EmInadimplencia            string     `gorm:"column:em_inadimplencia;default:N" json:"em_inadimplencia,omitempty"`
	EquipamentoTx              string     `gorm:"column:equipamento_tx" json:"equipamento_tx,omitempty"`
	EquipamentoModo            string     `gorm:"column:equipamento_modo" json:"equipamento_modo,omitempty"`
	EquipamentoIPLan           string     `gorm:"column:equipamento_ip_lan" json:"equipamento_ip_lan,omitempty"`
	EquipamentoPortaHTTP       string     `gorm:"column:equipamento_porta_http" json:"equipamento_porta_http,omitempty"`
	EquipamentoPortaSSH        string     `gorm:"column:equipamento_porta_ssh" json:"equipamento_porta_ssh,omitempty"`
	EquipamentoInfoDirect      string     `gorm:"column:equipamento_info_direct" json:"equipamento_info_direct,omitempty"`
	TmpSenhaMd5                string     `gorm:"column:tmp_senha_md5" json:"tmp_senha_md5,omitempty"`
	IgnoraBloqueio             string     `gorm:"column:ignora_bloqueio" json:"ignora_bloqueio,omitempty"`
	PortaDsl                   int32      `gorm:"column:porta_dsl" json:"porta_dsl,omitempty"`
	PortaDg                    int32      `gorm:"column:porta_dg" json:"porta_dg,omitempty"`
	PortaTelefone              int32      `gorm:"column:porta_telefone" json:"porta_telefone,omitempty"`
	Cep                        string     `gorm:"column:cep" json:"cep,omitempty"`
	Numero                     int32      `gorm:"column:numero" json:"numero,omitempty"`
	Complemento                string     `gorm:"column:complemento" json:"complemento,omitempty"`
	ProblemaInfoConcentrador   string     `gorm:"column:problema_info_concentrador" json:"problema_info_concentrador,omitempty"`
	IPAtualConcentrador        string     `gorm:"column:ip_atual_concentrador" json:"ip_atual_concentrador,omitempty"`
	UltimoDesbloqueio          *time.Time `gorm:"column:ultimo_desbloqueio" json:"ultimo_desbloqueio,omitempty"`
	UserEap                    string     `gorm:"column:user_eap" json:"user_eap,omitempty"`
	Mtu                        string     `gorm:"column:mtu" json:"mtu,omitempty"`
	PoolMk                     string     `gorm:"column:pool_mk" json:"pool_mk,omitempty"`
	GUIDTmp                    string     `gorm:"column:guid_tmp" json:"guid_tmp,omitempty"`
	Nasportidname              string     `gorm:"column:nasportidname" json:"nasportidname,omitempty"`
	CdIPGerencia               int32      `gorm:"column:cd_ip_gerencia" json:"cd_ip_gerencia,omitempty"`
	OnuSerial                  string     `gorm:"column:onu_serial" json:"onu_serial,omitempty"`
	Tecnologia                 string     `gorm:"column:tecnologia" json:"tecnologia,omitempty"`
	CdPortaCliente             int32      `gorm:"column:cd_porta_cliente" json:"cd_porta_cliente,omitempty"`
	FiberMetrosUsados          float64    `gorm:"column:fiber_metros_usados" json:"fiber_metros_usados,omitempty"`
	FiberTipoCabo              int32      `gorm:"column:fiber_tipo_cabo" json:"fiber_tipo_cabo,omitempty"`
	InterfaceOlt               string     `gorm:"column:interface_olt" json:"interface_olt,omitempty"`
	PortaCondominio            int32      `gorm:"column:porta_condominio" json:"porta_condominio,omitempty"`
	CdEndIPEquipamento         int32      `gorm:"column:cd_end_ip_equipamento" json:"cd_end_ip_equipamento,omitempty"`
	HotspotBypassed            string     `gorm:"column:hotspot_bypassed" json:"hotspot_bypassed,omitempty"`
	Obs                        string     `gorm:"column:obs" json:"obs,omitempty"`
	HashWebfilterMara          string     `gorm:"column:hash_webfilter_mara" json:"hash_webfilter_mara,omitempty"`
	WebFilter                  string     `gorm:"column:web_filter" json:"web_filter,omitempty"`
	CacheRota                  string     `gorm:"column:cache_rota" json:"cache_rota,omitempty"`
	SeraReduzido               string     `gorm:"column:sera_reduzido" json:"sera_reduzido,omitempty"`
	UltimaReducao              *time.Time `gorm:"column:ultima_reducao" json:"ultima_reducao,omitempty"`
	EstaReduzida               string     `gorm:"column:esta_reduzida;default:N" json:"esta_reduzida,omitempty"`
	ContaGerouReducao          int32      `gorm:"column:conta_gerou_reducao" json:"conta_gerou_reducao,omitempty"`
	HtHrCriacao                *time.Time `gorm:"column:ht_hr_criacao" json:"ht_hr_criacao,omitempty"`
	IDNetmaps                  string     `gorm:"column:id_netmaps" json:"id_netmaps,omitempty"`
	CdEmpresa                  int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	ClassificacaoAcesso        string     `gorm:"column:classificacao_acesso" json:"classificacao_acesso,omitempty"`
	ModeloMta                  string     `gorm:"column:modelo_mta" json:"modelo_mta,omitempty"`
	MacMta                     string     `gorm:"column:mac_mta" json:"mac_mta,omitempty"`
	MacCpe                     string     `gorm:"column:mac_cpe" json:"mac_cpe,omitempty"`
	PerfilProvCmts             int32      `gorm:"column:perfil_prov_cmts" json:"perfil_prov_cmts,omitempty"`
	TipoProvCmts               string     `gorm:"column:tipo_prov_cmts" json:"tipo_prov_cmts,omitempty"`
	DataHoraCriacao            *time.Time `gorm:"column:data_hora_criacao" json:"data_hora_criacao,omitempty"`
	CdCircuitoDesignador       string     `gorm:"column:cd_circuito_designador" json:"cd_circuito_designador,omitempty"`
	EndIgualCadastro           string     `gorm:"column:end_igual_cadastro" json:"end_igual_cadastro,omitempty"`
	CdBlocoIpv6                int32      `gorm:"column:cd_bloco_ipv6;comment:" json:"cd_bloco_ipv6,omitempty"`
	SeraInadimplente           string     `gorm:"column:sera_inadimplente" json:"sera_inadimplente,omitempty"`
	EstaInadimplente           string     `gorm:"column:esta_inadimplente;default:N" json:"esta_inadimplente,omitempty"`
	DtUltimaInadimplencia      *time.Time `gorm:"column:dt_ultima_inadimplencia" json:"dt_ultima_inadimplencia,omitempty"`
	CdContaFatInadimplencia    int32      `gorm:"column:cd_conta_fat_inadimplencia" json:"cd_conta_fat_inadimplencia,omitempty"`
	HashControleInsercao       string     `gorm:"column:hash_controle_insercao" json:"hash_controle_insercao,omitempty"`
	UsarIpv6                   string     `gorm:"column:usar_ipv6" json:"usar_ipv6,omitempty"`
	PrefixoIpv6                int32      `gorm:"column:prefixo_ipv6;comment:Este e o prefixo que o cliente recebe no tunel pppoe e esta vinculado a tabela de IPs" json:"prefixo_ipv6,omitempty"`              // Este e o prefixo que o cliente recebe no tunel pppoe e esta vinculado a tabela de IPs
	PrefixoDelegadoIpv6        int32      `gorm:"column:prefixo_delegado_ipv6;comment:Este e o prefixo que a LAN do cliente distribui e esta ligado a tabela de ips" json:"prefixo_delegado_ipv6,omitempty"`    // Este e o prefixo que a LAN do cliente distribui e esta ligado a tabela de ips
	RangeDelegadoIpv6          int32      `gorm:"column:range_delegado_ipv6;comment:Este e o range de prefixo delegado que o cliente esta vinculado na tabela de ranges." json:"range_delegado_ipv6,omitempty"` // Este e o range de prefixo delegado que o cliente esta vinculado na tabela de ranges.
	RangePrefixoIpv6           int32      `gorm:"column:range_prefixo_ipv6;comment:Este e o range de prefixo que esta associado a tabela de ranges." json:"range_prefixo_ipv6,omitempty"`                       // Este e o range de prefixo que esta associado a tabela de ranges.
	NivelSinalFtth             int32      `gorm:"column:nivel_sinal_ftth" json:"nivel_sinal_ftth,omitempty"`
	NivelSinalFtthAnterior     int32      `gorm:"column:nivel_sinal_ftth_anterior" json:"nivel_sinal_ftth_anterior,omitempty"`
	PortaHTTP                  int32      `gorm:"column:porta_http" json:"porta_http,omitempty"`
	NivelSinalFtthReal         float64    `gorm:"column:nivel_sinal_ftth_real" json:"nivel_sinal_ftth_real,omitempty"`
	NivelSinalFtthAnteriorReal float64    `gorm:"column:nivel_sinal_ftth_anterior_real" json:"nivel_sinal_ftth_anterior_real,omitempty"`
	UltimaAtualizacaoSinal     *time.Time `gorm:"column:ultima_atualizacao_sinal" json:"ultima_atualizacao_sinal,omitempty"`
	CdPlanoAutorizado          int32      `gorm:"column:cd_plano_autorizado" json:"cd_plano_autorizado,omitempty"`
	DtLimiteAutorizaPlano      *time.Time `gorm:"column:dt_limite_autoriza_plano" json:"dt_limite_autoriza_plano,omitempty"`
	CdLead                     int32      `gorm:"column:cd_lead" json:"cd_lead,omitempty"`
	Framedipv6prefix           string     `gorm:"column:framedipv6prefix" json:"framedipv6prefix,omitempty"`
	CdAteRefAbertura           int32      `gorm:"column:cd_ate_ref_abertura" json:"cd_ate_ref_abertura,omitempty"`
	SinalRxInicial             float64    `gorm:"column:sinal_rx_inicial" json:"sinal_rx_inicial,omitempty"`
	TemperaturaInicial         int32      `gorm:"column:temperatura_inicial" json:"temperatura_inicial,omitempty"`
	DistanciaInicial           int32      `gorm:"column:distancia_inicial" json:"distancia_inicial,omitempty"`
	SinalTxInicial             float64    `gorm:"column:sinal_tx_inicial" json:"sinal_tx_inicial,omitempty"`
	CdCaixa                    int32      `gorm:"column:cd_caixa" json:"cd_caixa,omitempty"`
	TxCcq                      string     `gorm:"column:tx_ccq" json:"tx_ccq,omitempty"`
	CdPorta                    int32      `gorm:"column:cd_porta" json:"cd_porta,omitempty"`
	CdPortaPon                 int32      `gorm:"column:cd_porta_pon" json:"cd_porta_pon,omitempty"`
	CodplanoAcesso1            int32      `gorm:"column:codplano_acesso1" json:"codplano_acesso1,omitempty"`
	Zona                       string     `gorm:"column:zona;default:U" json:"zona,omitempty"`
	Migra                      int32      `gorm:"column:migra" json:"migra,omitempty"`
	IDCircuito                 string     `gorm:"column:id_circuito" json:"id_circuito,omitempty"`
	Tags                       string     `gorm:"column:tags" json:"tags,omitempty"`
	ObsMigracao                string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
}

func (*MkConexao) TableName() string {
	return TableNameMkConexo
}
