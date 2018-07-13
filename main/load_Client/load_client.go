package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var url string
var rps string

func ProcessFlags() {
	flag.StringVar(&url, "url", "http://localhost:8080/?num=9000000031", "url with http, with port")
	flag.StringVar(&rps, "rps", "200", "количество запросов в секунду.")
	flag.Parse()
}

func doGet() {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.Status != "200 OK" {
		fmt.Print(response.Status)
	}
}

func load(wait int) {
	for {
		go doGet()
		time.Sleep(time.Duration(wait) * time.Millisecond)
		fmt.Print(".")
	}
}

func init() {
	ProcessFlags()
}

func CalcAverageAnswerTime() {
	t1 := time.Now()
	for i := 0; i < 1000; i++ {
		doGet()
	}
	fmt.Println("1000 requests =", time.Now().Sub(t1))
}

func main() {
	CalcAverageAnswerTime()
	fmt.Println("example usage http_client.exe -rps=1 -url=http://localhost:8080/?num=9000000031")
	fmt.Println("start with flags url=", url, "rps=", rps)
	fmt.Println("cpu count=", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	//SendPost()
	wait_ms, err := strconv.Atoi(rps)
	//fmt.Println("wait_ms=", wait_ms)
	if err != nil {
		fmt.Println("error converting rps to integer")
		panic(err)
	}
	wait := 1000 / wait_ms
	//fmt.Println(wait)
	go load(wait)
	fmt.Println("prees any key for stop generating")
	fmt.Scanln()
}
