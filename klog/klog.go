package klog

import (
    "fmt"
	"bytes"
    "log"
    "os/exec"
    "strconv"
    "strings"
    //"time"
)

type SysInfo struct {
    mem_used int
    mem_free int
    cpu_used float64
}

/*
* Printlog for linux 
*/
func Printlog(fname string) int {
	cmd := exec.Command("vmstat")
    var out bytes.Buffer
    var stderr bytes.Buffer

    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
       fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    }

   sysInfo := make([]*SysInfo, 0)
   for {
        line, err := out.ReadString('\n')
        if err!=nil {
            break;
        }
        /* split string into array */
        ft := make([]string, 0)
        tokens := strings.Split(line, " ")
        for _, t := range(tokens) {
            if t!="" && t!="\t" {
                ft = append(ft, t)
            }
        }
        //log.Println(len(ft), ft)

        /* mem_used : swapd + buffer + cached */
        swap, err := strconv.Atoi(ft[2])
        buf, err := strconv.Atoi(ft[4])
        cach, err := strconv.Atoi(ft[5])
        mem_used := swap + buf + cach;
        if err!=nil {
            continue
        }

        mem_free, err := strconv.Atoi(ft[3])
        if err!=nil {
            continue
        }

        cpu_used, err := strconv.ParseFloat(ft[12], 64)
        if err!=nil {
            continue
        }

        sysInfo = append(sysInfo, &SysInfo{mem_used, mem_free, cpu_used})
        for _, s := range(sysInfo) {
            log.Println("[", fname, "] ,mem-used : ", s.mem_used, ",mem-free : ", s.mem_free ,",cpu-used : ", s.cpu_used )
        }
    }

   return 0;
}