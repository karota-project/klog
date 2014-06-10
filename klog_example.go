package main

import (
	"./klog"
	"fmt"
)

func main() {
	s, err := klog.Printlog("main")
	if err != nil {
		fmt.Println(s, err)
	}

	s, err = klog.Printfile("main", "sample.log")
	if err != nil {
		fmt.Println(s, err)
	}
}
