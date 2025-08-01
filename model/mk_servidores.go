package model

import (
	"time"
)

const TableNameMkServidore = "mk_servidores"

type MkServidore struct {
	Codmovimento       int32      `gorm:"column:codmovimento;not null" json:"codmovimento,omitempty"`
	Descricao          string     `gorm:"column:descricao;not null" json:"descricao,omitempty"`
	Sistemaoperacional int32      `gorm:"column:sistemaoperacional;not null;comment:sistema operacional do servidor" json:"sistemaoperacional,omitempty"` // sistema operacional do servidor
	Localizacao        string     `gorm:"column:localizacao;comment:aonde esta localizado esse servidor" json:"localizacao,omitempty"`                    // aonde esta localizado esse servidor
	Codprofile         int32      `gorm:"column:codprofile;not null;comment:a qual profile obedecem as configuracoes" json:"codprofile,omitempty"`        // a qual profile obedecem as configuracoes
	SoftwareID         string     `gorm:"column:software_id;comment:mikrotik serial" json:"software_id,omitempty"`                                        // mikrotik serial
	IPComunicacao      string     `gorm:"column:ip_comunicacao;not null;comment:numero de ip de comunicacao" json:"ip_comunicacao,omitempty"`             // numero de ip de comunicacao
	AcessoUser         string     `gorm:"column:acesso_user;comment:nome de usuario permitido para acessar ao sistema" json:"acesso_user,omitempty"`      // nome de usuario permitido para acessar ao sistema
	AcessoPassword     string     `gorm:"column:acesso_password;comment:senha para acesso ftp ao servidor" json:"acesso_password,omitempty"`              // senha para acesso ftp ao servidor
	PortaFtp           string     `gorm:"column:porta_ftp;comment:numero da porta ftp configurada" json:"porta_ftp,omitempty"`                            // numero da porta ftp configurada
	HabilitaRd01       string     `gorm:"column:habilita_rd01;comment:se habilita o radius secundario 01" json:"habilita_rd01,omitempty"`                 // se habilita o radius secundario 01
	HabilitaRd02       string     `gorm:"column:habilita_rd02" json:"habilita_rd02,omitempty"`
	IntegraRadiusLocal string     `gorm:"column:integra_radius_local;comment:se vai integrar-se a tabela radius do banco de dados local." json:"integra_radius_local,omitempty"` // se vai integrar-se a tabela radius do banco de dados local.
	NasAddress         string     `gorm:"column:nas_address;comment:ip address de comunicacao com o radius server." json:"nas_address,omitempty"`                                // ip address de comunicacao com o radius server.
	Radius01           int32      `gorm:"column:radius01;comment:indicacao de servidor radius secundario." json:"radius01,omitempty"`                                            // indicacao de servidor radius secundario.
	Radius02           int32      `gorm:"column:radius02;comment:indicacao de servidor secundario de radius." json:"radius02,omitempty"`                                         // indicacao de servidor secundario de radius.
	IntegraInfo        string     `gorm:"column:integra_info" json:"integra_info,omitempty"`
	RegraPlantao       int32      `gorm:"column:regra_plantao" json:"regra_plantao,omitempty"`
	IntegraActive      string     `gorm:"column:integra_active" json:"integra_active,omitempty"`
	IntegraNeighboards string     `gorm:"column:integra_neighboards" json:"integra_neighboards,omitempty"`
	RemotoPortaSSH     string     `gorm:"column:remoto_porta_ssh" json:"remoto_porta_ssh,omitempty"`
	RemotoUserSSH      string     `gorm:"column:remoto_user_ssh" json:"remoto_user_ssh,omitempty"`
	HabilitaSSH        string     `gorm:"column:habilita_ssh" json:"habilita_ssh,omitempty"`
	BkpAuto            string     `gorm:"column:bkp_auto" json:"bkp_auto,omitempty"`
	PortaAPI           int32      `gorm:"column:porta_api" json:"porta_api,omitempty"`
	TempoConexao       float64    `gorm:"column:tempo_conexao" json:"tempo_conexao,omitempty"`
	ObterInfoTrafego   string     `gorm:"column:obter_info_trafego" json:"obter_info_trafego,omitempty"`
	APIStatus          string     `gorm:"column:api_status" json:"api_status,omitempty"`
	APIStatusUser      string     `gorm:"column:api_status_user" json:"api_status_user,omitempty"`
	APIStatusDt        *time.Time `gorm:"column:api_status_dt" json:"api_status_dt,omitempty"`
	APIStatusOk        string     `gorm:"column:api_status_ok" json:"api_status_ok,omitempty"`
	MtIdentify         string     `gorm:"column:mt_identify" json:"mt_identify,omitempty"`
	APIStatusAux       string     `gorm:"column:api_status_aux" json:"api_status_aux,omitempty"`
	APIStatusHr        string     `gorm:"column:api_status_hr" json:"api_status_hr,omitempty"`
	SSHStatusOk        string     `gorm:"column:ssh_status_ok" json:"ssh_status_ok,omitempty"`
	SSHUserTest        string     `gorm:"column:ssh_user_test" json:"ssh_user_test,omitempty"`
	SSHPortaTest       int32      `gorm:"column:ssh_porta_test" json:"ssh_porta_test,omitempty"`
	SSHDtTest          *time.Time `gorm:"column:ssh_dt_test" json:"ssh_dt_test,omitempty"`
	SSHHrTest          string     `gorm:"column:ssh_hr_test" json:"ssh_hr_test,omitempty"`
	SSHStatus          string     `gorm:"column:ssh_status" json:"ssh_status,omitempty"`
	NasPortID          string     `gorm:"column:nas_port_id" json:"nas_port_id,omitempty"`
	MonitoraCPU        string     `gorm:"column:monitora_cpu" json:"monitora_cpu,omitempty"`
	MonitoraMemoria    string     `gorm:"column:monitora_memoria" json:"monitora_memoria,omitempty"`
	SnmpPorta          int32      `gorm:"column:snmp_porta" json:"snmp_porta,omitempty"`
	SnmpComunidade     string     `gorm:"column:snmp_comunidade" json:"snmp_comunidade,omitempty"`
	SnmpTstHora        string     `gorm:"column:snmp_tst_hora" json:"snmp_tst_hora,omitempty"`
	SnmpTstData        *time.Time `gorm:"column:snmp_tst_data" json:"snmp_tst_data,omitempty"`
	SnmpTstStatus      string     `gorm:"column:snmp_tst_status" json:"snmp_tst_status,omitempty"`
	SnmpTstUser        string     `gorm:"column:snmp_tst_user" json:"snmp_tst_user,omitempty"`
	SnmpTstPorta       int32      `gorm:"column:snmp_tst_porta" json:"snmp_tst_porta,omitempty"`
	MonitoraQueue      string     `gorm:"column:monitora_queue" json:"monitora_queue,omitempty"`
	PortaRadclient     int32      `gorm:"column:porta_radclient" json:"porta_radclient,omitempty"`
	CdCluster          int32      `gorm:"column:cd_cluster" json:"cd_cluster,omitempty"`
	CdEmpresa          int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	CdPerfilCheckReply int32      `gorm:"column:cd_perfil_check_reply" json:"cd_perfil_check_reply,omitempty"`
	Codpopinfo         int32      `gorm:"column:codpopinfo" json:"codpopinfo,omitempty"`
}

func (*MkServidore) TableName() string {
	return TableNameMkServidore
}
