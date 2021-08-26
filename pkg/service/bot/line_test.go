package bot

import "testing"

func TestGetBotInfo(t *testing.T) {
	botClientInfo, err := botClient.GetBotInfo().Do()
	if err != nil {
		t.Fatal("Failed to get bot info:", err)
	}
	t.Logf("line bot get info: %v", botClientInfo)
}
