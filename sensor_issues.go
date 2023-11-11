package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	SeverityName       string `json:"severity_name"`
	Problem            string `json:"problem"`
	ProblemDescription string `json:"problem_description"`
	Solution           string `json:"solution"`
}
type IssueList struct {
	Issues []Issue `json:"data"`
}

func (c *Client) GetSensorIssues(ctx context.Context, sensor_id int) (*IssueList, error) {
	body := map[string]int{"sensor_id": sensor_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/dashboard/sensor_issues", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := IssueList{}
	if err := c.sendRequest(ctx, req, &res, "issue_details"); err != nil {
		return nil, err
	}
	return &res, nil
}
