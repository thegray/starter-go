package config

import (
	"starter-go/internal/pkg/logger"
)

type loggerConfig struct {
	EnableStdout   bool            `yaml:"enable_stdout"`
	EnableLogFile  bool            `yaml:"enable_logfile"`
	CallerSkipSet  bool            `yaml:"caller_skipset"`
	CallerSkip     int             `yaml:"caller_skip"`
	LogFileConfigs []logFileConfig `yaml:"logfile_configs"`
}

type logFileConfig struct {
	Levels           []string `yaml:"levels"`
	IsAccessLog      bool     `yaml:"is_access_log"`
	FullpathFilename string   `yaml:"fullpath_filename"`
	MaxSize          int      `yaml:"max_size"`
	MaxAge           int      `yaml:"max_age"`
	MaxBackups       int      `yaml:"max_backups"`
	LocalTime        bool     `yaml:"local_time"`
	Compress         bool     `yaml:"compress"`
}

func LoggerConfig() logger.LogConfig {
	var logFileConfigs []logger.LogFileConfig
	for _, fileConf := range cfg.Logger.LogFileConfigs {
		logFileConfigs = append(logFileConfigs, logger.LogFileConfig{
			Levels:           fileConf.Levels,
			IsAccessLog:      fileConf.IsAccessLog,
			FullpathFilename: fileConf.FullpathFilename,
			MaxSize:          fileConf.MaxSize,
			MaxAge:           fileConf.MaxAge,
			MaxBackups:       fileConf.MaxBackups,
			LocalTime:        fileConf.LocalTime,
			Compress:         fileConf.Compress,
		})
	}

	return logger.LogConfig{
		EnableStdout:   cfg.Logger.EnableStdout,
		EnableLogFile:  cfg.Logger.EnableLogFile,
		CallerSkipSet:  cfg.Logger.CallerSkipSet,
		CallerSkip:     cfg.Logger.CallerSkip,
		LogFileConfigs: logFileConfigs,
	}
}
