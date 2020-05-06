package fastly

import (
	"testing"
)

func TestClient_GetRealtimeStats_validation(t *testing.T) {
	var err error
	_, err = testStatsClient.GetRealtimeStats(&GetRealtimeStatsInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}
}

func TestStatsClient_GetRealtimeStats(t *testing.T) {
	t.Parallel()

	var err error

	// Get
	recordRealtimeStats(t, "realtime_stats/get", func(c *RTSClient) {
		_, err = c.GetRealtimeStats(&GetRealtimeStatsInput{
			Service:   testServiceID,
			Timestamp: 0,
			Limit:     3,
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStatsClient_GetRealtimeStatsJSON(t *testing.T) {
	t.Parallel()

	var ret struct {
		RenameTimestamp uint64 `json:"Timestamp"`
	}

	var err error
	recordRealtimeStats(t, "realtime_stats/get", func(c *RTSClient) {
		err = c.GetRealtimeStatsJSON(&ret, &GetRealtimeStatsInput{
			Service:   testServiceID,
			Timestamp: 0,
			Limit:     3,
		})
	})
	if err != nil {
		t.Fatal(err)
	}

	if ret.RenameTimestamp == 0 {
		t.Fatalf("got RenameTimestamp=%d, want nonzero", ret.RenameTimestamp)
	}
}
