package api

type AccessDetailResponse struct {
	TCPEndpoint string `json:"tcpEndpoint"`
	SslEndpoint string `json:"sslEndpoint"`
	// wssEndpoint string `json:"wssEndpoint"`
	Username string `json:"username"`
	Key      string `json:"key"`
}
