package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire метод будет вызываться каждый раз при записи лога
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

// Levels будет возвращать Level из каждого хука
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

// Logger своя структура для логгера - более стабильная работа при изменениях библиотек
// а также облегчает переход на другой логгер
type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func init() {
	// инициализируем логгер
	l := logrus.New()
	// включаем report
	l.SetReportCaller(true)

	l.Formatter = &logrus.TextFormatter{
		// укажем формат логгирования
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			// данный логгер будет возвращать функцию, имя файла, линию
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	// создаём каталог
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	// открываем файл
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard) // чтобы логи никуда не отправлялись

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	// установим максимально подробный уровень логирования
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
