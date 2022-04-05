package models

import "github.com/BurntSushi/toml"

const (
	MainConfigPlaceHolder = "config/api.toml"
)

// MainConfig - весь конфигурационный файл.
type MainConfig struct {
	HTTPserverAddr string         `toml:"http_server_addr"`
	BindAddr       string         `toml:"bind_addr"`
	LoggerLevel    string         `toml:"logger_level"`
	DocsAddr       string         `toml:"docs_addr"`
	Database       DatabaseConfig `toml:"database"`
}

// DatabaseConfig - только часть конфиг. файла с определённым в нём полем Database (подключение к БД).
type DatabaseConfig struct {
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Host         string `toml:"host"`
	DatabaseName string `toml:"database_name"`
}

// NewMainConfig - иcпользуем, чтобы получить весь конфиг. файл.
func NewMainConfig() (*MainConfig, error) {

	mainConfig := new(MainConfig)

	if _, err := toml.DecodeFile(MainConfigPlaceHolder, &mainConfig); err != nil {
		return mainConfig, err
	}

	return mainConfig, nil
}
