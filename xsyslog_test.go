//go:build !windows && !plan9

package xsyslog

import (
	"log/syslog"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()
	s, err := New(syslog.LOG_USER|syslog.LOG_ERR, "test")
	if err != nil {
		t.Fatalf("error: %v", s)
	}

	tests := []string{
		"write_test",
		"write_test 2",
	}

	for _, test := range tests {
		i, err := s.Write([]byte(test))
		if err != nil {
			t.Fatalf("error writing: %v", err)
		}
		t.Logf("wrote %d bytes", i)
	}
}
