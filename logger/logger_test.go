package logger

import (
	"fmt"
	"github.com/Streamer272/ddos/options"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLogger_Log(t *testing.T) {
	t.Run("Test logger output", func(t *testing.T) {
		os.Remove("test.log")

		log := NewLogger(options.Options{
			Delay:         0,
			MaxRetryCount: 0,
			RequestCount:  0,
			Address:       "www.google.com:443",
			Message:       "",
			OutputFile:    "test.log",
			LogLevel:      "INFO",
			Http:          false,
			IgnoreError:   false,
			NoColor:       false,
		})

		currentTime := time.Now().Format("15:04:05")

		log.Log("INFO", "This is test message 1", true)
		log.Log("INFO", "This is test message 2", true)

		expected := fmt.Sprintf("[INFO] %v: This is test message 1\n[INFO] %v: This is test message 2\n", currentTime, currentTime)

		content, err := ioutil.ReadFile("test.log")
		if err != nil {
			t.Errorf("Couldn't read file, %v...", err)
		}

		if string(content) != expected {
			t.Errorf("Content isn't equal, expected:\n%v\nactual:\n%v\n", expected, string(content))
		}

		os.Remove("test.log")
	})
}
