package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RfAnalyticsRadio0 struct {
	OutChannel          string `json:"out_channel"`
	AirtimeTotalPercent string `json:"airtime_total_percent"`
	MgmtPercent         string `json:"mgmt_percent"`
	CtrlPercent         string `json:"ctrl_percent"`
	DataPercent         string `json:"data_percent"`
	OthersPercent       string `json:"others_percent"`
	AvailablePercent    string `json:"available_percent"`
	Noise               string `json:"noise"`
	//ClientMacList       []string `json:"client_mac_list"`
	//ClientHostnameList      []string `json:"client_hostname_list"`
	ClientAirtimePercentage string `json:"client_airtime_percentage"`
}
type RfAnalyticsRadio struct {
	OutChannel          string `json:"out_channel"`
	AirtimeTotalPercent string `json:"airtime_total_percent"`
	MgmtPercent         string `json:"mgmt_percent"`
	CtrlPercent         string `json:"ctrl_percent"`
	DataPercent         string `json:"data_percent"`
	OthersPercent       string `json:"others_percent"`
	AvailablePercent    string `json:"available_percent"`
	Noise               string `json:"noise"`
	// ClientMacList       string `json:"Client_mac_list"`
	// ClientHostnameList      string `json:"Client_hostname_list"`
	ClientAirtimePercentage string `json:"client_airtime_percentage"`
}
type RfDetails struct {
	Radio0 RfAnalyticsRadio `json:"rf_analytics_radio_0"`
	Radio1 RfAnalyticsRadio `json:"rf_analytics_radio_1"`
}

type RfDetailsResponse struct {
	Data RfDetails `json:"data"`
}

func (c *Client) GetRfDetails(ctx context.Context, sensor_id int) (*RfDetailsResponse, error) {
	body := map[string]int{"sensor_id": sensor_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/dashboard/rf_analytics", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := RfDetailsResponse{}
	if err := c.sendRequest(ctx, req, &res, "rf_details"); err != nil {
		return nil, err
	}
	return &res, nil
}
