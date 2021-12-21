package options

type Options struct {
	Delay         int
	MaxRetryCount int
	WorkerCount   int
	Address       string
	Message       string
	OutputFile    string
	LogLevel      string
	Http          bool
	IgnoreError   bool
	NoColor       bool
}
