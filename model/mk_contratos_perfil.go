package model

import (
	"time"
)

const TableNameMkContratosPerfil = "mk_contratos_perfil"

type MkContratosPerfil struct {
	Codcontratoperfil int32      `gorm:"column:codcontratoperfil;not null" json:"codcontratoperfil,omitempty"`
	DescricaoPerfil   string     `gorm:"column:descricao_perfil;not null" json:"descricao_perfil,omitempty"`
	TextoContrato     string     `gorm:"column:texto_contrato" json:"texto_contrato,omitempty"`
	DataCriacao       *time.Time `gorm:"column:data_criacao" json:"data_criacao,omitempty"`
	MensagemCabecalho string     `gorm:"column:mensagem_cabecalho" json:"mensagem_cabecalho,omitempty"`
	MensagemRodape    string     `gorm:"column:mensagem_rodape" json:"mensagem_rodape,omitempty"`
	MensagemConclusao string     `gorm:"column:mensagem_conclusao" json:"mensagem_conclusao,omitempty"`
	PathWar           string     `gorm:"column:path_war" json:"path_war,omitempty"`
	MensagemSaudacao  string     `gorm:"column:mensagem_saudacao" json:"mensagem_saudacao,omitempty"`
	AdesaoTitulo      string     `gorm:"column:adesao_titulo" json:"adesao_titulo,omitempty"`
	AdesaoTexto       string     `gorm:"column:adesao_texto" json:"adesao_texto,omitempty"`
	SvaTexto          string     `gorm:"column:sva_texto" json:"sva_texto,omitempty"`
	RescisaoTexto     string     `gorm:"column:rescisao_texto" json:"rescisao_texto,omitempty"`
}

func (*MkContratosPerfil) TableName() string {
	return TableNameMkContratosPerfil
}
