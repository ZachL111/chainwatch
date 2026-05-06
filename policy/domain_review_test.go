package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 66, Slack: 26, Drag: 31, Confidence: 64}
	if got := DomainReviewScore(item); got != 129 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "watch" {
		t.Fatalf("lane = %s", got)
	}
}
