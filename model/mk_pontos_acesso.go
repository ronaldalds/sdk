package model

import (
	"time"
)

const TableNameMkPontosAcesso = "mk_pontos_acesso"

type MkPontosAcesso struct {
	Codpontoacesso         int32      `gorm:"column:codpontoacesso;not null" json:"codpontoacesso,omitempty"`
	Ssid                   string     `gorm:"column:ssid;not null;comment:nome de ssid da rede" json:"ssid,omitempty"`                                                       // nome de ssid da rede
	Localizacao            string     `gorm:"column:localizacao;comment:dados para localizacao do ponto" json:"localizacao,omitempty"`                                       // dados para localizacao do ponto
	SistemaOperacional     int32      `gorm:"column:sistema_operacional;not null;comment:sistema operacional do equipamento principal" json:"sistema_operacional,omitempty"` // sistema operacional do equipamento principal
	SoftwareID             string     `gorm:"column:software_id;comment:chave do software instalado" json:"software_id,omitempty"`                                           // chave do software instalado
	BloqueiaMac            string     `gorm:"column:bloqueia_mac;comment:se o ap bloqueia mac na passagem dos clientes" json:"bloqueia_mac,omitempty"`                       // se o ap bloqueia mac na passagem dos clientes
	MacPassagem            string     `gorm:"column:mac_passagem;comment:em caso de bloquear mac na bridge, qual a mac q deve ser lancada." json:"mac_passagem,omitempty"`   // em caso de bloquear mac na bridge, qual a mac q deve ser lancada.
	Funcao                 int32      `gorm:"column:funcao;comment:se tem a funcao de ap ou enlace" json:"funcao,omitempty"`                                                 // se tem a funcao de ap ou enlace
	RemotoUser             string     `gorm:"column:remoto_user;comment:nome de usuario permitido para acesso na porta de ftp" json:"remoto_user,omitempty"`                 // nome de usuario permitido para acesso na porta de ftp
	RemotoPassword         string     `gorm:"column:remoto_password;comment:senha para conceder ao acesso." json:"remoto_password,omitempty"`                                // senha para conceder ao acesso.
	RemotoPortaFtp         string     `gorm:"column:remoto_porta_ftp;comment:porta de acesso ao servico de ftp" json:"remoto_porta_ftp,omitempty"`                           // porta de acesso ao servico de ftp
	NasAddress             string     `gorm:"column:nas_address;comment:ip address NAS para comunicacao com o radius server." json:"nas_address,omitempty"`                  // ip address NAS para comunicacao com o radius server.
	RemotoIPAddress        string     `gorm:"column:remoto_ip_address;comment:endereco de ip para acesso remoto." json:"remoto_ip_address,omitempty"`                        // endereco de ip para acesso remoto.
	IntegraInfo            string     `gorm:"column:integra_info" json:"integra_info,omitempty"`
	RegraPlantao           int32      `gorm:"column:regra_plantao" json:"regra_plantao,omitempty"`
	IntegraConectividade   string     `gorm:"column:integra_conectividade" json:"integra_conectividade,omitempty"`
	IntegraNeighboards     string     `gorm:"column:integra_neighboards" json:"integra_neighboards,omitempty"`
	RemotoPortaSSH         string     `gorm:"column:remoto_porta_ssh" json:"remoto_porta_ssh,omitempty"`
	RemotoUserSSH          string     `gorm:"column:remoto_user_ssh" json:"remoto_user_ssh,omitempty"`
	HabilitaSSH            string     `gorm:"column:habilita_ssh" json:"habilita_ssh,omitempty"`
	AckAceito              int32      `gorm:"column:ack_aceito" json:"ack_aceito,omitempty"`
	AckBom                 int32      `gorm:"column:ack_bom" json:"ack_bom,omitempty"`
	SinalAceito            int32      `gorm:"column:sinal_aceito" json:"sinal_aceito,omitempty"`
	SinalBom               int32      `gorm:"column:sinal_bom" json:"sinal_bom,omitempty"`
	RespondePeloServidor   string     `gorm:"column:responde_pelo_servidor" json:"responde_pelo_servidor,omitempty"`
	IntegraActive          string     `gorm:"column:integra_active" json:"integra_active,omitempty"`
	NasPortID              string     `gorm:"column:nas_port_id" json:"nas_port_id,omitempty"`
	BkpAuto                string     `gorm:"column:bkp_auto" json:"bkp_auto,omitempty"`
	Longitude              string     `gorm:"column:longitude" json:"longitude,omitempty"`
	Latitude               string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Endereco               string     `gorm:"column:endereco" json:"endereco,omitempty"`
	Cidade                 string     `gorm:"column:cidade" json:"cidade,omitempty"`
	Estado                 string     `gorm:"column:estado" json:"estado,omitempty"`
	PortaAPI               int32      `gorm:"column:porta_api" json:"porta_api,omitempty"`
	MkboxConsiderarNas     string     `gorm:"column:mkbox_considerar_nas" json:"mkbox_considerar_nas,omitempty"`
	LimiteConexoes         int32      `gorm:"column:limite_conexoes" json:"limite_conexoes,omitempty"`
	TipoAvisoLimite        string     `gorm:"column:tipo_aviso_limite" json:"tipo_aviso_limite,omitempty"`
	ObterInfoTrafego       string     `gorm:"column:obter_info_trafego" json:"obter_info_trafego,omitempty"`
	ServidorIndicado       int32      `gorm:"column:servidor_indicado" json:"servidor_indicado,omitempty"`
	ObrigarServidor        string     `gorm:"column:obrigar_servidor" json:"obrigar_servidor,omitempty"`
	Tecnologia             string     `gorm:"column:tecnologia" json:"tecnologia,omitempty"`
	Inativo                string     `gorm:"column:inativo" json:"inativo,omitempty"`
	Bairro                 string     `gorm:"column:bairro" json:"bairro,omitempty"`
	Cep                    string     `gorm:"column:cep" json:"cep,omitempty"`
	Polarizacao            string     `gorm:"column:polarizacao" json:"polarizacao,omitempty"`
	Banda                  string     `gorm:"column:banda" json:"banda,omitempty"`
	Frequencia             int32      `gorm:"column:frequencia" json:"frequencia,omitempty"`
	TamanhoCanal           int32      `gorm:"column:tamanho_canal" json:"tamanho_canal,omitempty"`
	DescritivoAbrangência  string     `gorm:"column:descritivo_abrangência" json:"descritivo_abrangência,omitempty"`
	DescritivoEquipamentos string     `gorm:"column:descritivo_equipamentos" json:"descritivo_equipamentos,omitempty"`
	ComportamentoRestRuas  string     `gorm:"column:comportamento_rest_ruas" json:"comportamento_rest_ruas,omitempty"`
	Prioridade             int32      `gorm:"column:prioridade" json:"prioridade,omitempty"`
	APIStatus              string     `gorm:"column:api_status" json:"api_status,omitempty"`
	APIStatusUser          string     `gorm:"column:api_status_user" json:"api_status_user,omitempty"`
	APIStatusDt            *time.Time `gorm:"column:api_status_dt" json:"api_status_dt,omitempty"`
	APIStatusOk            string     `gorm:"column:api_status_ok" json:"api_status_ok,omitempty"`
	MtIdentify             string     `gorm:"column:mt_identify" json:"mt_identify,omitempty"`
	APIStatusAux           string     `gorm:"column:api_status_aux" json:"api_status_aux,omitempty"`
	APIStatusHr            string     `gorm:"column:api_status_hr" json:"api_status_hr,omitempty"`
	SSHStatusOk            string     `gorm:"column:ssh_status_ok" json:"ssh_status_ok,omitempty"`
	SSHUserTest            string     `gorm:"column:ssh_user_test" json:"ssh_user_test,omitempty"`
	SSHPortaTest           int32      `gorm:"column:ssh_porta_test" json:"ssh_porta_test,omitempty"`
	SSHDtTest              *time.Time `gorm:"column:ssh_dt_test" json:"ssh_dt_test,omitempty"`
	SSHHrTest              string     `gorm:"column:ssh_hr_test" json:"ssh_hr_test,omitempty"`
	SSHStatus              string     `gorm:"column:ssh_status" json:"ssh_status,omitempty"`
	IndicarNasAuthLogin    string     `gorm:"column:indicar_nas_auth_login" json:"indicar_nas_auth_login,omitempty"`
	NasAddAuthLogin        string     `gorm:"column:nas_add_auth_login" json:"nas_add_auth_login,omitempty"`
	NasPortAuthLogin       string     `gorm:"column:nas_port_auth_login" json:"nas_port_auth_login,omitempty"`
	CdArmario              int32      `gorm:"column:cd_armario" json:"cd_armario,omitempty"`
	Projeto                string     `gorm:"column:projeto" json:"projeto,omitempty"`
	CdPopTipo              int32      `gorm:"column:cd_pop_tipo" json:"cd_pop_tipo,omitempty"`
	CdProjeto              int32      `gorm:"column:cd_projeto" json:"cd_projeto,omitempty"`
	SnmpPorta              int32      `gorm:"column:snmp_porta" json:"snmp_porta,omitempty"`
	SnmpComunidade         string     `gorm:"column:snmp_comunidade" json:"snmp_comunidade,omitempty"`
	MonitoraCPU            string     `gorm:"column:monitora_cpu" json:"monitora_cpu,omitempty"`
	MonitoraInterface      string     `gorm:"column:monitora_interface" json:"monitora_interface,omitempty"`
	MonitoraMemoria        string     `gorm:"column:monitora_memoria" json:"monitora_memoria,omitempty"`
	MonitoraSinalOnu       string     `gorm:"column:monitora_sinal_onu" json:"monitora_sinal_onu,omitempty"`
	SnmpTstPorta           int32      `gorm:"column:snmp_tst_porta" json:"snmp_tst_porta,omitempty"`
	SnmpTstUser            string     `gorm:"column:snmp_tst_user" json:"snmp_tst_user,omitempty"`
	SnmpTstStatus          string     `gorm:"column:snmp_tst_status" json:"snmp_tst_status,omitempty"`
	SnmpTstData            *time.Time `gorm:"column:snmp_tst_data" json:"snmp_tst_data,omitempty"`
	SnmpTstHora            string     `gorm:"column:snmp_tst_hora" json:"snmp_tst_hora,omitempty"`
	MonitoraQueue          string     `gorm:"column:monitora_queue" json:"monitora_queue,omitempty"`
	TipoIPSugerido         int32      `gorm:"column:tipo_ip_sugerido" json:"tipo_ip_sugerido,omitempty"`
	PoolSugerido           int32      `gorm:"column:pool_sugerido" json:"pool_sugerido,omitempty"`
	IPAddressAnm           string     `gorm:"column:ip_address_anm" json:"ip_address_anm,omitempty"`
	PortaTelnet            int32      `gorm:"column:porta_telnet" json:"porta_telnet,omitempty"`
	CdEmpresa              int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	Associacao             int32      `gorm:"column:associacao;not null;comment:1= POP, 2=Armario" json:"associacao,omitempty"` // 1= POP, 2=Armario
	CdPop                  int32      `gorm:"column:cd_pop" json:"cd_pop,omitempty"`
	PotenciaSinal          float64    `gorm:"column:potencia_sinal" json:"potencia_sinal,omitempty"`
	Obs                    string     `gorm:"column:obs" json:"obs,omitempty"`
	CdPerfilMonitoramento  int32      `gorm:"column:cd_perfil_monitoramento" json:"cd_perfil_monitoramento,omitempty"`
	SnmpVersao             string     `gorm:"column:snmp_versao" json:"snmp_versao,omitempty"`
	CdPoste                int32      `gorm:"column:cd_poste" json:"cd_poste,omitempty"`
	CdCaixa                int32      `gorm:"column:cd_caixa" json:"cd_caixa,omitempty"`
	CdStatus               int32      `gorm:"column:cd_status" json:"cd_status,omitempty"`
	Dh                     *time.Time `gorm:"column:dh" json:"dh,omitempty"`
	IDOpe                  int32      `gorm:"column:id_ope" json:"id_ope,omitempty"`
	Descricao              string     `gorm:"column:descricao" json:"descricao,omitempty"`
	CdProjetoV3            int32      `gorm:"column:cd_projeto_v3" json:"cd_projeto_v3,omitempty"`
	CdPropriedade          int32      `gorm:"column:cd_propriedade" json:"cd_propriedade,omitempty"`
	HashInsert             string     `gorm:"column:hash_insert" json:"hash_insert,omitempty"`
	TipoPop                int32      `gorm:"column:tipo_pop;default:1" json:"tipo_pop,omitempty"`
	Valor                  float64    `gorm:"column:valor" json:"valor,omitempty"`
}

func (*MkPontosAcesso) TableName() string {
	return TableNameMkPontosAcesso
}
