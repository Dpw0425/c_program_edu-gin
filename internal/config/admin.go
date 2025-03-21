package config

type Admin struct {
	UserName string `json_utils:"user_name" yaml:"user_name" mapstructure:"user_name"`
	Password string `json_utils:"password" yaml:"password" mapstructure:"password"`
}
