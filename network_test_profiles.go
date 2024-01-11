package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkTestProfile struct {
	ProfileID      int    `json:"network_test_profile_id"`
	ProfileName    string `json:"network_test_profile_name"`
	SuiteID        int    `json:"network_test_suite_id"`
	SuiteName      string `json:"network_test_suite_name"`
	Ssid           string `json:"ssid"`
	ScheduleTypeID int    `json:"schedule_type_id"`
	Schedule       string `json:"schedule"`
	Enabled        bool   `json:"enabled"`
	IsValid        bool   `json:"is_valid"`
}
type NetworkTestProfilesResponse struct {
	Data []NetworkTestProfile `json:"data"`
}

func (c *Client) GetNetworkTestProfiles(ctx context.Context, location_id int) (*NetworkTestProfilesResponse, error) {
	body := map[string]int{"location_id": location_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/test/get_network_test_profiles", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := NetworkTestProfilesResponse{}
	if err := c.sendRequest(ctx, req, &res, "network_test_profiles"); err != nil {
		return nil, err
	}
	return &res, nil
}
