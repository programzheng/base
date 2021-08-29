package bot

import (
	"base/pkg/helper"
	"strings"
	"testing"

	underscore "github.com/ahl5esoft/golang-underscore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	if len(results) != 1 {
		t.Log("get results success")
	} else {
		t.Errorf("fail")
	}
}

func TestGetDistinctByUserID(t *testing.T) {
	lb := LineBilling{}
	where := make(map[string]interface{})
	not := make(map[string]interface{})
	results, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if len(results) != 1 {
		t.Log("get results success")
	} else {
		t.Errorf("fail")
	}
	dst := make([]LineBilling, 0)
	underscore.Chain(results).DistinctBy("UserID").Value(&dst)
	var testUserID string
	for _, v := range dst {
		if testUserID == v.UserID {
			t.Errorf("fail")
		}
		testUserID = v.UserID
	}
	t.Log("get distinct by UserID success")
}

func TestGetDistinctByUserIDAndLineMember(t *testing.T) {
	lb := LineBilling{}
	where := make(map[string]interface{})
	not := make(map[string]interface{})
	lbs, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if len(lbs) != 1 {
		t.Log("get results success")
	} else {
		t.Errorf("fail")
	}
	dstByUserID := make(map[string]string, 0)
	underscore.Chain(lbs).DistinctBy("UserID").SelectMany(func(lb LineBilling, _ int) map[string]string {
		lineMember, err := botClient.GetGroupMemberProfile(lb.GroupID, lb.UserID).Do()
		if err != nil {
			log.Fatal("line messaging api get group member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
		}
		dst := make(map[string]string)
		dst[lb.UserID] = lineMember.DisplayName
		return dst
	}).Value(&dstByUserID)
	t.Log(dstByUserID)
}

func TestGetLineBillingList(t *testing.T) {
	var testGroupID string
	testGroupID = viper.Get("TEST_GET_LINE_BILLING_LIST_GROUP_ID").(string)
	lb := LineBilling{}
	where := make(map[string]interface{})
	not := make(map[string]interface{})
	lbs, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if len(lbs) != 1 {
		t.Log("get results success")
	} else {
		t.Errorf("fail")
	}
	dstByUserID := make(map[string]string, 0)
	underscore.Chain(lbs).DistinctBy("UserID").SelectMany(func(lb LineBilling, _ int) map[string]string {
		lineMember, err := botClient.GetGroupMemberProfile(lb.GroupID, lb.UserID).Do()
		if err != nil {
			log.Fatal("line messaging api get group member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
		}
		dst := make(map[string]string)
		dst[lb.UserID] = lineMember.DisplayName
		return dst
	}).Value(&dstByUserID)
	var sb strings.Builder
	sb.Grow(len(lbs))
	for _, lb := range lbs {
		memberName := "Unknow"
		//check line member display name is exist
		if _, ok := dstByUserID[lb.UserID]; ok {
			memberName = dstByUserID[lb.UserID]
		}
		amountAvg, amountAvgBase := calculateAmount(testGroupID, helper.ConvertToFloat64(lb.Billing.Amount))
		text := lb.Billing.CreatedAt.Format(helper.Yyyymmddhhmmss) + " " +
			lb.Billing.Title + "|" + helper.ConvertToString(lb.Billing.Amount) + "/" + helper.ConvertToString(amountAvgBase) + " = " + helper.ConvertToString(amountAvg) + " |" + memberName + "|" + lb.Billing.Note + "\n"
		sb.WriteString(text)
	}
	t.Log(sb.String())
}

func TestPluck(t *testing.T) {
	lb := LineBilling{}
	where := make(map[string]interface{})
	not := make(map[string]interface{})
	results, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get pluck failed: %v", err)
	}
	result := helper.Pluck(results, "UserID")
	t.Logf("%+v\n", result)
}

func TestPluckSet(t *testing.T) {
	lb := LineBilling{}
	where := make(map[string]interface{})
	not := make(map[string]interface{})
	results, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get pluck failed: %v", err)
	}
	pluck := helper.Pluck(results, "UserID")
	set := helper.SliceToSet(pluck)
	t.Logf("%+v\n", set)
}
