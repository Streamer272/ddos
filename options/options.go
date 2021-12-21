package options

type Options struct {
	Delay         int
	MaxRetryCount int
	WorkerCount   int
	Address       string
	Message       string
	LogLevel      string
	OutputFile    string
	Http          bool
	IgnoreError   bool
	NoColor       bool
}
