package main

import (
	//"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	//"strconv"
	//"strings"
	"github.com/go-ping/ping"
)

func getLatency() int64 {
	pinger, err := ping.NewPinger("google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 5

	fmt.Println("Finding your current latency.")
	pinger.SetPrivileged(true)
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()

	return stats.AvgRtt.Microseconds()
}

func determineLatency(desiredLatency string) int {
	// Latency in milliseconds
	latency, _ := strconv.Atoi(desiredLatency)

	currentLatency := getLatency()

	fmt.Println("Found latency")
	latencyToAdd := int(float64(latency) - float64(currentLatency)/1000.0)

	if latencyToAdd > 0 {
		return latencyToAdd
	}

	return 0
}

func linuxAddLatency(latency string, jitter string, loss string) {
	//command := fmt.Sprintf("qdisc add dev " + netAdapter + " root netem delay " + fmt.Sprint(determineLatency()) + "ms")
	cmd := exec.Command("/usr/sbin/tc", "qdisc", "add", "dev", "enp10s0", "root", "netem", "delay", fmt.Sprint(determineLatency(latency))+"ms", jitter+"ms", loss+"%")
	fmt.Println(cmd)
	err := cmd.Run()

	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	latency := string(os.Args[1])
	jitter := string(os.Args[2])
	loss := string(os.Args[3])
	linuxAddLatency(latency, jitter, loss)
}
