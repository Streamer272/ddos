package options

type Options struct {
	Delay         int
	MaxRetryCount int
	Address       string
	Message       string
	LogLevel      string
	Http          bool
	IgnoreError   bool
}

func (opt Options) LogLevelToInt() int {
	switch opt.LogLevel {
	case "INFO":
		return 0
	case "WARN":
		return 1
	case "ERROR":
		return 2
	default:
		return 3
	}
}
