package model

import (
	"time"
)

const TableNameFrLogEvent = "fr_log_event"

type FrLogEvent struct {
	LogData      time.Time `gorm:"column:log_data" json:"log_data,omitempty"`
	LogHora      string    `gorm:"column:log_hora" json:"log_hora,omitempty"`
	LogCodform   int32     `gorm:"column:log_codform" json:"log_codform,omitempty"`
	LogDescform  string    `gorm:"column:log_descform" json:"log_descform,omitempty"`
	LogOperacao  string    `gorm:"column:log_operacao" json:"log_operacao,omitempty"`
	LogUsuario   string    `gorm:"column:log_usuario" json:"log_usuario,omitempty"`
	LogSistema   string    `gorm:"column:log_sistema" json:"log_sistema,omitempty"`
	LogChave     string    `gorm:"column:log_chave" json:"log_chave,omitempty"`
	LogChavecont string    `gorm:"column:log_chavecont" json:"log_chavecont,omitempty"`
	LogConteudo  string    `gorm:"column:log_conteudo" json:"log_conteudo,omitempty"`
	LogID        int32     `gorm:"column:log_id;not null" json:"log_id,omitempty"`
	LogIP        string    `gorm:"column:log_ip" json:"log_ip,omitempty"`
}

func (*FrLogEvent) TableName() string {
	return TableNameFrLogEvent
}
