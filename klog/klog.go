package klog

import (
	"bytes"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type SysInfo struct {
	mem_used int
	mem_free int
	cpu_used float64
}

type Priority int

const (
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

func convToSyslogPriority(p Priority) (_p syslog.Priority) {

	switch {
	case p == LOG_EMERG:
		_p = syslog.LOG_EMERG
	case p == LOG_ALERT:
		_p = syslog.LOG_ALERT
	case p == LOG_CRIT:
		_p = syslog.LOG_CRIT
	case p == LOG_ERR:
		_p = syslog.LOG_ERR
	case p == LOG_WARNING:
		_p = syslog.LOG_WARNING
	case p == LOG_NOTICE:
		_p = syslog.LOG_NOTICE
	case p == LOG_INFO:
		_p = syslog.LOG_INFO
	case p == LOG_DEBUG:
		_p = syslog.LOG_DEBUG
	}
	return _p
}

/*
* exec vmstat command
 */
func getSystemInfo() (sysInfo []*SysInfo, err error) {

	cmd := exec.Command("vmstat")
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	sysInfo = make([]*SysInfo, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		/* split string into array */
		ft := make([]string, 0)
		tokens := strings.Split(line, " ")
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		//log.Println(len(ft), ft)

		/* mem_used : swapd + buffer + cached */
		swap, err := strconv.Atoi(ft[2])
		buf, err := strconv.Atoi(ft[4])
		cach, err := strconv.Atoi(ft[5])
		mem_used := swap + buf + cach
		if err != nil {
			continue
		}

		mem_free, err := strconv.Atoi(ft[3])
		if err != nil {
			continue
		}

		cpu_used, err := strconv.ParseFloat(ft[12], 64)
		if err != nil {
			continue
		}

		sysInfo = append(sysInfo, &SysInfo{mem_used, mem_free, cpu_used})

	}

	return sysInfo, err
}

/*
* to convert a float number to a string
 */
func floattostr(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'g', 1, 64)
}

/*
* Write file
 */
func WriteFile(_func string, outfile string) (err error) {
	_, file, line, _ := runtime.Caller(1)
	_line := strconv.Itoa(line)

	t := time.Now()
	str := fmt.Sprintf("%v", t)

	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		str := []string{"[", str, "] ", file, "(line", _line, ") {\"func\" : \"", _func, "\" ,\"mem_used\" : ", strconv.Itoa(s.mem_used), ", \"mem_free\" : ", strconv.Itoa(s.mem_free), ", \"cpu_used\" : ", floattostr(s.cpu_used), "}\n"}
		strjoin := strings.Join(str, "")
		f.WriteString(strjoin)
	}

	defer f.Close()

	return nil
}

/*
* Printlog for linux
 */
func Stdout(_func string) (err error) {
	_, file, line, _ := runtime.Caller(1)

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		log.Println(file, "(line", line, ") {\"func\" : \"", _func, "\", \"mem_used\" : ", s.mem_used, ",\"mem_free\" : ", s.mem_free, ",\"cpu_used\" : ", s.cpu_used, "}")
	}

	return nil
}

/*
* Print syslog for unix
 */

func Syslog(p Priority, facility string) (err error) {

	_p := convToSyslogPriority(p)

	// Configure logger to write to the syslog. You could do this in init(), too.
	logwriter, err := syslog.New(_p, facility)
	if err != nil {
		return err
	}

	log.SetOutput(logwriter)

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		str := []string{"{\"mem_used\" : ", strconv.Itoa(s.mem_used), ", \"mem_free\" : ", strconv.Itoa(s.mem_free), ", \"cpu_used\" : ", floattostr(s.cpu_used), "}\n"}
		strjoin := strings.Join(str, "")
		// $ cat /var/log/system.log | grep karota
		log.Print(strjoin)
	}

	return nil
}
