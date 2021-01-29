package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-ping/ping"
)

func getLatency() int64 {
	pinger, err := ping.NewPinger("google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	fmt.Printf("\nAverage Latency: %v \nStandard Deviation: %v\n", stats.AvgRtt, stats.StdDevRtt)

	return stats.AvgRtt.Microseconds()
}

func determineLatency() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter desired latency to achieve in milliseconds: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	desiredLatency, err := strconv.Atoi(text)

	if err != nil {
		panic("Invalid latency")
	}

	fmt.Println("Finding your current latency.")
	currentLatency := getLatency()

	return int(float64(desiredLatency) - float64(currentLatency)/1000.0)
}

func windowsAddLatency() {
	// Implement later
}

func darwinAddLatency() {
	// Implement Later
}

func linuxAddLatency() {
	// Use tc command locally installed.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter name of network adapter: ")
	text, _ := reader.ReadString('\n')
	netAdapter := strings.TrimSpace(text)
	//command := fmt.Sprintf("qdisc add dev " + netAdapter + " root netem delay " + fmt.Sprint(determineLatency()) + "ms")

	cmd := exec.Command("/usr/sbin/tc", "qdisc", "add", "dev", netAdapter, "root", "netem", "delay", fmt.Sprint(determineLatency())+"ms")
	err := cmd.Run()

	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	var os string

	// Runs the correct add latency command for the platform
	// depending on the build flags.
	if os == "windows" {
		windowsAddLatency()
	} else if os == "darwin" {
		darwinAddLatency()
	} else {
		linuxAddLatency()
	}
}
