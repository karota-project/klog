package main

import (
	"./klog"
	"fmt"
)

func main() {
	err := klog.Stdout("main")
	if err != nil {
		fmt.Println(err)
	}

	err = klog.WriteFile("main", "sample.log")
	if err != nil {
		fmt.Println(err)
	}

	err = klog.Syslog(klog.LOG_NOTICE, "main")
	if err != nil {
		fmt.Println(err)
	}
}
