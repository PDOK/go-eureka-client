package eureka

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var logger *eurekaLogger
var debug_enabled bool = true

func SetLogger(l *log.Logger) {
	logger = &eurekaLogger{l}
}

func GetLogger() *log.Logger {
	return logger.log
}

func GetEurekaLogger() *eurekaLogger {
	return logger
}

func SetDebugEnabled(enabled bool) {
	debug_enabled = enabled
}

type eurekaLogger struct {
	log *log.Logger
}

func (p *eurekaLogger) Debug(args ...interface{}) {
	if (debug_enabled) {
		msg := "DEBUG: " + fmt.Sprint(args...)
		p.log.Println(msg)
	}
}

func (p *eurekaLogger) Debugf(f string, args ...interface{}) {
	if (debug_enabled) {
		msg := "DEBUG: " + fmt.Sprintf(f, args...)
		// Append newline if necessary
		if !strings.HasSuffix(msg, "\n") {
			msg = msg + "\n"
		}
		p.log.Print(msg)
	}
}

func (p *eurekaLogger) Warning(args ...interface{}) {
	msg := "WARNING: " + fmt.Sprint(args...)
	p.log.Println(msg)
}

func (p *eurekaLogger) Warningf(f string, args ...interface{}) {
	msg := "WARNING: " + fmt.Sprintf(f, args...)
	// Append newline if necessary
	if !strings.HasSuffix(msg, "\n") {
		msg = msg + "\n"
	}
	p.log.Print(msg)
}

func (p *eurekaLogger) Errorf(f string, args ...interface{}) {
	msg := "ERROR: " + fmt.Sprintf(f, args...)
	// Append newline if necessary
	if !strings.HasSuffix(msg, "\n") {
		msg = msg + "\n"
	}
	p.log.Print(msg)
}

func init() {
	// Default logger uses the go default log.
	SetLogger(log.New(os.Stdout, "go-eureka", log.LstdFlags))
}