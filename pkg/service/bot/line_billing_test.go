package bot

import (
	"base/pkg/helper"
	"fmt"
	"strings"
	"testing"

	underscore "github.com/ahl5esoft/golang-underscore"
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
		t.Logf("Add failed: %v", err)
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
		t.Error("fail")
	}
	dstByUserID := make(map[string]string, 0)
	underscore.Chain(lbs).DistinctBy("UserID").SelectMany(func(lb LineBilling, _ int) map[string]string {
		dst := make(map[string]string)
		lineMember, err := botClient.GetGroupMemberProfile(lb.GroupID, lb.UserID).Do()
		if err != nil {
			t.Log("line messaging api get group member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
			dst[lb.UserID] = "Unkonw"
			return dst
		}
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
		dst := make(map[string]string)
		lineMember, err := botClient.GetGroupMemberProfile(lb.GroupID, lb.UserID).Do()
		if err != nil {
			t.Log("line messaging api get group member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
			dst[lb.UserID] = "Unkonw"
			return dst
		}
		dst[lb.UserID] = lineMember.DisplayName
		return dst
	}).Value(&dstByUserID)
	lbUserIDAmount := make(map[string]float64, 0)
	var sb strings.Builder
	sb.Grow(len(lbs))
	for _, lb := range lbs {
		memberName := "Unknow"
		amountAvg, amountAvgBase := calculateAmount(testGroupID, helper.ConvertToFloat64(lb.Billing.Amount))
		//check line member display name is exist
		if _, ok := dstByUserID[lb.UserID]; ok {
			memberName = dstByUserID[lb.UserID]
			lbUserIDAmount[lb.UserID] = lbUserIDAmount[lb.UserID] + amountAvg
		}
		text := lb.Billing.CreatedAt.Format(helper.Yyyymmddhhmmss) + " " +
			lb.Billing.Title + "|" + helper.ConvertToString(lb.Billing.Amount) + "/" + helper.ConvertToString(amountAvgBase) + " = " + helper.ConvertToString(amountAvg) + " |" + memberName + "|" + lb.Billing.Note + "\n"
		sb.WriteString(text)
	}
	t.Logf("%v\n", lbUserIDAmount)
	t.Log(sb.String())
}

func TestGetLineBillingListTemplateText(t *testing.T) {
	var testGroupID string
	testGroupID = viper.Get("TEST_GET_LINE_BILLING_LIST_GROUP_ID").(string)
	lb := LineBilling{}
	where := make(map[string]interface{})
	where["group_id"] = testGroupID
	not := make(map[string]interface{})
	lbs, err := lb.Get(where, not)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if len(lbs) != 1 {
		t.Logf("%v\n", lbs)
	} else {
		t.Errorf("fail")
	}
	//user id line member display name
	dstByUserID := make(map[string]string, 0)
	//user id total amount
	lbUserIDAmount := make(map[string]float64, 0)
	underscore.Chain(lbs).DistinctBy("UserID").SelectMany(func(lb LineBilling, _ int) map[string]string {
		dst := make(map[string]string)
		lineMember, err := botClient.GetGroupMemberProfile(lb.GroupID, lb.UserID).Do()
		if err != nil {
			t.Log("line messaging api get group member profile group id:"+lb.GroupID+" user id:"+lb.UserID+" error:", err)
			dst[lb.UserID] = "Unkonw"
			return dst
		}
		dst[lb.UserID] = lineMember.DisplayName
		return dst
	}).Value(&dstByUserID)
	var sbList strings.Builder
	sbList.Grow(len(lbs))
	for _, lb := range lbs {
		var memberName string
		amountAvg, amountAvgBase := calculateAmount(testGroupID, helper.ConvertToFloat64(lb.Billing.Amount))
		//check line member display name is exist
		if _, ok := dstByUserID[lb.UserID]; ok {
			memberName = dstByUserID[lb.UserID]
			lbUserIDAmount[lb.UserID] = lbUserIDAmount[lb.UserID] + amountAvg
		}
		text := lb.Billing.CreatedAt.Format(helper.Yyyymmddhhmmss) + " " +
			lb.Billing.Title + "|" + helper.ConvertToString(lb.Billing.Amount) + "/" + helper.ConvertToString(amountAvgBase) + " = " + helper.ConvertToString(amountAvg) + " |" + memberName + "|" + lb.Billing.Note + "\n"
		sbList.WriteString(text)
	}
	//billing list string
	t.Log(sbList.String())
	var sbTotal strings.Builder
	sbTotal.Grow(len(dstByUserID))
	text := "總付款金額：\n"
	sbTotal.WriteString(text)
	for userID, name := range dstByUserID {
		text = fmt.Sprintf("%v: *%v*\n", name, helper.ConvertToString(lbUserIDAmount[userID]))
		sbTotal.WriteString(text)
	}
	//user id total string
	t.Logf(sbTotal.String())
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
