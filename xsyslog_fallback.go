//go:build !darwin

package xsyslog

import (
	"log"
	"log/syslog"
)

type Writer struct {
	*syslog.Writer
}

func New(priority syslog.Priority, tag string) (*Writer, error) {
	s, err := syslog.New(priority, tag)
	if err != nil {
		return nil, err
	}

	return &Writer{Writer: s}, nil
}

func NewLogger(p syslog.Priority, logFlag int) (*log.Logger, error) {
	return syslog.NewLogger(p, logFlag)
}
