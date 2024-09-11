package logger

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes logging
func Init() error {
	runLogFile, err := os.OpenFile(
		"app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		return fmt.Errorf("open file failed: %w", err)
	}

	multi := zerolog.MultiLevelWriter(runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()

	// Disable Console Color, don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	gin.DefaultWriter = multi
	gin.DefaultErrorWriter = multi

	return nil
}
