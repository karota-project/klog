package main

import (
	"./klog"
	"fmt"
)

func main() {
	s , err := klog.Printlog("function-name")
	if err != nil {
		fmt.Println(s, err)
	}

	s , err = klog.Printfile("function-name", "output-file-name")
	if err != nil {
		fmt.Println(s, err)
	}
}
