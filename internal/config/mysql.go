package config

type Mysql struct {
	Host       string `json_utils:"host" yaml:"host" mapstructure:"host"`
	Port       string `json_utils:"port" yaml:"port" mapstructure:"port"`
	Username   string `json_utils:"username" yaml:"username" mapstructure:"username"`
	Password   string `json_utils:"password" yaml:"password" mapstructure:"password"`
	DBname     string `json_utils:"dbname" yaml:"dbname" mapstructure:"dbname"`
	Parameters string `json_utils:"parameters" yaml:"parameters" mapstructure:"parameters"`
}

func (m *Mysql) DSN() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBname + "?" + m.Parameters
}
