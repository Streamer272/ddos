package main

import (
	"github.com/Streamer272/ddos/logger"
	"github.com/Streamer272/ddos/options"
	"testing"
)

func Test_ddos(t *testing.T) {
	tests := []struct {
		name    string
		opt     options.Options
		wantErr bool
	}{
		{
			name: "Test bad protocol",
			opt: options.Options{
				Delay:         0,
				MaxRetryCount: 1,
				RequestCount:  1,
				Address:       "localhost:-1",
				Message:       "",
				OutputFile:    "",
				LogLevel:      "INFO",
				Http:          false,
				ForceHttps:    false,
				IgnoreError:   false,
				NoColor:       true,
			},
			wantErr: true,
		},
		{
			name: "Test invalid IP",
			opt: options.Options{
				Delay:         0,
				MaxRetryCount: 1,
				RequestCount:  1,
				Address:       "192.256.1.1:8080",
				Message:       "",
				OutputFile:    "",
				LogLevel:      "INFO",
				Http:          false,
				ForceHttps:    false,
				IgnoreError:   false,
				NoColor:       true,
			},
			wantErr: true,
		},
		{
			name: "Test OK response",
			opt: options.Options{
				Delay:         0,
				MaxRetryCount: 1,
				RequestCount:  1,
				Address:       "www.google.com:443",
				Message:       "",
				OutputFile:    "",
				LogLevel:      "INFO",
				Http:          true,
				ForceHttps:    false,
				IgnoreError:   false,
				NoColor:       true,
			},
			wantErr: false,
		},
		{
			name: "Test regex",
			opt: options.Options{
				Delay:         0,
				MaxRetryCount: 1,
				RequestCount:  1,
				Address:       "www.google.com",
				Message:       "",
				OutputFile:    "",
				LogLevel:      "INFO",
				Http:          false,
				ForceHttps:    true,
				IgnoreError:   false,
				NoColor:       true,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			log := logger.NewLogger(test.opt)
			log.Disable()
			test.opt.Address = fixAddress(test.opt, log)
			if err := ddos(test.opt); (err != nil) != test.wantErr {
				t.Errorf("ddos(%v) error = %v, wantErr = %v", test.opt, err, test.wantErr)
			}
		})
	}
}
