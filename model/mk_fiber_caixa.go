package model

import (
	"time"
)

const TableNameMkFiberCaixa = "mk_fiber_caixa"

type MkFiberCaixa struct {
	Codcaixa                   int32      `gorm:"column:codcaixa;not null" json:"codcaixa,omitempty"`
	Identificacao              string     `gorm:"column:identificacao;not null" json:"identificacao,omitempty"`
	Tipo                       int32      `gorm:"column:tipo" json:"tipo,omitempty"`
	CdPoste                    int32      `gorm:"column:cd_poste" json:"cd_poste,omitempty"`
	IndicePerda                float64    `gorm:"column:indice_perda" json:"indice_perda,omitempty"`
	QtdePortas                 float64    `gorm:"column:qtde_portas" json:"qtde_portas,omitempty"`
	Latitude                   string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude                  string     `gorm:"column:longitude" json:"longitude,omitempty"`
	Projeto                    string     `gorm:"column:projeto" json:"projeto,omitempty"`
	DataInstalacao             *time.Time `gorm:"column:data_instalacao" json:"data_instalacao,omitempty"`
	HoraInstalacao             string     `gorm:"column:hora_instalacao" json:"hora_instalacao,omitempty"`
	ValorEstimado              float64    `gorm:"column:valor_estimado" json:"valor_estimado,omitempty"`
	CdProjeto                  int32      `gorm:"column:cd_projeto" json:"cd_projeto,omitempty"`
	Codtipocaixa               int32      `gorm:"column:codtipocaixa" json:"codtipocaixa,omitempty"`
	NumeroInicial              int32      `gorm:"column:numero_inicial" json:"numero_inicial,omitempty"`
	NumeroFinal                int32      `gorm:"column:numero_final" json:"numero_final,omitempty"`
	AutoAnuncio                string     `gorm:"column:auto_anuncio" json:"auto_anuncio,omitempty"`
	LimitePortasAnuncio        int32      `gorm:"column:limite_portas_anuncio" json:"limite_portas_anuncio,omitempty"`
	RaioAtendimento            float64    `gorm:"column:raio_atendimento" json:"raio_atendimento,omitempty"`
	CdArmario                  int32      `gorm:"column:cd_armario" json:"cd_armario,omitempty"`
	Desbalanceado              string     `gorm:"column:desbalanceado" json:"desbalanceado,omitempty"`
	RegistroKey                string     `gorm:"column:registro_key" json:"registro_key,omitempty"`
	Diagrama                   string     `gorm:"column:diagrama" json:"diagrama,omitempty"`
	Descricao                  string     `gorm:"column:descricao" json:"descricao,omitempty"`
	DiagramaHTML               string     `gorm:"column:diagrama_html" json:"diagrama_html,omitempty"`
	Prospectado                string     `gorm:"column:prospectado" json:"prospectado,omitempty"`
	QrCode                     []uint8    `gorm:"column:qr_code" json:"qr_code,omitempty"`
	JSONLigacoes               string     `gorm:"column:json_ligacoes" json:"json_ligacoes,omitempty"`
	GUIDCaixa                  string     `gorm:"column:guid_caixa" json:"guid_caixa,omitempty"`
	VersaoDiagramaMk           int32      `gorm:"column:versao_diagrama_mk" json:"versao_diagrama_mk,omitempty"`
	LocalDiagrama              string     `gorm:"column:local_diagrama" json:"local_diagrama,omitempty"`
	OutrasEspecificacoes       string     `gorm:"column:outras_especificacoes" json:"outras_especificacoes,omitempty"`
	AutoProvisionamento        string     `gorm:"column:auto_provisionamento" json:"auto_provisionamento,omitempty"`
	CdSistemaOperacional       int32      `gorm:"column:cd_sistema_operacional" json:"cd_sistema_operacional,omitempty"`
	Rascunho                   string     `gorm:"column:rascunho" json:"rascunho,omitempty"`
	CdPontoAcesso              int32      `gorm:"column:cd_ponto_acesso" json:"cd_ponto_acesso,omitempty"`
	CdPortaPop                 int32      `gorm:"column:cd_porta_pop" json:"cd_porta_pop,omitempty"`
	NivelSinalFtth             int32      `gorm:"column:nivel_sinal_ftth" json:"nivel_sinal_ftth,omitempty"`
	NivelSinalFtthAnterior     int32      `gorm:"column:nivel_sinal_ftth_anterior" json:"nivel_sinal_ftth_anterior,omitempty"`
	NivelSinalFtthReal         float64    `gorm:"column:nivel_sinal_ftth_real" json:"nivel_sinal_ftth_real,omitempty"`
	NivelSinalFtthAnteriorReal float64    `gorm:"column:nivel_sinal_ftth_anterior_real" json:"nivel_sinal_ftth_anterior_real,omitempty"`
	UltimaAtualizacao          *time.Time `gorm:"column:ultima_atualizacao" json:"ultima_atualizacao,omitempty"`
	TemperaturaCapturada       float64    `gorm:"column:temperatura_capturada" json:"temperatura_capturada,omitempty"`
	Serial                     string     `gorm:"column:serial" json:"serial,omitempty"`
	DistanciaCapturada         int32      `gorm:"column:distancia_capturada" json:"distancia_capturada,omitempty"`
	CaixaGenerica              string     `gorm:"column:caixa_generica" json:"caixa_generica,omitempty"`
	CdArmarioAnterior          int32      `gorm:"column:cd_armario_anterior" json:"cd_armario_anterior,omitempty"`
	DiagramaLegado             string     `gorm:"column:diagrama_legado" json:"diagrama_legado,omitempty"`
}

func (*MkFiberCaixa) TableName() string {
	return TableNameMkFiberCaixa
}
