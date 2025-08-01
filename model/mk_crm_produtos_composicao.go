package model

const TableNameMkCrmProdutosComposicao = "mk_crm_produtos_composicao"

type MkCrmProdutosComposicao struct {
	Codcrmprodcompo         int32   `gorm:"column:codcrmprodcompo;not null" json:"codcrmprodcompo,omitempty"`
	CdProduto               int32   `gorm:"column:cd_produto" json:"cd_produto,omitempty"`
	CdPlano                 int32   `gorm:"column:cd_plano" json:"cd_plano,omitempty"`
	VlrVenda                float64 `gorm:"column:vlr_venda" json:"vlr_venda,omitempty"`
	NumConexoes             int32   `gorm:"column:num_conexoes" json:"num_conexoes,omitempty"`
	VlrDesconto             float64 `gorm:"column:vlr_desconto" json:"vlr_desconto,omitempty"`
	DescricaoDesconto       string  `gorm:"column:descricao_desconto" json:"descricao_desconto,omitempty"`
	BloqueiaDescontoProduto string  `gorm:"column:bloqueia_desconto_produto" json:"bloqueia_desconto_produto,omitempty"`
}

func (*MkCrmProdutosComposicao) TableName() string {
	return TableNameMkCrmProdutosComposicao
}
