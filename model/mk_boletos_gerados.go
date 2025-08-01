package model

import (
	"time"
)

const TableNameMkBoletosGerado = "mk_boletos_gerados"

type MkBoletosGerado struct {
	Codgeracao               int32      `gorm:"column:codgeracao;not null" json:"codgeracao,omitempty"`
	Valor                    float64    `gorm:"column:valor" json:"valor,omitempty"`
	Vencimento               string     `gorm:"column:vencimento" json:"vencimento,omitempty"`
	NossoNumero              string     `gorm:"column:nosso_numero" json:"nosso_numero,omitempty"`
	NomeSacado               string     `gorm:"column:nome_sacado" json:"nome_sacado,omitempty"`
	Endereco                 string     `gorm:"column:endereco" json:"endereco,omitempty"`
	Cep                      string     `gorm:"column:cep" json:"cep,omitempty"`
	Cidade                   string     `gorm:"column:cidade" json:"cidade,omitempty"`
	Estado                   string     `gorm:"column:estado" json:"estado,omitempty"`
	Linha1                   string     `gorm:"column:linha1" json:"linha1,omitempty"`
	Linha2                   string     `gorm:"column:linha2" json:"linha2,omitempty"`
	Linha3                   string     `gorm:"column:linha3" json:"linha3,omitempty"`
	Linha4                   string     `gorm:"column:linha4" json:"linha4,omitempty"`
	Linha5                   string     `gorm:"column:linha5" json:"linha5,omitempty"`
	Linha6                   string     `gorm:"column:linha6" json:"linha6,omitempty"`
	Linha7                   string     `gorm:"column:linha7" json:"linha7,omitempty"`
	Linha8                   string     `gorm:"column:linha8" json:"linha8,omitempty"`
	NomeCedente              string     `gorm:"column:nome_cedente" json:"nome_cedente,omitempty"`
	LocalPgto                string     `gorm:"column:local_pgto" json:"local_pgto,omitempty"`
	TxAux1                   string     `gorm:"column:tx_aux_1" json:"tx_aux_1,omitempty"`
	TxAux2                   string     `gorm:"column:tx_aux_2" json:"tx_aux_2,omitempty"`
	TxAux3                   string     `gorm:"column:tx_aux_3" json:"tx_aux_3,omitempty"`
	TxAux4                   string     `gorm:"column:tx_aux_4" json:"tx_aux_4,omitempty"`
	TxAux5                   string     `gorm:"column:tx_aux_5" json:"tx_aux_5,omitempty"`
	TxAux6                   string     `gorm:"column:tx_aux_6" json:"tx_aux_6,omitempty"`
	TxAux7                   string     `gorm:"column:tx_aux_7" json:"tx_aux_7,omitempty"`
	Banco                    string     `gorm:"column:banco" json:"banco,omitempty"`
	Agencia                  string     `gorm:"column:agencia" json:"agencia,omitempty"`
	Aceite                   string     `gorm:"column:aceite" json:"aceite,omitempty"`
	CedenteBarra             string     `gorm:"column:cedente_barra;comment:Variavel obrigatoria para carteira especial unibanco" json:"cedente_barra,omitempty"`              // Variavel obrigatoria para carteira especial unibanco
	CodCliente               string     `gorm:"column:cod_cliente;comment:Obrigatoria para carteira CSB e HSBC e cobranca sem registro do Safra" json:"cod_cliente,omitempty"` // Obrigatoria para carteira CSB e HSBC e cobranca sem registro do Safra
	CnpjCedente              string     `gorm:"column:cnpj_cedente" json:"cnpj_cedente,omitempty"`
	CpfCnpjSacado            string     `gorm:"column:cpf_cnpj_sacado" json:"cpf_cnpj_sacado,omitempty"`
	DataProcessamento        string     `gorm:"column:data_processamento" json:"data_processamento,omitempty"`
	Desconto                 string     `gorm:"column:desconto" json:"desconto,omitempty"`
	Modalidade               string     `gorm:"column:modalidade" json:"modalidade,omitempty"`
	Multa                    string     `gorm:"column:multa" json:"multa,omitempty"`
	NumConvenio              string     `gorm:"column:num_convenio" json:"num_convenio,omitempty"`
	NumDoc                   string     `gorm:"column:num_doc" json:"num_doc,omitempty"`
	Parcela                  string     `gorm:"column:parcela" json:"parcela,omitempty"`
	Carteira                 string     `gorm:"column:carteira" json:"carteira,omitempty"`
	CodCedente               string     `gorm:"column:cod_cedente" json:"cod_cedente,omitempty"`
	OutrosAcrescimos         string     `gorm:"column:outros_acrescimos" json:"outros_acrescimos,omitempty"`
	ValorCobrado             string     `gorm:"column:valor_cobrado" json:"valor_cobrado,omitempty"`
	Codconta                 int32      `gorm:"column:codconta" json:"codconta,omitempty"`
	Profile                  int32      `gorm:"column:profile" json:"profile,omitempty"`
	Substituido              string     `gorm:"column:substituido;default:N" json:"substituido,omitempty"`
	DataRemocao              *time.Time `gorm:"column:data_remocao" json:"data_remocao,omitempty"`
	HoraRemocao              string     `gorm:"column:hora_remocao" json:"hora_remocao,omitempty"`
	UserRemocao              string     `gorm:"column:user_remocao" json:"user_remocao,omitempty"`
	Codcontrato              int32      `gorm:"column:codcontrato" json:"codcontrato,omitempty"`
	TxMensTitulo             string     `gorm:"column:tx_mens_titulo" json:"tx_mens_titulo,omitempty"`
	TxMensLinha1             string     `gorm:"column:tx_mens_linha1" json:"tx_mens_linha1,omitempty"`
	TxMensLinha2             string     `gorm:"column:tx_mens_linha2" json:"tx_mens_linha2,omitempty"`
	TxMensLinha3             string     `gorm:"column:tx_mens_linha3" json:"tx_mens_linha3,omitempty"`
	TxMensLinha4             string     `gorm:"column:tx_mens_linha4" json:"tx_mens_linha4,omitempty"`
	TxMensLinha5             string     `gorm:"column:tx_mens_linha5" json:"tx_mens_linha5,omitempty"`
	TxMensLinha6             string     `gorm:"column:tx_mens_linha6" json:"tx_mens_linha6,omitempty"`
	TxMensLinha7             string     `gorm:"column:tx_mens_linha7" json:"tx_mens_linha7,omitempty"`
	SMTP                     string     `gorm:"column:smtp" json:"smtp,omitempty"`
	Username                 string     `gorm:"column:username" json:"username,omitempty"`
	Password                 string     `gorm:"column:password" json:"password,omitempty"`
	EmailOrigem              string     `gorm:"column:email_origem" json:"email_origem,omitempty"`
	EmailDestino             string     `gorm:"column:email_destino" json:"email_destino,omitempty"`
	ContaEmail               int32      `gorm:"column:conta_email" json:"conta_email,omitempty"`
	VencimentoDtNormal       *time.Time `gorm:"column:vencimento_dt_normal" json:"vencimento_dt_normal,omitempty"`
	GUID                     string     `gorm:"column:guid" json:"guid,omitempty"`
	ArquivoRemessa           string     `gorm:"column:arquivo_remessa;comment:1- Boleto 2- Debito" json:"arquivo_remessa,omitempty"` // 1- Boleto 2- Debito
	GeradoTransmissao        string     `gorm:"column:gerado_transmissao" json:"gerado_transmissao,omitempty"`
	DtGerTransmissao         *time.Time `gorm:"column:dt_ger_transmissao" json:"dt_ger_transmissao,omitempty"`
	TmpTrans                 string     `gorm:"column:tmp_trans" json:"tmp_trans,omitempty"`
	NomeArquivo              string     `gorm:"column:nome_arquivo" json:"nome_arquivo,omitempty"`
	SeqArqRemessa            int32      `gorm:"column:seq_arq_remessa" json:"seq_arq_remessa,omitempty"`
	TmpMarcaConsulta         string     `gorm:"column:tmp_marca_consulta" json:"tmp_marca_consulta,omitempty"`
	GUIDEnvioEmail           string     `gorm:"column:guid_envio_email" json:"guid_envio_email,omitempty"`
	RetNossoNumBanco         string     `gorm:"column:ret_nosso_num_banco" json:"ret_nosso_num_banco,omitempty"`
	ProfileAntiga            int32      `gorm:"column:profile_antiga" json:"profile_antiga,omitempty"`
	EntradaConfirmada        string     `gorm:"column:entrada_confirmada" json:"entrada_confirmada,omitempty"`
	NomenclaturaIntegracao   string     `gorm:"column:nomenclatura_integracao" json:"nomenclatura_integracao,omitempty"`
	Codvinculado             int32      `gorm:"column:codvinculado" json:"codvinculado,omitempty"`
	NossoNumero2             string     `gorm:"column:nosso_numero2" json:"nosso_numero2,omitempty"`
	GUIDExibicao             string     `gorm:"column:guid_exibicao" json:"guid_exibicao,omitempty"`
	Suspenso                 string     `gorm:"column:suspenso;default:N" json:"suspenso,omitempty"`
	CdFatura                 int32      `gorm:"column:cd_fatura" json:"cd_fatura,omitempty"`
	ProfileRedefinida        string     `gorm:"column:profile_redefinida" json:"profile_redefinida,omitempty"`
	ProfileOriginal          int32      `gorm:"column:profile_original" json:"profile_original,omitempty"`
	CodigoBarras             string     `gorm:"column:codigo_barras" json:"codigo_barras,omitempty"`
	LinhaDigitavel           string     `gorm:"column:linha_digitavel" json:"linha_digitavel,omitempty"`
	NossoNumeroFormatado     string     `gorm:"column:nosso_numero_formatado" json:"nosso_numero_formatado,omitempty"`
	HashIntegridade          string     `gorm:"column:hash_integridade" json:"hash_integridade,omitempty"`
	Instrucoes               string     `gorm:"column:instrucoes" json:"instrucoes,omitempty"`
	IDUnificacao             int32      `gorm:"column:id_unificacao" json:"id_unificacao,omitempty"`
	CdBoletoReplicado        int32      `gorm:"column:cd_boleto_replicado" json:"cd_boleto_replicado,omitempty"`
	CdEmpresa                int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	TmpTransEmPendencia      string     `gorm:"column:tmp_trans_em_pendencia" json:"tmp_trans_em_pendencia,omitempty"`
	GUIDControleBack         string     `gorm:"column:guid_controle_back" json:"guid_controle_back,omitempty"`
	NomeFileBack             string     `gorm:"column:nome_file_back" json:"nome_file_back,omitempty"`
	ControleDeOrdenacao      int32      `gorm:"column:controle_de_ordenacao" json:"controle_de_ordenacao,omitempty"`
	NossoNumeroDv            string     `gorm:"column:nosso_numero_dv" json:"nosso_numero_dv,omitempty"`
	GUIDControleTrocaProfile string     `gorm:"column:guid_controle_troca_profile" json:"guid_controle_troca_profile,omitempty"`
	Dh                       *time.Time `gorm:"column:dh;default:now()" json:"dh,omitempty"`
	CdRemessaGerada          int32      `gorm:"column:cd_remessa_gerada" json:"cd_remessa_gerada,omitempty"`
	ParcelaICarne            int32      `gorm:"column:parcela_i_carne" json:"parcela_i_carne,omitempty"`
	ParcelaFCarne            int32      `gorm:"column:parcela_f_carne" json:"parcela_f_carne,omitempty"`
	QrcodePixImagem          string     `gorm:"column:qrcode_pix_imagem" json:"qrcode_pix_imagem,omitempty"`
	LinhaQrcodePix           string     `gorm:"column:linha_qrcode_pix" json:"linha_qrcode_pix,omitempty"`
	IDCobrancaMknext         string     `gorm:"column:id_cobranca_mknext" json:"id_cobranca_mknext,omitempty"`
	StatusCobrancaMknext     string     `gorm:"column:status_cobranca_mknext" json:"status_cobranca_mknext,omitempty"`
	DhStatusMknext           *time.Time `gorm:"column:dh_status_mknext" json:"dh_status_mknext,omitempty"`
	DhRegistroMknext         *time.Time `gorm:"column:dh_registro_mknext;comment:Data e hora que a cobrança foi registrada/cadastrada no banco." json:"dh_registro_mknext,omitempty"` // Data e hora que a cobrança foi registrada/cadastrada no banco.
	ComandoAlteracao         string     `gorm:"column:comando_alteracao" json:"comando_alteracao,omitempty"`
}

func (*MkBoletosGerado) TableName() string {
	return TableNameMkBoletosGerado
}
