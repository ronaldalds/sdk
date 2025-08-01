package model

const TableNameMkPlanosAcessoTipo = "mk_planos_acesso_tipos"

type MkPlanosAcessoTipo struct {
	Codplanoacessotipo int32  `gorm:"column:codplanoacessotipo;not null" json:"codplanoacessotipo,omitempty"`
	Descricao          string `gorm:"column:descricao" json:"descricao,omitempty"`
}

func (*MkPlanosAcessoTipo) TableName() string {
	return TableNameMkPlanosAcessoTipo
}
