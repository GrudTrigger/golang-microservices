package env

import "github.com/caarlos0/env/v11"

type loggerConfigEnv struct {
	Level  string `env:"LOGGER_LEVEL,required"`
	AsJson bool   `env:"LOGGER_AS_JSON,required"`
}

type loggerConfig struct {
	row loggerConfigEnv
}

func NewLoggerConfig() (*loggerConfig, error) {
	var row loggerConfigEnv
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &loggerConfig{row: row}, nil
}

func (l *loggerConfig) Level() string {
	return l.row.Level
}

func (l *loggerConfig) AsJson() bool {
	return l.row.AsJson
}
