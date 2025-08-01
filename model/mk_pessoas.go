package model

import (
	"time"
)

const TableNameMkPessoa = "mk_pessoas"

type MkPessoa struct {
	// Relationsships Init
	Conexoes           []MkConexao  `gorm:"foreignKey:codcliente;references:codpessoa" json:"conexoes,omitempty"`
	Contratos          []MkContrato `gorm:"foreignKey:cliente;references:codpessoa" json:"contratos,omitempty"`
	Logradouro         MkLogradouro `gorm:"foreignKey:codlogradouro;references:codlogradouro" json:"logradouro,omitempty"`
	Logradourocobranca MkLogradouro `gorm:"foreignKey:codlogradouro;references:codlogradourocobranca" json:"logradourocobranca,omitempty"`
	Revenda            MkRevenda    `gorm:"foreignKey:codrevenda;references:cd_revenda" json:"revenda,omitempty"`
	// Relationsships End
	Codpessoa                   int32      `gorm:"column:codpessoa;not null" json:"codpessoa,omitempty"`
	Codlogradouro               int32      `gorm:"column:codlogradouro;not null" json:"codlogradouro,omitempty"`
	Codlogradourocobranca       int32      `gorm:"column:codlogradourocobranca;not null" json:"codlogradourocobranca,omitempty"`
	NomeRazaosocial             string     `gorm:"column:nome_razaosocial;not null" json:"nome_razaosocial,omitempty"`
	Cpf                         string     `gorm:"column:cpf" json:"cpf,omitempty"`
	Cnpj                        string     `gorm:"column:cnpj" json:"cnpj,omitempty"`
	Rg                          string     `gorm:"column:rg" json:"rg,omitempty"`
	CdRevenda                   int32      `gorm:"column:cd_revenda" json:"cd_revenda,omitempty"`
	Ie                          string     `gorm:"column:ie" json:"ie,omitempty"`
	Nascimento                  *time.Time `gorm:"column:nascimento" json:"nascimento,omitempty"`
	Fundacao                    *time.Time `gorm:"column:fundacao" json:"fundacao,omitempty"`
	Cep                         string     `gorm:"column:cep" json:"cep,omitempty"`
	Cepcobranca                 string     `gorm:"column:cepcobranca" json:"cepcobranca,omitempty"`
	Codcidade                   int32      `gorm:"column:codcidade;not null" json:"codcidade,omitempty"`
	Codcidadecobranca           int32      `gorm:"column:codcidadecobranca;not null" json:"codcidadecobranca,omitempty"`
	Codbairro                   int32      `gorm:"column:codbairro;not null" json:"codbairro,omitempty"`
	Codbairrocobranca           int32      `gorm:"column:codbairrocobranca;not null" json:"codbairrocobranca,omitempty"`
	Codestado                   int32      `gorm:"column:codestado;not null" json:"codestado,omitempty"`
	Codestadocobranca           int32      `gorm:"column:codestadocobranca;not null" json:"codestadocobranca,omitempty"`
	Nomepai                     string     `gorm:"column:nomepai" json:"nomepai,omitempty"`
	Nomemae                     string     `gorm:"column:nomemae" json:"nomemae,omitempty"`
	Email                       string     `gorm:"column:email" json:"email,omitempty"`
	Contato                     string     `gorm:"column:contato;comment:pessoa de contato" json:"contato,omitempty"` // pessoa de contato
	Fone01                      string     `gorm:"column:fone01;not null" json:"fone01,omitempty"`
	Fone02                      string     `gorm:"column:fone02" json:"fone02,omitempty"`
	Fax                         string     `gorm:"column:fax" json:"fax,omitempty"`
	Complementoendereco         string     `gorm:"column:complementoendereco" json:"complementoendereco,omitempty"`
	Complementoenderecocobr     string     `gorm:"column:complementoenderecocobr" json:"complementoenderecocobr,omitempty"`
	Diapgto                     int32      `gorm:"column:diapgto;comment:dia preferivel para pgtos" json:"diapgto,omitempty"`                                   // dia preferivel para pgtos
	Tipopgtoobservacoes         string     `gorm:"column:tipopgtoobservacoes;comment:detalhes referentes ao tipo de pgto" json:"tipopgtoobservacoes,omitempty"` // detalhes referentes ao tipo de pgto
	Datacadastro                *time.Time `gorm:"column:datacadastro;not null" json:"datacadastro,omitempty"`
	Enderecoweb                 string     `gorm:"column:enderecoweb" json:"enderecoweb,omitempty"`
	Horarioparaatendimento      string     `gorm:"column:horarioparaatendimento;comment:horario preferivel para os atendimentos" json:"horarioparaatendimento,omitempty"`       // horario preferivel para os atendimentos
	Tipopessoa                  int32      `gorm:"column:tipopessoa;comment:1 fisica, 2 juridica" json:"tipopessoa,omitempty"`                                                  // 1 fisica, 2 juridica
	Classificacao               int32      `gorm:"column:classificacao;not null;comment:cliente, fornecedor, cliente e fornecedor, funcionario" json:"classificacao,omitempty"` // cliente, fornecedor, cliente e fornecedor, funcionario
	Tipopgto                    int32      `gorm:"column:tipopgto;comment:boleto bancario, debito em conta, cheque, outros" json:"tipopgto,omitempty"`                          // boleto bancario, debito em conta, cheque, outros
	Numero                      int32      `gorm:"column:numero;not null;comment:numero do endereco" json:"numero,omitempty"`                                                   // numero do endereco
	Numerocobranca              int32      `gorm:"column:numerocobranca;not null;comment:nuumero endereco cobranca" json:"numerocobranca,omitempty"`                            // nuumero endereco cobranca
	Observacoes                 string     `gorm:"column:observacoes" json:"observacoes,omitempty"`
	Igualresidencia             string     `gorm:"column:igualresidencia" json:"igualresidencia,omitempty"`
	AceitaEmails                string     `gorm:"column:aceita_emails" json:"aceita_emails,omitempty"`
	Foto                        []uint8    `gorm:"column:foto" json:"foto,omitempty"`
	AtualizadoSac               *time.Time `gorm:"column:atualizado_sac" json:"atualizado_sac,omitempty"`
	IDAlternativo               int32      `gorm:"column:id_alternativo" json:"id_alternativo,omitempty"`
	UsarIDAlternativo           string     `gorm:"column:usar_id_alternativo" json:"usar_id_alternativo,omitempty"`
	TmpAniversariante           string     `gorm:"column:tmp_aniversariante" json:"tmp_aniversariante,omitempty"`
	NomeExibicaoChat            string     `gorm:"column:nome_exibicao_chat" json:"nome_exibicao_chat,omitempty"`
	PasswdChat                  string     `gorm:"column:passwd_chat" json:"passwd_chat,omitempty"`
	MensagemExibicaoChat        string     `gorm:"column:mensagem_exibicao_chat" json:"mensagem_exibicao_chat,omitempty"`
	SenhaEmail                  string     `gorm:"column:senha_email" json:"senha_email,omitempty"`
	Inativo                     string     `gorm:"column:inativo;default:N" json:"inativo,omitempty"`
	MotivoInatividade           int32      `gorm:"column:motivo_inatividade" json:"motivo_inatividade,omitempty"`
	PlanoPadrao                 int32      `gorm:"column:plano_padrao" json:"plano_padrao,omitempty"`
	DiaVencimento               int32      `gorm:"column:dia_vencimento" json:"dia_vencimento,omitempty"`
	ProfilePadrao               int32      `gorm:"column:profile_padrao" json:"profile_padrao,omitempty"`
	MensagemNf                  string     `gorm:"column:mensagem_nf" json:"mensagem_nf,omitempty"`
	TmpSemBoletoGerado          string     `gorm:"column:tmp_sem_boleto_gerado" json:"tmp_sem_boleto_gerado,omitempty"`
	DtInatividade               *time.Time `gorm:"column:dt_inatividade" json:"dt_inatividade,omitempty"`
	Observacoes2                string     `gorm:"column:observacoes2" json:"observacoes2,omitempty"`
	ComConexao                  string     `gorm:"column:com_conexao" json:"com_conexao,omitempty"`
	RetornoSerasa               string     `gorm:"column:retorno_serasa" json:"retorno_serasa,omitempty"`
	MensgemNf2                  string     `gorm:"column:mensgem_nf2" json:"mensgem_nf2,omitempty"`
	AutorizadoAbrirSuporte      string     `gorm:"column:autorizado_abrir_suporte" json:"autorizado_abrir_suporte,omitempty"`
	AutorizadoAbrirSuporteCpf   string     `gorm:"column:autorizado_abrir_suporte_cpf" json:"autorizado_abrir_suporte_cpf,omitempty"`
	Cfop                        string     `gorm:"column:cfop" json:"cfop,omitempty"`
	Condominio                  int32      `gorm:"column:condominio" json:"condominio,omitempty"`
	UserSac                     string     `gorm:"column:user_sac" json:"user_sac,omitempty"`
	PassSac                     string     `gorm:"column:pass_sac" json:"pass_sac,omitempty"`
	AcessoSac                   string     `gorm:"column:acesso_sac" json:"acesso_sac,omitempty"`
	MensagemParaCliente         string     `gorm:"column:mensagem_para_cliente" json:"mensagem_para_cliente,omitempty"`
	DtMaxMensagemCliente        *time.Time `gorm:"column:dt_max_mensagem_cliente" json:"dt_max_mensagem_cliente,omitempty"`
	NaturezaJuridica            int32      `gorm:"column:natureza_juridica" json:"natureza_juridica,omitempty"`
	ControleAssinante           int32      `gorm:"column:controle_assinante" json:"controle_assinante,omitempty"`
	CdPessoaExterno             int32      `gorm:"column:cd_pessoa_externo" json:"cd_pessoa_externo,omitempty"`
	Im                          string     `gorm:"column:im" json:"im,omitempty"`
	GrupoAtendimento            int32      `gorm:"column:grupo_atendimento" json:"grupo_atendimento,omitempty"`
	OperadoraFone               string     `gorm:"column:operadora_fone" json:"operadora_fone,omitempty"`
	Whatsapp                    string     `gorm:"column:whatsapp" json:"whatsapp,omitempty"`
	NomeFantasia                string     `gorm:"column:nome_fantasia" json:"nome_fantasia,omitempty"`
	TokenGcm                    string     `gorm:"column:token_gcm" json:"token_gcm,omitempty"`
	TokenSac                    string     `gorm:"column:token_sac" json:"token_sac,omitempty"`
	LoginSacHotspot             string     `gorm:"column:login_sac_hotspot" json:"login_sac_hotspot,omitempty"`
	CdAgenteCobranca            int32      `gorm:"column:cd_agente_cobranca" json:"cd_agente_cobranca,omitempty"`
	LoginSacEspn                string     `gorm:"column:login_sac_espn" json:"login_sac_espn,omitempty"`
	FuncionarioTerceirizado     string     `gorm:"column:funcionario_terceirizado" json:"funcionario_terceirizado,omitempty"`
	CdEmpresa                   int32      `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	ClassificacaoTipoClienteCom string     `gorm:"column:classificacao_tipo_cliente_com" json:"classificacao_tipo_cliente_com,omitempty"`
	SincAutoMultiempresa        string     `gorm:"column:sinc_auto_multiempresa" json:"sinc_auto_multiempresa,omitempty"`
	IDTelegram                  string     `gorm:"column:id_telegram" json:"id_telegram,omitempty"`
	TxSpcBlackList              string     `gorm:"column:tx_spc_black_list;comment:" json:"tx_spc_black_list,omitempty"`                                                                   //
	ControleRestricaoCadastro   int32      `gorm:"column:controle_restricao_cadastro;comment:1 = led verde 2 = led amarelo 3 = led vermelho" json:"controle_restricao_cadastro,omitempty"` // 1 = led verde 2 = led amarelo 3 = led vermelho
	GUIDControleInsert          string     `gorm:"column:guid_controle_insert" json:"guid_controle_insert,omitempty"`
	TipoAssinante               int32      `gorm:"column:tipo_assinante" json:"tipo_assinante,omitempty"`
	NivelValidacaoEmail         int32      `gorm:"column:nivel_validacao_email" json:"nivel_validacao_email,omitempty"`
	SucessoValidaEmail          string     `gorm:"column:sucesso_valida_email" json:"sucesso_valida_email,omitempty"`
	EnderecoPrestacaoServico    string     `gorm:"column:endereco_prestacao_servico;default:C" json:"endereco_prestacao_servico,omitempty"`
	Latitude                    string     `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude                   string     `gorm:"column:longitude" json:"longitude,omitempty"`
	Sexo                        int32      `gorm:"column:sexo;comment:0- Nao informado 1- Feminino 2- Masculino 3- Indefinido" json:"sexo,omitempty"` // 0- Nao informado 1- Feminino 2- Masculino 3- Indefinido
	ChatIgnoraAteAuto           string     `gorm:"column:chat_ignora_ate_auto;default:N" json:"chat_ignora_ate_auto,omitempty"`
	CdPais                      int32      `gorm:"column:cd_pais;default:1058" json:"cd_pais,omitempty"`
	LicencaCodigoIntegracao     string     `gorm:"column:licenca_codigo_integracao" json:"licenca_codigo_integracao,omitempty"`
	IDEstrangeiro               string     `gorm:"column:id_estrangeiro" json:"id_estrangeiro,omitempty"`
	CelularWhats                string     `gorm:"column:celular_whats" json:"celular_whats,omitempty"`
	Migra                       int32      `gorm:"column:migra" json:"migra,omitempty"`
	MigraDuplo                  int32      `gorm:"column:migra_duplo" json:"migra_duplo,omitempty"`
	PinRecuperacao              string     `gorm:"column:pin_recuperacao" json:"pin_recuperacao,omitempty"`
	PinDataExpiracao            *time.Time `gorm:"column:pin_data_expiracao" json:"pin_data_expiracao,omitempty"`
	NomeSocial                  string     `gorm:"column:nome_social" json:"nome_social,omitempty"`
	Extrangeiro                 string     `gorm:"column:extrangeiro" json:"extrangeiro,omitempty"`
	TipoDocExt                  string     `gorm:"column:tipo_doc_ext" json:"tipo_doc_ext,omitempty"`
	NumeroDocExt                string     `gorm:"column:numero_doc_ext" json:"numero_doc_ext,omitempty"`
	IDClienteMknext             string     `gorm:"column:id_cliente_mknext" json:"id_cliente_mknext,omitempty"`
	EmailProtegido              string     `gorm:"column:email_protegido;comment:Criado para ativar ou não o envio de boleto em PDF protegido por senha." json:"email_protegido,omitempty"` // Criado para ativar ou não o envio de boleto em PDF protegido por senha.
	Pin2factor                  string     `gorm:"column:pin2factor" json:"pin2factor,omitempty"`
	Pin2factorExpireDate        *time.Time `gorm:"column:pin2factor_expire_date" json:"pin2factor_expire_date,omitempty"`
	Enable2FactorAuth           bool       `gorm:"column:enable_2factor_auth" json:"enable_2factor_auth,omitempty"`
	PinValidated                bool       `gorm:"column:pin_validated" json:"pin_validated,omitempty"`
	SacExibeMsgInativacaoCartao string     `gorm:"column:sac_exibe_msg_inativacao_cartao" json:"sac_exibe_msg_inativacao_cartao,omitempty"`
	ObsMigracao                 string     `gorm:"column:obs_migracao" json:"obs_migracao,omitempty"`
	DataExpiracaoToken          *time.Time `gorm:"column:data_expiracao_token" json:"data_expiracao_token,omitempty"`
	MknextSincronizado          string     `gorm:"column:mknext_sincronizado" json:"mknext_sincronizado,omitempty"`
	Referencia                  string     `gorm:"column:referencia" json:"referencia,omitempty"`
	Referenciacobr              string     `gorm:"column:referenciacobr" json:"referenciacobr,omitempty"`
	Horacadastro                *time.Time `gorm:"column:horacadastro;default:CURRENT_TIMESTAMP" json:"horacadastro,omitempty"`
	SistemaOperacional          string     `gorm:"column:sistema_operacional" json:"sistema_operacional,omitempty"`
	TermoAceite                 string     `gorm:"column:termo_aceite" json:"termo_aceite,omitempty"`
}

func (*MkPessoa) TableName() string {
	return TableNameMkPessoa
}
