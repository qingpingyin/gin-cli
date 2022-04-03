package config

type System struct {
	Env               string `mapstructure:"env" json:"env" yaml:"env"`                            // 环境值
	Addr              string `mapstructure:"addr" json:"addr" yaml:"addr"`                         // 地址
	Port              int    `mapstructure:"port" json:"port" yaml:"port"`                         // 端口值
	DbType            string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	TablePrefix       string `mapstructure:"table-prefix" json:"table-prefix" yaml:"table-prefix"` // 表前缀
	IsDebug           bool   `mapstructure:"is-debug" json:"is-debug" yaml:"is-debug"`             //sql是否开启debug
	MaxLifetime       int    `mapstructure:"max-life-time" json:"max-life-time" yaml:"max-life-time"`
	MaxOpenConns      int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	MaxIdleConns      int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	EnableAutoMigrate bool   `mapstructure:"enable-auto-migrate"json:"enable-auto-migrate"yaml:"enable-auto-migrate"`
}
