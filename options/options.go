package options

import "os"

type Options struct {
	Delay         int
	MaxRetryCount int
	WorkerCount   int
	Address       string
	Message       string
	LogLevel      string
	Http          bool
	IgnoreError   bool
	NoColor       bool
	OutputFile    *os.File
}
