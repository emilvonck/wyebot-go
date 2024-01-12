package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeviceMonResultDetails struct {
	SensorID                  int    `json:"sensor_id"`
	SensorName                string `json:"sensor_name"`
	NetworkTestProfileID      int    `json:"network_test_profile_id"`
	NetworkTestProfileName    string `json:"network_test_profile_name"`
	NetworkTestSuiteID        int    `json:"network_test_suite_id"`
	SuiteName                 string `json:"suite_name"`
	StartTime                 string `json:"start_time"`
	FinishTime                string `json:"finish_time"`
	DeviceMonitorStartTime    string `json:"device_monitor_start_time"`
	TotalDurationMsec         int    `json:"total_duration_msec"`
	DeviceMonitorDurationMsec int    `json:"device_monitor_duration_msec"`
}

type DeviceMonResultData struct {
	StatusID            int    `json:"status_id"`
	Hostname            string `json:"hostname"`
	IPAddress           string `json:"ip_address"`
	DeviceMonitorResult string `json:"device_monitor_result"`
}

type DeviceMonResults struct {
	ResultStatusID   int                    `json:"result_status_id"`
	ResultStatusName string                 `json:"result_status_name"`
	ScheduledTime    string                 `json:"scheduled_time"`
	Details          DeviceMonResultDetails `json:"result"`
	Data             []DeviceMonResultData  `json:"data"`
}
type DeviceMonResultsResponse struct {
	NetworkTestResults DeviceMonResults `json:"network_test_results"`
}

func (c *Client) GetDetailedDeviceMonResults(ctx context.Context, execution_id int, sensor_id int, net_test_profile_id int) (*DeviceMonResultsResponse, error) {
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
	res := DeviceMonResultsResponse{}
	if err := c.sendRequest(ctx, req, &res, "network_test_profile_detailed_result_log"); err != nil {
		return nil, err
	}
	return &res, nil
}
