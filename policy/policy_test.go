package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 76, Capacity: 87, Latency: 18, Risk: 16, Weight: 8}
	if got := Score(signal); got != 155 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 89, Capacity: 88, Latency: 27, Risk: 22, Weight: 10}
	if got := Score(signal); got != 142 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 66, Capacity: 74, Latency: 27, Risk: 19, Weight: 7}
	if got := Score(signal); got != 85 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
