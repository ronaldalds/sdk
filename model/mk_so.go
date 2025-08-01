package model

const TableNameMkSo = "mk_so"

type MkSo struct {
	Codso                  int32  `gorm:"column:codso;not null" json:"codso,omitempty"`
	Descricao              string `gorm:"column:descricao" json:"descricao,omitempty"`
	Concentrador           string `gorm:"column:concentrador" json:"concentrador,omitempty"`
	PontoAcesso            string `gorm:"column:ponto_acesso" json:"ponto_acesso,omitempty"`
	CdscriptPesquisaSerial int32  `gorm:"column:cdscript_pesquisa_serial" json:"cdscript_pesquisa_serial,omitempty"`
	Ftth                   string `gorm:"column:ftth;default:N" json:"ftth,omitempty"`
	FormatarSinal          string `gorm:"column:formatar_sinal;default:N" json:"formatar_sinal,omitempty"`
	CalculoSinal           string `gorm:"column:calculo_sinal" json:"calculo_sinal,omitempty"`
}

func (*MkSo) TableName() string {
	return TableNameMkSo
}
