package wyebot

import (
	"context"
	"fmt"
	"net/http"
)

type LocationList struct {
	Locations []Location `json:"data"`
}

type Location struct {
	LocationID   int    `json:"location_id"`
	LocationName string `json:"location_name"`
}

func (c *Client) GetLocations(ctx context.Context) (*LocationList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/external_api/org/get_locations", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := LocationList{}
	if err := c.sendRequest(ctx, req, &res, "location_details"); err != nil {
		return nil, err
	}
	return &res, nil
}
