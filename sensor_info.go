package wyebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Specification struct {
	Model              string `json:"model"`
	SerialNumber       string `json:"serial_number"`
	WirelessMacAddress string `json:"wireless_mac_address"`
	WiredMacAddress    string `json:"wired_mac_address"`
	LinkSpeed          string `json:"link_speed"`
	PowerSource        string `json:"power_source"`
}

type Service struct {
	Uptime          string `json:"uptime"`
	SoftwareVersion string `json:"software_version"`
	LicenseInfo     string `json:"license_info"`
}

type Vlan struct {
	VlanID string `json:"vlan-id"`
	Pvid   bool   `json:"pvid"`
}

type UnknownTlv struct {
	Subtype string `json:"subtype"`
	Oui     string `json:"oui"`
	Value   string `json:"value"`
	Len     string `json:"len"`
}

type UnknownTlvs struct {
	UnknownTlv UnknownTlv `json:"unknown-tlv"`
}

type Capability struct {
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type ID struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Sw01 struct {
	Capability Capability `json:"capability"`
	MgmtIP     string     `json:"mgmt-ip"`
	ID         ID         `json:"id"`
	Descr      string     `json:"descr"`
}

type Chassis struct {
	Sw01 Sw01 `json:"sw01"`
}

type Advertised struct {
	Type string `json:"type"`
	Fd   bool   `json:"fd"`
	Hd   bool   `json:"hd"`
}

type AutoNegotiation struct {
	Current    string       `json:"current"`
	Supported  bool         `json:"supported"`
	Enabled    bool         `json:"enabled"`
	Advertised []Advertised `json:"advertised"`
}

type Port struct {
	AutoNegotiation AutoNegotiation `json:"auto-negotiation"`
	Descr           string          `json:"descr"`
	ID              ID              `json:"id"`
}

type Eth0 struct {
	Via         string      `json:"via"`
	Age         string      `json:"age"`
	Vlan        Vlan        `json:"vlan"`
	UnknownTlvs UnknownTlvs `json:"unknown-tlvs"`
	Chassis     Chassis     `json:"chassis"`
	Rid         string      `json:"rid"`
	Port        Port        `json:"port"`
}

type Interface struct {
	Eth0 Eth0 `json:"eth0"`
}

type Lldp struct {
	Interface Interface `json:"interface"`
}

type LldpInfo struct {
	Lldp Lldp `json:"lldp"`
}

type HardwareDetails struct {
	Specification Specification `json:"specification"`
	Service       Service       `json:"service"`
	//LldpInfo      LldpInfo      `json:"lldp_info"`
}
type SensorInfo struct {
	SensorID         int             `json:"sensor_id"`
	SensorName       string          `json:"sensor_name"`
	SensorStatusName string          `json:"sensor_status_name"`
	HardwareDetails  HardwareDetails `json:"hardware_details"`
}
type SensorsInfoResponse struct {
	Data SensorInfo `json:"data"`
}

func (c *Client) GetSensorInfo(ctx context.Context, sensor_id int) (*SensorsInfoResponse, error) {
	body := map[string]int{"sensor_id": sensor_id}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/external_api/org/get_sensor_info", c.baseURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := SensorsInfoResponse{}
	if err := c.sendRequest(ctx, req, &res, "sensor_info"); err != nil {
		return nil, err
	}
	return &res, nil
}
