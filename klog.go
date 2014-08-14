package klog

import (
	"encoding/json"
	"log"
	"log/syslog"
	"os"
	"runtime"
	"text/template"
	"time"
)

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

type Priority int

type JsonStruct struct {
	Func    string `json:"func"`
	MemUsed int    `json:"mem_used"`
	MemFree int    `json:"mem_free"`
	CpuUsed int    `json:"cpu_used"`
}

// Printlog for linux
func Stdout(functionName string) (err error) {
	_, file, line, _ := runtime.Caller(1)

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		v, err := json.Marshal(&JsonStruct{
			Func:    functionName,
			MemUsed: s.Memory.Swapd + s.Memory.Buffer + s.Memory.Cache,
			MemFree: s.Memory.Free,
			CpuUsed: s.Cpu.Us,
		})
		if err != nil {
			return err
		}

		t, err := template.New("template").Parse("{{.File}} (line {{.Line}}) {{.Json}}\n")
		if err != nil {
			return err
		}

		t.Execute(os.Stdout, struct {
			File string
			Line int
			Json string
		}{file, line, string(v)})
	}

	return nil
}

// Write file
func WriteFile(functionName string, outputfile string) (err error) {
	_, file, line, _ := runtime.Caller(1)

	f, err := os.OpenFile(outputfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		v, err := json.Marshal(&JsonStruct{
			Func:    functionName,
			MemUsed: s.Memory.Swapd + s.Memory.Buffer + s.Memory.Cache,
			MemFree: s.Memory.Free,
			CpuUsed: s.Cpu.Us,
		})
		if err != nil {
			return err
		}

		t, err := template.New("template").Parse("[{{.Date}}] {{.File}} (line {{.Line}}) {{.Json}}\n")
		if err != nil {
			return err
		}

		t.Execute(f, struct {
			Date time.Time
			File string
			Line int
			Json string
		}{time.Now(), file, line, string(v)})
	}

	defer f.Close()

	return nil
}

// Print syslog for unix
func Syslog(priority Priority, facility string) (err error) {
	p := convToSyslogPriority(priority)

	// Configure logger to write to the syslog. You could do this in init(), too.
	logWriter, err := syslog.New(p, facility)
	if err != nil {
		return err
	}

	logger := log.New(logWriter, "", log.LstdFlags)

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		v, err := json.Marshal(&JsonStruct{
			Func:    "",
			MemUsed: s.Memory.Swapd + s.Memory.Buffer + s.Memory.Cache,
			MemFree: s.Memory.Free,
			CpuUsed: s.Cpu.Us,
		})

		if err != nil {
			return err
		}

		logger.Print(string(v))
	}

	return nil
}

func convToSyslogPriority(p Priority) (priority syslog.Priority) {
	switch p {
	case LOG_EMERG:
		return syslog.LOG_EMERG

	case LOG_ALERT:
		return syslog.LOG_ALERT

	case LOG_CRIT:
		return syslog.LOG_CRIT

	case LOG_ERR:
		return syslog.LOG_ERR

	case LOG_WARNING:
		return syslog.LOG_WARNING

	case LOG_NOTICE:
		return syslog.LOG_NOTICE

	case LOG_INFO:
		return syslog.LOG_INFO

	case LOG_DEBUG:
		return syslog.LOG_DEBUG

	default:
		return syslog.LOG_DEBUG
	}
}
