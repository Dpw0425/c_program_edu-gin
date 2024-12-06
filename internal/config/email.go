package config

type Email struct {
	Host     string `json_utils:"host" yaml:"host" mapstructure:"host"`
	Smtp     string `json_utils:"smtp" yaml:"smtp" mapstructure:"smtp"`
	Addr     string `json_utils:"addr" yaml:"addr" mapstructure:"addr"`
	Name     string `json_utils:"name" yaml:"name" mapstructure:"name"`
	Password string `json_utils:"password" yaml:"password" mapstructure:"password"`
	Expires  int    `json_utils:"expires" yaml:"expires" mapstructure:"expires"`
}
