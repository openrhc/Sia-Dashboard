package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sia/utils"
	"strings"
)

type Status struct {
	Cpu     utils.CpuInfo
	Mem     utils.MemInfo
	Net     []utils.NetInfo
	Loadavg utils.LoadAvgInfo
}

//go:embed static/*
var f embed.FS
var ADDR = flag.String("a", "0.0.0.0", "address")
var PORT = flag.String("p", "8000", "port")

var Version string

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		version := fmt.Sprintf(
			"Sia-Dashboard %s. Custom %s %s/%s\n%s",
			Version, runtime.Version(), runtime.GOOS, runtime.GOARCH,
			"A Simple Linux Monitoring Dashboard.",
		)
		fmt.Println(version)
		os.Exit(0)
	}
	flag.Parse()
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/cpu", handleCpuInfo)
	http.HandleFunc("/mem", handleMemInfo)
	http.HandleFunc("/net", handleNetInfo)
	http.HandleFunc("/loadavg", handleLoadavgInfo)
	http.HandleFunc("/status", handleAllInfo)
	println("System Info Api running @ http://" + *ADDR + ":" + *PORT)
	err := http.ListenAndServe(*ADDR+":"+*PORT, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Preprocessing(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RemoteAddr, r.RequestURI)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	Preprocessing(w, r)
	url, path := "static", r.URL.Path
	if path == "/" {
		url += "/index.html"
	} else {
		url += path
	}
	b, err := f.ReadFile(url)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	arr := strings.Split(url, ".")
	suffix := arr[len(arr)-1]
	header := w.Header()
	fmt.Println(suffix)
	switch suffix {
	case "css":
		header.Set("Content-Type", "text/css; charset=utf-8")
	case "js":
		header.Set("Content-Type", "text/javascript; charset=utf-8")
	case "html":
		header.Set("Content-Type", "text/html; charset=utf-8")
	case "svg":
		header.Set("Content-Type", "image/svg+xml; charset=utf-8")
	}
	w.Write(b)
}

func handleCpuInfo(w http.ResponseWriter, r *http.Request) {
	Preprocessing(w, r)
	cpuinfo := utils.GetCpuInfo()
	b, _ := json.Marshal(cpuinfo)
	w.Write(b)
}

func handleMemInfo(w http.ResponseWriter, r *http.Request) {
	Preprocessing(w, r)
	meminfo := utils.GetMemInfo()
	b, _ := json.Marshal(meminfo)
	w.Write(b)
}

func handleNetInfo(w http.ResponseWriter, r *http.Request) {
	Preprocessing(w, r)
	netinfo := utils.GetNetInfo()
	b, _ := json.Marshal(netinfo)
	w.Write(b)
}

func handleLoadavgInfo(w http.ResponseWriter, r *http.Request) {
	Preprocessing(w, r)
	loadavg := utils.GetLoadAvg()
	b, _ := json.Marshal(loadavg)
	w.Write(b)
}

func handleAllInfo(w http.ResponseWriter, r *http.Request) {
	Preprocessing(w, r)
	status := Status{
		Cpu:     utils.GetCpuInfo(),
		Mem:     utils.GetMemInfo(),
		Net:     utils.GetNetInfo(),
		Loadavg: utils.GetLoadAvg(),
	}
	b, _ := json.Marshal(status)
	w.Write(b)
}
