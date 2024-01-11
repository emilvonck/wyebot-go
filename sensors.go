package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SensorsResponse struct {
	Data []Sensor `json:"data"`
}

type Sensor struct {
	SensorID   int    `json:"sensor_id"`
	SensorName string `json:"sensor_name"`
}

func (c *Client) GetSensors(ctx context.Context, location_id int) (*SensorsResponse, error) {
	body := map[string]int{"location_id": location_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/org/get_sensors", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := SensorsResponse{}
	if err := c.sendRequest(ctx, req, &res, "sensor_details"); err != nil {
		return nil, err
	}
	return &res, nil
}
