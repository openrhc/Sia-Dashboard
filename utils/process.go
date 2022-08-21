package utils

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Process struct {
	Name  string
	State string
	Pid   int
	VmRSS int
}

func GetProcessList() (processList []Process) {
	path := "/proc"
	if runtime.GOOS != "linux" {
		path = "./data"
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		return processList
	}
	for _, item := range dirs {
		name := item.Name()
		if item.IsDir() && IsDigit(name) {
			processList = append(processList, getProcessInfo(name))
		}
	}
	return processList
}

func getProcessInfo(pid string) (process Process) {
	path := "/proc/" + pid + "/status"
	if runtime.GOOS != "linux" {
		path = "./data/233/status.txt"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return process
	}
	arr := strings.Split(strings.Trim(string(file), "\n"), "\n")
	for _, line := range arr {
		fields := strings.Fields(line)
		key := fields[0]
		switch key {
		case "Name:":
			process.Name = strings.Join(fields[1:], " ")
		case "State:":
			process.State = fields[1]
		case "Pid:":
			pid, _ := strconv.Atoi(fields[1])
			process.Pid = pid
		case "VmRSS:":
			vmrss, _ := strconv.Atoi(fields[1])
			process.VmRSS = vmrss
		}
	}
	return process
}
