package initialize

import (
	"database/sql"
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"template/internal/app/config"
	"template/internal/app/initialize/internal"
	"time"
)

func InitGorm(c *config.Config) (*gorm.DB, func(), error) {

	db, err := newInstance(c)
	if err != nil {
		panic(err)
	}
	if c.System.EnableAutoMigrate {
		//自动迁移
	}

	if c.System.IsDebug {
		db = db.Debug()
	}

	cleanFunc := func() {

	}

	return db, cleanFunc, nil
}

func newInstance(c *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch strings.ToLower(c.System.DbType) {
	case "mysql":
		// create database if not exists
		cfg, err := mysqlDriver.ParseDSN(c.Mysql.DNS())
		if err != nil {
			return nil, err
		}

		err = createDatabaseWithMySQL(cfg)
		if err != nil {
			return nil, err
		}

		dialector = mysql.Open(c.Mysql.DNS())
	case "postgres":
		//dialector = postgres.Open(c.DSN)
	default:
		//dialector = sqlite.Open(c.DSN)
		dialector = mysql.Open(c.Mysql.DNS())
	}

	db, err := gorm.Open(dialector, internal.Gorm.Config())
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	return db, nil
}

func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
