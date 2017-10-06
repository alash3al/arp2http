package main

import (
	"log"
	"net"
	"sync"
	"time"
)

import (
	"github.com/mdlayher/arp"
	"github.com/paulstuart/ping"
)

var (
	// Here is our ARP table that will be
	// monitored.
	ARPTable = new(sync.Map)
)

// Here we watch for all incoming ARP packets
// and analyze it to extract our required information.
func ArpWatcher() {
	iface, err := net.InterfaceByName(*IFACE)
	if err != nil {
		log.Fatal(err)
	}
	client, err := arp.Dial(iface)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	for {
		packet, _, err := client.Read()
		if err != nil {
			continue
		}
		if _, found := ARPTable.Load(packet.SenderHardwareAddr.String()); found {
			continue
		}
		ARPTable.Store(packet.SenderHardwareAddr.String(), packet.SenderIP.String())
		go TriggerEventConnected(packet.SenderHardwareAddr.String(), packet.SenderIP.String())
		go IPWatcher(packet.SenderHardwareAddr.String(), packet.SenderIP.String())
	}
}

// Here we watch the specified ip,
// to detect whether it is disconnected or not,
// if it were disconnected, then remove it from our ARP table.
func IPWatcher(mac, ip string) {
	for ping.Ping(ip, 1) {
		time.Sleep(time.Duration(*TIMEOUT) * time.Second)
	}
	ARPTable.Delete(mac)
	TriggerEventDisconnected(mac, ip)
}
