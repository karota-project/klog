package main

import (
	"./klog"
)

func main() {
	klog.Printlog("function-name")
	klog.Printfile("function-name", "output-file-name")
}
