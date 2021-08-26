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
	t.Log(lineBilling)
}

func TestGet(t *testing.T) {
	lb := LineBilling{}
	where := make(map[string]interface{})
	not := make(map[string]interface{})
	results, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	t.Logf("%+v\n", results)
}
