package joineer_sms_go_sdk

import (
	"errors"
	"net/http"
)

func (j *JoineerClient) BulKSend(phone []string, content string) error {
	payload := map[string]interface{}{
		"phone":   phone,
		"content": content,
	}

	rawResp, err := j.cfg.httpClient.R().
		SetBody(payload).
		Post("/v1/api/bulk_send")
	if err != nil {
		return err
	}
	if rawResp.StatusCode() != http.StatusOK {
		return errors.New(rawResp.Status())
	}

	return nil
}

func (j *JoineerClient) Send(phone []string, content string) {

}
