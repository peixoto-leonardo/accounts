package logger

import (
	"context"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05 -0700",
		FullTimestamp:   true,
	})

	level := resolveLogLevelFromEnv()

	log.SetLevel(level)
}

func WithPrefix(ctx context.Context, prefix string) *log.Entry {
	return log.WithField("prefix", prefix).WithContext(ctx)
}

func resolveLogLevelFromEnv() log.Level {
	level, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		return log.InfoLevel
	}

	l, err := log.ParseLevel(strings.ToLower(level))
	if err != nil {
		log.Warnf("provided LOG_LEVEL %s is invalid. Fallback to info.", os.Getenv("LOG_LEVEL"))
		return log.InfoLevel
	}

	return l
}
