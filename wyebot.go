package wyebot

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(appkey string, baseurl string) *Client {
	return &Client{
		baseURL: baseurl,
		apiKey:  appkey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type errorResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	ResponseCode int    `json:"responseCode"`
}

type successResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"results"`
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}, detail_key string) error {
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("api_key", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return errors.New(errRes.Message)
	}

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	// Try to unmarshall and normalize return for successResponse
	var parsedJsonReturn map[string]interface{}
	json.Unmarshal(responseData, &parsedJsonReturn)

	parsedJsonReturn["results"] = parsedJsonReturn[detail_key]
	delete(parsedJsonReturn, detail_key)
	newReturn, err := json.Marshal(parsedJsonReturn)
	if err != nil {
		return err
	}

	fullResponse := successResponse{
		Data: v,
	}

	if err = json.Unmarshal(newReturn, &fullResponse); err != nil {
		return err
	}

	return nil
}
