package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`                               // 地址
	Port          string `mapstructure:"port" json:"port" yaml:"port"`                               // 端口
	DbType        string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`                      // 数据库类型: postgresql
	OssType       string `mapstructure:"oss_type" json:"oss_type" yaml:"oss_type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use_multipoint" json:"use_multipoint" yaml:"use_multipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"use_redis" json:"use_redis" yaml:"use_redis"`                // 使用redis
	LimitCountIP  int    `mapstructure:"iplimit_count" json:"iplimit_count" yaml:"iplimit_count"`    // IP限制登录次数
	LimitTimeIP   int    `mapstructure:"iplimit_time" json:"iplimit_time" yaml:"iplimit_time"`       // IP限制登录时间
}
