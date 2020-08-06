package stdout

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger instantiates, preconfigures and returns a new logger.
func NewLogger(logLevel string) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("could not initialize loger: %s", err)
	}

	return &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyLevel: "level_name",
			},
		},
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}, nil
}
