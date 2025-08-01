package model

const TableNameMkPontosOffMotivo = "mk_pontos_off_motivos"

type MkPontosOffMotivo struct {
	Codmotivo                  int32  `gorm:"column:codmotivo;not null" json:"codmotivo,omitempty"`
	Descricao                  string `gorm:"column:descricao" json:"descricao,omitempty"`
	TipoParada                 string `gorm:"column:tipo_parada" json:"tipo_parada,omitempty"`
	GerarIndiceErroOperacional string `gorm:"column:gerar_indice_erro_operacional" json:"gerar_indice_erro_operacional,omitempty"`
	GerarIndicePaneEletrica    string `gorm:"column:gerar_indice_pane_eletrica" json:"gerar_indice_pane_eletrica,omitempty"`
	GerarIndiceProbHardware    string `gorm:"column:gerar_indice_prob_hardware" json:"gerar_indice_prob_hardware,omitempty"`
	GerarIndiceProbSoftware    string `gorm:"column:gerar_indice_prob_software" json:"gerar_indice_prob_software,omitempty"`
	GerarIndiceProbOperadora   string `gorm:"column:gerar_indice_prob_operadora" json:"gerar_indice_prob_operadora,omitempty"`
	GerarIndiceProbChuvas      string `gorm:"column:gerar_indice_prob_chuvas" json:"gerar_indice_prob_chuvas,omitempty"`
	GerarIndiceRompimentoFibra string `gorm:"column:gerar_indice_rompimento_fibra" json:"gerar_indice_rompimento_fibra,omitempty"`
	GerarIndiceProbTransporte  string `gorm:"column:gerar_indice_prob_transporte" json:"gerar_indice_prob_transporte,omitempty"`
	GerarIndiceOutros          string `gorm:"column:gerar_indice_outros" json:"gerar_indice_outros,omitempty"`
}

func (*MkPontosOffMotivo) TableName() string {
	return TableNameMkPontosOffMotivo
}
