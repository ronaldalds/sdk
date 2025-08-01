package model

import (
	"time"
)

const TableNameMkCrmProduto = "mk_crm_produtos"

type MkCrmProduto struct {
	Codcrmproduto           int32      `gorm:"column:codcrmproduto;not null" json:"codcrmproduto,omitempty"`
	NomeProduto             string     `gorm:"column:nome_produto" json:"nome_produto,omitempty"`
	CdPlanoPrincipal        int32      `gorm:"column:cd_plano_principal" json:"cd_plano_principal,omitempty"`
	DataInicio              *time.Time `gorm:"column:data_inicio" json:"data_inicio,omitempty"`
	DataFim                 *time.Time `gorm:"column:data_fim" json:"data_fim,omitempty"`
	VendaApp                string     `gorm:"column:venda_app" json:"venda_app,omitempty"`
	Tecnologia              string     `gorm:"column:tecnologia" json:"tecnologia,omitempty"`
	VlrFinal                float64    `gorm:"column:vlr_final" json:"vlr_final,omitempty"`
	InfoPlano               string     `gorm:"column:info_plano" json:"info_plano,omitempty"`
	AdmProduto              int32      `gorm:"column:adm_produto" json:"adm_produto,omitempty"`
	Ativo                   string     `gorm:"column:ativo" json:"ativo,omitempty"`
	VendaEmAreasPre         string     `gorm:"column:venda_em_areas_pre" json:"venda_em_areas_pre,omitempty"`
	Prioridade              int32      `gorm:"column:prioridade" json:"prioridade,omitempty"`
	VlrAdesao               float64    `gorm:"column:vlr_adesao" json:"vlr_adesao,omitempty"`
	AdesaoParcelasCartao    int32      `gorm:"column:adesao_parcelas_cartao" json:"adesao_parcelas_cartao,omitempty"`
	AdesaoParcelasBoleto    int32      `gorm:"column:adesao_parcelas_boleto" json:"adesao_parcelas_boleto,omitempty"`
	MensalidadeCartao       string     `gorm:"column:mensalidade_cartao" json:"mensalidade_cartao,omitempty"`
	MensalidadeBoleto       string     `gorm:"column:mensalidade_boleto" json:"mensalidade_boleto,omitempty"`
	DiasVctoPrimeiraParcela int32      `gorm:"column:dias_vcto_primeira_parcela" json:"dias_vcto_primeira_parcela,omitempty"`
	ComposicaoPlanoVlr      string     `gorm:"column:composicao_plano_vlr" json:"composicao_plano_vlr,omitempty"`
	Ncm                     string     `gorm:"column:ncm" json:"ncm,omitempty"`
	InteresseEmTv           bool       `gorm:"column:interesse_em_tv" json:"interesse_em_tv,omitempty"`
	InteresseEmTelefone     bool       `gorm:"column:interesse_em_telefone" json:"interesse_em_telefone,omitempty"`
	InteresseEmInternetF    bool       `gorm:"column:interesse_em_internet_f" json:"interesse_em_internet_f,omitempty"`
	InteresseEmInternetW    bool       `gorm:"column:interesse_em_internet_w" json:"interesse_em_internet_w,omitempty"`
	InteresseEmOutros       bool       `gorm:"column:interesse_em_outros" json:"interesse_em_outros,omitempty"`
	CdPerfilAssociado       int32      `gorm:"column:cd_perfil_associado" json:"cd_perfil_associado,omitempty"`
	PontuacaoVenda          int32      `gorm:"column:pontuacao_venda" json:"pontuacao_venda,omitempty"`
	PontuacaoUpgrade        int32      `gorm:"column:pontuacao_upgrade" json:"pontuacao_upgrade,omitempty"`
	PerfilContrato          int32      `gorm:"column:perfil_contrato" json:"perfil_contrato,omitempty"`
	CdEmpresa               int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	UnidadeFinanceira       string     `gorm:"column:unidade_financeira" json:"unidade_financeira,omitempty"`
	CodPerfilativacao       int32      `gorm:"column:cod_perfilativacao" json:"cod_perfilativacao,omitempty"`
	CodPerfilcancelamento   int32      `gorm:"column:cod_perfilcancelamento" json:"cod_perfilcancelamento,omitempty"`
	PontuacaoCancelamento   int32      `gorm:"column:pontuacao_cancelamento" json:"pontuacao_cancelamento,omitempty"`
	PontuacaoRevertido      int32      `gorm:"column:pontuacao_revertido" json:"pontuacao_revertido,omitempty"`
	QtdMensalidade          int32      `gorm:"column:qtd_mensalidade;default:12" json:"qtd_mensalidade,omitempty"`
	ProdutoPersonalizado    string     `gorm:"column:produto_personalizado" json:"produto_personalizado,omitempty"`
	VendaAppSac             string     `gorm:"column:venda_app_sac;default:N" json:"venda_app_sac,omitempty"`
}

func (*MkCrmProduto) TableName() string {
	return TableNameMkCrmProduto
}
