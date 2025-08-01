package model

const TableNameMkRange = "mk_ranges"

type MkRange struct {
	Codrange             int32  `gorm:"column:codrange;not null" json:"codrange,omitempty"`
	Codservidor          int32  `gorm:"column:codservidor;comment:cod do servidor associado" json:"codservidor,omitempty"` // cod do servidor associado
	Mascara              int32  `gorm:"column:mascara;comment:1 para /24 e 2 para /30" json:"mascara,omitempty"`           // 1 para /24 e 2 para /30
	Descricao            string `gorm:"column:descricao;not null" json:"descricao,omitempty"`
	Prefixo              string `gorm:"column:prefixo" json:"prefixo,omitempty"`
	ParametroInicial     int32  `gorm:"column:parametro_inicial" json:"parametro_inicial,omitempty"`
	ParametroFinal       int32  `gorm:"column:parametro_final" json:"parametro_final,omitempty"`
	PrefixoBloqueio      string `gorm:"column:prefixo_bloqueio;comment:prefixo da range de ips bloqueados." json:"prefixo_bloqueio,omitempty"` // prefixo da range de ips bloqueados.
	RangeLivre           string `gorm:"column:range_livre" json:"range_livre,omitempty"`
	LimitarPlanos        string `gorm:"column:limitar_planos" json:"limitar_planos,omitempty"`
	BloqueioPorPool      string `gorm:"column:bloqueio_por_pool" json:"bloqueio_por_pool,omitempty"`
	LimitarPontos        string `gorm:"column:limitar_pontos" json:"limitar_pontos,omitempty"`
	IPReferencia         string `gorm:"column:ip_referencia" json:"ip_referencia,omitempty"`
	IPReferenciaBloqueio string `gorm:"column:ip_referencia_bloqueio" json:"ip_referencia_bloqueio,omitempty"`
	PublicarPoolCentral  string `gorm:"column:publicar_pool_central" json:"publicar_pool_central,omitempty"`
	TipoRangeCentral     string `gorm:"column:tipo_range_central" json:"tipo_range_central,omitempty"`
	IPInicial            string `gorm:"column:ip_inicial" json:"ip_inicial,omitempty"`
	IPFinal              string `gorm:"column:ip_final" json:"ip_final,omitempty"`
	RangeGerencia        string `gorm:"column:range_gerencia" json:"range_gerencia,omitempty"`
	RangeGerenciaVoip    string `gorm:"column:range_gerencia_voip" json:"range_gerencia_voip,omitempty"`
	NomeConexaoAdd       string `gorm:"column:nome_conexao_add" json:"nome_conexao_add,omitempty"`
	CdEmpresa            int32  `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	Familia              int32  `gorm:"column:familia;comment:6= IPv6 4= IPv4" json:"familia,omitempty"` // 6= IPv6 4= IPv4
	PrefixoV6            string `gorm:"column:prefixo_v6" json:"prefixo_v6,omitempty"`
	TipoPrefixoV6        int32  `gorm:"column:tipo_prefixo_v6;comment:1= Prefixo de sessao PPPoE(WAN) 2= Prefixo de delegacao(LAN)" json:"tipo_prefixo_v6,omitempty"` // 1= Prefixo de sessao PPPoE(WAN) 2= Prefixo de delegacao(LAN)
	Ipv6Inicial          string `gorm:"column:ipv6_inicial" json:"ipv6_inicial,omitempty"`
	Ipv6Final            string `gorm:"column:ipv6_final" json:"ipv6_final,omitempty"`
	Ipv6MascQuebra       int32  `gorm:"column:ipv6_masc_quebra;comment:1 = Quebrar em /48 2 = Quebrar em /56 3 = Quebrar em 48 para PJ e 56 para PF" json:"ipv6_masc_quebra,omitempty"` // 1 = Quebrar em /48 2 = Quebrar em /56 3 = Quebrar em 48 para PJ e 56 para PF
	Rota                 string `gorm:"column:rota" json:"rota,omitempty"`
}

func (*MkRange) TableName() string {
	return TableNameMkRange
}
