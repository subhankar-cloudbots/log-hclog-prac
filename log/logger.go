package log

import "github.com/hashicorp/go-hclog"

func InitLogger(loggerOptions *hclog.LoggerOptions) hclog.Logger {
	return hclog.New(loggerOptions)
}

func SubLogger(log hclog.Logger, name string) hclog.Logger {
	return log.Named(name)
}
