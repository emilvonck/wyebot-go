//go:build integration
// +build integration

package wyebot

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	LocationID = 3094
)

func TestGetLocations(t *testing.T) {
	c := NewClient(os.Getenv("WYEBOT_INTEGRATION_API_KEY"), "https://eu-cloud.wyebot.com")

	res, err := c.GetLocations()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, LocationID, res.Locations[0].LocationID, "expecting correct LocationID")

	}
}
