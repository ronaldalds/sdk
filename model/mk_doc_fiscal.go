package model

import (
	"time"
)

const TableNameMkDocFiscal = "mk_doc_fiscal"

type MkDocFiscal struct {
	// Relationsships Init
	Pessoa          MkPessoa           `gorm:"foreignKey:codpessoa;references:cd_cliente" json:"pessoa,omitempty"`
	ContasFaturadas []MkContasFaturada `gorm:"foreignKey:cd_fatura;references:cd_fatura" json:"contas_faturadas,omitempty"`
	// Relationsships End
	Coddocfiscal int32 `gorm:"column:coddocfiscal;not null" json:"coddocfiscal,omitempty"`
	CdCliente    int32 `gorm:"column:cd_cliente" json:"cd_cliente,omitempty"`
	CdFatura     int32 `gorm:"column:cd_fatura" json:"cd_fatura,omitempty"`

	CdCompetencia          int32      `gorm:"column:cd_competencia" json:"cd_competencia,omitempty"`
	Dh                     *time.Time `gorm:"column:dh" json:"dh,omitempty"`
	Numero                 string     `gorm:"column:numero" json:"numero,omitempty"`
	IDOperdador            int32      `gorm:"column:id_operdador" json:"id_operdador,omitempty"`
	Status                 int32      `gorm:"column:status;comment:1- previo 2- linhas geradas 9- concluido (nota gerada)" json:"status,omitempty"` // 1- previo 2- linhas geradas 9- concluido (nota gerada)
	Vlr                    float64    `gorm:"column:vlr" json:"vlr,omitempty"`
	VlrBaseCalculo         float64    `gorm:"column:vlr_base_calculo" json:"vlr_base_calculo,omitempty"`
	VlrIcms                float64    `gorm:"column:vlr_icms" json:"vlr_icms,omitempty"`
	VlrIsento              float64    `gorm:"column:vlr_isento" json:"vlr_isento,omitempty"`
	VlrOutros              float64    `gorm:"column:vlr_outros" json:"vlr_outros,omitempty"`
	TipoDoc                string     `gorm:"column:tipo_doc" json:"tipo_doc,omitempty"`
	AutenticacaoDigital    string     `gorm:"column:autenticacao_digital" json:"autenticacao_digital,omitempty"`
	Situacao               string     `gorm:"column:situacao;comment:N - normal S - cancelada" json:"situacao,omitempty"` // N - normal S - cancelada
	VlrComercialFatura     float64    `gorm:"column:vlr_comercial_fatura" json:"vlr_comercial_fatura,omitempty"`
	DocCliente             string     `gorm:"column:doc_cliente" json:"doc_cliente,omitempty"`
	StrAutenticacaoDigital string     `gorm:"column:str_autenticacao_digital" json:"str_autenticacao_digital,omitempty"`
	LinhaMestre            string     `gorm:"column:linha_mestre" json:"linha_mestre,omitempty"`
	LinhaDestinatario      string     `gorm:"column:linha_destinatario" json:"linha_destinatario,omitempty"`
	GUIDTmp                string     `gorm:"column:guid_tmp" json:"guid_tmp,omitempty"`
	Cfop                   int32      `gorm:"column:cfop" json:"cfop,omitempty"`
	Ie                     string     `gorm:"column:ie" json:"ie,omitempty"`
	Endereco               string     `gorm:"column:endereco" json:"endereco,omitempty"`
	Bairro                 string     `gorm:"column:bairro" json:"bairro,omitempty"`
	Cidade                 string     `gorm:"column:cidade" json:"cidade,omitempty"`
	Uf                     string     `gorm:"column:uf" json:"uf,omitempty"`
	Cep                    string     `gorm:"column:cep" json:"cep,omitempty"`
	NumEnd                 int32      `gorm:"column:num_end" json:"num_end,omitempty"`
	Telefone               string     `gorm:"column:telefone" json:"telefone,omitempty"`
	NumeroInt              int32      `gorm:"column:numero_int" json:"numero_int,omitempty"`
	Obs1                   string     `gorm:"column:obs1" json:"obs1,omitempty"`
	Obs2                   string     `gorm:"column:obs2" json:"obs2,omitempty"`
	Obs3                   string     `gorm:"column:obs3" json:"obs3,omitempty"`
	ImpressaoGerada        string     `gorm:"column:impressao_gerada" json:"impressao_gerada,omitempty"`
	EnvioGerado            string     `gorm:"column:envio_gerado" json:"envio_gerado,omitempty"`
	CdGeracaoMassa         int32      `gorm:"column:cd_geracao_massa" json:"cd_geracao_massa,omitempty"`
	VlrPis                 float64    `gorm:"column:vlr_pis" json:"vlr_pis,omitempty"`
	VlrCofins              float64    `gorm:"column:vlr_cofins" json:"vlr_cofins,omitempty"`
	VlrFust                float64    `gorm:"column:vlr_fust" json:"vlr_fust,omitempty"`
	VlrFuntel              float64    `gorm:"column:vlr_funtel" json:"vlr_funtel,omitempty"`
	Ibge                   string     `gorm:"column:ibge" json:"ibge,omitempty"`
	DhSo                   *time.Time `gorm:"column:dh_so;default:now()" json:"dh_so,omitempty"`
	Complemento            string     `gorm:"column:complemento" json:"complemento,omitempty"`
	RazaoSocialCliente     string     `gorm:"column:razao_social_cliente" json:"razao_social_cliente,omitempty"`
	CidadeIbge             string     `gorm:"column:cidade_ibge" json:"cidade_ibge,omitempty"`
	Obs4                   string     `gorm:"column:obs4" json:"obs4,omitempty"`
	TipoAssinante          int32      `gorm:"column:tipo_assinante" json:"tipo_assinante,omitempty"`
	DocEmitente            string     `gorm:"column:doc_emitente" json:"doc_emitente,omitempty"`
	Release                int32      `gorm:"column:release" json:"release,omitempty"`
	VlrRetencoes           float64    `gorm:"column:vlr_retencoes" json:"vlr_retencoes,omitempty"`
	VlrIss                 float64    `gorm:"column:vlr_iss" json:"vlr_iss,omitempty"`
	CdNfse                 int32      `gorm:"column:cd_nfse" json:"cd_nfse,omitempty"`
	NomeEmitente           string     `gorm:"column:nome_emitente" json:"nome_emitente,omitempty"`
	EndEmitente            string     `gorm:"column:end_emitente" json:"end_emitente,omitempty"`
	CidadeEmitente         string     `gorm:"column:cidade_emitente" json:"cidade_emitente,omitempty"`
	CepEmitente            string     `gorm:"column:cep_emitente" json:"cep_emitente,omitempty"`
	FoneEmitente           string     `gorm:"column:fone_emitente" json:"fone_emitente,omitempty"`
	SiteEmitente           string     `gorm:"column:site_emitente" json:"site_emitente,omitempty"`
	EmailEmitente          string     `gorm:"column:email_emitente" json:"email_emitente,omitempty"`
	IeEmitente             string     `gorm:"column:ie_emitente" json:"ie_emitente,omitempty"`
	Logo                   string     `gorm:"column:logo" json:"logo,omitempty"`
	TipoUtilizacao         int32      `gorm:"column:tipo_utilizacao" json:"tipo_utilizacao,omitempty"`
	OrigemDoc              int32      `gorm:"column:origem_doc;comment:1- de contratos 2- lcto avulsos" json:"origem_doc,omitempty"` // 1- de contratos 2- lcto avulsos
	NumTerminalTelefonico  string     `gorm:"column:num_terminal_telefonico" json:"num_terminal_telefonico,omitempty"`
	SpedCodPart            string     `gorm:"column:sped_cod_part;comment:Campo COD_PART Sped Fiscal EFD ICMS/IPI" json:"sped_cod_part,omitempty"` // Campo COD_PART Sped Fiscal EFD ICMS/IPI
	ValorLiquido           float64    `gorm:"column:valor_liquido" json:"valor_liquido,omitempty"`
	VlrPisDestaque         float64    `gorm:"column:vlr_pis_destaque" json:"vlr_pis_destaque,omitempty"`
	VlrCofinsDestaque      float64    `gorm:"column:vlr_cofins_destaque" json:"vlr_cofins_destaque,omitempty"`
	VlrCsllDestaque        float64    `gorm:"column:vlr_csll_destaque" json:"vlr_csll_destaque,omitempty"`
	VlrIrpjDestaque        float64    `gorm:"column:vlr_irpj_destaque" json:"vlr_irpj_destaque,omitempty"`
	VlrIrDestaque          float64    `gorm:"column:vlr_ir_destaque" json:"vlr_ir_destaque,omitempty"`
}

func (*MkDocFiscal) TableName() string {
	return TableNameMkDocFiscal
}
