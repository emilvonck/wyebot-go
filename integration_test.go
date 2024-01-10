//go:build integration
// +build integration

package wyebot

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	LocationID = 3094
	SensorID   = 30940001
)

func TestGetLocations(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetLocations(ctx)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensors(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetSensors(ctx, LocationID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensorInfo(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetSensorInfo(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensorIssues(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetSensorIssues(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensorNetworkInfo(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetSensorNetworkInfo(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetAccessPointList(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetAccessPointList(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetNetworkTestProfiles(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetNetworkTestProfiles(ctx, LocationID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetRfDetails(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetRfDetails(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestNetworkTestResults(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	ctx := context.Background()
	res, err := c.GetNetworkTestResults(ctx, 3094, 3641, "2021-02-21 00:40:00", "2024-01-10 21:00:00")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
