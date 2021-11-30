package config

import (
	"os"
	"time"
)

const (
	LayoutDateLog = "2006-01-02 15:04:05"
	LayoutDate    = "2006-01-02"
	LayoutHour    = "15:04:05"
)

type ByteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	//ZB
	//YB
)

var (
	HOST_MAXBYTE         = 1 << 26 // 64MB
	HOST_MAXBYTE_DEFAULT = 1 << 26 // 64MB
)

// Config provides basic configuration
type Config struct {
	Host           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func TypeEnv() string {
	amb := os.Getenv("ENV_AMBI")
	if len(amb) == 0 {
		amb = "LOCAL"
	}
	return amb
}
