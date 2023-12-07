package joineer_sms_go_sdk

import (
	"errors"
	"math"
	"net/http"
)

func (j *JoineerClient) BulKSend(phone []string, content string) error {
	mobileListId, err := j.batchAddMobile(phone)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"mobile_list_id": mobileListId,
		"content":        content,
		"user_id":        j.cfg.userId,
	}
	resp := &JoineerError{}
	rawResp, err := j.cfg.httpClient.R().
		SetBody(payload).
		SetResult(resp).
		Post("/bulk_send")
	if err != nil {
		return err
	}
	if rawResp.StatusCode() != http.StatusOK {
		return errors.New(rawResp.Status())
	}
	if resp.ErrorCode != 0 {
		return errors.New(resp.ErrorMsg)
	}
	return nil
}

func (j *JoineerClient) batchAddMobile(phone []string) (string, error) {
	payload := map[string]interface{}{
		"mobile_list":    phone,
		"mobile_list_id": "",
	}
	totalItems := len(phone)
	chunkSize := 100
	chunks := int(math.Ceil(float64(totalItems) / float64(chunkSize)))
	for i := 0; i < chunks; i++ {
		start := i * chunkSize
		end := int(math.Min(float64((i+1)*chunkSize), float64(totalItems)))
		sublist := phone[start:end]
		payload["mobile_list"] = sublist
		type batchAddMobileResp struct {
			JoineerError
			MobileListId string `json:"mobile_list_id"`
		}
		resp := &batchAddMobileResp{}
		rawResp, err := j.cfg.httpClient.R().SetResult(resp).SetBody(payload).Post("/batch_add_mobile")
		if err != nil {
			return "", err
		}
		if rawResp.StatusCode() != http.StatusOK {
			return "", errors.New(rawResp.Status())
		}
		if resp.ErrorCode != 0 {
			return "", errors.New(resp.ErrorMsg)
		}
		if resp.MobileListId != "" {
			payload["mobile_list_id"] = resp.MobileListId
		}
	}
	return payload["mobile_list_id"].(string), nil
}

func (j *JoineerClient) Send(phone string, content string) error {
	payload := map[string]interface{}{
		"phone":   phone,
		"content": content,
		"user_id": j.cfg.userId,
	}
	resp := &JoineerError{}
	rawResp, err := j.cfg.httpClient.R().
		SetBody(payload).
		SetResult(resp).
		Post("/send")
	if err != nil {
		return err
	}
	if rawResp.StatusCode() != http.StatusOK {
		return errors.New(rawResp.Status())
	}
	if resp.ErrorCode != 0 {
		return errors.New(resp.ErrorMsg)
	}
	return nil
}
