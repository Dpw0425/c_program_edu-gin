package config

type Server struct {
	Http      int `json_utils:"http" yaml:"http" mapstructure:"http"`
	Websocket int `json_utils:"websocket" yaml:"websocket" mapstructure:"websocket"`
	Tcp       int `json_utils:"tcp" yaml:"tcp" mapstructure:"tcp"`
}
