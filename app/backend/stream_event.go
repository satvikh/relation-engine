package main

import (
	"encoding/json"
	"time"
)

// RawEvent represents a general ingested data stream event before payload-specific decoding.
type RawEvent struct {
	EventID    string          `json:"event_id"`
	Domain     string          `json:"domain"`     // "crypto"
	EventType  string          `json:"event_type"` // "trade_tick"
	Source     string          `json:"source"`     // "coinbase"
	EntityID   string          `json:"entity_id"`  // "BTC-USD"
	EventTime  time.Time       `json:"event_time"`
	IngestedAt time.Time       `json:"ingested_at"`
	Payload    json.RawMessage `json:"payload"`
}

// CryptoTickerPayload represents a decoded crypto ticker payload.
type CryptoTickerPayload struct {
	Price float64 `json:"price"`
}



