package config

import (
	"strconv"
)

type Mysql struct {
	Host       string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	DbName     string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username   string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password   string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Parameters string `mapstructure:"parameters" json:"parameters" yaml:"parameters"`
	LogMode    string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap     bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}

func (m *Mysql) DNS() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DbName + "?" + m.Parameters
}
