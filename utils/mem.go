package utils

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// MemUsed = MemTotal + Shmem - MemFree - Buffers - Cached - SReclaimable
type MemInfo struct {
	MemTotal     int
	Shmem        int
	MemFree      int
	Buffers      int
	Cached       int
	SReclaimable int
}

func GetMemInfo() (mem MemInfo) {
	path := "/proc/meminfo"
	if runtime.GOOS != "linux" {
		path = "./data/meminfo.txt"
	}
	file, err := os.Open(path)
	if err != nil {
		return mem
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fields := strings.Fields(string(line))
		number, _ := strconv.Atoi(fields[1])
		switch fields[0] {
		case "MemTotal:":
			mem.MemTotal = number
		case "Shmem:":
			mem.Shmem = number
		case "MemFree:":
			mem.MemFree = number
		case "Buffers:":
			mem.Buffers = number
		case "Cached:":
			mem.Cached = number
		case "SReclaimable:":
			mem.SReclaimable = number
		}
	}
	return mem
}
