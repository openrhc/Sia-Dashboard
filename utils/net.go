package utils

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Traffic struct {
	Receive  TrafficItem
	Transmit TrafficItem
}

type TrafficItem struct {
	Bytes   int
	Packets int
	Drop    int
}

type NetInfo struct {
	Name    string
	Traffic Traffic
}

func getLineInfo(str string) (net NetInfo) {
	arr1 := strings.Split(strings.Trim(str, " "), ":")
	arr2 := strings.Split(arr1[1], " ")
	var arr3 []string
	for _, item := range arr2 {
		if item != "" {
			arr3 = append(arr3, item)
		}
	}
	rBytes, _ := strconv.Atoi(arr3[0])
	rPackets, _ := strconv.Atoi(arr3[1])
	rDrop, _ := strconv.Atoi(arr3[3])
	tBytes, _ := strconv.Atoi(arr3[8])
	tPackets, _ := strconv.Atoi(arr3[9])
	tDrop, _ := strconv.Atoi(arr3[11])
	net = NetInfo{
		Name: arr1[0],
		Traffic: Traffic{
			Receive: TrafficItem{
				Bytes:   rBytes,
				Packets: rPackets,
				Drop:    rDrop,
			},
			Transmit: TrafficItem{
				Bytes:   tBytes,
				Packets: tPackets,
				Drop:    tDrop,
			},
		},
	}
	return net
}

func GetNetInfo() (nets []NetInfo) {
	path := "/proc/net/dev"
	if runtime.GOOS != "linux" {
		path = "./data/net.txt"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nets
	}
	lines := strings.Split(strings.Trim(string(file), "\n"), "\n")
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		nets = append(nets, getLineInfo(line))
	}
	return nets
}
