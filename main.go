package main

// go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git

import (
	"arpgate"
	"flag"
	"fmt"
)

var (
	isVersion bool
	isHelp    bool
	peers     []string
)

func init() {
	flag.BoolVar(&isHelp, "h", false, "show help")
	flag.BoolVar(&isVersion, "v", false, "show version")
}

func main() {
	flag.Parse()
	peers = flag.Args()

	if isVersion {
		fmt.Println(arpgate.VERSION)
		return
	}

	if isHelp {
		fmt.Println("-h help")
		fmt.Println("-v versionsInit")
		return
	}

	arpgate.Start()
}
