package model

const TableNameMkRevenda = "mk_revendas"

type MkRevenda struct {
	Codrevenda            int32  `gorm:"column:codrevenda;not null" json:"codrevenda,omitempty"`
	NomeRevenda           string `gorm:"column:nome_revenda" json:"nome_revenda,omitempty"`
	Responsavel           string `gorm:"column:responsavel" json:"responsavel,omitempty"`
	FoneCelular           string `gorm:"column:fone_celular" json:"fone_celular,omitempty"`
	Fone01                string `gorm:"column:fone_01" json:"fone_01,omitempty"`
	Fone02                string `gorm:"column:fone_02" json:"fone_02,omitempty"`
	EmailOperacional      string `gorm:"column:email_operacional" json:"email_operacional,omitempty"`
	EmailFinanceiro       string `gorm:"column:email_financeiro" json:"email_financeiro,omitempty"`
	EmailAdministrativo   string `gorm:"column:email_administrativo" json:"email_administrativo,omitempty"`
	Uf                    int32  `gorm:"column:uf" json:"uf,omitempty"`
	Cidade                int32  `gorm:"column:cidade" json:"cidade,omitempty"`
	Bairro                int32  `gorm:"column:bairro" json:"bairro,omitempty"`
	Logradouro            int32  `gorm:"column:logradouro" json:"logradouro,omitempty"`
	Num                   string `gorm:"column:num" json:"num,omitempty"`
	Complemento           string `gorm:"column:complemento" json:"complemento,omitempty"`
	Site                  string `gorm:"column:site" json:"site,omitempty"`
	RegraComissaoSugerida int32  `gorm:"column:regra_comissao_sugerida" json:"regra_comissao_sugerida,omitempty"`
	DescritivoAbrangencia string `gorm:"column:descritivo_abrangencia" json:"descritivo_abrangencia,omitempty"`
	Cnpj                  string `gorm:"column:cnpj" json:"cnpj,omitempty"`
	Cep                   string `gorm:"column:cep" json:"cep,omitempty"`
	PontoRecebimento      int32  `gorm:"column:ponto_recebimento" json:"ponto_recebimento,omitempty"`
	Ie                    string `gorm:"column:ie" json:"ie,omitempty"`
	CoringaExtra01        string `gorm:"column:coringa_extra_01" json:"coringa_extra_01,omitempty"`
	CoringaExtra02        string `gorm:"column:coringa_extra_02" json:"coringa_extra_02,omitempty"`
	CoringaExtra03        string `gorm:"column:coringa_extra_03" json:"coringa_extra_03,omitempty"`
	CoringaExtra04        string `gorm:"column:coringa_extra_04" json:"coringa_extra_04,omitempty"`
	CoringaExtra05        string `gorm:"column:coringa_extra_05" json:"coringa_extra_05,omitempty"`
	UsarInfoParaContratos string `gorm:"column:usar_info_para_contratos" json:"usar_info_para_contratos,omitempty"`
	TmpMensagem1          string `gorm:"column:tmp_mensagem1" json:"tmp_mensagem1,omitempty"`
	TmpMensagem2          string `gorm:"column:tmp_mensagem2" json:"tmp_mensagem2,omitempty"`
}

func (*MkRevenda) TableName() string {
	return TableNameMkRevenda
}
