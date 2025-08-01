package model

import (
	"time"
)

const TableNameFrUsuario = "fr_usuario"

type FrUsuario struct {
	UsrCodigo               int32      `gorm:"column:usr_codigo;not null" json:"usr_codigo,omitempty"`
	UsrLogin                string     `gorm:"column:usr_login;not null" json:"usr_login,omitempty"`
	UsrSenha                string     `gorm:"column:usr_senha" json:"usr_senha,omitempty"`
	UsrAdministrador        string     `gorm:"column:usr_administrador;not null;default:N" json:"usr_administrador,omitempty"`
	UsrTipoExpiracao        string     `gorm:"column:usr_tipo_expiracao;not null;default:N" json:"usr_tipo_expiracao,omitempty"`
	UsrDiasExpiracao        int32      `gorm:"column:usr_dias_expiracao" json:"usr_dias_expiracao,omitempty"`
	UsrImagemDigital        []uint8    `gorm:"column:usr_imagem_digital" json:"usr_imagem_digital,omitempty"`
	UsrFoto                 []uint8    `gorm:"column:usr_foto" json:"usr_foto,omitempty"`
	UsrNome                 string     `gorm:"column:usr_nome;not null" json:"usr_nome,omitempty"`
	UsrEmail                string     `gorm:"column:usr_email" json:"usr_email,omitempty"`
	UsrDigital              int64      `gorm:"column:usr_digital" json:"usr_digital,omitempty"`
	UsrInicioExpiracao      *time.Time `gorm:"column:usr_inicio_expiracao" json:"usr_inicio_expiracao,omitempty"`
	FoneCelular             string     `gorm:"column:fone_celular" json:"fone_celular,omitempty"`
	IDWhatsapp              string     `gorm:"column:id_whatsapp" json:"id_whatsapp,omitempty"`
	IDTelegram              string     `gorm:"column:id_telegram" json:"id_telegram,omitempty"`
	SetorAssociado          string     `gorm:"column:setor_associado" json:"setor_associado,omitempty"`
	CdPerfilAcesso          int32      `gorm:"column:cd_perfil_acesso" json:"cd_perfil_acesso,omitempty"`
	StrHostsAcesso          string     `gorm:"column:str_hosts_acesso" json:"str_hosts_acesso,omitempty"`
	DtAtuCadastro           *time.Time `gorm:"column:dt_atu_cadastro" json:"dt_atu_cadastro,omitempty"`
	TokenAcesso             string     `gorm:"column:token_acesso" json:"token_acesso,omitempty"`
	TokenGcm                string     `gorm:"column:token_gcm" json:"token_gcm,omitempty"`
	UsaMobOs                string     `gorm:"column:usa_mob_os" json:"usa_mob_os,omitempty"`
	TokenGcmManager         string     `gorm:"column:token_gcm_manager" json:"token_gcm_manager,omitempty"`
	TokenGcmMaps            string     `gorm:"column:token_gcm_maps" json:"token_gcm_maps,omitempty"`
	OcultarPainelOs         string     `gorm:"column:ocultar_painel_os" json:"ocultar_painel_os,omitempty"`
	TokenGcmCrm             string     `gorm:"column:token_gcm_crm" json:"token_gcm_crm,omitempty"`
	PermiteLogarTodos       string     `gorm:"column:permite_logar_todos" json:"permite_logar_todos,omitempty"`
	ChatUsuario             string     `gorm:"column:chat_usuario" json:"chat_usuario,omitempty"`
	ChatSenha               string     `gorm:"column:chat_senha" json:"chat_senha,omitempty"`
	UsrSac                  string     `gorm:"column:usr_sac;default:N" json:"usr_sac,omitempty"`
	CdRevenda               int32      `gorm:"column:cd_revenda" json:"cd_revenda,omitempty"`
	MkbotXmppSenha          string     `gorm:"column:mkbot_xmpp_senha" json:"mkbot_xmpp_senha,omitempty"`
	MkbotXmppUsuario        string     `gorm:"column:mkbot_xmpp_usuario" json:"mkbot_xmpp_usuario,omitempty"`
	MkbotChaveLiberacao     string     `gorm:"column:mkbot_chave_liberacao" json:"mkbot_chave_liberacao,omitempty"`
	ObrigInternetAppAgentes string     `gorm:"column:obrig_internet_app_agentes;default:N" json:"obrig_internet_app_agentes,omitempty"`
	MapsUltLocalizacao      string     `gorm:"column:maps_ult_localizacao" json:"maps_ult_localizacao,omitempty"`
	DataExpiracaoToken      *time.Time `gorm:"column:data_expiracao_token" json:"data_expiracao_token,omitempty"`
	MkbotAttGrade           bool       `gorm:"column:mkbot_att_grade;default:true" json:"mkbot_att_grade,omitempty"`
	UIDFirebase             string     `gorm:"column:uid_firebase" json:"uid_firebase,omitempty"`
	TokenJwtFirebase        string     `gorm:"column:token_jwt_firebase" json:"token_jwt_firebase,omitempty"`
}

func (*FrUsuario) TableName() string {
	return TableNameFrUsuario
}
