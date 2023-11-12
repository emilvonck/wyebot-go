package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SensorNetworkInfo struct {
	SensorID         int    `json:"sensor_id"`
	SensorName       string `json:"sensor_name"`
	SensorStatusName string `json:"sensor_status_name"`
	ConnectionType   string `json:"connection_type"`
	Dhcp             bool   `json:"dhcp"`
	Ipaddr           string `json:"ipaddr"`
	IPSubnet         string `json:"ip_subnet"`
	IPGateway        string `json:"ip_gateway"`
	DNS1             string `json:"dns1"`
	DNS2             string `json:"dns2"`
}
type SensorNetWorkInfoResponse struct {
	SensorNetworkInfo SensorNetworkInfo `json:"data"`
}

func (c *Client) GetSensorNetworkInfo(ctx context.Context, sensor_id int) (*SensorNetWorkInfoResponse, error) {
	body := map[string]int{"sensor_id": sensor_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/org/get_sensor_network_info", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := SensorNetWorkInfoResponse{}
	if err := c.sendRequest(ctx, req, &res, "sensor_network_info"); err != nil {
		return nil, err
	}
	return &res, nil
}
