package klog

import (
  "testing"
  "./klog"
)

func TestPrintlog(t *testing.T) {
	/* expected TRUE test */
    actual := klog.Printlog("test")
    expected := true
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestPrintfile(t *testing.T) {
    /* expected TRUE test */
    actual := klog.Printfile("test", "test")
    expected := true
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}