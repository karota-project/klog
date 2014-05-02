package klog

import (
    "fmt"
	"bytes"
    "log"
    "os/exec"
    "strconv"
    "strings"
    "os"
    //"time"
)

type SysInfo struct {
    mem_used int
    mem_free int
    cpu_used float64
}

/*
* exec vmstat command
*/
func getSystemInfo() []*SysInfo{

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

    }

    return sysInfo;
}

/*
* to convert a float number to a string
*/
func floattostr(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'g', 1, 64)
}

/*
* Printfile for linux 
*/
func Printfile(func_n string , file_n string) bool {
    //t := time.Now()
    fout, err := os.Create(file_n)
    if err != nil {
        fmt.Println(file_n, err)
        return false
    }

    sysInfo := getSystemInfo()

    for _, s := range(sysInfo) {
        str := []string{"[", func_n , "] ,mem-used : ", strconv.Itoa(s.mem_used) , ",mem-free : " , strconv.Itoa(s.mem_free) , ",cpu-used : ", floattostr(s.cpu_used),"\n"}
        strjoin :=  strings.Join(str, "")
        fout.WriteString(strjoin)
    }

    return true;
}

/*
* Printlog for linux 
*/
func Printlog(func_n string) bool {
    sysInfo := getSystemInfo()
    
    for _, s := range(sysInfo) {
        log.Println("[", func_n, "] ,mem-used : ", s.mem_used, ",mem-free : ", s.mem_free ,",cpu-used : ", s.cpu_used )
    }

    return true;
}