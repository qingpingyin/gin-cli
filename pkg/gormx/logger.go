package gormx

import (
	"gorm.io/gorm/logger"
	"template/internal/app/config"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(c *config.Config, message string, data ...interface{}) {
	var logZap bool
	switch c.System.DbType {
	case "mysql":
		logZap = c.Mysql.LogZap
	case "pgsql":
		//logZap = global.GVA_CONFIG.Pgsql.LogZap
	}
	if logZap {
		//..Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
