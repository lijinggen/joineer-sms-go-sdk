package joineer_sms_go_sdk

import "testing"

func TestBulkSend(t *testing.T) {
	client, err := NewJoineerClient("compay_s_api_key", "6212ec66caa03900fc1d178f43b3bae7485c295c1ca127d6a04181bcd2876edf4faf481129f6c0cd4323eff47132322c")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Send("13192819231", "xxx")
	if err != nil {
		t.Fatal(err)
	}
	Fat

}
