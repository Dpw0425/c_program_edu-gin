package config

type Jwt struct {
	Secret      string `json_utils:"secret" yaml:"secret" mapstructure:"secret"`                   // Jwt 秘钥
	ExpiresTime int64  `json_utils:"expires_time" yaml:"expires_time" mapstructure:"expires_time"` // 过期时间(单位秒)
	BufferTime  int64  `json_utils:"buffer_time" yaml:"buffer_time" mapstructure:"buffer_time"`    // 缓冲时间(单位秒)
}
