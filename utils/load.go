package utils

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

type LoadAvgInfo struct {
	Load1  float64
	Load5  float64
	Load15 float64
}

func GetLoadAvg() (loadavg LoadAvgInfo) {
	path := "/proc/loadavg"
	if runtime.GOOS != "linux" {
		path = "./data/loadavg.txt"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return loadavg
	}
	str := strings.Trim(string(file), "\n")
	arr := strings.Split(str, " ")
	load1, _ := strconv.ParseFloat(arr[0], 64)
	load5, _ := strconv.ParseFloat(arr[1], 64)
	load15, _ := strconv.ParseFloat(arr[2], 64)
	loadavg = LoadAvgInfo{
		Load1:  load1,
		Load5:  load5,
		Load15: load15,
	}
	return loadavg
}
