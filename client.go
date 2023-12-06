package joineer_sms_go_sdk

import (
	"github.com/go-resty/resty/v2"
)

type JoineerClient struct {
	cfg *JoineerConfig
}

type JoineerConfig struct {
	ApiSecret string

	httpClient *resty.Client
}

func NewJoineerClient(secret string) *JoineerClient {
	client := &JoineerClient{}
	cfg := &JoineerConfig{
		ApiSecret: secret,
	}
	c := resty.New()
	c.SetBaseURL("http://120.24.169.86:5173/v1/api")
	c.SetHeader("secret", secret)
	cfg.httpClient = c
	return client
}

type JoineerError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}
