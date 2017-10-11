ARP2HTTP
=========
A tiny daemon that will listen on `ARP` packets and detect whether the registered devices are connected/disconnected, then notify the specified `webhook`.

Why ?
=====
- I wanted to know more about networking and its layers.
- I need to know when a device is on/off and then notify another endpoint, so I can manage and do what I want.

Usage
======

`arp2http --iface=wlan0 --webhook=http://some.web/hook --verbose`

Download
=========
Goto [Releases Page](https://github.com/alash3al/arp2http/releases/tag/v1.0.0) and select your platform.

Compile From Source ?
======================
1. You must have `Golang` installed
2. `go get github.com/alash3al/arp2http`
3. have fun as a gopher ;)
