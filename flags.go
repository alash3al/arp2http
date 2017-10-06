package main

import "flag"

var (
	IFACE   = flag.String("iface", "lo", "the local interface to listen on")
	WEBHOOK = flag.String("webhook", "http://localhost/", "the webhook that will receive the notifications")
	TIMEOUT = flag.Int("timeout", 5, "the timeout used to detect that a device is disconnected")
	VERBOSE = flag.Bool("verbose", false, "whether to verbose the internal states or not")
)

// no comment :D
func ParseFlags() {
	flag.Parse()
}
