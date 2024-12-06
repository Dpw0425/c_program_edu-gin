package config

type App struct {
	Env        string   `json_utils:"env" yaml:"env" mapstructure:"env"`
	RunMode    string   `json_utils:"run_mode" yaml:"run_mode" mapstructure:"run_mode"`
	PublicKey  string   `json_utils:"-" yaml:"public_key"`
	PrivateKey string   `json_utils:"-" yaml:"private_key"`
	AdminEmail []string `json_utils:"admin_email"`
}
