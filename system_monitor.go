package log

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Process struct {
	R int `json:"r"`
	B int `json:"b"`
}

type Memory struct {
	Swapd  int `json:"swapd"`
	Free   int `json:"free"`
	Buffer int `json:"buffer"`
	Cache  int `json:"cache"`
}

type Swap struct {
	Si int `json:"si"`
	So int `json:"so"`
}

type Io struct {
	Bi int `json:"bi"`
	Bo int `json:"bo"`
}

type System struct {
	In int `json:"in"`
	Cs int `json:"cs"`
}

type Cpu struct {
	Us int `json:"us"`
	Sy int `json:"sy"`
	Id int `json:"id"`
	Wa int `json:"wa"`
}

type SysInfo struct {
	Process Process `json:"process"`
	Memory  Memory  `json:"memory"`
	Swap    Swap    `json:"swap"`
	Io      Io      `json:"io"`
	System  System  `json:"system"`
	Cpu     Cpu     `json:"cpu"`
}

const (
	_             = iota
	PROCESS_R     = iota
	PROCESS_B     = iota
	MEMORY_SWAPD  = iota
	MEMORY_FREE   = iota
	MEMORY_BUFFER = iota
	MEMORY_CACHE  = iota
	SWAP_SI       = iota
	SWAP_SO       = iota
	IO_BI         = iota
	IO_BO         = iota
	SYSTEM_IN     = iota
	SYSTEM_CS     = iota
	CPU_US        = iota
	CPU_SY        = iota
	CPU_ID        = iota
	CPU_WA        = iota
)

// exec vmstat command
func getSystemInfo() (sysInfo []SysInfo, err error) {
	cmd := exec.Command("vmstat")
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	assigned := regexp.MustCompile("\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)")

	sysInfo = make([]SysInfo, 0)

	for {
		line, err := stdout.ReadString('\n')
		if err != nil {
			break
		}

		lineByte := []byte(line)
		if !assigned.Match(lineByte) {
			continue
		}

		groups := make([]int, 0)

		for _, g := range assigned.FindSubmatch(lineByte) {
			num, err := strconv.Atoi(string(g))

			if err != nil {
				num = 0
			}
			groups = append(groups, num)
		}

		sysInfo = append(sysInfo, SysInfo{
			Process: Process{
				groups[PROCESS_R],
				groups[PROCESS_B],
			},
			Memory: Memory{
				groups[MEMORY_SWAPD],
				groups[MEMORY_FREE],
				groups[MEMORY_BUFFER],
				groups[MEMORY_CACHE],
			},
			Swap: Swap{
				groups[SWAP_SI],
				groups[SWAP_SO],
			},
			Io: Io{
				groups[IO_BI],
				groups[IO_BO],
			},
			System: System{
				groups[SYSTEM_IN],
				groups[SYSTEM_CS],
			},
			Cpu: Cpu{
				groups[CPU_US],
				groups[CPU_SY],
				groups[CPU_ID],
				groups[CPU_WA],
			},
		})
	}

	return sysInfo, err
}
