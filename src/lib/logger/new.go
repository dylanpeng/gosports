package logger

func NewLogger(config *Config) (*Logger, error) {
	logger := &Logger{
		Config: config,
	}
	err := logger.InitLogger()
	return logger, err
}
