package model

import (
	"time"
)

const TableNameMkFiberPoste = "mk_fiber_postes"

type MkFiberPoste struct {
	Codposte                int32      `gorm:"column:codposte;not null" json:"codposte,omitempty"`
	Localizacao             string     `gorm:"column:localizacao" json:"localizacao,omitempty"`
	Latitude                string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude               string     `gorm:"column:longitude" json:"longitude,omitempty"`
	CdUf                    int32      `gorm:"column:cd_uf" json:"cd_uf,omitempty"`
	CdCidade                int32      `gorm:"column:cd_cidade" json:"cd_cidade,omitempty"`
	CdBairro                int32      `gorm:"column:cd_bairro" json:"cd_bairro,omitempty"`
	CdLogradouro            int32      `gorm:"column:cd_logradouro" json:"cd_logradouro,omitempty"`
	Numero                  int32      `gorm:"column:numero" json:"numero,omitempty"`
	Complemento             string     `gorm:"column:complemento" json:"complemento,omitempty"`
	TipoPoste               string     `gorm:"column:tipo_poste;comment:C- Concreto M - Madeira F - Ferro" json:"tipo_poste,omitempty"` // C- Concreto M - Madeira F - Ferro
	Identificacao           string     `gorm:"column:identificacao" json:"identificacao,omitempty"`
	CdPostetiipo            int32      `gorm:"column:cd_postetiipo" json:"cd_postetiipo,omitempty"`
	Projeto                 string     `gorm:"column:projeto" json:"projeto,omitempty"`
	DataCriacao             *time.Time `gorm:"column:data_criacao" json:"data_criacao,omitempty"`
	HoraCriacao             string     `gorm:"column:hora_criacao" json:"hora_criacao,omitempty"`
	CdProjeto               int32      `gorm:"column:cd_projeto" json:"cd_projeto,omitempty"`
	ElementoMapaJSON        string     `gorm:"column:elemento_mapa_json" json:"elemento_mapa_json,omitempty"`
	OutrasInformacoes       string     `gorm:"column:outras_informacoes" json:"outras_informacoes,omitempty"`
	CdProprietario          int32      `gorm:"column:cd_proprietario" json:"cd_proprietario,omitempty"`
	OutrasInformacoesJSON   string     `gorm:"column:outras_informacoes_json" json:"outras_informacoes_json,omitempty"`
	DataHoraAtualizacao     *time.Time `gorm:"column:data_hora_atualizacao" json:"data_hora_atualizacao,omitempty"`
	CaixaSubterranea        string     `gorm:"column:caixa_subterranea" json:"caixa_subterranea,omitempty"`
	ElementoMapaJSONMobile  string     `gorm:"column:elemento_mapa_json_mobile" json:"elemento_mapa_json_mobile,omitempty"`
	LatitudeN               float64    `gorm:"column:latitude_n" json:"latitude_n,omitempty"`
	LongitudeN              float64    `gorm:"column:longitude_n" json:"longitude_n,omitempty"`
	RetornoConsultaEndereco string     `gorm:"column:retorno_consulta_endereco" json:"retorno_consulta_endereco,omitempty"`
	TipoElemento            int32      `gorm:"column:tipo_elemento;comment:1 - Poste  2 - Marcadores  3 - Armarios  4 - Pops Wireles  5 - Conexoes  6 - Marcadores  7 - Outos equipamentos" json:"tipo_elemento,omitempty"` // 1 - Poste  2 - Marcadores  3 - Armarios  4 - Pops Wireles  5 - Conexoes  6 - Marcadores  7 - Outos equipamentos
	CdConexao               int32      `gorm:"column:cd_conexao" json:"cd_conexao,omitempty"`
	CdPontoAcesso           int32      `gorm:"column:cd_ponto_acesso" json:"cd_ponto_acesso,omitempty"`
	TipoConexao             string     `gorm:"column:tipo_conexao" json:"tipo_conexao,omitempty"`
	URLIcone                string     `gorm:"column:url_icone" json:"url_icone,omitempty"`
	CdArmario               int32      `gorm:"column:cd_armario" json:"cd_armario,omitempty"`
	CdMarcador              int32      `gorm:"column:cd_marcador" json:"cd_marcador,omitempty"`
	CdUsuario               int32      `gorm:"column:cd_usuario" json:"cd_usuario,omitempty"`
	CdPopInfo               int32      `gorm:"column:cd_pop_info" json:"cd_pop_info,omitempty"`
}

func (*MkFiberPoste) TableName() string {
	return TableNameMkFiberPoste
}
