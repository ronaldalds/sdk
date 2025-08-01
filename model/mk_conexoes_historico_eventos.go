package model

import (
	"time"
)

const TableNameMkConexoesHistoricoEvento = "mk_conexoes_historico_eventos"

type MkConexoesHistoricoEvento struct {
	Codconhistevento   int32      `gorm:"column:codconhistevento;not null" json:"codconhistevento,omitempty"`
	CdConexao          int32      `gorm:"column:cd_conexao" json:"cd_conexao,omitempty"`
	HashControle       string     `gorm:"column:hash_controle" json:"hash_controle,omitempty"`
	TipoEvento         int32      `gorm:"column:tipo_evento;comment:em conexoes_operacoes" json:"tipo_evento,omitempty"` // em conexoes_operacoes
	Data               *time.Time `gorm:"column:data" json:"data,omitempty"`
	Hora               string     `gorm:"column:hora" json:"hora,omitempty"`
	Operador           string     `gorm:"column:operador" json:"operador,omitempty"`
	TipoDebito         int32      `gorm:"column:tipo_debito;comment:1-  conta 2 - fatura" json:"tipo_debito,omitempty"` // 1-  conta 2 - fatura
	DtVctoDebito       *time.Time `gorm:"column:dt_vcto_debito" json:"dt_vcto_debito,omitempty"`
	ValorDebito        float64    `gorm:"column:valor_debito" json:"valor_debito,omitempty"`
	DiasVencidos       int32      `gorm:"column:dias_vencidos" json:"dias_vencidos,omitempty"`
	Efetivado          string     `gorm:"column:efetivado;default:N" json:"efetivado,omitempty"`
	CdRegua            int32      `gorm:"column:cd_regua" json:"cd_regua,omitempty"`
	DescricaoExtra     string     `gorm:"column:descricao_extra" json:"descricao_extra,omitempty"`
	DtLimiteEvento     *time.Time `gorm:"column:dt_limite_evento" json:"dt_limite_evento,omitempty"`
	CdDebito           int32      `gorm:"column:cd_debito" json:"cd_debito,omitempty"`
	Dh                 *time.Time `gorm:"column:dh;default:now()" json:"dh,omitempty"`
	OperadorAuditoria  int32      `gorm:"column:operador_auditoria" json:"operador_auditoria,omitempty"`
	DhAuditoria        *time.Time `gorm:"column:dh_auditoria" json:"dh_auditoria,omitempty"`
	ObsAuditoria       string     `gorm:"column:obs_auditoria" json:"obs_auditoria,omitempty"`
	AuditoriaRealizada string     `gorm:"column:auditoria_realizada;default:N" json:"auditoria_realizada,omitempty"`
}

func (*MkConexoesHistoricoEvento) TableName() string {
	return TableNameMkConexoesHistoricoEvento
}
