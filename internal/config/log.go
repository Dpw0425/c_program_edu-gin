package config

type Log struct {
	Path string `json_utils:"path" yaml:"path" mapstructure:"path"`
}
