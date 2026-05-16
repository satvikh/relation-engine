package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

type marketState struct {
	symbol string
	price  float64
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	markets := []marketState{
		{symbol: "BTC-USD", price: 68500.00},
		{symbol: "ETH-USD", price: 3250.00},
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	defer signal.Stop(signals)

	emitAll(markets)

	for {
		select {
		case tickTime := <-ticker.C:
			for i := range markets {
				markets[i].price = nextPrice(markets[i].price, rng)
				emitTick(markets[i], tickTime)
			}
		case <-signals:
			fmt.Fprintln(os.Stderr, "shutting down price stream")
			return
		}
	}
}

func emitAll(markets []marketState) {
	now := time.Now().UTC()
	for _, market := range markets {
		emitTick(market, now)
	}
}

func emitTick(market marketState, eventTime time.Time) {
	payload, err := json.Marshal(CryptoTickerPayload{
		Price: market.price,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal payload for %s: %v\n", market.symbol, err)
		return
	}

	event := RawEvent{
		EventID:    fmt.Sprintf("%s-%d", market.symbol, eventTime.UnixNano()),
		Domain:     "crypto",
		EventType:  "trade_tick",
		Source:     "simulator",
		EntityID:   market.symbol,
		EventTime:  eventTime.UTC(),
		IngestedAt: time.Now().UTC(),
		Payload:    payload,
	}

	encoded, err := json.Marshal(event)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal event for %s: %v\n", market.symbol, err)
		return
	}

	fmt.Println(string(encoded))
}

func nextPrice(current float64, rng *rand.Rand) float64 {
	changePct := (rng.Float64() - 0.5) * 0.01
	return current * (1 + changePct)
}
