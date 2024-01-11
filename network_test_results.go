package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkTestResult struct {
	ResultStatusID       int    `json:"result_status_id"`
	ResultStatusName     string `json:"result_status_name"`
	SensorID             int    `json:"sensor_id"`
	SensorName           string `json:"sensor_name"`
	NetworkTestSuiteID   int    `json:"network_test_suite_id"`
	NetworkTestSuiteName string `json:"network_test_suite_name"`
	StartTime            int64  `json:"start_time"`
	ScheduledTime        string `json:"scheduled_time"`
	ExecutionID          int    `json:"execution_id"`
}
type NetworkTestResultsResponse struct {
	Data []NetworkTestResult `json:"data"`
}

func (c *Client) GetNetworkTestResults(ctx context.Context, location_id int, net_test_profile_id int, start string, end string) (*NetworkTestResultsResponse, error) {
	body := map[string]interface{}{
		"location_id":             location_id,
		"network_test_profile_id": net_test_profile_id,
		"data_range_start_time":   start,
		"data_range_end_time":     end,
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/test/get_network_test_results", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := NetworkTestResultsResponse{}
	if err := c.sendRequest(ctx, req, &res, "network_test_results"); err != nil {
		return nil, err
	}
	return &res, nil
}
