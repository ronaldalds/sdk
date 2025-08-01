package model

const TableNameMkContasFaturada = "mk_contas_faturadas"

type MkContasFaturada struct {
	// Relationsships Init
	PlanoConta []MkPlanoConta `gorm:"foreignKey:codconta;references:cd_conta" json:"plano_conta,omitempty"`
	// Relationsships End
	Codcontafaturada int32 `gorm:"column:codcontafaturada;not null" json:"codcontafaturada,omitempty"`
	CdFatura         int32 `gorm:"column:cd_fatura;not null" json:"cd_fatura,omitempty"`
	CdConta          int32 `gorm:"column:cd_conta;not null" json:"cd_conta,omitempty"`
}

func (*MkContasFaturada) TableName() string {
	return TableNameMkContasFaturada
}
