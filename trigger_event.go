package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// This function will call the webhook and send it the specified params.
func TriggerEvent(params map[string]interface{}) bool {
	data, err := json.Marshal(params)
	if err != nil {
		return false
	}
	resp, err := http.Post(*WEBHOOK, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return false
	}
	resp.Body.Close()
	return true
}

// This is just a shortcut
func TriggerEventConnected(mac, ip string) bool {
	verbose("[CONNECTED]", ip)
	return TriggerEvent(map[string]interface{}{
		"event": "connected",
		"source": map[string]string{
			"ip":  ip,
			"mac": mac,
		},
	})
}

// This is just a shortcut
func TriggerEventDisconnected(mac, ip string) bool {
	verbose("[DISCONNECTED]", ip)
	return TriggerEvent(map[string]interface{}{
		"event": "disconnected",
		"source": map[string]string{
			"ip":  ip,
			"mac": mac,
		},
	})
}
