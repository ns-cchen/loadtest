package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {

	targeter1 := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:4567/search/apiv1/search?starttime=0&endtime=1714010845&sync=1&query=&limit=10&offset=0&tenantid=1016&trid=cchentest&explain=1&eventtype=alert",
	})
	targeter2 := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:4567/search/apiv1/search?starttime=0&endtime=1714010845&sync=1&query=&limit=10&offset=0&tenantid=1016&trid=cchentest&explain=1&eventtype=network",
	})
	targeter3 := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:4567/search/apiv1/search?starttime=0&endtime=1714010845&sync=1&query=&limit=10&offset=0&tenantid=1016&trid=cchentest&explain=1&eventtype=audit",
	})
	targeter4 := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:4567/search/apiv1/search?starttime=0&endtime=1714010845&sync=1&query=&limit=10&offset=0&tenantid=1016&trid=cchentest&explain=1&eventtype=incident",
	})
	targeter5 := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:4567/search/apiv1/search?starttime=0&endtime=1714010845&sync=1&query=&limit=10&offset=0&tenantid=1016&trid=cchentest&explain=1&collection=watchlist_events",
	})
	targeters := []vegeta.Targeter{targeter1, targeter2, targeter3, targeter4, targeter5}

	attack(500, 60*time.Second, targeters)
	attack(1000, 60*time.Second, targeters)
	attack(2000, 60*time.Second, targeters)
}

func attack(frequency int, duration time.Duration, targeters []vegeta.Targeter) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics

	randomIndex := r.Intn(len(targeters))
	for result := range attacker.Attack(targeters[randomIndex], rate, duration, "Big Bang!") {
		metrics.Add(result)
	}

	metrics.Close()

	marshal, err := json.Marshal(metrics.Latencies)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Duration: %s\n", duration)
	fmt.Printf("Rate: %f (nanosecond)\n", metrics.Rate)
	fmt.Printf("Latencies: %s (nanosecond)\n", string(marshal))
}
