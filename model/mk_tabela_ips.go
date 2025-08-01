package model

const TableNameMkTabelaIP = "mk_tabela_ips"

type MkTabelaIP struct {
	Codip                int32  `gorm:"column:codip;not null" json:"codip,omitempty"`
	Codrange             int32  `gorm:"column:codrange" json:"codrange,omitempty"`
	NumeroIP             string `gorm:"column:numero_ip" json:"numero_ip,omitempty"`
	Status               string `gorm:"column:status;not null;comment:U para usado e L para disponivel" json:"status,omitempty"`                          // U para usado e L para disponivel
	Historico            string `gorm:"column:historico;comment:historico de ultimos clientes associados ao ip" json:"historico,omitempty"`               // historico de ultimos clientes associados ao ip
	AssociacaoAtual      int32  `gorm:"column:associacao_atual;comment:conexao que o ip esta associado atualmente." json:"associacao_atual,omitempty"`    // conexao que o ip esta associado atualmente.
	NumeroIPBloqueio     string `gorm:"column:numero_ip_bloqueio;comment:numero de ip recebido em caso de bloqueio." json:"numero_ip_bloqueio,omitempty"` // numero de ip recebido em caso de bloqueio.
	Ocultar              string `gorm:"column:ocultar" json:"ocultar,omitempty"`
	UsadoParaEquipamento string `gorm:"column:usado_para_equipamento;default:N" json:"usado_para_equipamento,omitempty"`
	CdConexaoEquip       int32  `gorm:"column:cd_conexao_equip" json:"cd_conexao_equip,omitempty"`
	UsadoParaRota        string `gorm:"column:usado_para_rota" json:"usado_para_rota,omitempty"`
	IgnoraDuplicidade    string `gorm:"column:ignora_duplicidade" json:"ignora_duplicidade,omitempty"`
	CdRota               int32  `gorm:"column:cd_rota" json:"cd_rota,omitempty"`
	GUIDReferencia       string `gorm:"column:guid_referencia" json:"guid_referencia,omitempty"`
	CdEmpresa            int32  `gorm:"column:cd_empresa" json:"cd_empresa,omitempty"`
	Familia              int32  `gorm:"column:familia;comment:6=IPv6 4=IPv4" json:"familia,omitempty"` // 6=IPv6 4=IPv4
	NumeroIpv6           string `gorm:"column:numero_ipv6" json:"numero_ipv6,omitempty"`
}

func (*MkTabelaIP) TableName() string {
	return TableNameMkTabelaIP
}
