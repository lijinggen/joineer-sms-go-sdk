package joineer_sms_go_sdk

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type JoineerClient struct {
	cfg *JoineerConfig
}

type JoineerConfig struct {
	ApiSecret string

	httpClient *resty.Client
	userId     string
}

func NewJoineerClient(key, secret string) (*JoineerClient, error) {
	client := &JoineerClient{}
	cfg := &JoineerConfig{
		ApiSecret: secret,
	}
	c := resty.New()
	c.SetBaseURL("http://localhost:5678/v1/api")
	c.SetHeader("secret", secret)
	c.SetHeader("key", key)
	c.SetHeader("from", "sdk")
	userId, err := getUserId(c, secret)
	if err != nil {
		return nil, err
	}
	c.SetHeader("user_id", userId)
	cfg.userId = userId
	cfg.httpClient = c
	client.cfg = cfg
	return client, nil
}

type JoineerError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func getUserId(httpClient *resty.Client, secret string) (string, error) {
	type getUserIdResp struct {
		JoineerError
		UserId string `json:"user_id"`
	}
	res := getUserIdResp{}
	get, err := httpClient.NewRequest().SetResult(&res).Get("/user/get_user_id_by_secret?secret=" + secret)
	if err != nil {
		return "", err
	}
	if get.StatusCode() != http.StatusOK {
		return "", errors.New(get.Status())
	}
	if res.ErrorCode != 0 {
		return "", errors.New(res.ErrorMsg)
	}
	return res.UserId, nil
}
