package bot

import (
	"base/pkg/helper"
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

	helper.GetJSON(lineBilling)
}
