package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type CpuModel struct {
	Name     string
	Count    int
	Features string
}

type CpuTimes struct {
	User int
	Nice int
	Sys  int
	Idle int
}

type CpuInfo struct {
	Temp  float64
	Os    string
	Arch  string
	Model CpuModel
	Times []CpuTimes
}

func GetCpuInfo() CpuInfo {
	cpu := CpuInfo{
		Os:    runtime.GOOS,
		Arch:  runtime.GOARCH,
		Temp:  getCpuTemp(),
		Model: getCpuModel(),
		Times: getCpuTimes(),
	}
	return cpu
}

// 获取cpu温度
func getCpuTemp() (temp float64) {
	path := "/sys/class/thermal/thermal_zone0/temp"
	if runtime.GOOS != "linux" {
		path = "./data/temp.txt"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	tempStr := strings.Trim(string(file), "\n")
	temp, err = strconv.ParseFloat(tempStr, 64)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	temp /= 1000
	return temp
}

// 获取cpu型号 数量
func getCpuModel() (model CpuModel) {
	model = CpuModel{Name: "Unknow", Count: 0, Features: ""}
	path := "/proc/cpuinfo"
	if runtime.GOOS != "linux" {
		path = "./data/cpuinfo.txt"
	}
	file, err := os.Open(path)
	if err != nil {
		return model
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		a := string(line)
		if strings.Contains(a, "model name") {
			model.Count++
		}
		if model.Name == "Unknow" && strings.Contains(a, "model name") {
			arr := strings.Split(a, ":")
			if len(arr) == 2 {
				model.Name = strings.Trim(arr[1], " ")
			}
		}
		if model.Features == "" && strings.Contains(a, "Features") {
			arr := strings.Split(a, ":")
			if len(arr) == 2 {
				model.Features = strings.Trim(arr[1], " ")
			}
		}
	}
	return model
}

func getCpuTimes() (stat []CpuTimes) {
	path := "/proc/stat"
	if runtime.GOOS != "linux" {
		path = "./data/stat.txt"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return stat
	}
	arr1 := strings.Split(strings.Trim(string(file), "\n"), "\n")
	for _, line := range arr1[1:5] {
		arr2 := strings.Split(line, " ")
		user, _ := strconv.Atoi(arr2[1])
		nice, _ := strconv.Atoi(arr2[2])
		sys, _ := strconv.Atoi(arr2[3])
		idle, _ := strconv.Atoi(arr2[4])
		stat = append(stat, CpuTimes{
			User: user,
			Nice: nice,
			Sys:  sys,
			Idle: idle,
		})
	}
	return stat
}
