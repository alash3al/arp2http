// arp2http
// a small utilty that listenes on arp packets and catch the newly (dis)connected
// devices on the network and notify another endpoint (webhook).
package main

// KISS ;)
func main() {
	ParseFlags()
	ArpWatcher()
}
