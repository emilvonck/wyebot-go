package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetDiscoveryResultDetails struct {
	SensorID                    int      `json:"sensor_id"`
	SensorName                  string   `json:"sensor_name"`
	NetworkTestProfileID        int      `json:"network_test_profile_id"`
	NetworkTestProfileName      string   `json:"network_test_profile_name"`
	NetworkTestSuiteID          int      `json:"network_test_suite_id"`
	NetworkTestSuiteName        string   `json:"network_test_suite_name"`
	StartTime                   string   `json:"start_time"`
	FinishTime                  string   `json:"finish_time"`
	ScheduledTime               string   `json:"scheduled_time"`
	ConnectionStartTime         string   `json:"connection_start_time"`
	ConnectivityResult          int      `json:"connectivity_result"`
	ConnectivityResultName      string   `json:"connectivity_result_name"`
	WirelessDurationMsec        int      `json:"wireless_duration_msec"`
	Channel                     int      `json:"channel"`
	Bssid                       string   `json:"bssid"`
	Ssid                        string   `json:"ssid"`
	Vendor                      string   `json:"vendor"`
	ConnectivityIPAddress       string   `json:"connectivity_ip_address"`
	PublicIPAddress             string   `json:"public_ip_address"`
	Gateway                     string   `json:"gateway"`
	SubnetMask                  string   `json:"subnet_mask"`
	DNS1Info                    []string `json:"dns1_info"`
	DhcpServer                  string   `json:"dhcp_server"`
	DNSResolveTime              string   `json:"dns_resolve_time"`
	CaptivePortalTimeMsec       int      `json:"captive_portal_time_msec"`
	AverageRssi                 int      `json:"average_rssi"`
	CurrentRssi                 int      `json:"current_rssi"`
	LinkSpeedMbps               int      `json:"link_speed_mbps"`
	DhcpRetrievalDurationMsec   int      `json:"dhcp_retrieval_duration_msec"`
	EapTimeMsec                 int      `json:"eap_time_msec"`
	DhcpLeaseTimeSeconds        int      `json:"dhcp_lease_time_seconds"`
	DhcpThresholdSeconds        int      `json:"dhcp_threshold_seconds"`
	RadiusThresholdSeconds      int      `json:"radius_threshold_seconds"`
	DNSThresholdMsec            int      `json:"dns_threshold_msec"`
	NetworkDiscoveryResultName  string   `json:"network_discovery_result_name"`
	DeviceDiscoveryStartTime    string   `json:"device_discovery_start_time"`
	IsWireless                  bool     `json:"is_wireless"`
	DeviceCount                 int      `json:"device_count"`
	TotalDurationMsec           int      `json:"total_duration_msec"`
	DeviceDiscoveryDurationMsec int      `json:"device_discovery_duration_msec"`
	LatencyUsec                 []int    `json:"latency_usec"`
}

type NetDiscoveryResultData struct {
	IPAddress  string `json:"ip_address"`
	Hostname   string `json:"hostname"`
	MacAddress string `json:"mac_address"`
	OutVendor  string `json:"out_vendor"`
}

type NetDiscoveryResults struct {
	ResultStatusID   int                       `json:"result_status_id"`
	ResultStatusName string                    `json:"result_status_name"`
	ScheduledTime    string                    `json:"scheduled_time"`
	Details          NetDiscoveryResultDetails `json:"result"`
	Data             []NetDiscoveryResultData  `json:"data"`
}
type NetDiscoveryResultsResponse struct {
	NetworkTestResults NetDiscoveryResults `json:"network_test_results"`
}

func (c *Client) GetDetailedNetDiscoveryResults(ctx context.Context, execution_id int, sensor_id int, net_test_profile_id int) (*NetDiscoveryResultsResponse, error) {
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
	res := NetDiscoveryResultsResponse{}
	if err := c.sendRequest(ctx, req, &res, "network_test_profile_detailed_result_log"); err != nil {
		return nil, err
	}
	return &res, nil
}
