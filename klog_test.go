package klog

import (
	"./klog"
	"testing"
)

func TestPrintlog(t *testing.T) {
	/* expected TRUE test */
	actual := klog.Printlog("TestPrintfile")
	expected := true
	if actual != expected {
		t.Errorf("got %v,want %v", actual, expected)
	}
}

func TestPrintfile(t *testing.T) {
	/* expected TRUE test */
	actual := klog.Printfile("TestPrintfile", "test")
	expected := true
	if actual != expected {
		t.Errorf("got %v,want %v", actual, expected)
	}
}
