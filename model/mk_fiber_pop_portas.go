package model

import (
	"time"
)

const TableNameMkFiberPopPorta = "mk_fiber_pop_portas"

type MkFiberPopPorta struct {
	Codpopporta       int32      `gorm:"column:codpopporta;not null" json:"codpopporta,omitempty"`
	IDPorta           int32      `gorm:"column:id_porta;comment:sequencial de 1 a X para cada POP" json:"id_porta,omitempty"`                                                                                                                                                                 // sequencial de 1 a X para cada POP
	CdPop             int32      `gorm:"column:cd_pop;comment:vinculado ao campo codponto_acesso na tabela mk_pontos_acesso O campo SSID traz o nome identificativo do POP e somente s?o pops de fibra os registro em que o campo tecnologia=C" json:"cd_pop,omitempty"`                      // vinculado ao campo codponto_acesso na tabela mk_pontos_acesso O campo SSID traz o nome identificativo do POP e somente s?o pops de fibra os registro em que o campo tecnologia=C
	CapacidadeCliente int32      `gorm:"column:capacidade_cliente;comment:n?mero identificativo do n?mero de clientes permitidos nessa porta." json:"capacidade_cliente,omitempty"`                                                                                                           // n?mero identificativo do n?mero de clientes permitidos nessa porta.
	CapacidadeMb      float64    `gorm:"column:capacidade_mb;comment:capacidade de tgrafego em MB permitido por porta" json:"capacidade_mb,omitempty"`                                                                                                                                        // capacidade de tgrafego em MB permitido por porta
	LimitarCapacidade string     `gorm:"column:limitar_capacidade;comment:S - para limitar na capacidade N- para n?o limitar a capacidade" json:"limitar_capacidade,omitempty"`                                                                                                               // S - para limitar na capacidade N- para n?o limitar a capacidade
	CorCabo           string     `gorm:"column:cor_cabo;comment:o cabo tem cores diferenciadas usadas pelo provedor. tamb?m podemos colocar essa mesma cor na linha do mapa. imagino que podemos uma tabela aschi para isso assim como ? com a agenda eletronica." json:"cor_cabo,omitempty"` // o cabo tem cores diferenciadas usadas pelo provedor. tamb?m podemos colocar essa mesma cor na linha do mapa. imagino que podemos uma tabela aschi para isso assim como ? com a agenda eletronica.
	IDPorta2          string     `gorm:"column:id_porta2" json:"id_porta2,omitempty"`
	Alias             string     `gorm:"column:alias" json:"alias,omitempty"`
	Vlan              string     `gorm:"column:vlan" json:"vlan,omitempty"`
	CdCabofibra       int32      `gorm:"column:cd_cabofibra" json:"cd_cabofibra,omitempty"`
	CdCabo            int32      `gorm:"column:cd_cabo" json:"cd_cabo,omitempty"`
	CdTuboLoose       int32      `gorm:"column:cd_tubo_loose" json:"cd_tubo_loose,omitempty"`
	PotenciaSinal     float64    `gorm:"column:potencia_sinal" json:"potencia_sinal,omitempty"`
	SinalBom          float64    `gorm:"column:sinal_bom" json:"sinal_bom,omitempty"`
	SinalMedio        float64    `gorm:"column:sinal_medio" json:"sinal_medio,omitempty"`
	HashInsert        string     `gorm:"column:hash_insert" json:"hash_insert,omitempty"`
	Dh                *time.Time `gorm:"column:dh" json:"dh,omitempty"`
	IDOpe             int32      `gorm:"column:id_ope" json:"id_ope,omitempty"`
}

func (*MkFiberPopPorta) TableName() string {
	return TableNameMkFiberPopPorta
}
