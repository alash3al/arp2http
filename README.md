ARP2HTTP
=========
A tiny daemon that will listen on `ARP` packets and detect whether the registered devices are connected/disconnected, then notify the specified `webhook`.

Usage
======

`arp2http --iface=wlan0 --webhook=http://some.web/hook --verbose`
