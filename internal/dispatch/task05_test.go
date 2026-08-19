package dispatch

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"riderguard/internal/domain"
)

func TestRuleEffectiveToBoundary(t *testing.T) {
	when := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	clk := clock.NewMock()
	clk.Set(when)
	rule := &domain.Rule{Version: 1, Name: "boundary", MatchCategory: "劳动报酬", LeadDepartment: "权益组", EffectiveFrom: when.Add(-time.Hour), EffectiveTo: when, Status: domain.RuleStatusActive}
	item := &domain.RightsCase{ID: "case-05", Category: "劳动报酬", RegisteredAt: when}
	ref, err := NewAdjudicator(clk).Adjudicate(context.Background(), item, []*domain.Rule{rule})
	if err != nil || ref.RuleVersion != 1 {
		t.Fatalf("expected rule active at boundary, ref=%+v err=%v", ref, err)
	}
}
