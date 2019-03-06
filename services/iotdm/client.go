package iotdm

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/YaoJunZhao/bce-sdk-go/services/iotdm/api"
)

type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}

	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:    endpoint,
		Region:      bce.DEFAULT_REGION,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption:  defaultSignOptions,
		Retry:       bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func (c *Client) CreateDevice(deviceName, describle, schemaID string) (*api.AccessDetailResponse, error) {
	return api.CreateDevice(c, deviceName, describle, schemaID)
}
