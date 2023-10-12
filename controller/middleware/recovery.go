package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-clean-monolith/pkg/logger"
	"os"
	"runtime"
	"unsafe"
)

// ---------------------------------------------------------------------------------------------------------------------

// RecoveryMiddleware description
type RecoveryMiddleware struct {
	logger logger.Logger
}

// NewRecoveryMiddleware description
func NewRecoveryMiddleware(logger logger.Logger) RecoveryMiddleware {
	return RecoveryMiddleware{
		logger: logger,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Setup description
func (m RecoveryMiddleware) Setup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Set("stack_trace", stack(3))

				httpResponse := gin.H{
					"error": err,
				}

				ctx.AbortWithStatusJSON(500, httpResponse)
			}
		}()
		ctx.Next()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

type Stack struct {
	File     string `json:"file_name"`
	Line     int    `json:"line"`
	FuncName string `json:"func_name"`
	FuncLine string `json:"func_line"`
}

type Stacks []Stack

func stack(skip int) (stacks Stacks) {
	var lines [][]byte
	var lastFile string

	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		if file != lastFile {
			data, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}

		f := function(pc)

		if string(f) == "HandlerWrapper.func1" {
			break
		}

		stacks = append(stacks, Stack{file, line, string(f), string(source(lines, line))})
	}

	return stacks
}
func source(lines [][]byte, n int) []byte {
	n-- // в stack индекс начинается от 1, у нас - 0
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	// конвертация из string в bytes без выделения памяти
	name := unsafe.Slice(unsafe.StringData(fn.Name()), len(fn.Name()))
	// Мы видим:
	//	runtime/debug.*T·ptrmethod
	// А хотим:
	//	*T.ptrmethod
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.ReplaceAll(name, centerDot, dot)
	return name
}

// ---------------------------------------------------------------------------------------------------------------------
