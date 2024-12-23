package logapp

import (
	"app/internal/connection"
	constant "app/internal/constants"
	"fmt"

	"go.uber.org/zap"
)

func Logger(title string, mess string, typeLog constant.LOGGER_TYPE) {
	data := fmt.Sprintf("%s: %s", title, mess)
	logger := connection.GetLogger().WithOptions(zap.AddCallerSkip(1))

	switch typeLog {
	case constant.INFO_LOG:
		logger.Info(data)
	case constant.ERROR_LOG:
		logger.Error(data)
	}
}
