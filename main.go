package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 10 * time.Second
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
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics

	randomIndex := rand.Intn(len(targeters))
	for result := range attacker.Attack(targeters[randomIndex], rate, duration, "Big Bang!") {
		metrics.Add(result)
	}

	metrics.Close()

	marshal, err := json.Marshal(metrics)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(marshal))
}
