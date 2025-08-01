package model

const TableNameMkNotificacaoParadaEquipamento = "mk_notificacao_parada_equipamentos"

type MkNotificacaoParadaEquipamento struct {
	Codnotifparadaequip int32 `gorm:"column:codnotifparadaequip;not null" json:"codnotifparadaequip,omitempty"`
	CdParada            int32 `gorm:"column:cd_parada" json:"cd_parada,omitempty"`
	CdEquipamento       int32 `gorm:"column:cd_equipamento" json:"cd_equipamento,omitempty"`
	TipoEquipamento     int32 `gorm:"column:tipo_equipamento;comment:1 - concentrador 2 - POP/OLT 3- PON 4- NAP" json:"tipo_equipamento,omitempty"` // 1 - concentrador 2 - POP/OLT 3- PON 4- NAP
	PortaPon            int32 `gorm:"column:porta_pon" json:"porta_pon,omitempty"`
}

func (*MkNotificacaoParadaEquipamento) TableName() string {
	return TableNameMkNotificacaoParadaEquipamento
}
