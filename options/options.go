package options

type Options struct {
	Delay         int
	MaxRetryCount int
	RequestCount  int
	Address       string
	Message       string
	OutputFile    string
	LogLevel      string
	Http          bool
	ForceHttps    bool
	IgnoreError   bool
	NoColor       bool
}
