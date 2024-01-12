//go:build integration
// +build integration

package wyebot

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var BaseUrl = os.Getenv("WYEBOT_URL")
var ApiKey = os.Getenv("WYEBOT_API_KEY")
var LocationID, _ = strconv.Atoi(os.Getenv("WYEBOT_LOCATION_ID"))
var SensorID, _ = strconv.Atoi(os.Getenv("WYEBOT_SENSOR_ID"))

func TestGetLocations(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetLocations(ctx)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensors(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetSensors(ctx, LocationID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensorInfo(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetSensorInfo(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensorIssues(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetSensorIssues(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetSensorNetworkInfo(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetSensorNetworkInfo(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetAccessPointList(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetAccessPointList(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetNetworkTestProfiles(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetNetworkTestProfiles(ctx, LocationID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetRfDetails(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetRfDetails(ctx, SensorID)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetNetworkTestResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetNetworkTestResults(ctx, LocationID, 3641, "2021-02-21 00:40:00", "2024-01-10 21:00:00")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedSpeedTestResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedSpeedTestResults(ctx, 1705054812, 30940001, 8816)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedPingResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedPingResults(ctx, 1705059065, 30940004, 8830)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedAppResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedAppResults(ctx, 1705059633, 30940001, 8853)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedVideoResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedVideoResults(ctx, 1705058501, 30940001, 8818)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedNetDiscoveryResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedNetDiscoveryResults(ctx, 1705063581, 30940002, 8912)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedIperfResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedIperfResults(ctx, 1705065045, 30940002, 8915)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetDetailedDnsResults(t *testing.T) {
	c := NewClient(ApiKey, BaseUrl)

	ctx := context.Background()
	res, err := c.GetDetailedDnsResults(ctx, 1705089791, 30940002, 8943)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
