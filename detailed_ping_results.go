package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PingResultDetails struct {
	SensorID                  int      `json:"sensor_id"`
	SensorName                string   `json:"sensor_name"`
	NetworkTestProfileID      int      `json:"network_test_profile_id"`
	NetworkTestProfileName    string   `json:"network_test_profile_name"`
	NetworkTestSuiteID        int      `json:"network_test_suite_id"`
	NetworkTestSuiteName      string   `json:"network_test_suite_name"`
	StartTime                 string   `json:"start_time"`
	FinishTime                string   `json:"finish_time"`
	ScheduledTime             string   `json:"scheduled_time"`
	ConnectionStartTime       string   `json:"connection_start_time"`
	ConnectivityResult        int      `json:"connectivity_result"`
	ConnectivityResultName    string   `json:"connectivity_result_name"`
	WirelessDurationMsec      int      `json:"wireless_duration_msec"`
	Channel                   int      `json:"channel"`
	Bssid                     string   `json:"bssid"`
	Ssid                      string   `json:"ssid"`
	Vendor                    string   `json:"vendor"`
	ConnectivityIPAddress     string   `json:"connectivity_ip_address"`
	PublicIPAddress           string   `json:"public_ip_address"`
	Gateway                   string   `json:"gateway"`
	SubnetMask                string   `json:"subnet_mask"`
	DNS1Info                  []string `json:"dns1_info"`
	DhcpServer                string   `json:"dhcp_server"`
	CaptivePortalTimeMsec     int      `json:"captive_portal_time_msec"`
	AverageRssi               int      `json:"average_rssi"`
	CurrentRssi               int      `json:"current_rssi"`
	LinkSpeedMbps             int      `json:"link_speed_mbps"`
	DhcpRetrievalDurationMsec int      `json:"dhcp_retrieval_duration_msec"`
	EapTime                   string   `json:"eap_time"`
	EapTimeMsec               int      `json:"eap_time_msec"`
	DhcpLeaseTimeSeconds      int      `json:"dhcp_lease_time_seconds"`
	DhcpThresholdSeconds      int      `json:"dhcp_threshold_seconds"`
	RadiusThresholdSeconds    int      `json:"radius_threshold_seconds"`
	DNSThresholdMsec          int      `json:"dns_threshold_msec"`
	PingStartTime             string   `json:"ping_start_time"`
	IsWireless                bool     `json:"is_wireless"`
	TotalDurationMsec         int      `json:"total_duration_msec"`
	PingDurationMsec          int      `json:"ping_duration_msec"`
}

type PingResultData struct {
	StatusID                 int    `json:"status_id"`
	Hostname                 string `json:"hostname"`
	IPAddress                string `json:"ip_address"`
	PingResultName           string `json:"ping_result_name"`
	Transmitted              int    `json:"transmitted"`
	Received                 int    `json:"received"`
	LostPercent              int    `json:"lost_percent"`
	RoundTripTimeMinUsec     int    `json:"round_trip_time_min_usec"`
	RoundTripTimeMaxUsec     int    `json:"round_trip_time_max_usec"`
	RoundTripTimeAverageUsec int    `json:"round_trip_time_average_usec"`
	DNSStatusID              int    `json:"dns_status_id"`
	PingDNSResolveTimeUsec   int    `json:"ping_dns_resolve_time_usec"`
}

type PingResults struct {
	ResultStatusID   int               `json:"result_status_id"`
	ResultStatusName string            `json:"result_status_name"`
	ScheduledTime    string            `json:"scheduled_time"`
	Details          PingResultDetails `json:"result"`
	Data             []PingResultData  `json:"data"`
}
type PingResultsResponse struct {
	NetworkTestResults PingResults `json:"network_test_results"`
}

func (c *Client) GetDetailedPingResults(ctx context.Context, execution_id int, sensor_id int, net_test_profile_id int) (*PingResultsResponse, error) {
	body := map[string]interface{}{
		"network_test_profile_id": net_test_profile_id,
		"sensor_id":               sensor_id,
		"execution_id":            execution_id,
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/test/get_detailed_network_test_results", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := PingResultsResponse{}
	if err := c.sendRequest(ctx, req, &res, "network_test_profile_detailed_result_log"); err != nil {
		return nil, err
	}
	return &res, nil
}
