package model

import (
	"time"
)

const TableNameMkNotificacaoParada = "mk_notificacao_parada"

type MkNotificacaoParada struct {
	Codnotifparada               int32      `gorm:"column:codnotifparada;not null" json:"codnotifparada,omitempty"`
	DtHrLcto                     *time.Time `gorm:"column:dt_hr_lcto;default:now()" json:"dt_hr_lcto,omitempty"`
	Descricao                    string     `gorm:"column:descricao" json:"descricao,omitempty"`
	Observacoes                  string     `gorm:"column:observacoes" json:"observacoes,omitempty"`
	DtHrParada                   *time.Time `gorm:"column:dt_hr_parada" json:"dt_hr_parada,omitempty"`
	DtHrRetorno                  *time.Time `gorm:"column:dt_hr_retorno" json:"dt_hr_retorno,omitempty"`
	Retornou                     string     `gorm:"column:retornou;default:N" json:"retornou,omitempty"`
	Argumentacao                 string     `gorm:"column:argumentacao" json:"argumentacao,omitempty"`
	OperadorLcto                 int32      `gorm:"column:operador_lcto" json:"operador_lcto,omitempty"`
	OperadorManutencao           int32      `gorm:"column:operador_manutencao" json:"operador_manutencao,omitempty"`
	CdMotivoParada               int32      `gorm:"column:cd_motivo_parada" json:"cd_motivo_parada,omitempty"`
	DtHrPrevRetorno              *time.Time `gorm:"column:dt_hr_prev_retorno" json:"dt_hr_prev_retorno,omitempty"`
	TempoParada                  int32      `gorm:"column:tempo_parada" json:"tempo_parada,omitempty"`
	EquipServidor                int32      `gorm:"column:equip_servidor;default:1" json:"equip_servidor,omitempty"`
	EquipPop                     int32      `gorm:"column:equip_pop;default:2" json:"equip_pop,omitempty"`
	EquipPon                     int32      `gorm:"column:equip_pon;default:3" json:"equip_pon,omitempty"`
	EquipNap                     int32      `gorm:"column:equip_nap;default:4" json:"equip_nap,omitempty"`
	DtHrLctoRetorno              *time.Time `gorm:"column:dt_hr_lcto_retorno" json:"dt_hr_lcto_retorno,omitempty"`
	OperadorLctoRetorno          int32      `gorm:"column:operador_lcto_retorno" json:"operador_lcto_retorno,omitempty"`
	HashEquipamentos             string     `gorm:"column:hash_equipamentos" json:"hash_equipamentos,omitempty"`
	TempoAtraso                  int32      `gorm:"column:tempo_atraso" json:"tempo_atraso,omitempty"`
	ObsEncerramento              string     `gorm:"column:obs_encerramento" json:"obs_encerramento,omitempty"`
	Codmkbotconta                int32      `gorm:"column:codmkbotconta" json:"codmkbotconta,omitempty"`
	Codmkbottemplate             int32      `gorm:"column:codmkbottemplate" json:"codmkbottemplate,omitempty"`
	EnviarWhats                  bool       `gorm:"column:enviar_whats" json:"enviar_whats,omitempty"`
	EnviarWhatsEncerramento      bool       `gorm:"column:enviar_whats_encerramento" json:"enviar_whats_encerramento,omitempty"`
	CodmkbottemplateEncerramento int32      `gorm:"column:codmkbottemplate_encerramento" json:"codmkbottemplate_encerramento,omitempty"`
	CodmkbotcontaEncerramento    int32      `gorm:"column:codmkbotconta_encerramento" json:"codmkbotconta_encerramento,omitempty"`
}

func (*MkNotificacaoParada) TableName() string {
	return TableNameMkNotificacaoParada
}
