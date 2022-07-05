package config

import "strconv"

type Pgsql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 数据库连接字符串
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口
	User         string `mapstructure:"user" json:"user" yaml:"user"`                               // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库名
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 最大空闲连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 最大打开连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
}

// Dsn 基于配置文件获取 dsn
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " port=" + strconv.Itoa(p.Port) + " user=" + p.User + " password=" + p.Password + " dbname=" + p.Dbname + " " + p.Config
}

// LinkDsn 根据 dbname 生成 dsn
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + " port=" + strconv.Itoa(p.Port) + " user=" + p.User + " password=" + p.Password + " dbname=" + dbname + " " + p.Config
}

func (m *Pgsql) GetLogMode() string {
	return m.LogMode
}
