package main

import (
	"./klog"
	"fmt"
)

func main() {
	s, err := klog.Stdout("main")
	if err != nil {
		fmt.Println(s, err)
	}

	s, err = klog.WriteFile("main", "sample.log")
	if err != nil {
		fmt.Println(s, err)
	}

	s, err = klog.Syslog(klog.LOG_NOTICE, "main")
	if err != nil {
		fmt.Println(s, err)
	}
}
