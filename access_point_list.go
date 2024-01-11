package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessPointDetails struct {
	MacAddress         string `json:"mac_address"`
	Hostname           string `json:"hostname"`
	HostnameTypeID     int    `json:"hostname_type_id"`
	Channel            string `json:"channel"`
	PhyType            string `json:"phy_type"`
	MaxDataRate        string `json:"max_data_rate"`
	SignalStrength     string `json:"signal_strength"`
	Vendor             string `json:"vendor"`
	ClassificationType int    `json:"classification_type"`
}
type AccessPointsDetailsResponse struct {
	Data []AccessPointDetails `json:"data"`
}

func (c *Client) GetAccessPointList(ctx context.Context, sensor_id int) (*AccessPointsDetailsResponse, error) {
	body := map[string]int{"sensor_id": sensor_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/dashboard/accesspointlist", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := AccessPointsDetailsResponse{}
	if err := c.sendRequest(ctx, req, &res, "access_point_details"); err != nil {
		return nil, err
	}
	return &res, nil
}
