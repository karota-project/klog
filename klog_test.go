package klog

import (
	"testing"
)

func TestStdout(t *testing.T) {
	/* expected TRUE test */
	actual := Stdout("TestPrintfile")
	expected := true

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestWriteFile(t *testing.T) {
	/* expected TRUE test */
	actual := klog.WriteFile("main", "sample.log")
	var expected error = nil

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestSyslog(t *testing.T) {
	/* expected TRUE test */
	actual := klog.Syslog(klog.LOG_NOTICE, "main")
	var expected error = nil
	if actual != expected {
		t.Errorf("got %v,want %v", actual, expected)
	}
}
