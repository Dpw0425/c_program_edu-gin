package config

type Redis struct {
	Host     string `json_utils:"host" yaml:"host" mapstructure:"host"`             // 服务器 ip 地址
	Port     string `json_utils:"port" yaml:"port" mapstructure:"port"`             // 服务器端口号
	Auth     string `json_utils:"auth" yaml:"auth" mapstructure:"auth"`             // 密码
	Database int    `json_utils:"database" yaml:"database" mapstructure:"database"` // 数据库
}
