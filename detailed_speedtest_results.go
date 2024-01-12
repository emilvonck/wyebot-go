package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SpeedTestResultDetails struct {
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
	DNSResolveTime            string   `json:"dns_resolve_time"`
	CaptivePortalTimeMsec     int      `json:"captive_portal_time_msec"`
	AverageRssi               int      `json:"average_rssi"`
	CurrentRssi               int      `json:"current_rssi"`
	LinkSpeedMbps             int      `json:"link_speed_mbps"`
	DhcpRetrievalDurationMsec int      `json:"dhcp_retrieval_duration_msec"`
	EapTimeMsec               int      `json:"eap_time_msec"`
	DhcpLeaseTimeHrs          int      `json:"dhcp_lease_time_hrs"`
	DhcpThresholdSeconds      int      `json:"dhcp_threshold_seconds"`
	RadiusThresholdSeconds    int      `json:"radius_threshold_seconds"`
	DNSThresholdMsec          int      `json:"dns_threshold_msec"`
	SpeedStartTime            string   `json:"speed_start_time"`
	IsWireless                bool     `json:"is_wireless"`
	TestResultStatusName      string   `json:"test_result_status_name"`
	ServerInfo                string   `json:"server_info"`
	TotalDurationSeconds      int      `json:"total_duration_seconds"`
	SpeedDurationSeconds      int      `json:"speed_duration_seconds"`
}

type SpeedTestResultData struct {
	DownloadSpeedMbps     int `json:"download_speed_mbps"`
	UploadSpeedMbps       int `json:"upload_speed_mbps"`
	DownloadThresholdMbps int `json:"download_threshold_mbps"`
	UploadThresholdMbps   int `json:"upload_threshold_mbps"`
	PingDelayMsec         int `json:"ping_delay_msec"`
}

type SpeedTestResults struct {
	ResultStatusID   int                    `json:"result_status_id"`
	ResultStatusName string                 `json:"result_status_name"`
	ScheduledTime    string                 `json:"scheduled_time"`
	Details          SpeedTestResultDetails `json:"result"`
	Data             SpeedTestResultData    `json:"data"`
}
type SpeedTestResultsResponse struct {
	NetworkTestResults SpeedTestResults `json:"network_test_results"`
}

func (c *Client) GetDetailedSpeedTestResults(ctx context.Context, execution_id int, sensor_id int, net_test_profile_id int) (*SpeedTestResultsResponse, error) {
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
	res := SpeedTestResultsResponse{}
	if err := c.sendRequest(ctx, req, &res, "network_test_profile_detailed_result_log"); err != nil {
		return nil, err
	}
	return &res, nil
}
