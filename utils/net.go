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
	fields := strings.Fields(str)
	rBytes, _ := strconv.Atoi(fields[1])
	rPackets, _ := strconv.Atoi(fields[2])
	rDrop, _ := strconv.Atoi(fields[4])
	tBytes, _ := strconv.Atoi(fields[9])
	tPackets, _ := strconv.Atoi(fields[10])
	tDrop, _ := strconv.Atoi(fields[12])
	net = NetInfo{
		Name: strings.Replace(fields[0], ":", "", 1),
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
		nets = append(nets, getLineInfo(lines[i]))
	}
	return nets
}
