package config

type ServerConfig interface {
	GetTimeZone() string
	GetLoglevel() string
	GetEnvironment() string
	GetBaseURL() string
	GetReadTimeout() uint
	GetWriteTimeout() uint
	GetIdleTimeout() uint
}

type serverConfig struct {
	TimeZone     string `yaml:"time_zone"`
	Loglevel     string `yaml:"loglevel"`
	Environment  string `yaml:"env"`
	BaseURL      string `yaml:"base_url"`
	ReadTimeout  uint   `yaml:"read_timeout"`
	WriteTimeout uint   `yaml:"write_timeout"`
	IdleTimeout  uint   `yaml:"idle_timeout"`
}

func Server() ServerConfig {
	return &cfg.Server
}

func (server *serverConfig) GetTimeZone() string {
	return server.TimeZone
}

func (server *serverConfig) GetLoglevel() string {
	return server.Loglevel
}

func (server *serverConfig) GetEnvironment() string {
	return server.Environment
}

func (server *serverConfig) GetBaseURL() string {
	return server.BaseURL
}

func (server *serverConfig) GetReadTimeout() uint {
	return server.ReadTimeout
}

func (server *serverConfig) GetWriteTimeout() uint {
	return server.WriteTimeout
}

func (server *serverConfig) GetIdleTimeout() uint {
	return server.IdleTimeout
}
