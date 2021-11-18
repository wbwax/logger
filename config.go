package logger

// Config defines the configuration of logger
type Config struct {
	MaxSize    int    `json:"max_size" yaml:"max_size"`       // unit: MB
	MaxAge     int    `json:"max_age" yaml:"max_age"`         // unit: day
	MaxBackups int    `json:"max_backups" yaml:"max_backups"` // unit: short
	Level      string `json:"level" yaml:"level"`             // log level
	Path       string `json:"path" yaml:"path"`               // path to hold log file
	Encoding   string `json:"encoding" yaml:"encoding"`       // json or console
}
