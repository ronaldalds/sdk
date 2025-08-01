package model

const TableNameMkFiberSplitterPorta = "mk_fiber_splitter_portas"

type MkFiberSplitterPorta struct {
	Codsplitterporta int32   `gorm:"column:codsplitterporta;not null" json:"codsplitterporta,omitempty"`
	CdSplitter       int32   `gorm:"column:cd_splitter;not null;comment:vinculado ao cadastro de splitter" json:"cd_splitter,omitempty"` // vinculado ao cadastro de splitter
	IDPorta          int32   `gorm:"column:id_porta;comment:sequencial de controle de portas" json:"id_porta,omitempty"`                 // sequencial de controle de portas
	Livre            string  `gorm:"column:livre;comment:S - se est? dispon?vel N - se est? sendo utilizada j?." json:"livre,omitempty"` // S - se est? dispon?vel N - se est? sendo utilizada j?.
	IndicePerda      float64 `gorm:"column:indice_perda" json:"indice_perda,omitempty"`
}

func (*MkFiberSplitterPorta) TableName() string {
	return TableNameMkFiberSplitterPorta
}
