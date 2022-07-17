package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	JWT    `json:"jwt" mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System `json:"system" mapstructure:"system" yaml:"system"`
	Mysql  `json:"mysql" mapstructure:"mysql" yaml:"mysql"`
	Zap    `json:"zap" mapstructure:"zap" yaml:"zap"`
}

func LoadConfig(cfgFile string) (*Config, error) {
	file, err := os.Open(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("open file failed, err: %v", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("close file failed, err: %v", err)
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("file is not found, err: %v", err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("read file failed, err: %v", err)
	}
	return parseConfig(buffer)
}

func parseConfig(buffer []byte) (*Config, error) {
	var cfg Config
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(buffer)); err != nil {
		return nil, fmt.Errorf("viper .ReadConfig failed, err: %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("viper can't unmarshal data,err :%v", err)
	}
	return &cfg, nil
}

func PrintWithJSON(c *Config) {
	b, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
		return
	}
	os.Stdout.WriteString(string(b) + "\n")
}
