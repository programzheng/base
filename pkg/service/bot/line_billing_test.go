package bot

import (
	"testing"
)

func TestAdd(t *testing.T) {
	lb := LineBilling{
		BillingID: 1,
		GroupID:   "test",
		UserID:    "test",
	}
	lineBilling, err := lb.Add()
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	t.Log(lineBilling)
}

func TestGet(t *testing.T) {
	lbs := LineBilling{}
	results, err := lbs.Get()
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	t.Log(results)
}
