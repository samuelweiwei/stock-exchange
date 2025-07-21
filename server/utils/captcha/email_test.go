package captcha

import "testing"

func Test_aokSend_Send(t *testing.T) {
	s := &aokSend{}
	err := s.Send("", "234455")
	t.Logf("send err = %v", err)
}
