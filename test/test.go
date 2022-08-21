package main

import (
	"fmt"
	"sia/utils"
)

func main() {
	fmt.Println(utils.GetProcessList())
	fmt.Println(utils.GetCpuInfo())
	fmt.Println(utils.GetMemInfo())
	fmt.Println(utils.GetNetInfo())
	fmt.Println(utils.GetLoadAvg())
}
