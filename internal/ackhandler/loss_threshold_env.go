package ackhandler

import (
	"os"
	"strconv"
	"sync"
)

const (
	envLossTimeThreshold   = "QUIC_GO_LOSS_TIME_THRESHOLD"   // float, e.g. 2.0
	envLossPacketThreshold = "QUIC_GO_LOSS_PACKET_THRESHOLD" // int, e.g. 25
)

var initLossThresholdsOnce sync.Once

func initLossThresholdsFromEnv() {
	initLossThresholdsOnce.Do(func() {
		// timeThreshold is a float multiplier (default is 9/8).
		if v := os.Getenv(envLossTimeThreshold); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				// Guardrails: don't allow nonsense that would break loss detection completely.
				// Tune these clamps if you want.
				if f < 1.0 {
					f = 1.0
				}
				if f > 10.0 {
					f = 10.0
				}
				timeThreshold = f
			}
		}

		// packetThreshold is an integer (default is 3 packets).
		if v := os.Getenv(envLossPacketThreshold); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				if n < 3 {
					n = 3
				}
				if n > 1024 {
					n = 1024
				}
				packetThreshold = n
			}
		}
	})
}
