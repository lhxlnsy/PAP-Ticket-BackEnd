package config

type JWT struct {
	SigningKey string `mapstructure:"signing_key" json:"signing_key" yaml:"signing_key"` // JWT签名密钥
	ExpireTime int64  `mapstructure:"expire_time" json:"expire_time" yaml:"expire_time"` // JWT过期时间
	BufferTime int64  `mapstructure:"buffer_time" json:"buffer_time" yaml:"buffer_time"` // JWT缓冲时间
	Issuer     string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                // JWT签发者
}
