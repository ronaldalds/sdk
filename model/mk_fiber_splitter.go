package model

const TableNameMkFiberSplitter = "mk_fiber_splitter"

type MkFiberSplitter struct {
	Codsplitter     int32  `gorm:"column:codsplitter;not null" json:"codsplitter,omitempty"`
	Latitude        string `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude       string `gorm:"column:longitude" json:"longitude,omitempty"`
	CdUf            int32  `gorm:"column:cd_uf" json:"cd_uf,omitempty"`
	CdCidade        int32  `gorm:"column:cd_cidade" json:"cd_cidade,omitempty"`
	CdBairro        int32  `gorm:"column:cd_bairro" json:"cd_bairro,omitempty"`
	CdLogradouro    int32  `gorm:"column:cd_logradouro" json:"cd_logradouro,omitempty"`
	Numero          int32  `gorm:"column:numero" json:"numero,omitempty"`
	Complemento     string `gorm:"column:complemento" json:"complemento,omitempty"`
	DescricaoModelo string `gorm:"column:descricao_modelo" json:"descricao_modelo,omitempty"`
	Portas          int32  `gorm:"column:portas;comment:n?mero de portas" json:"portas,omitempty"`                                                                                                                                                               // n?mero de portas
	IDSplitter      string `gorm:"column:id_splitter;comment:numero sequencial da caixa seguido de uma letra, inicialmente A. Se uma segunda caixa for instalada no mesmo poste, ela deve ter o mesmo n?mero por?m com a letra B." json:"id_splitter,omitempty"` // numero sequencial da caixa seguido de uma letra, inicialmente A. Se uma segunda caixa for instalada no mesmo poste, ela deve ter o mesmo n?mero por?m com a letra B.
	CdPop           int32  `gorm:"column:cd_pop;comment:codigo do ponto de acesso que est? ligada a caixa" json:"cd_pop,omitempty"`                                                                                                                              // codigo do ponto de acesso que est? ligada a caixa
	CdPopPorta      int32  `gorm:"column:cd_pop_porta;comment:numero da porta do pop onde esta conectada a caixa. (splitter)" json:"cd_pop_porta,omitempty"`                                                                                                     // numero da porta do pop onde esta conectada a caixa. (splitter)
	CdPoste         int32  `gorm:"column:cd_poste" json:"cd_poste,omitempty"`
	CdCaixa         int32  `gorm:"column:cd_caixa" json:"cd_caixa,omitempty"`
	Codtiposplitter int32  `gorm:"column:codtiposplitter" json:"codtiposplitter,omitempty"`
	CdPontoacesso   int32  `gorm:"column:cd_pontoacesso" json:"cd_pontoacesso,omitempty"`
	CdSplitter      int32  `gorm:"column:cd_splitter" json:"cd_splitter,omitempty"`
	CdPortaSplitter int32  `gorm:"column:cd_porta_splitter" json:"cd_porta_splitter,omitempty"`
}

func (*MkFiberSplitter) TableName() string {
	return TableNameMkFiberSplitter
}
