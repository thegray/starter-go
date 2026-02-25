package config

type DatabaseConfig interface {
	GetHost() string
	GetPort() uint
	GetUser() string
	GetPassword() string
	GetName() string
}

type databaseConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     uint   `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Name     string `yaml:"name" mapstructure:"name"`
}

func Database() DatabaseConfig {
	return &cfg.DB
}

func (db *databaseConfig) GetHost() string {
	return db.Host
}

func (db *databaseConfig) GetPort() uint {
	return db.Port
}

func (db *databaseConfig) GetUser() string {
	return db.User
}

func (db *databaseConfig) GetPassword() string {
	return db.Password
}

func (db *databaseConfig) GetName() string {
	return db.Name
}
