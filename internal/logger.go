package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"payment-simulator/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//Loggers ...
type Loggers interface {
	// Initiate(filename string, maxsize int, maxBackups int, compress bool)
}

type Logging struct {
	Filename   string `mapstructure:"filename"`
	maxsize    int    `mapstructure:"maxsize"`
	maxBackups int    `mapstructure:"maxBackups"`
	compress   bool   `mapstructure:"compress"`
}

var (
	sugarLogger *zap.SugaredLogger

	logger *zap.Logger
	// filename    string
	// maxsize     int
	// maxBackups  int
	// maxAge      int
	// compress    bool
	// proddev     string
)

func Initiate(pfilename string, pmaxsize int, pmaxBackup int, pcompress bool) {

	//check if pfilename set already or not.
	if pfilename == "" {
		path, err := utils.GetPathNow()
		if err != nil {
			fmt.Println("error when get working directory")
		}
		pfilename = path + "/log/testlog.log"

	}

	Loggs := &Logging{
		Filename:   pfilename,
		maxsize:    pmaxsize,
		maxBackups: pmaxBackup,
		compress:   pcompress,
	}
	fmt.Println("Init Logger")
	writerSyncer := Loggs.getLogWriter()
	encoder := Loggs.getEncoderFile()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	zap.ReplaceGlobals(logger)
}

func (l Logging) getEncoderFile() zapcore.Encoder {
	// encoderConfig := zap.NewDevelopmentEncoderConfig()
	// if proddev == "0" {
	// 	encoderConfig = zap.NewProductionEncoderConfig()
	// }
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// return zapcore.NewConsoleEncoder(encoderConfig)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "msg"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "file"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2 Jan 2001 15:04:05"))
}

// rotated. It defaults to 100 megabytes.
func (l Logging) getLogWriter() zapcore.WriteSyncer {
	logpath := filepath.Dir(l.Filename)
	if err := os.MkdirAll(logpath, 0744); err != nil {
		fmt.Println("cant create log directory")
		return nil
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   l.Filename,
		MaxSize:    l.maxsize,
		MaxBackups: l.maxBackups,
		MaxAge:     10,
		Compress:   l.compress,
	}
	fmt.Println("LumberJackLogger ", lumberJackLogger)
	return zapcore.AddSync(lumberJackLogger)
}

func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func GetLoggerEncodeFile() *zap.SugaredLogger {
	return sugarLogger
}

func GinLogger() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		path := c.Request.URL.Path

		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)
		zap.S().Infow(path,

			zap.Int("status", c.Writer.Status()),

			zap.String("method", c.Request.Method),

			zap.String("path", path),

			zap.String("query", query),

			zap.String("ip", c.ClientIP()),

			zap.String("user-agent", c.Request.UserAgent()),

			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),

			zap.Duration("cost", cost))

	}

}
