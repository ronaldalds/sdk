package model

import (
	"time"
)

const TableNameMkDocFiscalCompetencium = "mk_doc_fiscal_competencia"

type MkDocFiscalCompetencium struct {
	Coddocfiscalcomp   int32      `gorm:"column:coddocfiscalcomp;not null" json:"coddocfiscalcomp,omitempty"`
	IDOperadorAbertura int32      `gorm:"column:id_operador_abertura" json:"id_operador_abertura,omitempty"`
	IDOperadorFecha    int32      `gorm:"column:id_operador_fecha" json:"id_operador_fecha,omitempty"`
	DhAbertura         *time.Time `gorm:"column:dh_abertura" json:"dh_abertura,omitempty"`
	DhFechamento       *time.Time `gorm:"column:dh_fechamento" json:"dh_fechamento,omitempty"`
	Finalizado         string     `gorm:"column:finalizado;default:N" json:"finalizado,omitempty"`
	Volume             int32      `gorm:"column:volume" json:"volume,omitempty"`
	Serie              string     `gorm:"column:serie" json:"serie,omitempty"`
	CdEmpresa          int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	CdProvedor         int32      `gorm:"column:cd_provedor" json:"cd_provedor,omitempty"`
	TotalNormais       float64    `gorm:"column:total_normais" json:"total_normais,omitempty"`
	TotalCancelados    float64    `gorm:"column:total_cancelados" json:"total_cancelados,omitempty"`
	ContNormais        int32      `gorm:"column:cont_normais" json:"cont_normais,omitempty"`
	ContCanceladas     int32      `gorm:"column:cont_canceladas" json:"cont_canceladas,omitempty"`
	PrimeiroSequencial int32      `gorm:"column:primeiro_sequencial" json:"primeiro_sequencial,omitempty"`
	Mes                int32      `gorm:"column:mes" json:"mes,omitempty"`
	Ano                int32      `gorm:"column:ano" json:"ano,omitempty"`
	MsgRodape          string     `gorm:"column:msg_rodape" json:"msg_rodape,omitempty"`
	UltimoSequencial   int32      `gorm:"column:ultimo_sequencial" json:"ultimo_sequencial,omitempty"`
	Modelo             string     `gorm:"column:modelo" json:"modelo,omitempty"`
	ControleSequencial int32      `gorm:"column:controle_sequencial" json:"controle_sequencial,omitempty"`
	Excluida           string     `gorm:"column:excluida;default:N" json:"excluida,omitempty"`
	IDOperadorExclusao int32      `gorm:"column:id_operador_exclusao" json:"id_operador_exclusao,omitempty"`
	DhExclusao         *time.Time `gorm:"column:dh_exclusao" json:"dh_exclusao,omitempty"`
	FileM              string     `gorm:"column:file_m" json:"file_m,omitempty"`
	FileI              string     `gorm:"column:file_i" json:"file_i,omitempty"`
	FileD              string     `gorm:"column:file_d" json:"file_d,omitempty"`
	ArquivosGerados    string     `gorm:"column:arquivos_gerados;default:N" json:"arquivos_gerados,omitempty"`
	FileMNome          string     `gorm:"column:file_m_nome" json:"file_m_nome,omitempty"`
	FileINome          string     `gorm:"column:file_i_nome" json:"file_i_nome,omitempty"`
	FileDNome          string     `gorm:"column:file_d_nome" json:"file_d_nome,omitempty"`
	CdEmitente         int32      `gorm:"column:cd_emitente" json:"cd_emitente,omitempty"`
	FileZip            string     `gorm:"column:file_zip" json:"file_zip,omitempty"`
	FileZipConteudo    []uint8    `gorm:"column:file_zip_conteudo" json:"file_zip_conteudo,omitempty"`
	VerEscrita         string     `gorm:"column:ver_escrita" json:"ver_escrita,omitempty"`
}

func (*MkDocFiscalCompetencium) TableName() string {
	return TableNameMkDocFiscalCompetencium
}
