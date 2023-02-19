package alarm_ex1

import (
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"testing"
)

func TestService1(t *testing.T) {
	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient1()
		service := Service{client}
		err := service.Send("01022334444")
		assert.Nil(t, err)
	})
}

func TestService2(t *testing.T) {
	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient2()
		service := Service{client}

		receiver := "01022334444"
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.Nil(t, err)
	})

	t.Run("문자 전송에 실패한다. - 네트워크 통신 에러", func(t *testing.T) {
		expect := net.DNSError{Err: "dns error"}
		client := NewMockSMSClient2()
		service := Service{client}

		receiver := "0101234139"

		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{}, &expect)

		err := service.Send(receiver)
		assert.ErrorIs(t, err, &expect)

	})

	t.Run("문자 전송에 실패한다. - 전화번호가 유효하지 않을 때", func(t *testing.T) {
		/*
			Code: 400,
			Message: invalid phone number
		*/

		client := NewMockSMSClient2()
		service := Service{client}

		receiver := "01011"
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{Code: 400, Message: "invalid phone number"}, nil)

		err := service.Send(receiver)
		assert.ErrorIs(t, err, ErrSMSFail)
	})
}
