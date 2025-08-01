package model

const TableNameMkOsTipo = "mk_os_tipo"

type MkOsTipo struct {
	Codostipo                 int32  `gorm:"column:codostipo;not null" json:"codostipo,omitempty"`
	Descricao                 string `gorm:"column:descricao;not null" json:"descricao,omitempty"`
	Observacoes               string `gorm:"column:observacoes" json:"observacoes,omitempty"`
	DiasFechamento            int32  `gorm:"column:dias_fechamento;comment:numero de dias para o atendimento" json:"dias_fechamento,omitempty"`  // numero de dias para o atendimento
	AssociarConexoes          string `gorm:"column:associar_conexoes;comment:associar a tabela de conexoes." json:"associar_conexoes,omitempty"` // associar a tabela de conexoes.
	HorasFechamento           int32  `gorm:"column:horas_fechamento" json:"horas_fechamento,omitempty"`
	ModeloOs                  int32  `gorm:"column:modelo_os" json:"modelo_os,omitempty"`
	TxExtraRecibo             string `gorm:"column:tx_extra_recibo" json:"tx_extra_recibo,omitempty"`
	EmailNotificacao          string `gorm:"column:email_notificacao" json:"email_notificacao,omitempty"`
	TxTitulo                  string `gorm:"column:tx_titulo" json:"tx_titulo,omitempty"`
	CorFundo                  int32  `gorm:"column:cor_fundo" json:"cor_fundo,omitempty"`
	CorBorda                  int32  `gorm:"column:cor_borda" json:"cor_borda,omitempty"`
	CorLetra                  int32  `gorm:"column:cor_letra" json:"cor_letra,omitempty"`
	ModeloOsPersonal          int32  `gorm:"column:modelo_os_personal" json:"modelo_os_personal,omitempty"`
	UnidadeFinanceiraSugerida string `gorm:"column:unidade_financeira_sugerida" json:"unidade_financeira_sugerida,omitempty"`
	MobNumImagens             int32  `gorm:"column:mob_num_imagens" json:"mob_num_imagens,omitempty"`
	MobDefaultTalk            string `gorm:"column:mob_default_talk" json:"mob_default_talk,omitempty"`
	GrupoAvisoEncerramento    int32  `gorm:"column:grupo_aviso_encerramento" json:"grupo_aviso_encerramento,omitempty"`
	GrupoAvisoEncTecnico      int32  `gorm:"column:grupo_aviso_enc_tecnico" json:"grupo_aviso_enc_tecnico,omitempty"`
	EnviarOsEncCliente        string `gorm:"column:enviar_os_enc_cliente" json:"enviar_os_enc_cliente,omitempty"`
	TxEncerramentoCliente     string `gorm:"column:tx_encerramento_cliente" json:"tx_encerramento_cliente,omitempty"`
	ModeloEmailEncerramento   int32  `gorm:"column:modelo_email_encerramento" json:"modelo_email_encerramento,omitempty"`
	AssuntoEmailEncerramento  string `gorm:"column:assunto_email_encerramento" json:"assunto_email_encerramento,omitempty"`
	TempoMedio                int32  `gorm:"column:tempo_medio;not null;default:30" json:"tempo_medio,omitempty"`
	TemperaturaMinima         int32  `gorm:"column:temperatura_minima;not null;default:99" json:"temperatura_minima,omitempty"`
	PermiteChuva              string `gorm:"column:permite_chuva;not null;default:S" json:"permite_chuva,omitempty"`
	Inativo                   string `gorm:"column:inativo;default:N" json:"inativo,omitempty"`
	Sigla                     string `gorm:"column:sigla" json:"sigla,omitempty"`
	ClassificacaoInterna      int32  `gorm:"column:classificacao_interna" json:"classificacao_interna,omitempty"`
	ObrigarItemOsFinalizar    string `gorm:"column:obrigar_item_os_finalizar" json:"obrigar_item_os_finalizar,omitempty"`
}

func (*MkOsTipo) TableName() string {
	return TableNameMkOsTipo
}
