package logger

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"go-clean-monolith/config"
	"go-clean-monolith/constants"
	"go.uber.org/fx/fxevent"
	gormLogger "gorm.io/gorm/logger"
	"io"
	"os"
	"time"
)

// =====================================================================================================================

func getLogFile(path string) (*os.File, error) {
	file, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE,
		0664,
	)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func newLogger(logMode []string, logFileName string) zerolog.Logger {
	var writers []io.Writer
	for _, mode := range logMode {
		switch mode {
		case constants.EnvCmdLogMode:
			writers = append(writers, zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: "15:04:05 02.01.2006",
			})
		case constants.EnvJCmdLogMode:
			writers = append(writers, os.Stdout)
		case constants.EnvFileLogMode:
			file, _ := getLogFile(logFileName)
			writers = append(writers, file)
		}
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	multi := zerolog.MultiLevelWriter(writers...)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	return logger
}

// =====================================================================================================================

// FxLogger structure
type FxLogger struct {
	zerolog.Logger
}

func (l FxLogger) LogEvent(event fxevent.Event) {
	//fmt.Println(reflect2.TypeOf(event), event)
	//switch e := event.(type) {
	//case *fxevent.Provided:
	//	fmt.Println(e.ConstructorName, e.ModuleName, e.OutputTypeNames)
	//case *fxevent.Invoked:
	//	fmt.Println(e)
	//case *fxevent.Run:
	//	fmt.Println(e)
	//case *fxevent.Stopped:
	//	fmt.Println(2, e)
	//}
}

func NewFxLogger() FxLogger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return FxLogger{
		logger,
	}
}

// =====================================================================================================================

// GormLogger structure
type GormLogger struct {
	zerolog.Logger
	gormLogger.Config
}

// LogMode set log mode
func (l GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newlogger := l
	newlogger.LogLevel = level
	return &newlogger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		fmt.Println("GORM INFO:", str, args)
	}
}

// Warn prints warn messages
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		fmt.Println("GORM WARN:", str, args)
	}

}

// Error prints error messages
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		fmt.Println("GORM ERROR:", str, args)
	}
}

// Trace prints trace messages
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	if l.LogLevel >= gormLogger.Info {
		sql, rows := fc()
		fmt.Printf("GORM INFO: [%d ms, %d rows] sql -> %s \n", elapsed.Milliseconds(), rows, sql)
		return
	}

	if l.LogLevel >= gormLogger.Warn {
		sql, rows := fc()
		fmt.Printf("GORM WARN: [%d ms, %d rows] sql -> %s \n", elapsed.Milliseconds(), rows, sql)
		return
	}

	if l.LogLevel >= gormLogger.Error {
		sql, rows := fc()
		fmt.Printf("GORM ERROR: [%d ms, %d rows] sql -> %s \n", elapsed.Milliseconds(), rows, sql)
		return
	}
}

func newGormLogger() GormLogger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return GormLogger{
		logger,
		gormLogger.Config{
			LogLevel: gormLogger.Info,
		},
	}
}

// =====================================================================================================================

// Logger structure
type Logger struct {
	zerolog.Logger
	Gorm GormLogger
}

func NewLogger(env config.Env) Logger {
	return Logger{
		newLogger(env.LogMode, env.LogFileName),
		newGormLogger(),
	}
}

// =====================================================================================================================
