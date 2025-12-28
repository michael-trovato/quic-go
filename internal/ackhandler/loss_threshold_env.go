package ackhandler

import (
    "os"
    "strconv"
    "sync"

    "github.com/quic-go/quic-go/internal/protocol"
)

const (
    envLossTimeThreshold   = "QUIC_GO_LOSS_TIME_THRESHOLD"   // float
    envLossPacketThreshold = "QUIC_GO_LOSS_PACKET_THRESHOLD" // int
)

var once sync.Once

func initLossThresholdsFromEnv() {
    once.Do(func() {
        if v := os.Getenv(envLossTimeThreshold); v != "" {
            if f, err := strconv.ParseFloat(v, 64); err == nil {
                if f < 1.0 { f = 1.0 }
                if f > 10.0 { f = 10.0 }
                timeThreshold = f
            }
        }

        if v := os.Getenv(envLossPacketThreshold); v != "" {
            if n, err := strconv.Atoi(v); err == nil {
                if n < 3 { n = 3 }
                if n > 1024 { n = 1024 }
                packetThreshold = protocol.PacketNumber(n)
            }
        }
    })
}
