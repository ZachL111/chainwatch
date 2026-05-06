package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 76, Capacity: 87, Latency: 18, Risk: 16, Weight: 8}, wantScore: 155, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 89, Capacity: 88, Latency: 27, Risk: 22, Weight: 10}, wantScore: 142, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 66, Capacity: 74, Latency: 27, Risk: 19, Weight: 7}, wantScore: 85, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
