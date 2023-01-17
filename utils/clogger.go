package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)


func CustomLogger(filename string) *os.File {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0664,)
    if err != nil {
        panic(err)
    }
    return file
}

func InitLogger(filename string, loglevel int) (zerolog.LevelWriter, error) {
    zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
    zerolog.TimeFieldFormat = time.RFC3339Nano
    loglevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
    if err != nil {
        loglevel = int(zerolog.InfoLevel) // default to Info
    }
    fileLogger := &lumberjack.Logger{
            Filename:   filename,
            MaxSize:    10, // megabytes
            MaxBackups: 10, // files
            MaxAge:     14,   // days
            Compress:   true, // disabled by default
        }
        output := zerolog.MultiLevelWriter(os.Stderr, fileLogger)
        return output, nil
}
