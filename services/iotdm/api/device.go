package api

import (
	"encoding/json"

	"../../../bce"
	"../../../http"
)

type deviceInfo struct {
	DeviceName  string `json:"deviceName"`
	Description string `json:"description"`
	Schemaid    string `json:"schemaId"`
}

func CreateDevice(cli bce.Client, deviceName, describle, schemaID string) (*AccessDetailResponse, error) {

	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	req.SetUri("/v3/iot/management/device")

	info := &deviceInfo{deviceName, describle, schemaID}
	data, _ := json.Marshal(info)

	body, _ := bce.NewBodyFromString(string(data))
	req.SetBody(body)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	result := &AccessDetailResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}
