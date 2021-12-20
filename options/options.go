package options

type Options struct {
	Delay         int
	MaxRetryCount int
	Address       string
	Message       string
	LogLevel      string
	Http          bool
	IgnoreError   bool
	NoColor       bool
}
