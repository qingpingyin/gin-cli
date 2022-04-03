package app

import (
	"database/sql"
	"fmt"
	"gin-cli/internal/app/config"
	"gin-cli/internal/app/repo"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)

func InitGorm(c *config.Config) (*gorm.DB, func(), error) {

	db, err := dailGorm(c)
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	if c.System.EnableAutoMigrate {
		if err = repo.AutoMigrate(db); err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}

func dailGorm(c *config.Config) (*gorm.DB, error) {
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
	//case "postgres":
	//	dialector = postgres.Open(c.DSN)
	default:
		//dialector = sqlite.Open(c.DSN)
	}

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.System.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	if c.System.IsDebug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.System.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.System.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.System.MaxLifetime) * time.Second)

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
